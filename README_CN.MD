Version 3.21.8

新特性：
1. 新增桶加密相关接口：设置桶加密配置(SetBucketEncryption), 获取桶加密配置(GetBucketEncryption), 删除桶加密配置(DeleteBucketEncryption);
2. 新增追加上传接口(AppendObject);
3. 新增修改写接口(ModifyObject);
4. 列举桶接口(ListBuckets)支持指定桶类型列举;

资料&demo:
1. 补充桶加密相关接口描述；
2. 补充追加上传接口描述；
3. 补充修改写接口描述；
4. 补充列举桶新增参数说明；

修复问题：

-----------------------------------------------------------------------------------

Version 3.21.1

新特性：

资料&demo:

修复问题：
1. 修复代码在指定重试次数内重试成功时仍会返回失败错误信息的问题；
2. 修复代码在解析响应失败时未返回错误信息的问题；

-----------------------------------------------------------------------------------

Version 3.20.7.1

新特性：
1. 新增断点续传上传接口(ObsClient.UploadFile)及断点续传下载接口(ObsClient.DownloadFile)；

资料&demo:
1. 开发指南上传对象章节新增断点续传上传接口的相关描述，下载对象章节新增断点续传下载接口的相关描述；
2. 接口参考其他接口章节新增断点续传上传及断点续传下载的相关描述；

修复问题：
1. 优化部分代码实现；

-----------------------------------------------------------------------------------

Version 3.20.7

新特性：

资料&demo:

修复问题：
1. 优化部分代码实现；

-----------------------------------------------------------------------------------
Version 3.20.3

新特性：
1. 新增判断对象是否存在的接口(ObsClient.HeadObject)；
2. 新增设置桶请求者付费配置和获取桶请求者付费配置的接口(ObsClient.SetBucketRequestPayment/ObsClient.GetBucketRequestPayment)；
3. 新增对请求者付费头域的支持(obs.WithReqPaymentHeader)，可以在调用接口时指定配置携带请求者付费头域；

资料&demo:
1. 开发指南文档管理对象章节新增判断对象是否存在，接口参考文档对象相关接口章节新增判断对象是否存在，数据类型章节新增判断对象是否存在请求参数；
2. 开发指南文档管理桶章节新增管理桶请求者付费，接口参考文档桶相关接口章节新增设置桶请求者付费配置和获取桶请求者付费配置，数据类型章节新增设置桶请求者付费配置请求参数及获取桶请求者付费配置响应结果；
3. 接口参考文档初始化章节新增SDK接口拓展配置；
4. 接口参考文档枚举常量章节新增请求者付费配置；
5. 接口参考文档各接口方法定义中新增对拓展参数的描述；

修复问题：
1. 优化部分代码实现；
2. 修复CreateBucket接口在未指定StorageClass参数情况下指定Epid参数不生效的问题；
3. CreateBucket接口支持指定az-redundancy参数；

-----------------------------------------------------------------------------------

Version 3.20.1

新特性：

资料&demo:

修复问题：
1. 优化部分代码实现；

-----------------------------------------------------------------------------------

Version 3.19.11
更新发布版本号，新的版本号命名方式：主版本号.年标识.月标识。

新特性：

资料&demo:
1. 接口参考ObsClient初始化章节，新增WithMaxRedirectCount及WithSecurityToken的相关描述；
2. 开发指南配置OBS客户端章节，新增WithMaxRedirectCount及WithSecurityToken的相关描述；
3. 开发指南常见问题章节，新增HTTP状态码为405时的相关处理；

修复问题：
1.  修复携带DisPlayName参数可能造成返回400错误的问题；
2.  修复SetBucketNotification接口返回400错误的问题；
3.  修复SetObjectAcl接口不支持指定Delivered参数的问题；
4.  修复SetBucketLoggingConfiguration接口在不同协议下对Agency字段的兼容问题；
5.  修复SetBucketLoggingConfiguration接口对Delivered字段的处理问题；
6.  修复UploadPart接口设置ContentMD5不生效的问题；
7.  修复SetBucketStoragePolicy接口在不同协议下对错误存储类型处理行为不一致的问题；
8.  修复CreateBucket接口未区分协议造成返回400错误的问题；
9.  修复getBucketAclObs接口的语法问题；
10. 修复SetBucketWebsiteConfiguration接口的错误；
11. 修复临时鉴权接口在不同协议下的兼容问题；
12. 修复在GET请求返回302情况下进行重定向会添加鉴权信息头的问题；
13. 修复根据文件名后缀获取content-type类型时不支持大写文件名后缀的问题；
14. 修改svg格式后缀的文件对应获取的content-type类型为image/svg+xml；
15. 修复examples/bucket_operations_sample.go中删除桶示例调用错误接口的问题。
16. 修复通过临时访问密钥访问存在的计算鉴权问题；
17. 修复匿名访问方式下，部分响应字段为空的问题。

-----------------------------------------------------------------------------------

Version 3.1.2

新特性：

资料&demo:

修复问题：
1. 修复Obs.InitiateMultipartUpload设置ContentType不生效的问题；

-----------------------------------------------------------------------------------

Version 3.1.1

新特性：
1. 支持配置自定义域名访问OBS服务（obs.WithCustomDomainName）；
2. 上传对象ObsClient.PutObject支持自动识别更广泛的MIME类型；

修复问题：
1. 修复ObsClient.GetBucketStoragePolicy报错空指针的问题；
2. 修复ObsClient.SetBucketStoragePolicy报错的400的问题；
