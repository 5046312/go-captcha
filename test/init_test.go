package test

import (
	"fmt"
	"go-captcha"
	"testing"
)

func Test_Init(t *testing.T) {
	c := captcha.New()
	fmt.Println(c.Base64())
}
