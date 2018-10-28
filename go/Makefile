GOPATH := $(PWD)

all: lambda docker

deps:
	go get "github.com/aws/aws-lambda-go/events"
	go get "github.com/aws/aws-lambda-go/lambda"

lambda: deps
	GOOS=linux go build mandelbrot_handler.go
	zip mandelbrot_handler.zip ./mandelbrot_handler

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mandelbrot mandelbrot.go 
	docker build -t mandelbrot -f Dockerfile.mandelbrot .

run-docker:
	docker run -p 9999:8080 mandelbrot
	
