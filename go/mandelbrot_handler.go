package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"
	"strconv"

	mandelbrot "./lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func forQueryParam(r *events.APIGatewayProxyRequest, param string, f func(value float64)) {
	value, ok := r.QueryStringParameters[param]
	if ok {
		fval, err := strconv.ParseFloat(value, 64)
		if err == nil {
			f(fval)
		}
	}
}

func HandleRequest(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m := mandelbrot.Mandelbrot{
		Xstart:     -1.6,
		Xend:       -0.2,
		Ystart:     -1.2,
		Yend:       1.2,
		Iterations: 100,
		Width:      400,
		Height:     400,
	}
	forQueryParam(&r, "xstart", func(value float64) { m.Xstart = value })
	forQueryParam(&r, "xend", func(value float64) { m.Xend = value })
	forQueryParam(&r, "ystart", func(value float64) { m.Ystart = value })
	forQueryParam(&r, "yend", func(value float64) { m.Yend = value })
	forQueryParam(&r, "iterations", func(value float64) { m.Iterations = int(value) })
	image := m.Draw()
	buf := new(bytes.Buffer)
	png.Encode(buf, image)
	headers := make(map[string]string)
	headers["Content-Type"] = "text/html"
	body := fmt.Sprintf(`<!DOCTYPE html>
	<html>
  		<head>
    		<title>Mandelbrot</title>
  		</head>
		<body>
		  <img src="data:image/png;base64, %s" alt="Mandelbrot" />
		</body>
	</html>`, base64.StdEncoding.EncodeToString(buf.Bytes()))
	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            body,
		Headers:         headers,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
