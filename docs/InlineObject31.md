# InlineObject31

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchType** | **int64** | 匹配类型，0-离线匹配 1-实时匹配。离线匹配，不会实时返回结果，最大上传1w个数据，通过/poi/supplier/match/query/接口查询匹配结果； 在线匹配，实时返回结果，最大上传100个数据，需要申请授权。 | [optional] 
**MatchDataList** | [**[]PoiSupplierMatchMatchDataList**](_poi_supplier_match_match_data_list.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


