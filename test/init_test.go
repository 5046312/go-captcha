package test

import (
	"fmt"
	"go-captcha"
	"testing"
)

func Test_Init(t *testing.T) {
	base64 := captcha.New()
	fmt.Println(base64)
}
