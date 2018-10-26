package main

import (
	"net/http"
)

var stopChannel = make(chan int)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", health)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
