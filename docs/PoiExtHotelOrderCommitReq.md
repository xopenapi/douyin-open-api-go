# PoiExtHotelOrderCommitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | 订单支付状态。0 - 未支付, 1 - 已支付 | [optional] 
**CheckIn** | **string** | 入住时间 yyyyMMdd | [optional] 
**CheckOut** | **string** | 离店时间 yyyyMMdd | [optional] 
**CustomerName** | **string** | 预订人姓名 | [optional] 
**CustomerPhone** | **int64** | 预订人电话 | [optional] 
**OrderId** | **string** | 抖音订单号 | [optional] 
**Remark** | **string** | 备注 | [optional] 
**TotalPrice** | **int64** | 总价, 单位人民币分 | [optional] 
**ReserveAmount** | **int64** | 预定数量 | [optional] 
**SpuExtId** | **string** | 接入方房型ID | [optional] 
**DatePrice** | [**[]PoiExtHotelOrderCommitReqDatePrice**](PoiExtHotelOrderCommitReq_date_price.md) |  | [optional] 
**CustomerList** | [**[]PoiExtHotelOrderCommitReqCustomerList**](PoiExtHotelOrderCommitReq_customer_list.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


