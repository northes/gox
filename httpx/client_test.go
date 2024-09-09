package httpx_test

import (
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/northes/gox/httpx"
	"github.com/northes/gox/httpx/httpxutils"
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

func TestCustomJsonEncoderAndDecoder(t *testing.T) {
	type response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Ip        string    `json:"ip"`
			Country   string    `json:"country"`
			Province  string    `json:"province"`
			City      string    `json:"city"`
			District  string    `json:"district"`
			Isp       string    `json:"isp"`
			Location  string    `json:"location"`
			Source    string    `json:"source"`
			CacheTime time.Time `json:"cache_time"`
		} `json:"data"`
	}

	type st struct {
		name        string
		decoder     httpxutils.JSONUnmarshal
		wantSuccess bool
	}

	tests := []st{
		{
			name:        "json decoder",
			decoder:     json.Unmarshal,
			wantSuccess: true,
		},
		{
			name:        "nil decoder",
			decoder:     nil,
			wantSuccess: true,
		},
		{
			name: "custom decoder",
			decoder: func(data []byte, v interface{}) error {
				return json.Unmarshal(data, v)
			},
			wantSuccess: true,
		},
		{
			name: "error custom decoder",
			decoder: func(data []byte, v interface{}) error {
				return nil
			},
			wantSuccess: false,
		},
	}

	for _, tt := range tests {
		t.Logf("name: %s", tt.name)
		resp, err := httpx.NewClient("https://apihut.co/ip",
			httpx.WithJsonDecoder(tt.decoder),
		).Get()
		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("response: %s", resp.String())

		info := new(response)
		err = resp.Unmarshal(info)
		if err != nil {
			t.Fatal(err)
		}

		if tt.wantSuccess {
			require.NotEqual(t, info.Code, 0, tt.name)
			require.NotEqual(t, info.Data.Ip, "", tt.name)
		} else {
			require.Equal(t, info.Code, 0)
			require.Equal(t, info.Data.Ip, "")
		}
	}

}
