Version 3.25.4

New Features:
1. Supports passing metadata when use UploadFile method.

Documentation & Demo:

Resolved Issues:

-----------------------------------------------------------------------------------

Version 3.25.3

New Features:
1. Added BPA feature.
2. Added intelligent tiering feature.

Documentation & Demo:
1. Added descriptions about BPA APIs.

Resolved Issues:

-----------------------------------------------------------------------------------

Version 3.24.9

New Features:
1. Added list POSIX Object API.

Documentation & Demo:

Resolved Issues:
1. Optimize error log printing.

-----------------------------------------------------------------------------------

Version 3.24.6

New Features:
1. Supports passing metadata when use CopyObject and InitiateMultipartUpload method.

Documentation & Demo:

Resolved Issues:
1. Fixed the issue where bucket acl configuration failed in some scenarios.
2. Fixed the issue where data transmission exception in retry scenarios.

-----------------------------------------------------------------------------------

Version 3.24.3

New Features:
1. Added directory accesslable-related APIs, including SetDirAccesslabel, GetDirAccesslabel, and DeleteDirAccesslabel.

Documentation & Demo:
1. Added descriptions about directory accesslable-related APIs.

Resolved Issues:

-----------------------------------------------------------------------------------

Version 3.23.12

New Features:
1. Added obs.WithDisableKeepAlive method to implement short connection function.
2. Added deep archive storage-class.

Documentation & Demo:

Resolved Issues:
1. Fixed the issue where bucket life cycle configuration failed in some scenarios.

-----------------------------------------------------------------------------------

Version 3.23.9

New Features:
1. Added the obs.WithProgress method to support progress bar callback.
2. Added fragment expiration time in configuration life cycle management rules.

Documentation & Demo:

-----------------------------------------------------------------------------------

Version 3.23.4

New Features:
1. Added the obs.WithCustomHeader method when upload.
2. Added bucket customDomain-related APIs, including SetBucketCustomDomain, GetBucketCustomDomain, and DeleteBucketCustomDomain.
3. Added bucket mirrorBackToSource-related APIs, including SetBucketMirrorBackToSource, GetBucketMirrorBackToSource, and DeleteBucketMirrorBackToSource.

Documentation & Demo:
1. Added descriptions about bucket customDomain APIs.
2. Added descriptions about bucket mirrorBackToSource APIs.

-----------------------------------------------------------------------------------
Version 3.23.3

New Features:
1. Added the obs.WithCallbackHeader method when upload.

Documentation & Demo:
1. Added descriptions about callback when upload.
2. Added descriptions about extended configurations of SDK APIs.

Resolved Issues:
1. Optimized some code.

-----------------------------------------------------------------------------------
Version 3.21.12

New Features:
1. Added obs.WithTrafficLimitHeader method to realize single-connection bandwidth throttling.
   
Documentation & Demo:
1. Added descriptions about single-connection bandwidth throttling.
2. Added descriptions about the SDK interface extension configuration.

Resolved Issues:
1. Optimized part of the code.

-----------------------------------------------------------------------------------

Version 3.21.8

New Features:
1. Added bucket encryption-related APIs, including SetBucketEncryption, GetBucketEncryption, and DeleteBucketEncryption.
2. Added the AppendObject API.
3. Added the ModifyObject API.
4. Allowed you to specify a type in the ListBuckets API to list the buckets of this type.

Documentation & Demo:
1. Added descriptions about bucket encryption APIs.
2. Added descriptions about the AppendObject API.
3. Added descriptions about the ModifyObject API.
4. Added descriptions about the new parameter for specifying a bucket type in the ListBuckets API.

Resolved Issues:

-----------------------------------------------------------------------------------

Version 3.21.1

New Features:

Documentation & Demo:

Resolved Issues:
1. Fixed the issue that a failure message is still returned even if code is successfully retried within allowed attempts.
2. Fixed the issue that no error messages are returned when a response message fails to be obtained.

-----------------------------------------------------------------------------------

Version 3.20.7.1

New Features:
1. Added the APIs for resumable upload (ObsClient.UploadFile) and resumable download (ObsClient.DownloadFile).

Documentation & Demo:
1. Added the description of the API for resumable upload to section "Object Upload" and the description of the API for resumable download to section "Object Download" in the OBS Go SDK Developer Guide.
2. Added the topics of resumable upload and resumable download to section "Other APIs" in OBS Go SDK API Reference.

Resolved Issues:
1. Optimized part of the code.

-----------------------------------------------------------------------------------

Version 3.20.7

New Features:

Documentation & Demo:

Resolved Issues:
1. Optimized part of the code.

