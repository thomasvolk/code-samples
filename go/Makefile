
all:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver webserver.go 
	docker build -t go-webserver -f Dockerfile.webserver .
	docker run -p 9999:8080 go-webserver
	
