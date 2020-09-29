# InlineObject22

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualAmount** | **int64** | 实际金额(单位分) | [optional] 
**CodeType** | **int64** | 0券码生成的方式;1系统生成;2自定义上传 | [optional] 
**H5Url** | **string** | 团购活动详情页链接 | [optional] 
**OrderLimit** | **int64** | 单用户购买数量上限 | [optional] 
**AuditMsg** | **string** | 审核失败原因 | [optional] 
**GrouponId** | **string** | 团购活动Id，审核失败修改用 | [optional] 
**StartTime** | **int64** | 活动开始时间 unix time | [optional] 
**Stock** | **int64** | 团购活动库存总数 | [optional] 
**Title** | **string** | 卡券标题 | [optional] 
**EndTime** | **int64** | 活动截止时间 unix time | [optional] 
**MerchantName** | **string** | 商户名称 | [optional] 
**ServiceNumber** | **string** | 联系电话 | [optional] 
**Status** | **int64** | 活动状态 创建时可以忽略 1有效 2审核中 3审核失败 4中止 | [optional] 
**Notification** | **string** | 团购须知 | [optional] 
**OriginalAmount** | **int64** | 原价(单位分) | [optional] 
**SoldCount** | **int64** | 已售出数量 | [optional] 
**UseType** | **int64** | 团购使用方式 1 到店核销 | [optional] 
**GrouponGoods** | [**[]EnterpriseGrouponSaveGrouponGoods**](_enterprise_groupon_save_groupon_goods.md) | 团购商品 | [optional] 
**CoverImages** | **[]string** | 封面图 | [optional] 
**PoiIds** | **[]string** | 绑定的POI 列表 默认展示全部门店 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


