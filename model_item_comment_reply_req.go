/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// ItemCommentReplyReq
type ItemCommentReplyReq struct {
	CommentId string `json:"comment_id,omitempty"`
	Content   string `json:"content,omitempty"`
	ItemId    string `json:"item_id,omitempty"`
}