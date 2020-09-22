package douyin_test

import(
	"context"
	"github.com/xopenapi/douyin-open-api-go"
	"testing"
)

func TestClient(t *testing.T) {
	cfg := douyin.NewConfiguration()
	client := douyin.NewAPIClient(cfg)

	rsp, r, err := client.OauthApi.Connect(context.Background(), 	"",
		"",
		"",
		"",
		"",
		"自定义",
)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	t.Log(rsp)
}
