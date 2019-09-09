package captcha

import "encoding/base64"

type Captcha struct {
	Background *Background
	Content    *Content
}

func New() string {
	c := &Captcha{}
	c.Background = NewBackground()
	c.Content = NewContent()

	bg := c.Background.encodedPNG()
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(bg)
	return base64Str
}
