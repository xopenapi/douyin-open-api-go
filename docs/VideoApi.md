# \VideoApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DouyinVideoCreate**](VideoApi.md#DouyinVideoCreate) | **Post** /video/create | 创建抖音视频
[**DouyinVideoData**](VideoApi.md#DouyinVideoData) | **Post** /video/data | 查询指定视频数据
[**DouyinVideoDelete**](VideoApi.md#DouyinVideoDelete) | **Post** /video/delete | 删除授权用户发布的视频
[**DouyinVideoList**](VideoApi.md#DouyinVideoList) | **Get** /video/list | 查询授权账号视频数据
[**DouyinVideoPartComplete**](VideoApi.md#DouyinVideoPartComplete) | **Post** /video/part/complete | 视频分片完成上传
[**DouyinVideoPartInit**](VideoApi.md#DouyinVideoPartInit) | **Post** /video/part/init | 抖音分片初始化上传
[**DouyinVideoPartUpload**](VideoApi.md#DouyinVideoPartUpload) | **Post** /video/part/upload | 上传视频分片到文件服务器
[**DouyinVideoPositionSearch**](VideoApi.md#DouyinVideoPositionSearch) | **Get** /poi/search/keyword | 查询POI信息
[**DouyinVideoShareId**](VideoApi.md#DouyinVideoShareId) | **Get** /share-id/ | 获取share-id
[**DouyinVideoUpload**](VideoApi.md#DouyinVideoUpload) | **Post** /video/upload | 抖音上传视频到文件服务器
[**ToutiaoVideoCreate**](VideoApi.md#ToutiaoVideoCreate) | **Post** /toutiao/video/create | 创建头条视频
[**ToutiaoVideoData**](VideoApi.md#ToutiaoVideoData) | **Post** /toutiao/video/data | 查询头条指定视频数据
[**ToutiaoVideoList**](VideoApi.md#ToutiaoVideoList) | **Get** /toutiao/video/list | 查询授权账号视频数据
[**ToutiaoVideoPartComplete**](VideoApi.md#ToutiaoVideoPartComplete) | **Post** /toutiao/video/part/complete/ | 头条视频分片完成上传
[**ToutiaoVideoPartInit**](VideoApi.md#ToutiaoVideoPartInit) | **Post** /toutiao/video/part/init | 头条分片初始化上传
[**ToutiaoVideoPartUpload**](VideoApi.md#ToutiaoVideoPartUpload) | **Post** /toutiao/video/part/upload | 上传视频分片到文件服务器
[**ToutiaoVideoUpload**](VideoApi.md#ToutiaoVideoUpload) | **Post** /toutiao/video/upload | 头条上传视频到文件服务器
[**XiguaVideoCreate**](VideoApi.md#XiguaVideoCreate) | **Post** /xigua/video/create | 创建西瓜视频
[**XiguaVideoData**](VideoApi.md#XiguaVideoData) | **Post** /xigua/video/data | 西瓜查询指定视频数据
[**XiguaVideoList**](VideoApi.md#XiguaVideoList) | **Get** /xigua/video/list | 西瓜查询授权账号视频数据
[**XiguaVideoPartComplete**](VideoApi.md#XiguaVideoPartComplete) | **Post** /xigua/video/part/complete | 视频分片完成上传
[**XiguaVideoPartInit**](VideoApi.md#XiguaVideoPartInit) | **Post** /xigua/video/part/init | 西瓜分片初始化上传
[**XiguaVideoPartUpload**](VideoApi.md#XiguaVideoPartUpload) | **Post** /xigua/video/part/upload | 上传视频分片到文件服务器
[**XiguaVideoUpload**](VideoApi.md#XiguaVideoUpload) | **Post** /xigua/video/upload | 西瓜上传视频到文件服务器



## DouyinVideoCreate

> DouyinVideoCreateRsp DouyinVideoCreate(ctx, openId, accessToken, body)

创建抖音视频

创建抖音视频

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**body** | [**DouyinVideoCreateReq**](DouyinVideoCreateReq.md)|  | 

### Return type

[**DouyinVideoCreateRsp**](DouyinVideoCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoData

> DouyinVideoData DouyinVideoData(ctx, openId, accessToken, body)

查询指定视频数据

查询指定视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**body** | [**InlineObject2**](InlineObject2.md)|  | 

### Return type

[**DouyinVideoData**](DouyinVideoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoDelete

> DouyinVideoDeleteRsp DouyinVideoDelete(ctx, openId, accessToken, body)

删除授权用户发布的视频

删除授权用户发布的视频

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**body** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**DouyinVideoDeleteRsp**](DouyinVideoDeleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoList

> DouyinVideoListRsp DouyinVideoList(ctx, openId, accessToken, count, optional)

查询授权账号视频数据

查询授权账号视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***DouyinVideoListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DouyinVideoListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**DouyinVideoListRsp**](DouyinVideoListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoPartComplete

> DouyinVideoPartCompleteRsp DouyinVideoPartComplete(ctx, openId, accessToken, uploadId)

视频分片完成上传

视频分片完成上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uploadId** | **string**| 分片上传的标记。有限时间为2小时 | 

### Return type

[**DouyinVideoPartCompleteRsp**](DouyinVideoPartCompleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoPartInit

> DouyinVideoPartInitRsp DouyinVideoPartInit(ctx, openId, accessToken)

抖音分片初始化上传

抖音分片初始化上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 

### Return type

[**DouyinVideoPartInitRsp**](DouyinVideoPartInitRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoPartUpload

> DouyinVideoPartUploadRsp DouyinVideoPartUpload(ctx, openId, accessToken, uploadId, partNumber, video)

上传视频分片到文件服务器

上传视频分片到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uploadId** | **string**| 分片上传的标记。有限时间为2小时 | 
**partNumber** | **int64**| 第几个分片，从1开始 | 
**video** | [**interface{}**](interface{}.md)|  | 

### Return type

[**DouyinVideoPartUploadRsp**](DouyinVideoPartUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoPositionSearch

> DouyinVideoPositionSearchRsp DouyinVideoPositionSearch(ctx, accessToken, count, keyword, optional)

查询POI信息

查询POI信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权。 | 
**count** | **int64**| 每页数量。 | 
**keyword** | **string**| 查询关键字，例如美食 | 
 **optional** | ***DouyinVideoPositionSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DouyinVideoPositionSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 
 **city** | **optional.String**| 查询城市，例如上海、北京。 | 

### Return type

[**DouyinVideoPositionSearchRsp**](DouyinVideoPositionSearchRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoShareId

> DouyinVideoShareIdRsp DouyinVideoShareId(ctx, accessToken, optional)

获取share-id

获取share-id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/client_token/生成的token，此token不需要用户授权。 | 
 **optional** | ***DouyinVideoShareIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DouyinVideoShareIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **needCallback** | **optional.Bool**| 如果需要知道视频分享成功的结果，need_callback设置为true | 
 **sourceStyleId** | **optional.String**| 多来源样式id（暂未开放）。 | 
 **defaultHashtag** | **optional.String**| 追踪分享默认hashtag | 
 **linkParam** | **optional.String**| 分享来源url附加参数（暂未开放）。 | 

### Return type

[**DouyinVideoShareIdRsp**](DouyinVideoShareIdRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DouyinVideoUpload

> DouyinVideoUploadRsp DouyinVideoUpload(ctx, openId, accessToken, video)

抖音上传视频到文件服务器

抖音上传视频到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**video** | ***os.File*****os.File**|  | 

### Return type

[**DouyinVideoUploadRsp**](DouyinVideoUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoCreate

> TouTiaoVideoCreateRsp ToutiaoVideoCreate(ctx, openId, accessToken, optional)

创建头条视频

创建头条视频

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***ToutiaoVideoCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ToutiaoVideoCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **text** | **optional.String**| 视频标题不要超过30个字符 | 
 **videoId** | **optional.String**| video_id, 通过/toutiao/video/upload/接口得到 | 
 **extra** | [**optional.Interface of ToutiaoVideoCreateExtra**](_toutiao_video_create_extra.md)|  | 

### Return type

[**TouTiaoVideoCreateRsp**](TouTiaoVideoCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoData

> TouTiaoVideoData ToutiaoVideoData(ctx, openId, accessToken, optional)

查询头条指定视频数据

查询头条指定视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***ToutiaoVideoDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ToutiaoVideoDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemIds** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**TouTiaoVideoData**](TouTiaoVideoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoList

> TouTiaoVideoListRsp ToutiaoVideoList(ctx, openId, accessToken, count, optional)

查询授权账号视频数据

查询授权账号视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***ToutiaoVideoListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ToutiaoVideoListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**TouTiaoVideoListRsp**](TouTiaoVideoListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoPartComplete

> TouTiaoVideoPartCompleteRsp ToutiaoVideoPartComplete(ctx, openId, accessToken, uploadId)

头条视频分片完成上传

头条视频分片完成上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uploadId** | **string**| 分片上传的标记。有限时间为2小时 | 

### Return type

[**TouTiaoVideoPartCompleteRsp**](TouTiaoVideoPartCompleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoPartInit

> TouTiaoVideoPartInitRsp ToutiaoVideoPartInit(ctx, openId, accessToken)

头条分片初始化上传

头条分片初始化上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 

### Return type

[**TouTiaoVideoPartInitRsp**](TouTiaoVideoPartInitRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoPartUpload

> TouTiaoVideoPartUploadRsp ToutiaoVideoPartUpload(ctx, openId, accessToken, uploadId, partNumber, optional)

上传视频分片到文件服务器

上传视频分片到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uploadId** | **string**| 分片上传的标记。有限时间为2小时 | 
**partNumber** | **int64**| 第几个分片，从1开始 | 
 **optional** | ***ToutiaoVideoPartUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ToutiaoVideoPartUploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **video** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**TouTiaoVideoPartUploadRsp**](TouTiaoVideoPartUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToutiaoVideoUpload

> ToutiaoVideoUploadRsp ToutiaoVideoUpload(ctx, openId, accessToken, optional)

头条上传视频到文件服务器

头条上传视频到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***ToutiaoVideoUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ToutiaoVideoUploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **video** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**ToutiaoVideoUploadRsp**](ToutiaoVideoUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoCreate

> XiGuaVideoCreateRsp XiguaVideoCreate(ctx, openId, accessToken, optional)

创建西瓜视频

创建西瓜视频

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***XiguaVideoCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a XiguaVideoCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **abstract** | **optional.String**| 视频简介，400字以内 | 
 **claimOrigin** | **optional.String**| 是否声明原创（授权账号需要在西瓜视频端内开通「实名认证」） | 
 **coverTsp** | **optional.String**| 从视频中截取封面的时间，用该帧作为封面（单位为毫秒） | 
 **praise** | **optional.String**| 是否给视频开通可以赞赏的入口（授权账号需要在西瓜视频端内开通「实名认证」） | 
 **text** | **optional.String**| 标题长度应该在5-30字之间 | 
 **videoId** | **optional.String**| video_id, 通过/xigua/video/upload/接口得到 | 

### Return type

[**XiGuaVideoCreateRsp**](XiGuaVideoCreateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoData

> XiGuaVideoData XiguaVideoData(ctx, openId, accessToken, optional)

西瓜查询指定视频数据

西瓜ig查询指定视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***XiguaVideoDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a XiguaVideoDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemIds** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**XiGuaVideoData**](XiGuaVideoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoList

> XiGuaVideoListRsp XiguaVideoList(ctx, openId, accessToken, count, optional)

西瓜查询授权账号视频数据

西瓜查询授权账号视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***XiguaVideoListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a XiguaVideoListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**XiGuaVideoListRsp**](XiGuaVideoListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoPartComplete

> XiGuaVideoPartCompleteRsp XiguaVideoPartComplete(ctx, openId, accessToken, uploadId)

视频分片完成上传

视频分片完成上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uploadId** | **string**| 分片上传的标记。有限时间为2小时 | 

### Return type

[**XiGuaVideoPartCompleteRsp**](XiGuaVideoPartCompleteRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoPartInit

> XiGuaVideoPartInitRsp XiguaVideoPartInit(ctx, openId, accessToken)

西瓜分片初始化上传

西瓜分片初始化上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 

### Return type

[**XiGuaVideoPartInitRsp**](XiGuaVideoPartInitRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoPartUpload

> XiGuaVideoPartUploadRsp XiguaVideoPartUpload(ctx, openId, accessToken, uploadId, partNumber, optional)

上传视频分片到文件服务器

上传视频分片到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**uploadId** | **string**| 分片上传的标记。有限时间为2小时 | 
**partNumber** | **int64**| 第几个分片，从1开始 | 
 **optional** | ***XiguaVideoPartUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a XiguaVideoPartUploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **video** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**XiGuaVideoPartUploadRsp**](XiGuaVideoPartUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XiguaVideoUpload

> XiGuaVideoUploadRsp XiguaVideoUpload(ctx, openId, accessToken, optional)

西瓜上传视频到文件服务器

西瓜上传视频到文件服务器

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***XiguaVideoUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a XiguaVideoUploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **video** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**XiGuaVideoUploadRsp**](XiGuaVideoUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

