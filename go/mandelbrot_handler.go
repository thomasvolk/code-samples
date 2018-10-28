package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"

	mandelbrot "./lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m := mandelbrot.Mandelbrot{
		Xmin:       -2.0,
		Ymin:       -2.0,
		Step:       0.01,
		Iterations: 100,
		Width:      400,
		Height:     400,
	}
	image := m.Draw()
	buf := new(bytes.Buffer)
	png.Encode(buf, image)
	headers := make(map[string]string)
	headers["Content-Type"] = "image/png"
	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            base64.StdEncoding.EncodeToString(buf.Bytes()),
		Headers:         headers,
		IsBase64Encoded: true,
	}, nil
}

/*
API-GAteway Binary Support:
  Under the selected API in the primary navigation panel, choose Settings.
  In the Settings pane, choose Add Binary Media Type in the Binary Media Types section.

  image/png

Test:
  curl -H "Accept: image/png" https://XXXX.execute-api.us-east-1.amazonaws.com/STAGE/ -o test.png
*/
func main() {
	lambda.Start(HandleRequest)
}
