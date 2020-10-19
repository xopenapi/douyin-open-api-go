# \PoiApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoiQuery**](PoiApi.md#PoiQuery) | **Get** /poi/query | 获取抖音POI ID
[**PoiSupplierMatch**](PoiApi.md#PoiSupplierMatch) | **Post** /poi/supplier/match | 店铺匹配
[**PoiSupplierMatchQuery**](PoiApi.md#PoiSupplierMatchQuery) | **Get** /poi/supplier/match/query | 店铺匹配结果查询
[**PoiSupplierQuery**](PoiApi.md#PoiSupplierQuery) | **Get** /poi/supplier/query | 查询店铺
[**PoiSupplierSync**](PoiApi.md#PoiSupplierSync) | **Post** /poi/supplier/sync | 商铺同步



## PoiQuery

> PoiQueryRsp PoiQuery(ctx, accessToken, amapId)

获取抖音POI ID

获取抖音POI ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**amapId** | **string**|  | 

### Return type

[**PoiQueryRsp**](PoiQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSupplierMatch

> PoiSupplierMatchRsp PoiSupplierMatch(ctx, accessToken, optional)

店铺匹配

店铺匹配

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiSupplierMatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiSupplierMatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PoiSupplierMatchReq**](PoiSupplierMatchReq.md)|  | 

### Return type

[**PoiSupplierMatchRsp**](PoiSupplierMatchRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSupplierMatchQuery

> PoiSupplierMatchQueryRsp PoiSupplierMatchQuery(ctx, accessToken, optional)

店铺匹配结果查询

店铺匹配结果查询

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiSupplierMatchQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiSupplierMatchQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InlineObject7**](InlineObject7.md)|  | 

### Return type

[**PoiSupplierMatchQueryRsp**](PoiSupplierMatchQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSupplierQuery

> PoiSupplierQueryRsp PoiSupplierQuery(ctx, accessToken, supplierExtId)

查询店铺

查询店铺

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**supplierExtId** | **string**|  | 

### Return type

[**PoiSupplierQueryRsp**](PoiSupplierQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSupplierSync

> PoiSupplierSyncRsp PoiSupplierSync(ctx, accessToken, optional)

商铺同步

商铺同步

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiSupplierSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiSupplierSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PoiSupplier**](PoiSupplier.md)|  | 

### Return type

[**PoiSupplierSyncRsp**](PoiSupplierSyncRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

