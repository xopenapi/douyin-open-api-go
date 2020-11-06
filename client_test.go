package douyin_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/xopenapi/douyin-open-api-go"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
	"unsafe"
)

var clientKey1 = "awpamlod9sp7tgqv"
var clientSecret1 = "2050893f72b14fbf2d16dbf63e236553"
var openId1 = "3100890d-de49-4906-aefe-77ea8652afcb"
var code1 = "ZHWHvbCIhweyVTSn7eFsV10gRMrdAzrRYFJf"
var accessToken1 = "act.646888359106e8ffdf98308d72e58715a7oDDAVwCgEFWRj4A3Pb1jjIDWZC"

var clientKey = "awustlseg02cmzdb"
var clientSecret = "da6277aa187c17369919a991c2f95940"
var openId = "79c367dd-1a8f-4587-9bcc-210906ae0481"
var code = "DdsJkTr7L25cYtnKioAjvAzv6c2Nk2Pi83EH"
var accessToken = "act.5f6be9d0e225707354b86582900ea3c8cSETBPKhXsjskmHpV4OblyEpV6Vx"

// qiqi 抖音号
var to_openId = "28637e5d-a462-4c96-92d1-533b9d0df93b"

var refreshToken = "rft.bac75f7df8678f22432283be3a8a4c37dKmsX49rM0CMew98R730BY3XGBaq"
var scope = "user_info"
var grantType = "authorization_code"
var responseType = "code"
var state = "111111111111"

var redirectUrl = "http://proxy-admin.dev.lucfish.com"
var count = int64(10)
var cursor = int64(0)
var itemId = ""

func TestLauchHttpServer(t *testing.T) {
	http.HandleFunc("/douyin/login", func(w http.ResponseWriter, r *http.Request) {
		redirectURI := "http://proxy-admin.dev.lucfish.com:7898/douyin/login/redirect"
		redirectUrl := url.QueryEscape(redirectURI)
		//https://open.douyin.com/platform/oauth/connect/?client_key=awpamlod9sp7tgqv&response_type=code&scope=user_info&state=dy&redirect_uri=http://proxy-admin.dev.lucfish.com:7898/douyin/login/redirect
		//url := fmt.Sprintf("https://open.douyin.com/platform/oauth/connect/?client_key=%s&response_type=%s&scope=%s&redirect_uri=%s&state=%s",
		redirectOauthURL := "https://open.douyin.com/platform/oauth/connect?client_key=%s&response_type=%s&scope=%s&redirect_uri=%s&state=%s"
		url := fmt.Sprintf(redirectOauthURL, clientKey, responseType, scope, redirectUrl, state)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/douyin/login/redirect", func(w http.ResponseWriter, r *http.Request) {
		log.Print("redirect")
		r.ParseForm()
		fmt.Println(r.Form)
	})

	http.ListenAndServe(":8888", nil)
}

