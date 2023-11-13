package gox

import "github.com/bwmarrin/snowflake"

func NewSnowID(n int64) (snowflake.ID, error) {
	node, err := snowflake.NewNode(n)
	if err != nil {
		return 0, err
	}
	id := node.Generate()
	return id, nil
}

func NewSnowIDString(n int64) string {
	id, _ := NewSnowID(n)
	return id.String()
}

func NewSnowIDInt64(n int64) int64 {
	id, _ := NewSnowID(n)
	return id.Int64()
}
