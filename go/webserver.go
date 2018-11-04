package main

import (
	"flag"
	"fmt"
	"net/http"
)

var stopChannel = make(chan int)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "http port")
	flag.Parse()

	fmt.Printf("start server listen to port %d\n", port)
	http.HandleFunc("/", health)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