// 7	1	测试通过	账号授权
func TestOauthApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	/*	rsp1, r, err := client.OauthApi.Connect(context.Background(), clientKey,
			responseType,
			scope,
			redirectUrl,
			nil,)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		t.Log(rsp1)

		rsp2, r, err := client.OauthApi.Authorize(context.Background(), clientKey,
			responseType,
			scope,
			redirectUrl,
			nil,)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		t.Log(rsp2)*/

	rsp3, r, err := client.OauthApi.AccessToken(context.Background(), clientKey, clientSecret, code, grantType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3.Data.AccessToken)

	rsp4, r, err := client.OauthApi.RenewRefreshToken(context.Background(), clientKey, refreshToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	rsp5, r, err := client.OauthApi.ClientToken(context.Background(), clientKey, clientSecret, grantType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

	rsp6, r, err := client.OauthApi.RefreshToken(context.Background(), clientKey, refreshToken, grantType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp6)

	// 请求的域名跟其他接口也不同，是https://aweme.snssdk.com/
	rsp7, r, err := client.OauthApi.AuthorizeV2(context.Background(), clientKey, responseType, scope, redirectUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp7)

}

// 3	2	测试通过	用户管理
func TestUserinfoApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp1, r, err := client.UserinfoApi.Userinfo(context.Background(), openId, accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.UserinfoApi.FansList(context.Background(), openId, accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.UserinfoApi.FollowingList(context.Background(), openId, accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

}

// 7	4	互动管理	评论管理（普通用户）	测试通过
func TestCommentApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	itemId = "@9VwRxeeeT8w3byWibcI5Qs783WzoNPGFPJx0qQyhLlEXa/P660zdRmYqig357zEBaHA377QPTe+/dH3iP7hRVA=="
	rsp1, r, err := client.CommentApi.ItemCommentList(context.Background(), openId, accessToken, count, itemId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1.Data.List[0].CommentId)

	commentId := rsp1.Data.List[0].CommentId
	rsp2, r, err := client.CommentApi.ItemCommentReplyList(context.Background(), openId, accessToken, count, itemId, commentId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	var body douyin.ItemCommentReplyReq
	body.CommentId = rsp1.Data.List[0].CommentId
	body.Content = "我爱你"
	body.ItemId = itemId
	rsp3, r, err := client.CommentApi.ItemCommentReply(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

}

// 4	4-2 评论管理（企业号）	测试通过
func TestCommentEnterprise(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	//首先获取视频列表，拿到1个视频的itemId
	var itemIds = getVideos(t)

	// 该api已自动将 itemId 进行 urlEncode 了
	itemId = "@9VwU1/6eU81pLGf2dN8sVs783WLuOvGHP5N2rA2lK1AWb/n960zdRmYqig357zEBpRdBtTEjuK+IZVM2DbFJqA=="
	itemId = itemIds[0]
	// 获取该视频的评论列表
	rsp4, r, err := client.CommentApi.VideoCommentList(context.Background(), openId, accessToken, count, itemId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	// 获取某个评论 下的回复列表
	var commentId = rsp4.Data.List[0].CommentId
	rsp5, r, err := client.CommentApi.VideoCommentReplyList(context.Background(), openId, accessToken, count, itemId, commentId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

	// 回复某个评论
	var videoCommentReplyReq douyin.VideoCommentReplyReq
	videoCommentReplyReq.CommentId = commentId
	videoCommentReplyReq.ItemId = itemId
	videoCommentReplyReq.Content = "这是测试数据1"
	rsp6, r, err := client.CommentApi.VideoCommentReply(context.Background(), openId, accessToken, videoCommentReplyReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp6)

	// 置顶评论，注意，该评论不能是 回复列表的 comment_id
	var videoCommentTopReq douyin.VideoCommentTopReq
	videoCommentTopReq.ItemId = itemId
	videoCommentTopReq.CommentId = rsp4.Data.List[len(rsp4.Data.List)-1].CommentId
	videoCommentTopReq.Top = true
	rsp7, r, err := client.CommentApi.VideoCommentTop(context.Background(), openId, accessToken, videoCommentTopReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp7)
}

// 13	用户数据
func TestDataApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	dateType := int64(7)
	endDate := ""
	rsp1, r, err := client.DataApi.DataExternalUserItem(context.Background(), accessToken, openId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1.Data.ResultList)

	rsp2, r, err := client.DataApi.DataExternalUserFans(context.Background(), accessToken, openId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2.Data.ResultList)

	rsp3, r, err := client.DataApi.DataExternalUserLike(context.Background(), accessToken, openId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3.Data.ResultList)

	rsp4, r, err := client.DataApi.DataExternalUserComment(context.Background(), accessToken, openId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4.Data.ResultList)

	rsp5, r, err := client.DataApi.DataExternalUserShare(context.Background(), accessToken, openId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5.Data.ResultList)

	rsp6, r, err := client.DataApi.DataExternalUserProfile(context.Background(), accessToken, openId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp6.Data.ResultList)

	rsp7, r, err := client.DataApi.DataExternalItemBase(context.Background(), accessToken, openId, itemId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp7)

	rsp8, r, err := client.DataApi.DataExternalItemLike(context.Background(), accessToken, openId, itemId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp8)

	rsp9, r, err := client.DataApi.DataExternalItemComment(context.Background(), accessToken, openId, itemId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp9)

	rsp10, r, err := client.DataApi.DataExternalItemPlay(context.Background(), accessToken, openId, itemId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp10)

	rsp11, r, err := client.DataApi.DataExternalItemShare(context.Background(), accessToken, openId, itemId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp11)

	// 通过
	rsp12, r, err := client.DataApi.FansData(context.Background(), accessToken, openId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp12)

	rsp13, r, err := client.DataApi.DataExternalSdkShare(context.Background(), accessToken, endDate, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp13)

}

// 7
func TestDataPoiApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	poiId := ""
	beginDate := ""
	endDate := ""
	dateType := int64(1)
	billboardType := int64(1)
	amapId := ""
	rsp1, r, err := client.DataPoiApi.DataExternalPoiBase(context.Background(), accessToken, poiId, beginDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.DataPoiApi.DataExternalPoiUser(context.Background(), accessToken, poiId, dateType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.DataPoiApi.DataExternalPoiServiceBase(context.Background(), accessToken, poiId, beginDate, endDate, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	rsp4, r, err := client.DataPoiApi.DataExternalPoiServiceUser(context.Background(), accessToken, poiId, dateType, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	rsp5, r, err := client.DataPoiApi.DataExternalPoiBillboard(context.Background(), accessToken, billboardType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

	rsp6, r, err := client.DataPoiApi.DataExternalPoiClaimList(context.Background(), accessToken, openId, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp6)

	rsp7, r, err := client.DataPoiApi.PoiBaseQueryAmap(context.Background(), accessToken, amapId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp7)

}

// 1
func TestDevtoolApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	micappId := ""
	rsp1, r, err := client.DevtoolApi.DevtoolMicappIsLegal(context.Background(), accessToken, micappId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

}

// 21 企业号开放能力-管理意向用户	测试通过
func TestEnterpriseApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	startTime := int64(1)
	endTime := time.Now().UnixNano()
	actionType := int64(1)
	var userId = ""
	tagId := int64(1)
	// 获取意向用户列表
	rsp1, r, err := client.EnterpriseApi.EnterpriseLeadsUserList(context.Background(), accessToken, openId, count, startTime, endTime, actionType, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1.Data.List)
	t.Log(rsp1.Data.Users[1].OpenId)

	// 获取意向用户详情
	userId = rsp1.Data.Users[1].OpenId
	rsp2, r, err := client.EnterpriseApi.EnterpriseLeadsUserDetail(context.Background(), accessToken, openId, userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	// 获取意向用户互动记录
	rsp3, r, err := client.EnterpriseApi.EnterpriseLeadsUserActionList(context.Background(), accessToken, openId, count, userId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	// 获取标签列表
	rsp4, r, err := client.EnterpriseApi.EnterpriseLeadsTagList(context.Background(), accessToken, openId, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	// 获取打标签的用户列表
	tagId = 71783
	rsp5, r, err := client.EnterpriseApi.EnterpriseLeadsTagUserList(context.Background(), accessToken, openId, count, tagId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

	// 创建标签
	var enterpriseLeadsTagCreateReq douyin.EnterpriseLeadsTagCreateReq
	enterpriseLeadsTagCreateReq.TagName = "美女"
	rsp6, r, err := client.EnterpriseApi.EnterpriseLeadsTagCreate(context.Background(), accessToken, openId, enterpriseLeadsTagCreateReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp6.Data.TagId)
	tagId = rsp6.Data.TagId

	// 编辑标签
	var enterpriseLeadsTagUpdateReq douyin.EnterpriseLeadsTagUpdateReq
	enterpriseLeadsTagUpdateReq.TagName = "美女啊1"
	enterpriseLeadsTagUpdateReq.TagId = tagId
	rsp7, r, err := client.EnterpriseApi.EnterpriseLeadsTagUpdate(context.Background(), accessToken, openId, enterpriseLeadsTagUpdateReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp7)

	// 给用户设置标签
	var enterpriseLeadsTagUserUpdateReq douyin.EnterpriseLeadsTagUserUpdateReq
	enterpriseLeadsTagUserUpdateReq.TagId = tagId
	enterpriseLeadsTagUserUpdateReq.UserId = to_openId
	enterpriseLeadsTagUserUpdateReq.Bind = true
	rsp9, r, err := client.EnterpriseApi.EnterpriseLeadsTagUserUpdate(context.Background(), accessToken, openId, enterpriseLeadsTagUserUpdateReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp9)

	// 删除标签
	var enterpriseLeadsTagDeleteReq douyin.EnterpriseLeadsTagDeleteReq
	enterpriseLeadsTagDeleteReq.TagId = tagId
	rsp8, r, err := client.EnterpriseApi.EnterpriseLeadsTagDelete(context.Background(), accessToken, openId, enterpriseLeadsTagDeleteReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp8)

}

// 4
func TestEnterpriseGrouponApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp1, r, err := client.EnterpriseGrouponApi.EnterpriseGrouponSave(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.EnterpriseGrouponApi.EnterpriseGrouponList(context.Background(), accessToken, openId, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.EnterpriseGrouponApi.EnterpriseGrouponDetail(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	rsp4, r, err := client.EnterpriseGrouponApi.EnterpriseGrouponOffline(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

}

// 5
func TestEnterpriseImApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	// 创建客服
	var enterpriseImPersonCreateReq douyin.EnterpriseImPersonCreateReq
	enterpriseImPersonCreateReq.Nickname = "白天鹅"
	rsp1, r, err := client.EnterpriseImApi.EnterpriseImPersonCreate(context.Background(), accessToken, openId, enterpriseImPersonCreateReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1.Data.PersonaId)
	var personaId = rsp1.Data.PersonaId

	// 获取客服列表
	rsp3, r, err := client.EnterpriseImApi.EnterpriseImPersonList(context.Background(), accessToken, openId, cursor, count)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3.Data.Personas)

	//创建客服会话
	var enterpriseImPersonConversationCreateReq douyin.EnterpriseImPersonConversationCreateReq
	enterpriseImPersonConversationCreateReq.PersonaId = personaId
	enterpriseImPersonConversationCreateReq.ToUserId = to_openId
	rsp4, r, err := client.EnterpriseImApi.EnterpriseImPersonConversationCreate(context.Background(), accessToken, openId, enterpriseImPersonConversationCreateReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4.Data.ErrorCode)

	// 客服会话结束
	var enterpriseImPersonConversationDeleteReq douyin.EnterpriseImPersonConversationDeleteReq
	enterpriseImPersonConversationDeleteReq.PersonaId = personaId
	enterpriseImPersonConversationDeleteReq.ToUserId = to_openId
	rsp5, r, err := client.EnterpriseImApi.EnterpriseImPersonConversationDelete(context.Background(), accessToken, openId, enterpriseImPersonConversationDeleteReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

	//删除客服
	var enterpriseImPersonDeleteReq douyin.EnterpriseImPersonDeleteReq
	enterpriseImPersonDeleteReq.PersonaId = personaId
	rsp2, r, err := client.EnterpriseImApi.EnterpriseImPersonDelete(context.Background(), accessToken, openId, enterpriseImPersonDeleteReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

}

// 3
func TestEnterpriseImCardApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp1, r, err := client.EnterpriseImCardApi.EnterpriseImCardSave(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.EnterpriseImCardApi.EnterpriseImCardList(context.Background(), accessToken, openId, cursor, count)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.EnterpriseImCardApi.EnterpriseImCardDelete(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

}

// 5
func TestGroupOrderApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	startTime := int64(1)
	endTime := int64(1)
	orderId := ""

	rsp1, r, err := client.GrouponOrderApi.EnterpriseGrouponOrderRefundApplyList(context.Background(), accessToken, openId, count, startTime, endTime, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.GrouponOrderApi.EnterpriseGrouponOrderDetail(context.Background(), accessToken, openId, orderId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.GrouponOrderApi.EnterpriseGrouponOrderList(context.Background(), accessToken, openId, count, startTime, endTime, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	rsp4, r, err := client.GrouponOrderApi.EnterpriseGrouponOrderOverview(context.Background(), accessToken, openId, startTime, endTime)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	rsp5, r, err := client.GrouponOrderApi.EnterpriseGrouponOrderRefundConfirm(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

}

// 3	数据开放服务	热点视频数据
func TestHotsearchApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)
	hotSentence := ""

	rsp1, r, err := client.HotsearchApi.HotsearchSentences(context.Background(), accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.HotsearchApi.HotsearchTrendingSentences(context.Background(), accessToken, cursor, count)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.HotsearchApi.HotsearchVideos(context.Background(), accessToken, hotSentence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

}

// 2-3-1	测试通过，视频和文字、图片 通过
func TestImApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	var body douyin.EnterpriseImMessageSendReq
	//item_id: @9VwRxeeeT8w3byWibcI5Qs783WLuOvGHP5N2rA2lK1AWb/n960zdRmYqig357zEB+t0DB6zzc5FI301LfBxA/Q==
	//media_id:tos-cn-i-0813/8fde29664f5a4b6e8bc0aed7afbc3175
	//text:
	type image struct {
		MediaId string `json:"media_id"`
	}
	//转换成JSON字符串
	// @9VwRxeeeT8w3byWibcI5Qs783W3hP/+BOZxyrQmiJlUbZvb+60zdRmYqig357zEB1Fe4IiK/1Kcqd1+ELA7AWw==
	// tos-cn-i-0813/d4198b72a6c4476e95b9345a874e5992
	imageTest := image{
		MediaId: "8SPjP5rgbJAoRxZnmHqfjXK7njTYNhAW8yPHupjsCXO2xcffdxr2Wfr6d6o=",
	}
	imagejsons, errs := json.Marshal(imageTest) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(imagejsons)) //byte[]转换成string 输出

	// 视频
	type video struct {
		ItemId string `json:"item_id"`
	}
	//转换成JSON字符串
	videoTest := video{
		ItemId: "@9VwRxeeeT8w3byWibcI5Qs783WLuOvGHP5N2rA2lK1AWb/n960zdRmYqig357zEB+t0DB6zzc5FI301LfBxA/Q==",
	}
	videojsons, errs := json.Marshal(videoTest) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(videojsons)) //byte[]转换成string 输出

	// 文字
	type text struct {
		Text string `json:"text"`
	}
	//转换成JSON字符串
	textTest := text{
		Text: "ssssssss",
	}
	textjsons, errs := json.Marshal(textTest) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(textjsons)) //byte[]转换成string 输出

	body.Content = string(imagejsons)
	body.MessageType = "image"
	body.ToUserId = to_openId
	rsp1, r, err := client.ImApi.EnterpriseImMessageSend(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

}

// 2 测试通过
func TestImageApiClient(t *testing.T) {

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	var image, err = os.Open("/Users/mac/Downloads/logo.png")
	if err != nil {
		t.Log(err)
	}
	// 上传图片
	rsp1, r, err := client.ImageApi.DouyinImageUpload(context.Background(), openId, accessToken, image)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

	var imageId = rsp1.Data.Image.ImageId
	t.Log(imageId)

	// 发布图片
	var body douyin.InlineObject
	body.ImageId = imageId

	rsp2, r, err := client.ImageApi.DouyinImageCreate(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2.Data.ItemId)
}

// 1
func TestJsApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp1, r, err := client.JsApi.JsGetticket(context.Background(), accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

}

// 4 测试通过	素材
func TestMediaApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	filename := "/Users/mac/Downloads/111111111.png"
	media, err := os.Open(filename)
	if err != nil {
		t.Log(err)
	}

	rsp1, r, err := client.MediaApi.EnterpriseMediaUpload(context.Background(), accessToken, openId, media)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1.Data.Media.MediaId)

	rsp2, r, err := client.MediaApi.EnterpriseMediaTempUpload(context.Background(), accessToken, openId, media)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2.Data.Media.MediaId)

	rsp3, r, err := client.MediaApi.EnterpriseMediaList(context.Background(), accessToken, openId, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3.Data.Medias[0].MediaId)

	var body douyin.EnterpriseMediaDeleteReq
	body.MediaId = ""
	rsp4, r, err := client.MediaApi.EnterpriseMediaDelete(context.Background(), accessToken, openId, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

}

// 3
func TestPayApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)
	merchantId := int64(1)
	liveId := int64(1)

	var body douyin.DouyinPayAccountTransReq
	rsp1, r, err := client.PayApi.DouyinPayAccountTrans(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.PayApi.DouyinPayOrderQuery(context.Background(), accessToken, merchantId, liveId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.PayApi.DouyinPayAccountQuery(context.Background(), accessToken, merchantId, liveId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

}

// 5
func TestPoiApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	supplierExtId := ""
	amapId := ""

	rsp1, r, err := client.PoiApi.PoiSupplierSync(context.Background(), accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.PoiApi.PoiSupplierQuery(context.Background(), accessToken, supplierExtId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.PoiApi.PoiQuery(context.Background(), accessToken, amapId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	rsp4, r, err := client.PoiApi.PoiSupplierMatch(context.Background(), accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	rsp5, r, err := client.PoiApi.PoiSupplierMatchQuery(context.Background(), accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

}

// 5
func TestPoiOrderApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	var body1 douyin.PoiOrderStatusReq
	rsp1, r, err := client.PoiOrderApi.PoiOrderStatus(context.Background(), accessToken, body1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	var body2 douyin.PoiExtHotelOrderCommitReq
	rsp2, r, err := client.PoiOrderApi.PoiExtHotelOrderCommit(context.Background(), body2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	var poiExtHotelOrderStatusReq douyin.PoiExtHotelOrderStatusReq
	rsp3, r, err := client.PoiOrderApi.PoiExtHotelOrderStatus(context.Background(), poiExtHotelOrderStatusReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	var poiExtHotelOrderCancelReq douyin.PoiExtHotelOrderCancelReq
	rsp4, r, err := client.PoiOrderApi.PoiExtHotelOrderCancel(context.Background(), poiExtHotelOrderCancelReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	rsp5, r, err := client.PoiOrderApi.PoiOrderSync(context.Background(), accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)

}

// 2
func TestRankApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	myType := int32(1)

	rsp1, r, err := client.RankApi.DiscoveryEntRankItem(context.Background(), accessToken, myType, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.RankApi.DiscoveryEntRankVersion(context.Background(), accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

}

// 1
func TestSandboxApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	var sandboxWebhookEventEendReq douyin.SandboxWebhookEventEendReq
	rsp1, r, err := client.SandboxApi.SandboxWebhookEventSend(context.Background(), accessToken, sandboxWebhookEventEendReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

}

// 2
func TestSkuApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	startDate := ""
	endDate := ""
	spuExtId := []string{}

	rsp1, r, err := client.SkuApi.PoiSkuSync(context.Background(), accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.SkuApi.PoiExtHotelSku(context.Background(), spuExtId, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

}

// 2
func TestSpuApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	spuExtId := ""

	rsp1, r, err := client.SpuApi.PoiSpuSync(context.Background(), accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.SpuApi.PoiSpuQuery(context.Background(), accessToken, spuExtId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

}

// 3
func TestStarApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	hotListType := int64(1)
	uniqueId := ""

	rsp1, r, err := client.StarApi.StarHostList(context.Background(), accessToken, hotListType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.StarApi.StarAuthorScore(context.Background(), accessToken, openId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	rsp3, r, err := client.StarApi.StarAuthorScoreV2(context.Background(), accessToken, uniqueId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

}

// 读取文件到[]byte中
func file2Bytes(filename string) ([]byte, error) {

	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}

// 抖音 单次发布视频，视频文件小 测试通过
func TestPublishVideoOnce(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	filename := "/Users/mac/Downloads/2.mp4"
	video, err := os.Open(filename)
	if err != nil {
		t.Log(err)
	}

	// 单次上传并发布，测试通过
	rsp1, r, err := client.VideoApi.DouyinVideoUpload(context.Background(), openId, accessToken, video)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	videoId := rsp1.Data.Video.VideoId
	t.Log(rsp1.Data.Video.VideoId)
	var body douyin.DouyinVideoCreateReq
	body.VideoId = videoId
	// 发布
	rsp5_0, r, err := client.VideoApi.DouyinVideoCreate(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp5_0.Data.ItemId)
}

// 抖音分片上传，初步估计 swagger处理 byte型数据时有误，导致生成的代码有问题
func TestUploadVideoPart(t *testing.T) {
	uploadId := "@9VwU1/6eU81pLGf2dN8sVs76hmHqbP2EaZwirVikJ1YQbvj80iGwKwkhgQby5DoKCJvK0tgyE3SNQJptBWjjYw=="
	partNum := int64(1)
	uploadOneVideo(uploadId, partNum)
}

type DouyinVideoPartUploadRsp struct {
	Data  douyin.DouyinVideoPartUploadRspData `json:"data,omitempty"`
	Extra douyin.FollowingListRspDataExtra    `json:"extra,omitempty"`
}

// 上传单个分片： 非自动生成代码
func uploadOneVideo(uploadId string, partNum int64) {
	urlRaw := "https://open.douyin.com/video/part/upload/?open_id=%s&access_token=%s&upload_id=%s&part_number=%d"
	var url = fmt.Sprintf(urlRaw, openId, accessToken, uploadId, partNum)
	fmt.Println(url)
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, errFile1 := file2Bytes("/Users/mac/Downloads/1.mp4")
	part1,
		//errFile1 := writer.CreateFormFile("video",filepath.Base("/Users/mac/Downloads/1.mp4"))
		errFile1 := writer.CreateFormFile("video", "j")

	part1.Write(file)
	if errFile1 != nil {
		fmt.Println(errFile1)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	// 设置了边界
	req.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println(writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	var resw = *(**DouyinVideoPartUploadRsp)(unsafe.Pointer(&body))
	fmt.Println("-------------------------------------")
	fmt.Println(resw.Extra.Logid)
	fmt.Println("-------------------------------------")
}

// 抖音 分片发布视频文件，视频文件大
func TestPublishVideoPart(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	//分片初始化
	rsp2, r, err := client.VideoApi.DouyinVideoPartInit(context.Background(), openId, accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2.Data.UploadId)

	uploadId := rsp2.Data.UploadId
	partNumber := int64(1)
	filename1 := "/Users/mac/Downloads/2.mp4"
	video1, err := file2Bytes(filename1)
	if err != nil {
		t.Log(err)
	}
	// 分片上传
	rsp3, r, err := client.VideoApi.DouyinVideoPartUpload(context.Background(), openId, accessToken, uploadId, partNumber, video1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	rsp4, r, err := client.VideoApi.DouyinVideoPartComplete(context.Background(), openId, accessToken, uploadId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

	var body douyin.DouyinVideoCreateReq
	videoId_2 := rsp4.Data.Video.VideoId
	body.VideoId = string(videoId_2)
	//pass: postman、接口
	rsp5, r, err := client.VideoApi.DouyinVideoCreate(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp5)
}

// 抖音 获取视频列表
func TestGetVideos(t *testing.T) {
	getVideos(t)
}
func getVideos(t *testing.T) []string {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)
	count := int64(10)
	rsp7, r, err := client.VideoApi.DouyinVideoList(context.Background(), openId, accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	var list = rsp7.Data.List
	var items = []string{}
	for _, item := range list {
		t.Log(item.ItemId)
		items = append(items, item.ItemId)
	}
	t.Log(items)
	return items
}

// 抖音 删除视频
func TestDeleteVideos(t *testing.T) {
	// 先获取视频 itemId
	var itemIds = getVideos(t)
	var itemId = itemIds[0]

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)
	var option douyin.InlineObject1
	option.ItemId = itemId

	rsp6, r, err := client.VideoApi.DouyinVideoDelete(context.Background(), openId, accessToken, option)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp6)
}

// 查询指定 itemid 的视频
func TestGetSpecVideos(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	var itemIds = getVideos(t)
	var body douyin.InlineObject2
	body.ItemIds = itemIds

	rsp8, r, err := client.VideoApi.DouyinVideoData(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(len(rsp8.Data.List))
}

// 头条 单次发布视频，视频文件小 测试中
func TestTouTiaoPublishVideoOnce(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	filename := "/Users/mac/Downloads/2.mp4"
	video, err := os.Open(filename)
	if err != nil {
		t.Log(err)
	}

	// 单次上传并发布，测试通过
	rsp1, r, err := client.VideoApi.ToutiaoVideoUpload(context.Background(), openId, accessToken, video)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	videoId := rsp1.Data.Video.VideoId
	t.Log(rsp1.Data.Video.VideoId)
	var body douyin.DouyinVideoCreateReq
	body.VideoId = string(videoId)
	// 发布
	rsp5_0, r, err := client.VideoApi.ToutiaoVideoCreate(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp5_0.Data.ItemId)
}

// 头条 单次发布视频，视频文件小 测试中
func TestDouyinVideoShareId(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	filename := "/Users/mac/Downloads/2.mp4"
	video, err := os.Open(filename)
	if err != nil {
		t.Log(err)
	}

	// 单次上传并发布，测试通过
	rsp1, r, err := client.VideoApi.ToutiaoVideoUpload(context.Background(), openId, accessToken, video)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	videoId := rsp1.Data.Video.VideoId
	t.Log(rsp1.Data.Video.VideoId)
	var body douyin.DouyinVideoCreateReq
	body.VideoId = string(videoId)
	// 发布
	rsp5_0, r, err := client.VideoApi.ToutiaoVideoCreate(context.Background(), openId, accessToken, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp5_0.Data.ItemId)

	/*	var needCallback = true
		rsp9, r, err := client.VideoApi.DouyinVideoShareId(context.Background(), accessToken, needCallback, nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		t.Log(rsp9)*/

	keyword := "美食"
	var localOptions *douyin.DouyinVideoPositionSearchOpts
	(localOptions.City) = optional.NewString("北京")
	rsp10, r, err := client.VideoApi.DouyinVideoPositionSearch(context.Background(), accessToken, count, keyword, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp10)

}

// 21	视频测试
func TestVideoApiClient(t *testing.T) {
	// 抖音 单次发布视频
	TestPublishVideoOnce(t)

	// 抖音 分片发布视频
	TestPublishVideoPart(t)

	// 抖音 获取视频列表
	TestGetVideos(t)

	//删除单个视频
	TestDeleteVideos(t)

	//视频分享结果、位置信息
	TestDouyinVideoShareId(t)

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp11, r, err := client.VideoApi.ToutiaoVideoUpload(context.Background(), openId, accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp11)

	/*	rsp12, r, err := client.VideoApi.ToutiaoVideoCreate(context.Background(), openId, accessToken, nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		t.Log(rsp12)*/

	rsp13, r, err := client.VideoApi.ToutiaoVideoPartInit(context.Background(), openId, accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp13)

	uploadId := ""
	partNumber := int64(1)
	rsp14, r, err := client.VideoApi.ToutiaoVideoPartUpload(context.Background(), openId, accessToken, uploadId, partNumber, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp14)

	rsp15, r, err := client.VideoApi.ToutiaoVideoPartComplete(context.Background(), openId, accessToken, uploadId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp15)

	rsp16, r, err := client.VideoApi.ToutiaoVideoList(context.Background(), openId, accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp16)

	rsp17, r, err := client.VideoApi.ToutiaoVideoData(context.Background(), openId, accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp17)

	rsp18, r, err := client.VideoApi.XiguaVideoUpload(context.Background(), openId, accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp18)

	rsp19, r, err := client.VideoApi.XiguaVideoPartInit(context.Background(), openId, accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp19)

	rsp20, r, err := client.VideoApi.XiguaVideoPartUpload(context.Background(), openId, accessToken, uploadId, partNumber, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp20)

	rsp21, r, err := client.VideoApi.XiguaVideoPartComplete(context.Background(), openId, accessToken, uploadId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp21)

	rsp22, r, err := client.VideoApi.XiguaVideoCreate(context.Background(), openId, accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp22)

	rsp23, r, err := client.VideoApi.XiguaVideoList(context.Background(), openId, accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp23)

	rsp24, r, err := client.VideoApi.XiguaVideoData(context.Background(), openId, accessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp24)

}

// 4	搜索管理
func TestVideoSearchApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	keyword := "爱"
	commentId := "@9VwRxeeeT8w3byWibcI5Qs783WzoNPGKPZxyrg2gJlETb/H060zdRmYqig357zEBO1xTQcVkr0z4CLgIEuBWRA=="

	rsp1, r, err := client.VideoSearchApi.VideoSearch(context.Background(), openId, accessToken, count, keyword, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)
	//t.Log(rsp1.Data.List[0].SecItemId)

	var secItemId = "@9VwRxeeeT8w3byWibcI5Qs783WzoNPGFPJx0qQyhLlEXa/P660zdRmYqig357zEBaHA377QPTe+/dH3iP7hRVA=="
	rsp2, r, err := client.VideoSearchApi.VideoSearchCommentList(context.Background(), secItemId, accessToken, count, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

	var videoSearchCommentReplyReq douyin.VideoSearchCommentReplyReq
	rsp3, r, err := client.VideoSearchApi.VideoSearchCommentReply(context.Background(), accessToken, openId, videoSearchCommentReplyReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp3)

	rsp4, r, err := client.VideoSearchApi.VideoSearchCommentReplyList(context.Background(), accessToken, count, secItemId, commentId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp4)

}

// 2
func TestGrouponCodeApiClient(t *testing.T) {
	//TestLauchHttpServer()

	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp1, r, err := client.GrouponCodeApi.EnterpriseGrouponCodeVerification(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp1)

	rsp2, r, err := client.GrouponCodeApi.EnterpriseGrouponCodeStatus(context.Background(), accessToken, openId, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp2)

}
