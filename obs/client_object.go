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
	"io"
	"os"
	"strings"
)

// ListObjects lists objects in a bucket.
//
// You can use this API to list objects in a bucket. By default, a maximum of 1000 objects are listed.
func (obsClient ObsClient) ListObjects(input *ListObjectsInput, extensions ...ExtensionOptions) (output *ListObjectsOutput, err error) {
	if input == nil {
		return nil, errors.New("ListObjectsInput is nil")
	}
	output = &ListObjectsOutput{}
	err = obsClient.doActionWithBucket("ListObjects", HTTP_GET, input.Bucket, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		if location, ok := output.ResponseHeaders[HEADER_BUCKET_REGION]; ok {
			output.Location = location[0]
		}
		if output.EncodingType == "url" {
			err = decodeListObjectsOutput(output)
			if err != nil {
				doLog(LEVEL_ERROR, "Failed to get ListObjectsOutput with error: %v.", err)
				output = nil
			}
		}
	}
	return
}

// ListVersions lists versioning objects in a bucket.
//
// You can use this API to list versioning objects in a bucket. By default, a maximum of 1000 versioning objects are listed.
func (obsClient ObsClient) ListVersions(input *ListVersionsInput, extensions ...ExtensionOptions) (output *ListVersionsOutput, err error) {
	if input == nil {
		return nil, errors.New("ListVersionsInput is nil")
	}
	output = &ListVersionsOutput{}
	err = obsClient.doActionWithBucket("ListVersions", HTTP_GET, input.Bucket, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		if location, ok := output.ResponseHeaders[HEADER_BUCKET_REGION]; ok {
			output.Location = location[0]
		}
		if output.EncodingType == "url" {
			err = decodeListVersionsOutput(output)
			if err != nil {
				doLog(LEVEL_ERROR, "Failed to get ListVersionsOutput with error: %v.", err)
				output = nil
			}
		}
	}
	return
}

// HeadObject checks whether an object exists.
//
// You can use this API to check whether an object exists.
func (obsClient ObsClient) HeadObject(input *HeadObjectInput, extensions ...ExtensionOptions) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("HeadObjectInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucketAndKey("HeadObject", HTTP_HEAD, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	}
	return
}

// SetObjectMetadata sets object metadata.
func (obsClient ObsClient) SetObjectMetadata(input *SetObjectMetadataInput, extensions ...ExtensionOptions) (output *SetObjectMetadataOutput, err error) {
	output = &SetObjectMetadataOutput{}
	err = obsClient.doActionWithBucketAndKey("SetObjectMetadata", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseSetObjectMetadataOutput(output)
	}
	return
}

// DeleteObject deletes an object.
//
// You can use this API to delete an object from a specified bucket.
func (obsClient ObsClient) DeleteObject(input *DeleteObjectInput, extensions ...ExtensionOptions) (output *DeleteObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("DeleteObjectInput is nil")
	}
	output = &DeleteObjectOutput{}
	err = obsClient.doActionWithBucketAndKey("DeleteObject", HTTP_DELETE, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseDeleteObjectOutput(output)
	}
	return
}

// DeleteObjects deletes objects in a batch.
//
// You can use this API to batch delete objects from a specified bucket.
func (obsClient ObsClient) DeleteObjects(input *DeleteObjectsInput, extensions ...ExtensionOptions) (output *DeleteObjectsOutput, err error) {
	if input == nil {
		return nil, errors.New("DeleteObjectsInput is nil")
	}
	output = &DeleteObjectsOutput{}
	err = obsClient.doActionWithBucket("DeleteObjects", HTTP_POST, input.Bucket, input, output, extensions)
	if err != nil {
		output = nil
	} else if output.EncodingType == "url" {
		err = decodeDeleteObjectsOutput(output)
		if err != nil {
			doLog(LEVEL_ERROR, "Failed to get DeleteObjectsOutput with error: %v.", err)
			output = nil
		}
	}
	return
}

// SetObjectAcl sets ACL for an object.
//
// You can use this API to set the ACL for an object in a specified bucket.
func (obsClient ObsClient) SetObjectAcl(input *SetObjectAclInput, extensions ...ExtensionOptions) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetObjectAclInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucketAndKey("SetObjectAcl", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	}
	return
}

// GetObjectAcl gets the ACL of an object.
//
// You can use this API to obtain the ACL of an object in a specified bucket.
func (obsClient ObsClient) GetObjectAcl(input *GetObjectAclInput, extensions ...ExtensionOptions) (output *GetObjectAclOutput, err error) {
	if input == nil {
		return nil, errors.New("GetObjectAclInput is nil")
	}
	output = &GetObjectAclOutput{}
	err = obsClient.doActionWithBucketAndKey("GetObjectAcl", HTTP_GET, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		if versionID, ok := output.ResponseHeaders[HEADER_VERSION_ID]; ok {
			output.VersionId = versionID[0]
		}
	}
	return
}

