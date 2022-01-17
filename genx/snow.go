package genx

import "github.com/bwmarrin/snowflake"

func NewSnow(n int64) (snowflake.ID, error) {
	node, err := snowflake.NewNode(n)
	if err != nil {
		return 0, err
	}
	id := node.Generate()
	return id, nil
}

func SnowString(n int64) string {
	id, _ := NewSnow(n)
	return id.String()
}

func SnowInt64(n int64) int64 {
	id, _ := NewSnow(n)
	return id.Int64()
}
