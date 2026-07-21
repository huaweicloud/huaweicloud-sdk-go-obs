版本号：V3.26.6

新特性：

1. 新增桶清单（Bucket Inventory）相关接口：设置桶清单（SetBucketInventory）、获取桶清单（GetBucketInventory）、列举桶清单（ListBucketInventory）、删除桶清单（DeleteBucketInventory）。
2. 新增桶复制（Replication）相关接口：设置桶复制规则（SetBucketReplication）、获取桶复制规则（GetBucketReplication）、删除桶复制规则（DeleteBucketReplication）。
3. 新增直读冷存（DirectColdAccess）相关接口：设置直读冷存（SetBucketDirectColdAccess）、获取直读冷存（GetBucketDirectColdAccess）、删除直读冷存（DeleteBucketDirectColdAccess）。
4. 新增桶 DIS 策略（DisPolicy）相关接口：设置 DIS 策略（SetDisPolicy）、获取 DIS 策略（GetDisPolicy）、删除 DIS 策略（DeleteDisPolicy）。
5. 新增桶解压策略（Decompression）相关接口：设置桶解压策略（SetBucketDecompression）、获取桶解压策略（GetBucketDecompression）、删除桶解压策略（DeleteBucketDecompression）。
6. 新增桶 WORM 配置相关接口：设置桶 WORM 配置（SetBucketWormConfiguration）、获取桶 WORM 配置（GetBucketWormConfiguration）。
7. 新增对象 WORM（对象锁/对象 retention）接口：设置对象 retention（PutObjectRetention）。
8. 新增异步断点续传上传和下载接口（ObsClient.UploadFileAsync 和 ObsClient.DownloadFileAsync），返回 TransferTask，支持 Pause（暂停）、Resume（恢复）、Cancel（取消）、Status（状态查询）、GetResult（获取结果）。
9. 后台断点续传： 断点续传上传（UploadFile）和断点续传下载（DownloadFile）的检查点（checkpoint）文件改为由后台 AsyncWriter 协程异步持久化，取代原先加锁同步写入的方式，降低锁竞争，提升传输吞吐量。

资料 & demo：

1. 补充桶清单（Bucket Inventory）相关接口描述。
2. 补充桶复制（Replication）相关接口描述。
3. 补充直读冷存（DirectColdAccess）相关接口描述。
4. 补充桶 DIS 策略（DisPolicy）相关接口描述。
5. 补充桶解压策略（Decompression）相关接口描述。
6. 补充桶 WORM 配置相关接口描述。
7. 补充对象 retention 相关接口描述。
8. 补充桶高级操作示例描述，包括直读冷存、桶复制、桶 WORM 配置、对象 retention、桶解压相关接口示例。

修复问题：

-----------------------------------------------------------------------------------

版本号：V3.26.3

新特性：

1. 对象标签： 新增对象标签（Object Tagging）相关接口，支持对数据的分类管理。
2. 桶加密配置： 创建桶（Bucket）时支持自主配置加密方式，提升数据安全性。

修复问题：

1. 上传逻辑优化： 修复了在分块（Chunked）模式下上传 0KB 对象时，Content-Length 错误设置为 0 的问题。

-----------------------------------------------------------------------------------

Version 3.25.9

新特性：

资料 & demo：

修复问题：
1. 修复range无法下载首字节的问题
-----------------------------------------------------------------------------------

Version 3.25.4

新特性：
1. 支持在使用UploadFile方法时传递元数据信息。

资料 & demo：

修复问题：

-----------------------------------------------------------------------------------

Version 3.25.3

新特性：
1. 新增BPA相关接口。
2. 支持智能分级特性。

资料 & demo：
1. 补充BPA相关接口描述。

修复问题：

-----------------------------------------------------------------------------------

Version 3.24.9

新特性：
1. 新增listPosixObject接口。

资料 & demo：

修复问题：

1. 优化错误日志打印。

-----------------------------------------------------------------------------------

Version 3.24.6

新特性：
1. 支持在使用CopyObject和InitiateMultipartUpload方法时传递元数据信息。

资料 & demo：

修复问题：

1. 修复部分场景下bucket acl配置失败的问题。
2. 修复重试场景下数据传输异常的问题。

-----------------------------------------------------------------------------------

Version 3.24.3

新特性：
1. 新增accesslabel相关接口：设置对象accesslabel(SetDirAccesslabel), 获取对象accesslabel(GetDirAccesslabel), 删除对象accesslabel(DeleteDirAccesslabel);

资料&demo:
1. 补充accesslabel相关接口描述；

修复问题：

-----------------------------------------------------------------------------------

Version 3.23.12

新特性：
1. 新增obs.WithDisableKeepAlive方法实现使用短连接功能。
2. 新增深度归档存储类别。

资料&demo:

修复问题：
1. 修复部分场景下设置桶的生命周期配置失败的问题。
-----------------------------------------------------------------------------------

Version 3.23.9

新特性：
1. 新增obs.WithProgress方法支持进度条回调。
2. 新增配置生命周期管理规则中的碎片过期时间。

资料&demo:

-----------------------------------------------------------------------------------
Version 3.23.4

新特性：
1. 新增obs.WithCustomHeader方法。
2. 新增自定义域名相关API，包括SetBucketCustomDomain,GetBucketCustomDomain,DeleteBucketCustomDomain。
3. 新增镜像回源相关API，包括SetBucketMirrorBackToSource、GetBucketMirrorBackToSource、DeleteBucketMirrorBackToSource。

资料&demo:
1. 补充自定义域名相关API的描述。
2. 补充镜像回源相关API的描述。

-----------------------------------------------------------------------------------
Version 3.23.3

新特性：
1. 新增obs.WithCallbackHeader方法；

资料&demo:
1. 补充上传时回调的说明。
2. 补充SDK API扩展配置说明。

修复问题：
1. 优化部分代码实现

-----------------------------------------------------------------------------------
Version 3.21.12

新特性：
1. 新增obs.WithTrafficLimitHeader方法实现单链接限速功能;

资料&demo:
1. 补充单链接限速描述；
2. 补充SDK接口扩展配置描述；

修复问题：
1. 优化部分代码实现

-----------------------------------------------------------------------------------

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
