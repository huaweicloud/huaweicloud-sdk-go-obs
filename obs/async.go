/*
 * Copyright (C) 2024 Huawei Device Co., Ltd.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package obs

import (
	"errors"
	"sync"
)

// AsyncItem is the interface for applying data to target
type AsyncItem interface {
	Apply(target interface{})
}

// AsyncFlusher is the interface for flushing data to storage
type AsyncFlusher interface {
	// Flush is called after Apply has been called on the item
	Flush(item AsyncItem) error
	// FlushCheckpoint is called when a batch of items should be flushed to disk
	// Returns true if a flush actually happened
	FlushCheckpoint() bool
}

// AsyncWriter is a generic async writer
// Receives items from channel and calls Flusher to write in batches
// Flush operations are performed asynchronously to avoid blocking the caller
type AsyncWriter struct {
	ch               chan AsyncItem
	done             chan struct{}
	flushReq         chan chan error
	target           interface{}
	flusher          AsyncFlusher
	wg               sync.WaitGroup
	closeOnce        sync.Once
	batchSize        int
	applyImmediately bool

	// Async flush coordination
	flushWg     sync.WaitGroup // tracks in-flight async flush operations
	flushSignal chan struct{}   // buffered channel: each signal = one pending flush
	flushErr    error           // first async flush error (safe to read after flushWg.Wait)
}

// AsyncWriterOption is the configuration option for AsyncWriter
type AsyncWriterOption func(*AsyncWriter)

// WithChannelCapacity sets the channel capacity (default 10000)
func WithChannelCapacity(n int) AsyncWriterOption {
	return func(aw *AsyncWriter) {
		if n > 0 {
			aw.ch = make(chan AsyncItem, n)
		}
	}
}

// WithBatchSize enables batch mode with specified batch size
// When batch size > 1, items are accumulated and flushed together
func WithBatchSize(n int) AsyncWriterOption {
	return func(aw *AsyncWriter) {
		if n > 0 {
			aw.batchSize = n
		}
	}
}

// WithApplyImmediately enables immediate apply mode
// When enabled, Apply is called on the item before Flush
func WithApplyImmediately() AsyncWriterOption {
	return func(aw *AsyncWriter) {
		aw.applyImmediately = true
	}
}

// NewAsyncWriter creates an AsyncWriter
// The writer runs two background goroutines: one for processing items, one for async flushing
func NewAsyncWriter(target interface{}, flusher AsyncFlusher, opts ...AsyncWriterOption) *AsyncWriter {
	aw := &AsyncWriter{
		ch:          make(chan AsyncItem, 10000),
		done:        make(chan struct{}),
		flushReq:    make(chan chan error),
		target:      target,
		flusher:     flusher,
		batchSize:   1, // Default: flush every item
		flushSignal: make(chan struct{}, 128),
	}
	for _, opt := range opts {
		opt(aw)
	}
	aw.wg.Add(2) // one for run(), one for flushProcessor()
	go aw.run()
	go aw.flushProcessor()
	return aw
}

// Submit submits an item. It returns immediately without blocking.
func (aw *AsyncWriter) Submit(item AsyncItem) {
	select {
	case aw.ch <- item:
	case <-aw.done:
		doLog(LEVEL_WARN, "async writer already shut down, dropping item")
	}
}

// Shutdown gracefully shuts down and waits for all pending writes
func (aw *AsyncWriter) Shutdown() {
	aw.closeOnce.Do(func() {
		close(aw.done)
		aw.wg.Wait()
	})
}

// SyncFlush forces a synchronous flush of all pending items to disk.
// Called before shutdown to ensure no data loss on crash.
func (aw *AsyncWriter) SyncFlush() error {
	resultCh := make(chan error, 1)
	select {
	case aw.flushReq <- resultCh:
		return <-resultCh
	case <-aw.done:
		return errors.New("async writer already shut down")
	}
}

// processItem handles a single item: apply + flush
func (aw *AsyncWriter) processItem(item AsyncItem) {
	if aw.applyImmediately {
		item.Apply(aw.target)
	}
	if err := aw.flusher.Flush(item); err != nil {
		doLog(LEVEL_WARN, "flush item failed: %v", err)
	}
}

// drainChannel drains all items from the channel and processes them
func (aw *AsyncWriter) drainChannel() int {
	count := 0
	for {
		select {
		case item := <-aw.ch:
			aw.processItem(item)
			count++
		default:
			return count
		}
	}
}

func (aw *AsyncWriter) run() {
	defer aw.wg.Done()

	batchCount := 0

	for {
		select {
		case item := <-aw.ch:
			aw.processItem(item)
			batchCount++
			if aw.batchSize > 1 && batchCount >= aw.batchSize {
				aw.triggerAsyncFlush()
				batchCount = 0
			}

		case resultCh := <-aw.flushReq:
			batchCount += aw.drainChannel()
			aw.flushWg.Wait()
			if aw.flushErr != nil {
				resultCh <- aw.flushErr
				return
			}
			if batchCount > 0 {
				if !aw.flusher.FlushCheckpoint() {
					resultCh <- errors.New("checkpoint flush failed")
					return
				}
				batchCount = 0
			}
			resultCh <- nil

		case <-aw.done:
			batchCount += aw.drainChannel()
			aw.flushWg.Wait()
			if batchCount > 0 {
				aw.flusher.FlushCheckpoint()
			}
			return
		}
	}
}

// triggerAsyncFlush signals the flush processor to perform a flush.
// This returns immediately without waiting for the actual disk I/O.
func (aw *AsyncWriter) triggerAsyncFlush() {
	aw.flushWg.Add(1)
	select {
	case aw.flushSignal <- struct{}{}:
	default:
		// Channel full (capacity 128) — batch flushes are coalesced.
		// The existing signal in the channel will trigger one FlushCheckpoint
		// call, which persists the current state including this batch's data.
		// Undo the Add since no corresponding Done() will fire.
		aw.flushWg.Done()
	}
}

// flushProcessor runs in a separate goroutine to handle disk I/O asynchronously.
// This is only used when batchSize > 1 for background batch flushing.
func (aw *AsyncWriter) flushProcessor() {
	defer aw.wg.Done()

	for {
		select {
		case _, ok := <-aw.flushSignal:
			if !ok {
				return
			}
			if !aw.flusher.FlushCheckpoint() && aw.flushErr == nil {
				aw.flushErr = errors.New("async checkpoint flush failed")
			}
			aw.flushWg.Done()

		case <-aw.done:
			// Drain remaining flush signals before exit to prevent flushWg deadlock
			for {
				select {
				case _, ok := <-aw.flushSignal:
					if !ok {
						return
					}
					aw.flusher.FlushCheckpoint()
					aw.flushWg.Done()
				default:
					return
				}
			}
		}
	}
}
