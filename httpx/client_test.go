package httpx_test

import (
	"io"
	"testing"

	"github.com/northes/gox/httpx"
)

func TestClientGet(t *testing.T) {
	resp, err := httpx.NewClient("https://apihut.co").
		AddParam("hello", "hi").
		Get()
	if err != nil {
		t.Fatalf("Client.Get(): %v", err)
	}
	t.Logf("Resp: %s", resp)
}

func TestClientPostJson(t *testing.T) {
	type st struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	resp, err := httpx.NewClient("https://apihut.co/", httpx.WithDebug(true)).
		SetBody(&st{
			Name: "northes",
			Age:  18,
		}).Post()
	if err != nil {
		t.Fatal(err)
	}

	stResp := new(st)
	err = resp.Unmarshal(stResp)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", stResp)
}

func TestClientPostString(t *testing.T) {
	resp, err := httpx.NewClient("https://apihut.co",
		httpx.WithDebug(true),
	).SetBody("666").Post()
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Raw().Body)

	t.Logf("%v", resp.Raw().Body)
}
