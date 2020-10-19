# PoiOrderSyncReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderType** | **int64** | 订单类型- 201 预约点餐订单, 202 餐厅预定订单, 203 餐厅排队订单, 9001 景区门票订单, 9101 团购券订单, 9201 在线预约订单, 9301 外卖订单, 140 住宿订单 | [optional] 
**DateTime** | **int64** | 订单状态。0 - 未支付, 1 - 已支付 | [optional] 
**OrderDetail** | **string** | 订单的细节，不同的订单业务有不同的结构体，请具体询问业务方字段结构 | [optional] 
**ExtShopInfo** | [**PoiOrderSyncReqExtShopInfo**](PoiOrderSyncReq_ext_shop_info.md) |  | [optional] 
**MiniApp** | [**PoiOrderSyncReqMiniApp**](PoiOrderSyncReq_mini_app.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


