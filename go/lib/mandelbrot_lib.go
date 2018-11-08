package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

type Mandelbrot struct {
	Xstart     float64
	Xend       float64
	Ystart     float64
	Yend       float64
	Iterations int
	Width      int
	Height     int
}

func (mandelbrot *Mandelbrot) Draw() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, mandelbrot.Width, mandelbrot.Height))
	var wg sync.WaitGroup
	xStep := (mandelbrot.Xend - mandelbrot.Xstart) / float64(mandelbrot.Width)
	yStep := (mandelbrot.Yend - mandelbrot.Ystart) / float64(mandelbrot.Height)
	wg.Add(mandelbrot.Width * mandelbrot.Height)
	for x := 0; x < mandelbrot.Width; x++ {
		for y := 0; y < mandelbrot.Height; y++ {
			go func(image *image.RGBA, px int, py int) {
				defer wg.Done()
				mandelbrot.drawPoint(image, px, py, xStep, yStep)
			}(img, x, y)
		}
	}
	wg.Wait()
	return img
}

func (mandelbrot *Mandelbrot) drawPoint(img *image.RGBA, x int, y int, xStep float64, yStep float64) {
	color := color.RGBA{0x00, 0x00, 0x00, 0xff}
	point := complex((float64(x)*xStep)+mandelbrot.Xstart,
		(float64(y)*yStep)+mandelbrot.Ystart)
	if cmplx.Abs(point) < 2 {
		nextPoint := 0 + 0i
		for i := 0; i < mandelbrot.Iterations; i++ {
			nextPoint = nextPoint*nextPoint + point
			if cmplx.Abs(nextPoint) >= 2 {
				break
			}
			color = mandelbrot.calculateColor(i)
		}
	}
	img.Set(int(x), int(y), color)
}

func (mandelbrot *Mandelbrot) colorStep() float64 {
	return 255.0 / float64(mandelbrot.Iterations)
}

func (mandelbrot *Mandelbrot) calculateColor(i int) color.RGBA {
	blue := 0.0
	green := 2.0 * float64(i) * mandelbrot.colorStep()
	red := 0.0
	if i >= mandelbrot.Iterations/2 {
		blue = (float64(i) - float64(mandelbrot.Iterations)/2.0) * mandelbrot.colorStep()
	}
	if i > mandelbrot.Iterations/2 {
		green = float64(2*(mandelbrot.Iterations-i)) * mandelbrot.colorStep()
	}
	if i <= mandelbrot.Iterations/2 {
		red = 255.0 - 2.0*float64(i)*mandelbrot.colorStep()
	}
	return color.RGBA{uint8(red), uint8(green), uint8(blue), 0xff}
}
