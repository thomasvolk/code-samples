GOPATH := $(PWD)

all: lambda docker

deps:
	go get "github.com/aws/aws-lambda-go/events"
	go get "github.com/aws/aws-lambda-go/lambda"

lambda: deps
	GOOS=linux go build mandelbrot_handler.go
	zip mandelbrot_handler.zip ./mandelbrot_handler

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver webserver.go 
	docker build -t go-webserver -f Dockerfile.webserver .

run-docker:
	docker run -p 9999:8080 go-webserver
	
