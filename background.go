package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/png"
	"math/rand"
	"time"
)

var circleCount = 20

type Background struct {
	Image  *image.Paletted
	Width  int
	Height int
	Color  []color.Color // 背景颜色会从中随机取
}

func NewBackground() *Background {
	bg := &Background{}
	bg.Width = 120
	bg.Height = 80
	bg.Color = []color.Color{color.White}
	return bg
}

func (bg *Background) Build() *image.Paletted {
	bg.Image = image.NewPaletted(image.Rect(0, 0, bg.Width, bg.Height), palette.WebSafe)
	// 绘制背景
	bg.DrawPanel()
	return bg.Image
}

func (bg *Background) encodedPNG() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, bg.Build()); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

// 绘制背景
func (bg *Background) DrawPanel() {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	//填充主背景色
	index := ra.Intn(len(bg.Color))
	bkg := image.NewUniform(bg.Color[index])
	draw.Draw(bg.Image, bg.Image.Bounds(), bkg, image.ZP, draw.Over)
}
