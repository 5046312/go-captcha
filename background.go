package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
)

var circleCount = 20

type Background struct {
	Width  int
	Height int
}

func NewBackground() *Background {
	return &Background{}
}

func (bg *Background) Build() *image.Paletted {
	bg.Width = 220
	bg.Height = 120
	img := image.NewPaletted(image.Rect(0, 0, bg.Width, bg.Height), palette.WebSafe)
	for x := 0; x < bg.Width; x++ {
		for y := 0; y < bg.Height; y++ {
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 255, 255})
		}
	}
	return img
}

func (bg *Background) encodedPNG() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, bg.Build()); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}
