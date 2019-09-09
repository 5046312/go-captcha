package captcha

import (
	"math/rand"
	"time"
)

type Captcha struct {
	Background *Background
	Content    *Content
}

func init() {
	rand.Seed(time.Now().UnixNano() - 66)
}
