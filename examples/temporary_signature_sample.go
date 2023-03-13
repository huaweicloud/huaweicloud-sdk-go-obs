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
 * This sample demonstrates how to do common operations in temporary signature way
 * on OBS using the OBS SDK for Go.
 */
package examples

import (
	"fmt"
	"io/ioutil"
	"obs"
	"os"
	"path/filepath"
	"strings"
)

type TemporarySignatureSample struct {
	bucketName string
	objectKey  string
	location   string
	obsClient  *obs.ObsClient
}

func newTemporarySignatureSample(ak, sk, endpoint, bucketName, objectKey, location string) *TemporarySignatureSample {
	obsClient, err := obs.New(ak, sk, endpoint)
	if err != nil {
		fmt.Println(err)
	}
	return &TemporarySignatureSample{obsClient: obsClient, bucketName: bucketName, objectKey: objectKey, location: location}
}

func (sample TemporarySignatureSample) CreateBucket() {
	input := &obs.CreateSignedUrlInput{}
	input.Bucket = sample.bucketName
	input.Method = obs.HttpMethodPut
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "CreateBucket")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	data := strings.NewReader(fmt.Sprintf("<CreateBucketConfiguration><LocationConstraint>%s</LocationConstraint></CreateBucketConfiguration>", sample.location))

	_, err = sample.obsClient.CreateBucketWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders, data)
	if err == nil {
		fmt.Printf("Create bucket:%s successfully!\n", sample.bucketName)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func (sample TemporarySignatureSample) ListBuckets() {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodGet
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "ListBuckets")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	listBucketsOutput, err := sample.obsClient.ListBucketsWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err == nil {
		fmt.Printf("Owner.DisplayName:%s, Owner.ID:%s\n", listBucketsOutput.Owner.DisplayName, listBucketsOutput.Owner.ID)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	for index, val := range listBucketsOutput.Buckets {
		fmt.Printf("Bucket[%d]-Name:%s,CreationDate:%s\n", index, val.Name, val.CreationDate)
	}
	fmt.Println()
}

func (sample TemporarySignatureSample) DoBucketCors() {

	rawData := "<CORSConfiguration>" +
		"<CORSRule>" +
		"<AllowedOrigin>http://www.a.com</AllowedOrigin>" +
		"<AllowedMethod>PUT</AllowedMethod>" +
		"<AllowedMethod>POST</AllowedMethod>" +
		"<AllowedMethod>DELETE</AllowedMethod>" +
		"<AllowedHeader>*</AllowedHeader>" +
		"</CORSRule>" +
		"<CORSRule>" +
		"<AllowedOrigin>http://www.b.com</AllowedOrigin>" +
		"<AllowedMethod>GET</AllowedMethod>" +
		"</CORSRule>" +
		"</CORSConfiguration>"

	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodPut
	input.Bucket = sample.bucketName
	input.SubResource = obs.SubResourceCors
	input.Expires = 3600
	input.Headers = map[string]string{obs.HEADER_MD5_CAMEL: obs.Base64Md5([]byte(rawData))}
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "SetBucketCors")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	data := strings.NewReader(rawData)
	_, err = sample.obsClient.SetBucketCorsWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders, data)
	if err == nil {
		fmt.Printf("Set bucket cors:%s successfully!\n", sample.bucketName)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	input.Method = obs.HttpMethodGet
	output, err = sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "GetBucketCors")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	getBucketCorsOutput, err := sample.obsClient.GetBucketCorsWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
	for index, corsRule := range getBucketCorsOutput.CorsRules {
		fmt.Printf("CorsRule[%d]\n", index)
		fmt.Printf("ID:%s, AllowedOrigin:%s, AllowedMethod:%s, AllowedHeader:%s, MaxAgeSeconds:%d, ExposeHeader:%s\n",
			corsRule.ID, strings.Join(corsRule.AllowedOrigin, "|"), strings.Join(corsRule.AllowedMethod, "|"),
			strings.Join(corsRule.AllowedHeader, "|"), corsRule.MaxAgeSeconds, strings.Join(corsRule.ExposeHeader, "|"))
	}
	fmt.Println()

	input.Method = obs.HttpMethodDelete
	output, err = sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "DeleteBucketCors")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	_, err = sample.obsClient.DeleteBucketCorsWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err == nil {
		fmt.Println("Delete bucket cors successfully!")
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func (sample TemporarySignatureSample) PutObject() {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodPut
	input.Bucket = sample.bucketName
	input.Key = sample.objectKey
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "PutObject")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	data := strings.NewReader("Hello OBS")
	_, err = sample.obsClient.PutObjectWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders, data)
	if err == nil {
		fmt.Printf("Put object:%s successfully!\n", sample.objectKey)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func (TemporarySignatureSample) createSampleFile(sampleFilePath string) {
	if err := os.MkdirAll(filepath.Dir(sampleFilePath), os.ModePerm); err != nil {
		fmt.Println(err)
	}

	if err := ioutil.WriteFile(sampleFilePath, []byte("Hello OBS from file"), os.ModePerm); err != nil {
		fmt.Println(err)
	}
}

func (sample TemporarySignatureSample) PutFile(sampleFilePath string) {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodPut
	input.Bucket = sample.bucketName
	input.Key = sample.objectKey
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "PutFile")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	_, err = sample.obsClient.PutFileWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders, sampleFilePath)
	if err == nil {
		fmt.Printf("Put file:%s successfully!\n", sample.objectKey)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func (sample TemporarySignatureSample) GetObject() {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodGet
	input.Bucket = sample.bucketName
	input.Key = sample.objectKey
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "GetObject")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	getObjectOutput, err := sample.obsClient.GetObjectWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		errMsg := getObjectOutput.Body.Close()
		if errMsg != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Object content:")
	body, err := ioutil.ReadAll(getObjectOutput.Body)
	if err == nil {
		fmt.Println(string(body))
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func (sample TemporarySignatureSample) DoObjectAcl() {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodPut
	input.Bucket = sample.bucketName
	input.Key = sample.objectKey
	input.SubResource = obs.SubResourceAcl
	input.Expires = 3600
	input.Headers = map[string]string{obs.HEADER_ACL_AMZ: string(obs.AclPublicRead)}
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "SetObjectAcl")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	_, err = sample.obsClient.SetObjectAclWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders, nil)
	if err == nil {
		fmt.Printf("Set object acl:%s successfully!\n", sample.objectKey)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	input.Method = obs.HttpMethodGet
	output, err = sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "GetObjectAcl")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	getObjectAclOutput, err := sample.obsClient.GetObjectAclWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err == nil {
		fmt.Printf("Object owner - ownerId:%s, ownerName:%s\n", getObjectAclOutput.Owner.ID, getObjectAclOutput.Owner.DisplayName)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
	for index, grant := range getObjectAclOutput.Grants {
		fmt.Printf("Grant[%d]\n", index)
		fmt.Printf("GranteeUri:%s, GranteeId:%s, GranteeName:%s\n", grant.Grantee.URI, grant.Grantee.ID, grant.Grantee.DisplayName)
		fmt.Printf("Permission:%s\n", grant.Permission)
	}
	fmt.Println()
}

func (sample TemporarySignatureSample) DeleteObject() {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodDelete
	input.Bucket = sample.bucketName
	input.Key = sample.objectKey
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "DeleteObject")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	_, err = sample.obsClient.DeleteObjectWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err == nil {
		fmt.Printf("Delete object:%s successfully!\n", sample.objectKey)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func (sample TemporarySignatureSample) DeleteBucket() {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodDelete
	input.Bucket = sample.bucketName
	input.Expires = 3600
	output, err := sample.obsClient.CreateSignedUrl(input)
	if err == nil {
		fmt.Printf("%s using temporary signature url:\n", "DeleteBucket")
		fmt.Println(output.SignedUrl)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	_, err = sample.obsClient.DeleteBucketWithSignedUrl(output.SignedUrl, output.ActualSignedRequestHeaders)
	if err == nil {
		fmt.Printf("Delete bucket:%s successfully!\n", sample.bucketName)
		fmt.Println()
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println("Code:" + obsError.Code)
			fmt.Println("Message:" + obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

}

func RunTemporarySignatureSample() {
	const (
		endpoint   = "https://your-endpoint"
		ak         = "*** Provide your Access Key ***"
		sk         = "*** Provide your Secret Key ***"
		bucketName = "bucket-test"
		objectKey  = "object-test"
		location   = "yourbucketlocation"
	)

	sample := newTemporarySignatureSample(ak, sk, endpoint, bucketName, objectKey, location)

	// Create bucket
	sample.CreateBucket()

	// List buckets
	sample.ListBuckets()

	// Set/Get/Delete bucket cors
	sample.DoBucketCors()

	// Put object
	sample.PutObject()

	// Get object
	sample.GetObject()

	// Put file
	sampleFilePath := "/temp/sampleText.txt"
	sample.createSampleFile(sampleFilePath)

	sample.PutFile(sampleFilePath)
	// Get object
	sample.GetObject()

	// Set/Get object acl
	sample.DoObjectAcl()

	// Delete object
	sample.DeleteObject()

	// Delete bucket
	sample.DeleteBucket()
}