-----------------------------------------------------------------------------------

Version 3.20.3

New Features:
1. Added the ObsClient.HeadObject API for determining whether an object exists.
2. Added the ObsClient.SetBucketRequestPayment and ObsClient.GetBucketRequestPayment APIs respectively for configuring the Requester Pays function and obtaining related configuration.
3. Supports configuration of the Requester Pays header (obs.WithReqPaymentHeader) when calling an API.

Documentation & Demo:
1. Added the topic of checking whether an object exists to section "Object Management" in OBS Go SDK Developer Guide; added the API for checking whether an object exists to section "Bucket-Related APIs" in OBS Go SDK API Reference; added the request parameter for checking whether an object exists to section "Data Types" in OBS Go SDK API Reference.
2. Added the topic of Requester Pays to section "Bucket Management" in OBS Go SDK Developer Guide; added the APIs for configuring the Requester Pays function and obtaining related configuration to section "Bucket-Related APIs" in OBS Go SDK API Reference; added the request parameter for configuring Requester Pays and the response result of obtaining Requester Pays configuration to section "Data Types" in OBS Go SDK API Reference.
3. Added the topic of SDK API extended configurations to section "Initialization" in OBS Go SDK API Reference.
4. Added the topic of Requester Pays configuration to section "Enumeration Constants" in OBS Go SDK API Reference.
5. Added the description of extended parameters to the API method definitions in OBS Go SDK API Reference.

Resolved Issues:
1. Optimized part of the code.
2. Fixed the issue that the Epid parameter does not take effect if StorageClass is not specified when calling the CreateBucket API.
3. Supports the az-redundancy parameter when calling the CreateBucket API.

-----------------------------------------------------------------------------------

Version 3.20.1

New Features:

Documentation & Demo:

Resolved Issues:
1. Optimized part of the code.

-----------------------------------------------------------------------------------

Version 3.19.11
Updated the version ID format. The new version ID is named in the following format: Main version ID.Year ID.Month ID.

New Features:

Documentation & Demo:
1. Added the description of WithMaxRedirectCount and WithSecurityToken in section "ObsClient Initialization" in the API Reference.
2. Added descriptions about WithMaxRedirectCount and WithSecurityToken in section "Configuring an Instance of OBSClient" in the Developer Guide.
3. Added methods for handling HTTP status code 405 in section "FAQs" in the Developer Guide.

Resolved Issues:
1. Fixed the issue that error code 400 is returned when DisplayName is carried in some circumstances.
2. Fixed the issue that error code 400 is returned when the SetBucketNotification API is called.
3. Fixed the issue that the SetObjectAcl API does not support the specified Delivered parameter.
4. Fixed the issue that the SetBucketLoggingConfiguration API is incompatible with the Agency field in different protocols.
5. Fixed the issue that the SetBucketLoggingConfiguration API incorrectly processes the Delivered field.
6. Fixed the issue that the ContentMD5 configuration does not take effect when the UploadPart API is called.
7. Fixed the issue that the SetBucketStoragePolicy API processes incorrect storage classes inconsistently in different protocols.
8. Fixed the issue that error code 400 is returned because the CreateBucket API does not distinguish protocols.
9. Fixed the syntax issue of the getBucketAclObs API.
10. Rectified the error of the SetBucketWebsiteConfiguration API.
11. Fixed the compatibility issue of the temporary authentication API in different protocols.
12. Fixed the issue that the authentication information header is added when redirection is performed upon a 302 response returned for a GET request.
13. Fixed the issue that the content-type cannot be obtained based on the file name extension if the extension is in uppercase.
14. Changed the content-type of files with the .svg extension to image/svg+xml.
15. Fixed the issue that an incorrect API is called in the sample code for deleting a bucket in examples/bucket_operations_sample.go.
16. Fixed the issue in calculating the authentication value during the access using temporary access keys.
17. Fixed the issue that some response fields are empty in anonymous access.

-----------------------------------------------------------------------------------

Version 3.1.2

New Features:

Documentation & Demo:

Resolved Issues:
1. Fixed the issue that the configuration of ContentType does not take effect when Obs.InitiateMultipartUpload is called.

-----------------------------------------------------------------------------------

Version 3.1.1

New Features:
1. Supports access to OBS using a user-defined domain name (obs.WithCustomDomainName).
2. The API for object upload (ObsClient.PutObject) can automatically identify a wider MIME type.

Resolved Issues:
1. Fixed the issue that a null pointer error is reported when ObsClient.GetBucketStoragePolicy is called.
2. Fixed the 400 error reported by ObsClient.SetBucketStoragePolicy.
