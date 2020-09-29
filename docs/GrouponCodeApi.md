# \GrouponCodeApi

All URIs are relative to *https://open.douyin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseGrouponCodeStatus**](GrouponCodeApi.md#EnterpriseGrouponCodeStatus) | **Post** /enterprise/groupon/code/status | 查看券码状态
[**EnterpriseGrouponCodeVerification**](GrouponCodeApi.md#EnterpriseGrouponCodeVerification) | **Post** /enterprise/groupon/code/verification | 券码核销



## EnterpriseGrouponCodeStatus

> EnterpriseGrouponCodeStatusRsp EnterpriseGrouponCodeStatus(ctx, accessToken, openId, optional)

查看券码状态

查看券码状态

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseGrouponCodeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponCodeStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **code** | [**optional.Interface of []string**](string.md)| 券码的列表 | 
 **grouponId** | **optional.String**| 团购活动的Id | 

### Return type

[**EnterpriseGrouponCodeStatusRsp**](EnterpriseGrouponCodeStatusRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseGrouponCodeVerification

> EnterpriseGrouponCodeVerificationRsp EnterpriseGrouponCodeVerification(ctx, accessToken, openId, optional)

券码核销

券码核销

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 调用/oauth/access_token/生成的token，此token需要用户授权。 | 
**openId** | **string**| 通过/oauth/access_token/获取，用户唯一标志 | 
 **optional** | ***EnterpriseGrouponCodeVerificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnterpriseGrouponCodeVerificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **code** | **optional.String**| 券码的列表 | 
 **grouponId** | **optional.String**| 团购活动的Id | 

### Return type

[**EnterpriseGrouponCodeVerificationRsp**](EnterpriseGrouponCodeVerificationRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