// RestoreObject restores an object.
func (obsClient ObsClient) RestoreObject(input *RestoreObjectInput, extensions ...ExtensionOptions) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("RestoreObjectInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucketAndKey("RestoreObject", HTTP_POST, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	}
	return
}

// GetObjectMetadata gets object metadata.
//
// You can use this API to send a HEAD request to the object of a specified bucket to obtain its metadata.
func (obsClient ObsClient) GetObjectMetadata(input *GetObjectMetadataInput, extensions ...ExtensionOptions) (output *GetObjectMetadataOutput, err error) {
	if input == nil {
		return nil, errors.New("GetObjectMetadataInput is nil")
	}
	output = &GetObjectMetadataOutput{}
	err = obsClient.doActionWithBucketAndKey("GetObjectMetadata", HTTP_HEAD, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseGetObjectMetadataOutput(output)
	}
	return
}

func (obsClient ObsClient) GetAttribute(input *GetAttributeInput, extensions ...ExtensionOptions) (output *GetAttributeOutput, err error) {
	if input == nil {
		return nil, errors.New("GetAttributeInput is nil")
	}
	output = &GetAttributeOutput{}
	err = obsClient.doActionWithBucketAndKey("GetAttribute", HTTP_HEAD, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseGetAttributeOutput(output)
	}
	return
}

// GetObject downloads object.
//
// You can use this API to download an object in a specified bucket.
func (obsClient ObsClient) GetObject(input *GetObjectInput, extensions ...ExtensionOptions) (output *GetObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("GetObjectInput is nil")
	}
	output = &GetObjectOutput{}
	err = obsClient.doActionWithBucketAndKey("GetObject", HTTP_GET, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseGetObjectOutput(output)
	}
	return
}

// PutObject uploads an object to the specified bucket.
func (obsClient ObsClient) PutObject(input *PutObjectInput, extensions ...ExtensionOptions) (output *PutObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("PutObjectInput is nil")
	}

	if input.ContentType == "" && input.Key != "" {
		if contentType, ok := mimeTypes[strings.ToLower(input.Key[strings.LastIndex(input.Key, ".")+1:])]; ok {
			input.ContentType = contentType
		}
	}
	output = &PutObjectOutput{}
	var repeatable bool
	if input.Body != nil {
		if _, ok := input.Body.(*strings.Reader); ok {
			repeatable = true
		}
		if input.ContentLength > 0 {
			input.Body = &readerWrapper{reader: input.Body, totalCount: input.ContentLength}
		}
	}
	if repeatable {
		err = obsClient.doActionWithBucketAndKey("PutObject", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	} else {
		err = obsClient.doActionWithBucketAndKeyUnRepeatable("PutObject", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	}
	if err != nil {
		output = nil
	} else {
		ParsePutObjectOutput(output)
	}
	return
}

func (obsClient ObsClient) getContentType(input *PutObjectInput, sourceFile string) (contentType string) {
	if contentType, ok := mimeTypes[strings.ToLower(input.Key[strings.LastIndex(input.Key, ".")+1:])]; ok {
		return contentType
	}
	if contentType, ok := mimeTypes[strings.ToLower(sourceFile[strings.LastIndex(sourceFile, ".")+1:])]; ok {
		return contentType
	}
	return
}

func (obsClient ObsClient) isGetContentType(input *PutObjectInput) bool {
	if input.ContentType == "" && input.Key != "" {
		return true
	}
	return false
}

func (obsClient ObsClient) NewFolder(input *NewFolderInput, extensions ...ExtensionOptions) (output *NewFolderOutput, err error) {
	if input == nil {
		return nil, errors.New("NewFolderInput is nil")
	}

	if !strings.HasSuffix(input.Key, "/") {
		input.Key += "/"
	}

	output = &NewFolderOutput{}
	err = obsClient.doActionWithBucketAndKey("NewFolder", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseNewFolderOutput(output)
		output.ObjectUrl = fmt.Sprintf("%s/%s/%s", obsClient.conf.endpoint, input.Bucket, input.Key)
	}
	return
}

// PutFile uploads a file to the specified bucket.
func (obsClient ObsClient) PutFile(input *PutFileInput, extensions ...ExtensionOptions) (output *PutObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("PutFileInput is nil")
	}

	var body io.Reader
	sourceFile := strings.TrimSpace(input.SourceFile)
	if sourceFile != "" {
		fd, _err := os.Open(sourceFile)
		if _err != nil {
			err = _err
			return nil, err
		}
		defer func() {
			errMsg := fd.Close()
			if errMsg != nil {
				doLog(LEVEL_WARN, "Failed to close file with reason: %v", errMsg)
			}
		}()

		stat, _err := fd.Stat()
		if _err != nil {
			err = _err
			return nil, err
		}
		fileReaderWrapper := &fileReaderWrapper{filePath: sourceFile}
		fileReaderWrapper.reader = fd
		if input.ContentLength > 0 {
			if input.ContentLength > stat.Size() {
				input.ContentLength = stat.Size()
			}
			fileReaderWrapper.totalCount = input.ContentLength
		} else {
			fileReaderWrapper.totalCount = stat.Size()
		}
		body = fileReaderWrapper
	}

	_input := &PutObjectInput{}
	_input.PutObjectBasicInput = input.PutObjectBasicInput
	_input.Body = body

	if obsClient.isGetContentType(_input) {
		_input.ContentType = obsClient.getContentType(_input, sourceFile)
	}

	output = &PutObjectOutput{}
	err = obsClient.doActionWithBucketAndKey("PutFile", HTTP_PUT, _input.Bucket, _input.Key, _input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParsePutObjectOutput(output)
	}
	return
}

