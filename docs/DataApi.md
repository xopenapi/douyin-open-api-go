# \DataApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataExternalItemBase**](DataApi.md#DataExternalItemBase) | **Get** /data/external/item/base | 获取视频基础数据
[**DataExternalItemComment**](DataApi.md#DataExternalItemComment) | **Get** /data/external/item/comment | 获取视频评论数据
[**DataExternalItemLike**](DataApi.md#DataExternalItemLike) | **Get** /data/external/item/like | 获取视频点赞数据
[**DataExternalItemPlay**](DataApi.md#DataExternalItemPlay) | **Get** /data/external/item/play | 获取视频播放数据
[**DataExternalItemShare**](DataApi.md#DataExternalItemShare) | **Get** /data/external/item/share | 获取视频分享数据
[**DataExternalSdkShare**](DataApi.md#DataExternalSdkShare) | **Get** /data/external/sdk_share | 获取SDK分享视频数据
[**DataExternalUserComment**](DataApi.md#DataExternalUserComment) | **Get** /data/external/user/comment | 获取用户评论数
[**DataExternalUserFans**](DataApi.md#DataExternalUserFans) | **Get** /data/external/user/fans | 获取用户粉丝数
[**DataExternalUserItem**](DataApi.md#DataExternalUserItem) | **Get** /data/external/user/item | 获取用户视频情况
[**DataExternalUserLike**](DataApi.md#DataExternalUserLike) | **Get** /data/external/user/like | 获取用户点赞数
[**DataExternalUserProfile**](DataApi.md#DataExternalUserProfile) | **Get** /data/external/user/profile | 获取用户主页访问数
[**DataExternalUserShare**](DataApi.md#DataExternalUserShare) | **Get** /data/external/user/share | 获取用户分享数
[**FansData**](DataApi.md#FansData) | **Get** /fans/data | 获取用户粉丝数据



## DataExternalItemBase

> DataExternalItemBaseRsp DataExternalItemBase(ctx, accessToken, openId, itemId)

获取视频基础数据

获取视频基础数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**itemId** | **string**| item_id，仅能查询access_token对应用户上传的视频 | 

### Return type

[**DataExternalItemBaseRsp**](DataExternalItemBaseRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalItemComment

> DataExternalItemCommentRsp DataExternalItemComment(ctx, accessToken, openId, itemId, dateType)

获取视频评论数据

获取视频评论数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**itemId** | **string**| item_id，仅能查询access_token对应用户上传的视频 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天 | 

### Return type

[**DataExternalItemCommentRsp**](DataExternalItemCommentRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalItemLike

> DataExternalItemLikeRsp DataExternalItemLike(ctx, accessToken, openId, itemId, dateType)

获取视频点赞数据

获取视频点赞数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**itemId** | **string**| item_id，仅能查询access_token对应用户上传的视频 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天 | 

### Return type

[**DataExternalItemLikeRsp**](DataExternalItemLikeRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalItemPlay

> DataExternalItemPlayRsp DataExternalItemPlay(ctx, accessToken, openId, itemId, dateType)

获取视频播放数据

获取视频播放数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**itemId** | **string**| item_id，仅能查询access_token对应用户上传的视频 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天 | 

### Return type

[**DataExternalItemPlayRsp**](DataExternalItemPlayRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalItemShare

> DataExternalItemShareRsp DataExternalItemShare(ctx, accessToken, openId, itemId, dateType)

获取视频分享数据

获取视频分享数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**itemId** | **string**| item_id，仅能查询access_token对应用户上传的视频 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天 | 

### Return type

[**DataExternalItemShareRsp**](DataExternalItemShareRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalSdkShare

> DataExternalSdkShareRsp DataExternalSdkShare(ctx, accessToken, endDate, optional)

获取SDK分享视频数据

获取SDK分享视频数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**endDate** | **string**| 最近30天，结束日期(yyyy-MM-dd) | 
 **optional** | ***DataExternalSdkShareOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DataExternalSdkShareOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **beginDate** | **optional.String**| 最近30天，开始日期(yyyy-MM-dd)。 | 

### Return type

[**DataExternalSdkShareRsp**](DataExternalSdkShareRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalUserComment

> DataExternalUserCommentRsp DataExternalUserComment(ctx, accessToken, openId, dateType)

获取用户评论数

获取用户评论数

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天、30代表30天 | 

### Return type

[**DataExternalUserCommentRsp**](DataExternalUserCommentRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalUserFans

> DataExternalUserFansRsp DataExternalUserFans(ctx, accessToken, openId, dateType)

获取用户粉丝数

获取用户粉丝数

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天、30代表30天 | 

### Return type

[**DataExternalUserFansRsp**](DataExternalUserFansRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalUserItem

> DataExternalUserItemRsp DataExternalUserItem(ctx, accessToken, openId, dateType)

获取用户视频情况

获取用户视频情况

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天、30代表30天 | 

### Return type

[**DataExternalUserItemRsp**](DataExternalUserItemRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalUserLike

> DataExternalUserLikeRsp DataExternalUserLike(ctx, accessToken, openId, dateType)

获取用户点赞数

获取用户点赞数

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天、30代表30天 | 

### Return type

[**DataExternalUserLikeRsp**](DataExternalUserLikeRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalUserProfile

> DataExternalUserProfileRsp DataExternalUserProfile(ctx, accessToken, openId, dateType)

获取用户主页访问数

获取用户主页访问数

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天、30代表30天 | 

### Return type

[**DataExternalUserProfileRsp**](DataExternalUserProfileRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataExternalUserShare

> DataExternalUserShareRsp DataExternalUserShare(ctx, accessToken, openId, dateType)

获取用户分享数

获取用户分享数

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**dateType** | **int64**| 近7/15天；输入7代表7天、15代表15天、30代表30天 | 

### Return type

[**DataExternalUserShareRsp**](DataExternalUserShareRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FansData

> FansDataRsp FansData(ctx, accessToken, openId)

获取用户粉丝数据

获取用户粉丝数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 

### Return type

[**FansDataRsp**](FansDataRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

