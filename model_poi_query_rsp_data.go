/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiQueryRspData struct for PoiQueryRspData
type PoiQueryRspData struct {
	// 纬度
	Latitude float32 `json:"latitude,omitempty"`
	// 经度
	Longitude float32 `json:"longitude,omitempty"`
	// 抖音POI ID
	PoiId string `json:"poi_id,omitempty"`
	// POI 名称
	PoiName string `json:"poi_name,omitempty"`
	// POI地址
	Address string `json:"address,omitempty"`
	// 高德POI ID
	AmapId string `json:"amap_id,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	// POI所在城市
	City string `json:"city,omitempty"`
	// 错误码描述
	ErrorCode int64 `json:"error_code,omitempty"`
}
