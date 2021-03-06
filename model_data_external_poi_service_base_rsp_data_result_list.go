/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DataExternalPoiServiceBaseRspDataResultList struct for DataExternalPoiServiceBaseRspDataResultList
type DataExternalPoiServiceBaseRspDataResultList struct {
	// 点击uv
	ClickUv int64 `json:"click_uv,omitempty"`
	// 日期
	Date string `json:"date,omitempty"`
	// 曝光uv
	ExposureUv int64 `json:"exposure_uv,omitempty"`
	// 订单成交次数
	SuccessOrderCnt int64 `json:"success_order_cnt,omitempty"`
}