// CopyObject creates a copy for an existing object.
//
// You can use this API to create a copy for an object in a specified bucket.
func (obsClient ObsClient) CopyObject(input *CopyObjectInput, extensions ...ExtensionOptions) (output *CopyObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("CopyObjectInput is nil")
	}

	if strings.TrimSpace(input.CopySourceBucket) == "" {
		return nil, errors.New("Source bucket is empty")
	}
	if strings.TrimSpace(input.CopySourceKey) == "" {
		return nil, errors.New("Source key is empty")
	}

	output = &CopyObjectOutput{}
	err = obsClient.doActionWithBucketAndKey("CopyObject", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	} else {
		ParseCopyObjectOutput(output)
	}
	return
}

func (obsClient ObsClient) AppendObject(input *AppendObjectInput, extensions ...ExtensionOptions) (output *AppendObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("AppendObjectInput is nil")
	}

	if input.ContentType == "" && input.Key != "" {
		if contentType, ok := mimeTypes[strings.ToLower(input.Key[strings.LastIndex(input.Key, ".")+1:])]; ok {
			input.ContentType = contentType
		}
	}
	output = &AppendObjectOutput{}
	var repeatable bool
	if input.Body != nil {
		if _, ok := input.Body.(*strings.Reader); ok {
			repeatable = true
		}
		if input.ContentLength > 0 {
			input.Body = &readerWrapper{reader: input.Body, totalCount: input.ContentLength}
		}
	}
	if repeatable {
		err = obsClient.doActionWithBucketAndKey("AppendObject", HTTP_POST, input.Bucket, input.Key, input, output, extensions)
	} else {
		err = obsClient.doActionWithBucketAndKeyUnRepeatable("AppendObject", HTTP_POST, input.Bucket, input.Key, input, output, extensions)
	}
	if err != nil {
		output = nil
	} else {
		if err = ParseAppendObjectOutput(output); err != nil {
			output = nil
		}
	}
	return
}

func (obsClient ObsClient) ModifyObject(input *ModifyObjectInput, extensions ...ExtensionOptions) (output *ModifyObjectOutput, err error) {
	if input == nil {
		return nil, errors.New("ModifyObjectInput is nil")
	}

	output = &ModifyObjectOutput{}
	var repeatable bool
	if input.Body != nil {
		if _, ok := input.Body.(*strings.Reader); ok {
			repeatable = true
		}
		if input.ContentLength > 0 {
			input.Body = &readerWrapper{reader: input.Body, totalCount: input.ContentLength}
		}
	}
	if repeatable {
		err = obsClient.doActionWithBucketAndKey("ModifyObject", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	} else {
		err = obsClient.doActionWithBucketAndKeyUnRepeatable("ModifyObject", HTTP_PUT, input.Bucket, input.Key, input, output, extensions)
	}
	if err != nil {
		output = nil
	} else {
		ParseModifyObjectOutput(output)
	}
	return
}

func (obsClient ObsClient) RenameFile(input *RenameFileInput, extensions ...ExtensionOptions) (output *RenameFileOutput, err error) {
	if input == nil {
		return nil, errors.New("RenameFileInput is nil")
	}

	output = &RenameFileOutput{}
	err = obsClient.doActionWithBucketAndKey("RenameFile", HTTP_POST, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	}
	return
}

func (obsClient ObsClient) RenameFolder(input *RenameFolderInput, extensions ...ExtensionOptions) (output *RenameFolderOutput, err error) {
	if input == nil {
		return nil, errors.New("RenameFolderInput is nil")
	}

	if !strings.HasSuffix(input.Key, "/") {
		input.Key += "/"
	}
	if !strings.HasSuffix(input.NewObjectKey, "/") {
		input.NewObjectKey += "/"
	}
	output = &RenameFolderOutput{}
	err = obsClient.doActionWithBucketAndKey("RenameFolder", HTTP_POST, input.Bucket, input.Key, input, output, extensions)
	if err != nil {
		output = nil
	}
	return
}
