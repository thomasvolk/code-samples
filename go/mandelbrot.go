package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Mandelbrot struct {
	xmin       float64
	ymin       float64
	step       float64
	iterations int
	width      int
	height     int
}

func (mandelbrot *Mandelbrot) draw() *image.RGBA {
	var img = image.NewRGBA(image.Rect(0, 0, mandelbrot.width, mandelbrot.height))
	var wg sync.WaitGroup
	wg.Add(mandelbrot.width * mandelbrot.height)
	for x := 0; x < mandelbrot.width; x++ {
		for y := 0; y < mandelbrot.height; y++ {
			go func(image *image.RGBA, px int, py int) {
				defer wg.Done()
				mandelbrot.drawPoint(image, px, py)
			}(img, x, y)
		}
	}
	wg.Wait()
	return img
}

func (mandelbrot *Mandelbrot) drawPoint(img *image.RGBA, x int, y int) {
	color := color.RGBA{0x00, 0x00, 0x00, 0xff}
	point := complex(mandelbrot.xmin+float64(x)*mandelbrot.step,
		mandelbrot.ymin+float64(y)*mandelbrot.step)
	if cmplx.Abs(point) < 2 {
		nextPoint := 0 + 0i
		for i := 0; i < mandelbrot.iterations; i++ {
			nextPoint = nextPoint*nextPoint + point
			if cmplx.Abs(nextPoint) < 2 {
				color = mandelbrot.calculateColor(i)
			}
		}
	}
	img.Set(int(x), int(y), color)
}

func (mandelbrot *Mandelbrot) colorStep() int {
	return 255 / mandelbrot.iterations
}

func (mandelbrot *Mandelbrot) calculateColor(i int) color.RGBA {
	blue := 0
	green := 2 * i * mandelbrot.colorStep()
	red := 0
	if i >= mandelbrot.iterations/2 {
		blue = (i - mandelbrot.iterations/2) * mandelbrot.colorStep()
	}
	if i > mandelbrot.iterations/2 {
		green = 2 * (mandelbrot.iterations - i) * mandelbrot.colorStep()
	}
	if i <= mandelbrot.iterations/2 {
		red = 255 - 2*i*mandelbrot.colorStep()
	}
	return color.RGBA{uint8(red), uint8(green), uint8(blue), 0xff}
}

func writeFile(outputfile string, image *image.RGBA) {
	f, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, image)
}

func ForParam(r *http.Request, param string, f func(value float64)) {
	values, ok := r.URL.Query()[param]
	if ok {
		fval, err := strconv.ParseFloat(values[0], 64)
		if err == nil {
			f(fval)
		}
	}
}

func drawHandler(m Mandelbrot) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		ForParam(r, "xmin", func(value float64) { m.xmin = value })
		ForParam(r, "ymin", func(value float64) { m.ymin = value })
		ForParam(r, "step", func(value float64) { m.step = value })
		image := m.draw()
		png.Encode(w, image)
	}
}

func main() {
	var xmin float64
	var ymin float64
	var step float64
	var iterations int
	var width int
	var height int
	var outputfile string
	var serve bool
	var port int

	flag.Float64Var(&xmin, "xmin", -2, "xmin")
	flag.Float64Var(&ymin, "ymin", -2, "ymin")
	flag.Float64Var(&step, "step", 0.01, "step")
	flag.IntVar(&iterations, "iterations", 100, "iterations")
	flag.IntVar(&width, "width", 400, "width")
	flag.IntVar(&height, "height", 400, "height")
	flag.StringVar(&outputfile, "outputfile", "mandelbrot.png", "outputfile")
	flag.BoolVar(&serve, "serve", false, "start http server")
	flag.IntVar(&port, "port", 8080, "http port")

	flag.Parse()

	m := Mandelbrot{
		xmin:       xmin,
		ymin:       ymin,
		step:       step,
		iterations: iterations,
		width:      width,
		height:     height,
	}

	if serve {
		http.HandleFunc("/", drawHandler(m))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			panic(err)
		}
	} else {
		image := m.draw()
		writeFile(outputfile, image)
	}
}
