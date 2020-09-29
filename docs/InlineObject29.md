# InlineObject29

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | 卡片id，创建时不传；更新时必传。同一个用户的卡片id不可重复 | [optional] 
**CardType** | **string** | 卡片类型 &#x60;question_list&#x60; - 问题列表 | [optional] 
**Content** | **string** | 卡片内容字段 json序列化成string， { \&quot;question_list\&quot; { \&quot;text\&quot; \&quot;有什么疑问呢\&quot;, \&quot;questions\&quot; [{ \&quot;name\&quot; \&quot;问题1\&quot;, \&quot;text\&quot; \&quot;关键词1\&quot; }, { \&quot;name\&quot; \&quot;问题2\&quot;, \&quot;text\&quot; \&quot;关键词2\&quot; } ] } } | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


