/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiSupplier
type PoiSupplier struct {
	// 联系手机号
	ContactPhone string `json:"contact_phone,omitempty"`
	// 店铺名称
	Name string `json:"name,omitempty"`
	// 抖音poi id, 三方如果使用高德poi id可以通过/poi/query/接口转换，其它三方poi id走poi匹配功能进行抖音poi id获取
	PoiId string `json:"poi_id,omitempty"`
	// 在线状态 1 - 在线; 2 - 下线
	Status int64 `json:"status,omitempty"`
	// 店铺地址
	Address string `json:"address,omitempty"`
	// 联系座机号
	ContactTel string `json:"contact_tel,omitempty"`
	// 店铺介绍(<=500字)
	Description string `json:"description,omitempty"`
	// 接入方店铺id
	SupplierExtId string `json:"supplier_ext_id,omitempty"`
	// 店铺类型 1 - 酒店民宿 2 - 餐饮 3 - 景区 4 - 电商 5 - 教育 6 - 丽人 7 - 爱车 8 - 亲子 9 - 宠物 10 - 家装 11 - 娱乐场所 12 - 图文快印
	Type int64 `json:"type,omitempty"`
	// 店铺图片
	Images []string `json:"images,omitempty"`
	// 店铺提供的服务列表
	Services   []PoiSupplierServices `json:"services,omitempty"`
	Attributes PoiSupplierAttributes `json:"attributes,omitempty"`
}