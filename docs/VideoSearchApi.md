# \VideoSearchApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideoSearch**](VideoSearchApi.md#VideoSearch) | **Get** /video/search | 关键词视频搜索
[**VideoSearchCommentList**](VideoSearchApi.md#VideoSearchCommentList) | **Get** /video/search/comment/list | 关键词视频评论列表
[**VideoSearchCommentReply**](VideoSearchApi.md#VideoSearchCommentReply) | **Post** /video/search/comment/reply | 关键词视频评论回复
[**VideoSearchCommentReplyList**](VideoSearchApi.md#VideoSearchCommentReplyList) | **Get** /video/search/comment/reply/list | 评论回复列表



## VideoSearch

> VideoSearchRsp VideoSearch(ctx, openId, accessToken, count, keyword, optional)

关键词视频搜索

关键词视频搜索

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量。 | 
**keyword** | **string**| 关键词 | 
 **optional** | ***VideoSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**VideoSearchRsp**](VideoSearchRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoSearchCommentList

> VideoSearchCommentListRsp VideoSearchCommentList(ctx, secItemId, accessToken, count, optional)

关键词视频评论列表

关键词视频评论列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secItemId** | **string**| 视频搜索接口返回的加密的视频id | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量。 | 
 **optional** | ***VideoSearchCommentListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoSearchCommentListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**VideoSearchCommentListRsp**](VideoSearchCommentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoSearchCommentReply

> VideoSearchCommentReplyRsp VideoSearchCommentReply(ctx, accessToken, openId, body)

关键词视频评论回复

关键词视频评论回复

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志。 | 
**body** | [**VideoSearchCommentReplyReq**](VideoSearchCommentReplyReq.md)|  | 

### Return type

[**VideoSearchCommentReplyRsp**](VideoSearchCommentReplyRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoSearchCommentReplyList

> VideoSearchCommentReplyListRsp VideoSearchCommentReplyList(ctx, accessToken, count, secItemId, commentId, optional)

评论回复列表

评论回复列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
**secItemId** | **string**| 视频搜索接口返回的加密的视频id | 
**commentId** | **string**| 评论id | 
 **optional** | ***VideoSearchCommentReplyListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoSearchCommentReplyListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 

### Return type

[**VideoSearchCommentReplyListRsp**](VideoSearchCommentReplyListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

