/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// VideoSearchCommentReplyReq
type VideoSearchCommentReplyReq struct {
	// 视频搜索接口返回的加密的视频id
	SecItemId string `json:"sec_item_id,omitempty"`
	// 需要回复的评论id（如果需要回复的是视频不传此字段）
	CommentId string `json:"comment_id,omitempty"`
	// 评论内容
	Content string `json:"content,omitempty"`
}
