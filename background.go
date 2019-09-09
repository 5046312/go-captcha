package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"time"
)

var circleCount = 20

type Background struct {
	Image      *image.Paletted
	Width      int
	Height     int
	Color      []color.Color // 背景颜色会从中随机取
	NoisyColor []color.Color // 噪点颜色
}

func NewBackground() *Background {
	bg := &Background{}
	bg.Width = 120
	bg.Height = 80
	bg.Color = []color.Color{color.White}
	bg.NoisyColor = []color.Color{color.Black}
	return bg
}

func (bg *Background) Build() *Background {
	bg.Image = image.NewPaletted(image.Rect(0, 0, bg.Width, bg.Height), palette.WebSafe)
	// 绘制背景
	bg.DrawPanel()
	bg.DrawNoisy()
	return bg
}

func (bg *Background) EncodedPNG() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, bg.Image); err != nil {
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

// 绘制干扰点
func (bg *Background) DrawNoisy() {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))

	dlen := 4
	// 绘制干扰斑点
	for i := 0; i < dlen; i++ {
		x := ra.Intn(bg.Width)
		y := ra.Intn(bg.Height)
		r := ra.Intn(bg.Height/20) + 1
		colorindex := ra.Intn(len(bg.NoisyColor))
		bg.DrawCircle(x, y, r, i%4 != 0, bg.NoisyColor[colorindex])
	}

	// 绘制干扰线
	for i := 0; i < dlen; i++ {
		x := ra.Intn(bg.Width)
		y := ra.Intn(bg.Height)
		o := int(math.Pow(-1, float64(i)))
		w := ra.Intn(bg.Height) * o
		h := ra.Intn(bg.Height/10) * o
		colorindex := ra.Intn(len(bg.NoisyColor))
		bg.DrawLine(x, y, x+w, y+h, bg.NoisyColor[colorindex])
		colorindex++
	}
}

// xc,yc 圆心坐标 r 半径 fill是否填充颜色
func (bg *Background) DrawCircle(xc, yc, r int, fill bool, c color.Color) {
	x, y, d := 0, r, 3-2*r
	for x <= y {
		if fill {
			for yi := x; yi <= y; yi++ {
				bg.drawCircle8(xc, yc, x, yi, c)
			}
		} else {
			bg.drawCircle8(xc, yc, x, y, c)
		}
		if d < 0 {
			d = d + 4*x + 6
		} else {
			d = d + 4*(x-y) + 10
			y--
		}
		x++
	}
}

// 画线
func (bg *Background) DrawLine(x1, y1, x2, y2 int, c color.Color) {
	dx, dy, flag := int(math.Abs(float64(x2-x1))), int(math.Abs(float64(y2-y1))), false
	if dy > dx {
		flag = true
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		dx, dy = dy, dx
	}
	ix, iy := sign(x2-x1), sign(y2-y1)
	n2dy := dy * 2
	n2dydx := (dy - dx) * 2
	d := n2dy - dx
	for x1 != x2 {
		if d < 0 {
			d += n2dy
		} else {
			y1 += iy
			d += n2dydx
		}
		if flag {
			bg.Image.Set(y1, x1, c)
		} else {
			bg.Image.Set(x1, y1, c)
		}
		x1 += ix
	}
}

func (bg *Background) drawCircle8(xc, yc, x, y int, c color.Color) {
	bg.Image.Set(xc+x, yc+y, c)
	bg.Image.Set(xc-x, yc+y, c)
	bg.Image.Set(xc+x, yc-y, c)
	bg.Image.Set(xc-x, yc-y, c)
	bg.Image.Set(xc+y, yc+x, c)
	bg.Image.Set(xc-y, yc+x, c)
	bg.Image.Set(xc+y, yc-x, c)
	bg.Image.Set(xc-y, yc-x, c)
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}
