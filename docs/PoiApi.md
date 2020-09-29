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

 **matchType** | **optional.Int64**| 匹配类型，0-离线匹配 1-实时匹配。离线匹配，不会实时返回结果，最大上传1w个数据，通过/poi/supplier/match/query/接口查询匹配结果； 在线匹配，实时返回结果，最大上传100个数据，需要申请授权。 | 
 **matchDataList** | [**optional.Interface of []PoiSupplierMatchMatchDataList**](PoiSupplierMatchMatchDataList.md)|  | 

### Return type

[**PoiSupplierMatchRsp**](PoiSupplierMatchRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
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

 **supplierExtIds** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**PoiSupplierMatchQueryRsp**](PoiSupplierMatchQueryRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
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

 **contactPhone** | **optional.String**| 联系手机号 | 
 **name** | **optional.String**| 店铺名称 | 
 **poiId** | **optional.String**| 抖音poi id, 三方如果使用高德poi id可以通过/poi/query/接口转换，其它三方poi id走poi匹配功能进行抖音poi id获取 | 
 **status** | **optional.Int64**| 在线状态 1 - 在线; 2 - 下线 | 
 **address** | **optional.String**| 店铺地址 | 
 **contactTel** | **optional.String**| 联系座机号 | 
 **description** | **optional.String**| 店铺介绍(&lt;&#x3D;500字) | 
 **supplierExtId** | **optional.String**| 接入方店铺id | 
 **type_** | **optional.Int64**| 店铺类型 1 - 酒店民宿 2 - 餐饮 3 - 景区 4 - 电商 5 - 教育 6 - 丽人 7 - 爱车 8 - 亲子 9 - 宠物 10 - 家装 11 - 娱乐场所 12 - 图文快印 | 
 **images** | [**optional.Interface of []string**](string.md)| 店铺图片 | 
 **services** | [**optional.Interface of []PoiSupplierServices**](PoiSupplierServices.md)| 店铺提供的服务列表 | 
 **attributes** | [**optional.Interface of PoiSupplierAttributes**](PoiSupplier_attributes.md)|  | 

### Return type

[**PoiSupplierSyncRsp**](PoiSupplierSyncRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

