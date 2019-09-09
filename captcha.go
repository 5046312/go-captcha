package captcha

import (
	"encoding/base64"
	"image"
	"math"
)

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

func (c *Captcha) build() *Captcha {
	return c
}

// 绘制文字
func (c *Captcha) drawString() {

}

func (c *Captcha) Base64() string {

	bg := c.Background.EncodedPNG()
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(bg)
	return base64Str
}

func (c *Captcha) Distort(amplude float64, period float64) {
	w := c.Background.Width
	h := c.Background.Height

	oldm := c.Background.Image
	newm := image.NewPaletted(image.Rect(0, 0, w, h), oldm.Palette)

	dx := 1.4 * math.Pi / period
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			xo := amplude * math.Sin(float64(y)*dx)
			yo := amplude * math.Cos(float64(x)*dx)
			newm.SetColorIndex(x, y, oldm.ColorIndexAt(x+int(xo), y+int(yo)))
		}
	}
	c.Background.Image = newm
}
