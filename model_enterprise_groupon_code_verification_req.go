/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// EnterpriseGrouponCodeVerificationReq
type EnterpriseGrouponCodeVerificationReq struct {
	// 券码的列表
	Code string `json:"code,omitempty"`
	// 团购活动的Id
	GrouponId string `json:"groupon_id,omitempty"`
}