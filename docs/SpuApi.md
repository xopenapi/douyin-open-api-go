# \SpuApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoiSpuQuery**](SpuApi.md#PoiSpuQuery) | **Get** /poi/spu/query | 查询SPU
[**PoiSpuSync**](SpuApi.md#PoiSpuSync) | **Post** /poi/spu/sync | SPU同步



## PoiSpuQuery

> PoiSpuQueryRsp PoiSpuQuery(ctx, accessToken, spuExtId)

查询SPU

查询SPU

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**spuExtId** | **string**|  | 

### Return type

[**PoiSpuQueryRsp**](PoiSpuQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoiSpuSync

> PoiSpuSyncRsp PoiSpuSync(ctx, accessToken, optional)

SPU同步

SPU同步

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
 **optional** | ***PoiSpuSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoiSpuSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**| SPU描述 | 
 **name** | **optional.String**| SPU名称 | 
 **order** | **optional.Int64**| SPU展示顺序,降序 | 
 **spuExtId** | **optional.String**| 接入方SPU ID | 
 **spuType** | **optional.Int64**| spu类型号，1-酒店民宿房型，90-景区门票，91-团购券 20 电商实体商品 21 电商虚拟商品 | 
 **status** | **optional.Int64**| 在线状态 1 - 在线; 2 - 下线 | 
 **supplierExtId** | **optional.String**| 接入方店铺ID | 
 **attributes** | [**optional.Interface of PoiSupplier**](PoiSupplier.md)|  | 

### Return type

[**PoiSpuSyncRsp**](PoiSpuSyncRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

