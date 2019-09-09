package captcha

import "encoding/base64"

type Captcha struct {
	Background *Background
	Content    *Content
}

func New() *Captcha {
	c := &Captcha{}
	c.Background = NewBackground()
	c.Content = NewContent()
	return c
}

func (c *Captcha) Base64() string {
	bg := c.Background.encodedPNG()
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(bg)
	return base64Str
}
