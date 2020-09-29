/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// TouTiaoVideoListRspDataList struct for TouTiaoVideoListRspDataList
type TouTiaoVideoListRspDataList struct {
	Cover      string                            `json:"cover,omitempty"`
	CreateTime int64                             `json:"create_time,omitempty"`
	ItemId     string                            `json:"item_id,omitempty"`
	Title      string                            `json:"title,omitempty"`
	Statistics TouTiaoVideoListRspDataStatistics `json:"statistics,omitempty"`
}