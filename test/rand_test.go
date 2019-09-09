package test

import (
	"fmt"
	"go-captcha"
	"testing"
)

func Test_Rand(t *testing.T) {
	for i := 25; i > 0; i-- {
		r := captcha.RandRange(3)
		fmt.Println(r)
	}
}

func Test_BuildContent(t *testing.T) {
	c := captcha.NewContent()
	s := c.Build()
	fmt.Println(s)
}