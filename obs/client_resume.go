// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

package obs

import (
	"errors"
	"fmt"
	"os"
)

// UploadFile resume uploads.
//
// This API is an encapsulated and enhanced version of multipart upload, and aims to eliminate large file
// upload failures caused by poor network conditions and program breakdowns.
func (obsClient ObsClient) UploadFile(input *UploadFileInput, extensions ...extensionOptions) (output *CompleteMultipartUploadOutput, err error) {
	if input == nil {
		return nil, errors.New("UploadFileInput is nil")
	}
	if input.EnableCheckpoint && input.CheckpointFile == "" {
		input.CheckpointFile = input.UploadFile + ".uploadfile_record"
	}

	if input.TaskNum <= 0 {
		input.TaskNum = 1
	}
	if input.PartSize < MIN_PART_SIZE {
		input.PartSize = MIN_PART_SIZE
	} else if input.PartSize > MAX_PART_SIZE {
		input.PartSize = MAX_PART_SIZE
	}

	output, err = obsClient.resumeUpload(input, extensions)
	return
}

// DownloadFile resume downloads.
//
// This API is an encapsulated and enhanced version of partial download, and aims to eliminate large file
// download failures caused by poor network conditions and program breakdowns.
func (obsClient ObsClient) DownloadFile(input *DownloadFileInput, extensions ...extensionOptions) (output *GetObjectMetadataOutput, err error) {
	if input == nil {
		return nil, errors.New("DownloadFileInput is nil")
	}
	if input.DownloadFile == "" {
		input.DownloadFile = input.Key
	}

	if input.EnableCheckpoint && input.CheckpointFile == "" {
		input.CheckpointFile = input.DownloadFile + ".downloadfile_record"
	}

	if input.TaskNum <= 0 {
		input.TaskNum = 1
	}
	if input.PartSize <= 0 {
		input.PartSize = DEFAULT_PART_SIZE
	}

	output, err = obsClient.resumeDownload(input, extensions)
	return
}

// UploadFileAsync asynchronously resume uploads a file, returning a controllable TransferTask.
//
// This API is an encapsulated and enhanced version of multipart upload, and aims to eliminate large file
// upload failures caused by poor network conditions and program breakdowns.
// The returned TransferTask supports pause, resume, and cancel operations.
func (obsClient ObsClient) UploadFileAsync(input *UploadFileInput, extensions ...extensionOptions) (task *TransferTask, err error) {
	if input == nil {
		err = errors.New("UploadFileInput is nil")
		return
	}
	if input.EnableCheckpoint && input.CheckpointFile == "" {
		input.CheckpointFile = input.UploadFile + ".uploadfile_record"
	}
	if input.TaskNum <= 0 {
		input.TaskNum = 1
	}
	if input.PartSize < MIN_PART_SIZE {
		input.PartSize = MIN_PART_SIZE
	} else if input.PartSize > MAX_PART_SIZE {
		input.PartSize = MAX_PART_SIZE
	}

	uploadFileStat, err := os.Stat(input.UploadFile)
	if err != nil {
		doLog(LEVEL_ERROR, fmt.Sprintf("Failed to stat uploadFile with error: [%v].", err))
		return
	}
	if uploadFileStat.IsDir() {
		doLog(LEVEL_ERROR, "UploadFile can not be a folder.")
		err = errors.New("uploadFile can not be a folder")
		return
	}

	ufc := &UploadCheckpoint{}
	var needCheckpoint = true
	var checkpointFilePath = input.CheckpointFile
	var enableCheckpoint = input.EnableCheckpoint
	if enableCheckpoint {
		needCheckpoint, err = getCheckpointFile(ufc, uploadFileStat, input, &obsClient, extensions)
		if err != nil {
			return
		}
	}
	if needCheckpoint {
		err = prepareUpload(ufc, uploadFileStat, input, &obsClient, extensions)
		if err != nil {
			return
		}
		if enableCheckpoint {
			err = updateCheckpointFile(ufc, checkpointFilePath)
			if err != nil {
				doLog(LEVEL_ERROR, "Failed to update checkpoint file with error [%v].", err)
				_err := abortTask(ufc.Bucket, ufc.Key, ufc.UploadId, &obsClient, extensions)
				if _err != nil {
					doLog(LEVEL_WARN, "Failed to abort task [%s].", ufc.UploadId)
				}
				return
			}
		}
	}

	task = newTransferTask(func() {
		_err := abortTask(ufc.Bucket, ufc.Key, ufc.UploadId, &obsClient, extensions)
		if _err != nil {
			doLog(LEVEL_WARN, "Failed to abort task [%s].", ufc.UploadId)
		}
		if enableCheckpoint {
			_err = os.Remove(checkpointFilePath)
			if _err != nil {
				doLog(LEVEL_WARN, "Failed to remove checkpoint file with error [%v].", _err)
			}
		}
	})

	go func() {
		uploadPartError := obsClient.uploadPartConcurrentWithTask(ufc, checkpointFilePath, input, extensions, task)

		if task.isCancelled() {
			task.cancelCleanup()
			task.finish(nil, errors.New("upload task cancelled"))
			return
		}

		if uploadPartError != nil {
			if !enableCheckpoint {
				_err := abortTask(ufc.Bucket, ufc.Key, ufc.UploadId, &obsClient, extensions)
				if _err != nil {
					doLog(LEVEL_WARN, "Failed to abort task [%s].", ufc.UploadId)
				}
			}
			task.finish(nil, uploadPartError)
			return
		}

		output, completeErr := completeParts(ufc, enableCheckpoint, checkpointFilePath, &obsClient, input.EncodingType, extensions)
		task.finish(output, completeErr)
	}()

	return
}

