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

/**
 * This sample demonstrates how to do advanced bucket-related operations
 * (such as direct cold access, cross-region replication, WORM, decompression)
 * on OBS using the OBS SDK for Go.
 */
package examples

import (
	"fmt"
	"time"

	"obs-sdk-go/obs"
)

type BucketAdvancedOperationsSample struct {
	bucketName string
	objectKey  string
	location   string
	obsClient  *obs.ObsClient
}

func newBucketAdvancedOperationsSample(ak, sk, endpoint, bucketName, objectKey, location string) *BucketAdvancedOperationsSample {
	obsClient, err := obs.New(ak, sk, endpoint)
	if err != nil {
		panic(err)
	}
	return &BucketAdvancedOperationsSample{obsClient: obsClient, bucketName: bucketName, objectKey: objectKey, location: location}
}

// DoDirectColdAccessOperation demonstrates how to enable direct cold access for archived objects.
// When direct cold access is enabled, archived objects can be downloaded without restore.
func (sample BucketAdvancedOperationsSample) DoDirectColdAccessOperation() {
	// 1. Enable direct cold access
	setInput := &obs.SetBucketDirectColdAccessInput{}
	setInput.Bucket = sample.bucketName
	setInput.DirectColdAccessConfiguration.Status = obs.DirectColdAccessStatusEnabled
	_, err := sample.obsClient.SetBucketDirectColdAccess(setInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enable direct cold access successfully!")
	fmt.Println()

	// 2. Get direct cold access configuration
	output, err := sample.obsClient.GetBucketDirectColdAccess(sample.bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Direct cold access status: %s\n", output.Status)
	fmt.Println()

	// 3. Disable direct cold access
	setInput.DirectColdAccessConfiguration.Status = obs.DirectColdAccessStatusDisabled
	_, err = sample.obsClient.SetBucketDirectColdAccess(setInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Disable direct cold access successfully!")
	fmt.Println()

	// 4. Get direct cold access configuration again
	output, err = sample.obsClient.GetBucketDirectColdAccess(sample.bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Direct cold access status: %s\n", output.Status)
	fmt.Println()

	// 5. Delete direct cold access configuration
	_, err = sample.obsClient.DeleteBucketDirectColdAccess(sample.bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete direct cold access configuration successfully!")
	fmt.Println()
}

// DoBucketReplicationOperation demonstrates how to configure cross-region replication for a bucket.
func (sample BucketAdvancedOperationsSample) DoBucketReplicationOperation() {
	// 1. Set bucket replication configuration
	setInput := &obs.SetBucketReplicationInput{}
	setInput.Bucket = sample.bucketName
	setInput.Agency = "your-agency-name"

	var rules []obs.ReplicationRule
	rule := obs.ReplicationRule{}
	rule.ID = "rule-1"
	rule.Prefix = "prefix"
	rule.Status = "Enabled"
	rule.Destination.Bucket = "target-bucket-name"
	rule.Destination.StorageClass = "STANDARD"
	rule.HistoricalObjectReplication = "Enabled"
	rules = append(rules, rule)
	setInput.Rules = rules

	_, err := sample.obsClient.SetBucketReplication(setInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Set bucket replication configuration successfully!")
	fmt.Println()

	// 2. Get bucket replication configuration
	output, err := sample.obsClient.GetBucketReplication(sample.bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Replication agency: %s\n", output.Agency)
	for index, rule := range output.Rules {
		fmt.Printf("Rule[%d] - ID: %s, Prefix: %s, Status: %s\n", index, rule.ID, rule.Prefix, rule.Status)
		fmt.Printf("  Destination bucket: %s, StorageClass: %s\n", rule.Destination.Bucket, rule.Destination.StorageClass)
	}
	fmt.Println()

	// 3. Delete bucket replication configuration
	_, err = sample.obsClient.DeleteBucketReplication(sample.bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete bucket replication configuration successfully!")
	fmt.Println()
}

// DoBucketWormConfigurationOperation demonstrates how to configure bucket-level default WORM policy.
func (sample BucketAdvancedOperationsSample) DoBucketWormConfigurationOperation() {
	// 1. Set bucket WORM configuration (enable object lock)
	setInput := &obs.SetBucketWormConfigurationInput{}
	setInput.Bucket = sample.bucketName
	setInput.ObjectLockConfiguration.ObjectLockEnabled = "Enabled"
	setInput.ObjectLockConfiguration.Rule = &obs.Rule{}
	setInput.ObjectLockConfiguration.Rule.DefaultRetention = &obs.DefaultRetention{}
	setInput.ObjectLockConfiguration.Rule.DefaultRetention.Mode = "COMPLIANCE"
	setInput.ObjectLockConfiguration.Rule.DefaultRetention.Days = 30

	_, err := sample.obsClient.SetBucketWormConfiguration(setInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Set bucket WORM configuration successfully!")
	fmt.Println()

	// 2. Get bucket WORM configuration
	output, err := sample.obsClient.GetBucketWormConfiguration(sample.bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ObjectLockEnabled: %s\n", output.ObjectLockEnabled)
	if output.Rule != nil && output.Rule.DefaultRetention != nil {
		fmt.Printf("DefaultRetention Mode: %s\n", output.Rule.DefaultRetention.Mode)
		fmt.Printf("DefaultRetention Days: %d\n", output.Rule.DefaultRetention.Days)
	}
	fmt.Println()
}

// DoObjectRetentionOperation demonstrates how to configure object-level WORM protection.
func (sample BucketAdvancedOperationsSample) DoObjectRetentionOperation() {
	// 1. First, ensure the bucket has WORM enabled
	setWormInput := &obs.SetBucketWormConfigurationInput{}
	setWormInput.Bucket = sample.bucketName
	setWormInput.ObjectLockConfiguration.ObjectLockEnabled = "Enabled"
	_, err := sample.obsClient.SetBucketWormConfiguration(setWormInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enable bucket WORM successfully!")
	fmt.Println()

	// 2. Put object retention (WORM protection for specific object)
	// RetainUntilDate is the Unix timestamp in seconds when the retention period expires
	retainUntilDate := time.Now().Add(30 * 24 * time.Hour).Unix()
	putRetentionInput := &obs.PutObjectRetentionInput{}
	putRetentionInput.Bucket = sample.bucketName
	putRetentionInput.Key = sample.objectKey
	putRetentionInput.RetentionMode = string(obs.WormRetentionModeCompliance)
	putRetentionInput.RetainUntilDate = retainUntilDate

	_, err = sample.obsClient.PutObjectRetention(putRetentionInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Set object retention successfully!")
	fmt.Printf("Retention Mode: %s, Retain Until: %s\n", obs.WormRetentionModeCompliance, time.Unix(retainUntilDate, 0).Format(time.RFC3339))
	fmt.Println()
}

// DoBucketDecompressionOperation demonstrates how to configure decompression policy for a bucket.
func (sample BucketAdvancedOperationsSample) DoBucketDecompressionOperation() {
	// 1. Set bucket decompression policy
	setInput := &obs.SetBucketDecompressionInput{}
	setInput.Bucket = sample.bucketName

	var rules []obs.DecompressionRule
	rule := obs.DecompressionRule{}
	rule.ID = fmt.Sprintf("rule-%d", time.Now().Unix())
	rule.Project = "your-project-id"
	rule.Agency = "your-agency-name"
	rule.Events = []string{"ObjectCreated:Post"}
	rule.Prefix = "compressed/"
	rule.Suffix = ".zip"
	rule.Overwrite = 1
	rule.DecompressPath = "extracted/"
	rule.PolicyType = "decompress"
	rules = append(rules, rule)
	setInput.Rules = rules

	_, err := sample.obsClient.SetBucketDecompression(setInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Set bucket decompression policy successfully!")
	fmt.Println()

	// 2. Get bucket decompression policy
	getInput := &obs.GetBucketDecompressionInput{}
	getInput.Bucket = sample.bucketName

	output, err := sample.obsClient.GetBucketDecompression(getInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decompression rules count: %d\n", len(output.Rules))
	for index, rule := range output.Rules {
		fmt.Printf("Rule[%d] - ID: %s, Prefix: %s, Suffix: %s, PolicyType: %s\n",
			index, rule.ID, rule.Prefix, rule.Suffix, rule.PolicyType)
	}
	fmt.Println()

	// 3. Delete bucket decompression policy
	deleteInput := &obs.DeleteBucketDecompressionInput{}
	deleteInput.Bucket = sample.bucketName

	_, err = sample.obsClient.DeleteBucketDecompression(deleteInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete bucket decompression policy successfully!")
	fmt.Println()
}

func RunBucketAdvancedOperationsSample() {
	const (
		endpoint   = "https://your-endpoint"
		ak         = "*** Provide your Access Key ***"
		sk         = "*** Provide your Secret Key ***"
		bucketName = "bucket-test"
		objectKey  = "object-test"
		location   = "yourbucketlocation"
	)

	sample := newBucketAdvancedOperationsSample(ak, sk, endpoint, bucketName, objectKey, location)

	// Direct cold access operations (archive direct read)
	sample.DoDirectColdAccessOperation()

	// Cross-region replication operations
	sample.DoBucketReplicationOperation()

	// Bucket-level WORM configuration
	sample.DoBucketWormConfigurationOperation()

	// Object-level WORM protection
	sample.DoObjectRetentionOperation()

	// Decompression policy operations
	sample.DoBucketDecompressionOperation()
}