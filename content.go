package captcha

import (
	"math/rand"
)

const Default_Length = 4

var Default_Charset = []byte("abcdefghijkmnpqrstuvwxyz123456789ABCDEFGHJKLMNPQRSTUVWXYZ")

type Content struct {
	Length  int
	Charset []byte
}

func NewContent() *Content {
	c := &Content{}
	return c
}

func (c *Content) SetLength(len int) *Content {
	c.Length = len
	return c
}
func (c *Content) SetCharset(b []byte) *Content {
	c.Charset = b
	return c
}

// 在Charset中生成Length长度的随机码
func (c *Content) Build() string {
	if c.Length <= 0 {
		c.Length = Default_Length
	}
	if len(c.Charset) == 0 {
		c.Charset = Default_Charset
	}
	charset := ""
	for i := c.Length; i > 0; i-- {
		charset += string(c.Charset[i])
	}
	return charset
}

// 获取范围内的随机数
func RandRange(max int) int {
	return rand.Intn(max)
}
