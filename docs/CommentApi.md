# \CommentApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ItemCommentList**](CommentApi.md#ItemCommentList) | **Get** /item/comment/list | 评论列表
[**ItemCommentReply**](CommentApi.md#ItemCommentReply) | **Post** /item/comment/reply | 评论回复列表
[**ItemCommentReplyList**](CommentApi.md#ItemCommentReplyList) | **Get** /item/comment/reply/list | 评论回复列表
[**VideoCommentList**](CommentApi.md#VideoCommentList) | **Get** /video/comment/list | 评论列表
[**VideoCommentReply**](CommentApi.md#VideoCommentReply) | **Post** /video/comment/reply | (企业号)回复视频评论
[**VideoCommentReplyList**](CommentApi.md#VideoCommentReplyList) | **Get** /video/comment/reply/list | 评论列表
[**VideoCommentTop**](CommentApi.md#VideoCommentTop) | **Post** /video/comment/top | (企业号)置顶视频评论



## ItemCommentList

> ItemCommentListRsp ItemCommentList(ctx, openId, accessToken, count, itemId, optional)

评论列表

评论列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
**itemId** | **string**| 视频id | 
 **optional** | ***ItemCommentListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ItemCommentListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据 | 

### Return type

[**ItemCommentListRsp**](ItemCommentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemCommentReply

> ItemCommentReplyRsp ItemCommentReply(ctx, openId, accessToken, optional)

评论回复列表

评论回复列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***ItemCommentReplyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ItemCommentReplyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commentId** | **optional.String**|  | 
 **content** | **optional.String**|  | 
 **itemId** | **optional.String**|  | 

### Return type

[**ItemCommentReplyRsp**](ItemCommentReplyRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemCommentReplyList

> ItemCommentReplyListRsp ItemCommentReplyList(ctx, openId, accessToken, count, itemId, commentId, optional)

评论回复列表

评论回复列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
**itemId** | **string**| 视频id | 
**commentId** | **string**| 评论id | 
 **optional** | ***ItemCommentReplyListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ItemCommentReplyListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据 | 

### Return type

[**ItemCommentReplyListRsp**](ItemCommentReplyListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoCommentList

> VideoCommentListRsp VideoCommentList(ctx, openId, accessToken, count, itemId, optional)

评论列表

评论列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
**itemId** | **string**| 视频id | 
 **optional** | ***VideoCommentListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoCommentListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据 | 

### Return type

[**VideoCommentListRsp**](VideoCommentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoCommentReply

> VideoCommentReplyRsp VideoCommentReply(ctx, openId, accessToken, optional)

(企业号)回复视频评论

(企业号)回复视频评论

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***VideoCommentReplyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoCommentReplyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commentId** | **optional.String**| 需要回复的评论id | 
 **content** | **optional.String**| 评论内容 | 
 **itemId** | **optional.String**| 视频id | 

### Return type

[**VideoCommentReplyRsp**](VideoCommentReplyRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoCommentReplyList

> VideoCommentReplyListRsp VideoCommentReplyList(ctx, openId, accessToken, count, itemId, commentId, optional)

评论列表

评论列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**count** | **int64**| 每页数量 | 
**itemId** | **string**| 视频id | 
**commentId** | **string**| 评论id | 
 **optional** | ***VideoCommentReplyListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoCommentReplyListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据 | 

### Return type

[**VideoCommentReplyListRsp**](VideoCommentReplyListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoCommentTop

> VideoCommentTopRsp VideoCommentTop(ctx, openId, accessToken, optional)

(企业号)置顶视频评论

(企业号)置顶视频评论

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***VideoCommentTopOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VideoCommentTopOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Bool**| true置顶, false取消置顶 | 
 **commentId** | **optional.String**| 需要回复的评论id | 
 **itemId** | **optional.String**| 视频id | 

### Return type

[**VideoCommentTopRsp**](VideoCommentTopRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

