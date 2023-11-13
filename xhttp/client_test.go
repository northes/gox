package xhttp

import "testing"

func TestClientGet(t *testing.T) {
	cli, err := NewClient("https://apihut.co",
		WithParams(map[string]string{
			"hello": "hi",
		}),
	)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	resp, err := cli.Get()
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

	cli, err := NewClient("https://apihut.co", WithBody(
		&st{
			Name: "northes",
			Age:  18,
		},
	))
	if err != nil {
		t.Fatal(err)
	}

	r, err := cli.Post()
	if err != nil {
		t.Fatal(err)
	}

	resp := new(st)

	err = r.Unmarshal(resp)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", resp)
}

func TestClientPostString(t *testing.T) {
	cli, err := NewClient("https://apihut.co", WithBody("666"))
	if err != nil {
		t.Fatal(err)
	}

	r, err := cli.Post()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", r)
}
