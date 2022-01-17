package genx

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// UUID 遵循 RFC4122 标准，UUID为128 bit (16 字节)
type UUID [16]byte

// `rand.Reader`是一个全局、共享的密码用强随机数生成器
var rander = rand.Reader

// Nil 定义一个类型为UUID的空值
var Nil UUID

func NewUUID() UUID {
	return Must(NewRandom())
}

// Must 发生异常时触发 panic
// V1版本此处不触发panic，而是返回error
func Must(uuid UUID, err error) UUID {
	if err != nil {
		panic(err)
	}
	return uuid
}

func NewRandom() (UUID, error) {
	return NewRandomFromReader(rander)
}

// NewRandomFromReader `io.ReadFull` 从 `rand.Reader` 精确地读取len(uuid)字节数据填充进uuid
func NewRandomFromReader(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return Nil, err
	}
	// 设置uuid版本信息
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

// 按照 8-4-4-4-12 的规则将 uuid 分段编码，使用 - 连接
func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
