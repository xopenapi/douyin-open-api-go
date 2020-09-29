# \GrouponApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseGrouponDetail**](GrouponApi.md#EnterpriseGrouponDetail) | **Get** /enterprise/groupon/detail | 团购活动详情
[**EnterpriseGrouponList**](GrouponApi.md#EnterpriseGrouponList) | **Get** /enterprise/groupon/list | 团购活动列表
[**EnterpriseGrouponOffline**](GrouponApi.md#EnterpriseGrouponOffline) | **Post** /enterprise/groupon/offline | 团购活动下线
[**EnterpriseGrouponSave**](GrouponApi.md#EnterpriseGrouponSave) | **Post** /enterprise/groupon/save | 创建团购活动



## EnterpriseGrouponDetail

> EnterpriseGrouponDetailRsp EnterpriseGrouponDetail(ctx, accessToken, openId, optional)

团购活动详情

团购活动详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseGrouponDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponDetailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **grouponIds** | [**optional.Interface of []string**](string.md)| 团购活动的Id | 

### Return type

[**EnterpriseGrouponDetailRsp**](EnterpriseGrouponDetailRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponList

> EnterpriseGrouponListRsp EnterpriseGrouponList(ctx, accessToken, openId, count, optional)

团购活动列表

团购活动列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
**count** | **int64**| 每页数量 | 
 **optional** | ***EnterpriseGrouponListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.Int64**| 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。 | 
 **status** | **optional.Int64**| 状态筛选 1 有效 2 审核中 3 审核失败 4 中止 | 

### Return type

[**EnterpriseGrouponListRsp**](EnterpriseGrouponListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponOffline

> EnterpriseGrouponOfflineRsp EnterpriseGrouponOffline(ctx, accessToken, openId, optional)

团购活动下线

团购活动下线

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseGrouponOfflineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponOfflineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **grouponId** | **optional.String**| 团购活动Id | 

### Return type

[**EnterpriseGrouponOfflineRsp**](EnterpriseGrouponOfflineRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponSave

> EnterpriseGrouponSaveRsp EnterpriseGrouponSave(ctx, accessToken, openId, optional)

创建团购活动

创建团购活动

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseGrouponSaveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponSaveOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actualAmount** | **optional.Int64**| 实际金额(单位分) | 
 **codeType** | **optional.Int64**| 0券码生成的方式;1系统生成;2自定义上传 | 
 **h5Url** | **optional.String**| 团购活动详情页链接 | 
 **orderLimit** | **optional.Int64**| 单用户购买数量上限 | 
 **auditMsg** | **optional.String**| 审核失败原因 | 
 **grouponId** | **optional.String**| 团购活动Id，审核失败修改用 | 
 **startTime** | **optional.Int64**| 活动开始时间 unix time | 
 **stock** | **optional.Int64**| 团购活动库存总数 | 
 **title** | **optional.String**| 卡券标题 | 
 **endTime** | **optional.Int64**| 活动截止时间 unix time | 
 **merchantName** | **optional.String**| 商户名称 | 
 **serviceNumber** | **optional.String**| 联系电话 | 
 **status** | **optional.Int64**| 活动状态 创建时可以忽略 1有效 2审核中 3审核失败 4中止 | 
 **notification** | **optional.String**| 团购须知 | 
 **originalAmount** | **optional.Int64**| 原价(单位分) | 
 **soldCount** | **optional.Int64**| 已售出数量 | 
 **useType** | **optional.Int64**| 团购使用方式 1 到店核销 | 
 **grouponGoods** | [**optional.Interface of []EnterpriseGrouponSaveGrouponGoods**](EnterpriseGrouponSaveGrouponGoods.md)| 团购商品 | 
 **coverImages** | [**optional.Interface of []string**](string.md)| 封面图 | 
 **poiIds** | [**optional.Interface of []string**](string.md)| 绑定的POI 列表 默认展示全部门店 | 

### Return type

[**EnterpriseGrouponSaveRsp**](EnterpriseGrouponSaveRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

