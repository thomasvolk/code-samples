package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strconv"

	mandelbrot "./mandelbrot"
)

func writeFile(outputfile string, image *image.RGBA) {
	f, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, image)
}

func forParam(r *http.Request, param string, f func(value float64)) {
	values, ok := r.URL.Query()[param]
	if ok {
		fval, err := strconv.ParseFloat(values[0], 64)
		if err == nil {
			f(fval)
		}
	}
}

func drawHandler(m mandelbrot.Mandelbrot) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		forParam(r, "xmin", func(value float64) { m.Xmin = value })
		forParam(r, "ymin", func(value float64) { m.Ymin = value })
		forParam(r, "step", func(value float64) { m.Step = value })
		image := m.Draw()
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

	m := mandelbrot.Mandelbrot{
		Xmin:       xmin,
		Ymin:       ymin,
		Step:       step,
		Iterations: iterations,
		Width:      width,
		Height:     height,
	}

	if serve {
		http.HandleFunc("/", drawHandler(m))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			panic(err)
		}
	} else {
		image := m.Draw()
		writeFile(outputfile, image)
	}
}
