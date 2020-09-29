# PoiSupplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactPhone** | **string** | 联系手机号 | [optional] 
**Name** | **string** | 店铺名称 | [optional] 
**PoiId** | **string** | 抖音poi id, 三方如果使用高德poi id可以通过/poi/query/接口转换，其它三方poi id走poi匹配功能进行抖音poi id获取 | [optional] 
**Status** | **int64** | 在线状态 1 - 在线; 2 - 下线 | [optional] 
**Address** | **string** | 店铺地址 | [optional] 
**ContactTel** | **string** | 联系座机号 | [optional] 
**Description** | **string** | 店铺介绍(&lt;&#x3D;500字) | [optional] 
**SupplierExtId** | **string** | 接入方店铺id | [optional] 
**Type** | **int64** | 店铺类型 1 - 酒店民宿 2 - 餐饮 3 - 景区 4 - 电商 5 - 教育 6 - 丽人 7 - 爱车 8 - 亲子 9 - 宠物 10 - 家装 11 - 娱乐场所 12 - 图文快印 | [optional] 
**Images** | **[]string** | 店铺图片 | [optional] 
**Services** | [**[]PoiSupplierServices**](PoiSupplier_services.md) | 店铺提供的服务列表 | [optional] 
**Attributes** | [**PoiSupplierAttributes**](PoiSupplier_attributes.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


