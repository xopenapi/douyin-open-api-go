package douyin_test

import(
	"context"
	"github.com/xopenapi/douyin-open-api-go"
	"testing"
)

func TestClient(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	req := douyin.ConnectReq{
		ClientKey: "",
		ResponseType: "",
		Scope: "",
		OptionalScope: "",
		RedirectUri: "",
		State: "自定义",
	}

	rsp, r, err := client.OauthApi.Connect(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp)
}