// DownloadFileAsync asynchronously resume downloads a file, returning a controllable TransferTask.
//
// This API is an encapsulated and enhanced version of partial download, and aims to eliminate large file
// download failures caused by poor network conditions and program breakdowns.
// The returned TransferTask supports pause, resume, and cancel operations.
func (obsClient ObsClient) DownloadFileAsync(input *DownloadFileInput, extensions ...extensionOptions) (task *TransferTask, err error) {
	if input == nil {
		err = errors.New("DownloadFileInput is nil")
		return
	}
	if input.DownloadFile == "" {
		input.DownloadFile = input.Key
	}
	if input.EnableCheckpoint && input.CheckpointFile == "" {
		input.CheckpointFile = input.DownloadFile + ".downloadfile_record"
	}
	if input.TaskNum <= 0 {
		input.TaskNum = 1
	}
	if input.PartSize <= 0 {
		input.PartSize = DEFAULT_PART_SIZE
	}

	getObjectmetaOutput, err := getObjectInfo(input, &obsClient, extensions)
	if err != nil {
		return
	}

	objectSize := getObjectmetaOutput.ContentLength
	partSize := input.PartSize
	dfc := &DownloadCheckpoint{}
	var needCheckpoint = true
	var checkpointFilePath = input.CheckpointFile
	var enableCheckpoint = input.EnableCheckpoint
	if enableCheckpoint {
		needCheckpoint, err = getDownloadCheckpointFile(dfc, input, getObjectmetaOutput)
		if err != nil {
			return
		}
	}
	if needCheckpoint {
		dfc.Bucket = input.Bucket
		dfc.Key = input.Key
		dfc.VersionId = input.VersionId
		dfc.DownloadFile = input.DownloadFile
		dfc.ObjectInfo = ObjectInfo{}
		dfc.ObjectInfo.LastModified = getObjectmetaOutput.LastModified.Unix()
		dfc.ObjectInfo.Size = getObjectmetaOutput.ContentLength
		dfc.ObjectInfo.ETag = getObjectmetaOutput.ETag
		dfc.TempFileInfo = TempFileInfo{}
		dfc.TempFileInfo.TempFileUrl = input.DownloadFile + ".tmp"
		dfc.TempFileInfo.Size = getObjectmetaOutput.ContentLength

		sliceObject(objectSize, partSize, dfc)
		_err := prepareTempFile(dfc.TempFileInfo.TempFileUrl, dfc.TempFileInfo.Size)
		if _err != nil {
			err = _err
			return
		}
		if enableCheckpoint {
			_err := updateCheckpointFile(dfc, checkpointFilePath)
			if _err != nil {
				doLog(LEVEL_ERROR, "Failed to update checkpoint file with error [%v].", _err)
				_errMsg := os.Remove(dfc.TempFileInfo.TempFileUrl)
				if _errMsg != nil {
					doLog(LEVEL_WARN, "Failed to remove temp download file with error [%v].", _errMsg)
				}
				err = _err
				return
			}
		}
	}

	task = newTransferTask(func() {
		if dfc.TempFileInfo.TempFileUrl != "" {
			_err := os.Remove(dfc.TempFileInfo.TempFileUrl)
			if _err != nil {
				doLog(LEVEL_WARN, "Failed to remove temp download file with error [%v].", _err)
			}
		}
		if enableCheckpoint {
			_err := os.Remove(checkpointFilePath)
			if _err != nil {
				doLog(LEVEL_WARN, "Failed to remove checkpoint file with error [%v].", _err)
			}
		}
	})

	go func() {
		downloadFileError := obsClient.downloadFileConcurrentWithTask(input, dfc, extensions, task)

		if task.isCancelled() {
			task.cancelCleanup()
			task.finish(nil, errors.New("download task cancelled"))
			return
		}

		if downloadFileError != nil {
			if !enableCheckpoint {
				_err := os.Remove(dfc.TempFileInfo.TempFileUrl)
				if _err != nil {
					doLog(LEVEL_WARN, "Failed to remove temp download file with error [%v].", _err)
				}
			}
			task.finish(nil, downloadFileError)
			return
		}

		renameErr := os.Rename(dfc.TempFileInfo.TempFileUrl, input.DownloadFile)
		if renameErr != nil {
			doLog(LEVEL_ERROR, "Failed to rename temp download file [%s] to download file [%s] with error [%v].", dfc.TempFileInfo.TempFileUrl, input.DownloadFile, renameErr)
			task.finish(nil, renameErr)
			return
		}
		if enableCheckpoint {
			_err := os.Remove(checkpointFilePath)
			if _err != nil {
				doLog(LEVEL_WARN, "Download file successfully, but remove checkpoint file failed with error [%v].", _err)
			}
		}
		task.finish(getObjectmetaOutput, nil)
	}()

	return
}
