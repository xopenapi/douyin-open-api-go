# ConnectReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientKey** | **string** | 应用唯一标识 | [optional] 
**ResponseType** | **string** | 填写code | [optional] 
**Scope** | **string** | 应用授权作用域,多个授权作用域以英文逗号（,）分隔 | [optional] 
**OptionalScope** | **string** | 应用授权可选作用域,多个授权作用域以英文逗号（,）分隔，每一个授权作用域后需要加上一个是否默认勾选的参数，1为默认勾选，0为默认不勾选 | [optional] 
**RedirectUri** | **string** | 授权成功后的回调地址，必须以http/https开头。域名必须对应申请应用时填写的域名，如不清楚请联系应用申请人。 | [optional] 
**State** | **string** | 用于保持请求和回调的状态 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


