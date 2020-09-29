# DouyinVideoCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** | 文章ID，暂不开放 | [optional] 
**Text** | **string** | 视频标题， 可以带话题,@用户。 如title1#话题1 | [optional] 
**TimelinessLabel** | **int64** | 时效新闻标签，1表示使用。暂不开放 | [optional] 
**ArticleTitle** | **string** | 文章自定义标题，暂不开放 | [optional] 
**CoverTsp** | **float32** | 将传入的指定时间点对应帧设置为视频封面（单位：秒） | [optional] 
**MicroAppTitle** | **string** | 小程序标题 | [optional] 
**VideoId** | **string** | video_id, 通过/video/upload/接口得到。注意每次调用/video/create/都要调用/video/upload/生成新的video_id | [optional] 
**GameId** | **string** | 游戏id。暂不开放 | [optional] 
**MicroAppId** | **string** | 小程序id | [optional] 
**PoiId** | **string** | 地理位置id | [optional] 
**GameContent** | **string** | 游戏个性化参数 | [optional] 
**PoiName** | **string** | 地理位置名称 | [optional] 
**TimelinessKeyword** | **string** | 最多可添加3个，用&#x60;\\|&#x60;隔开。暂不开放 | [optional] 
**AtUsers** | **[]string** | 如果需要at其他用户。将text中@nickname对应的open_id放到这里。 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


