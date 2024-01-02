package gox

import "encoding/json"

// JsonMarshalToStringX 强制转string，发生错误直接panic
func JsonMarshalToStringX(t any) string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(b)
}
