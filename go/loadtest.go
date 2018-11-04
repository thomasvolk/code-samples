package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	POISON_PILL = -1
)

func (w Worker) get(url string) string {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err)
	}

	defer resp.Body.Close()
	return resp.Status
}

type Worker struct {
	url          string
	workerBuffer chan int
}

func (w Worker) run(i int) {
	for v := range w.workerBuffer {
		if v == POISON_PILL {
			break
		} else {
			status := w.get(w.url)
			log.Printf("Worker[%d] request[%d] open %s => %s\n", i, v, w.url, status)
		}
	}
}

func main() {
	var worker int
	var url string
	var buffer int
	var count int
	flag.IntVar(&worker, "worker", 4, "worker")
	flag.IntVar(&buffer, "buffer", 10, "buffer")
	flag.StringVar(&url, "url", "", "url")
	flag.IntVar(&count, "count", 10, "count")
	flag.Parse()

	if url == "" {
		fmt.Printf("usage %s -url <url>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	workerBuffer := make(chan int, buffer)
	defer close(workerBuffer)

	var workerWg sync.WaitGroup
	workerWg.Add(worker)

	for i := 0; i < worker; i++ {
		w := Worker{url, workerBuffer}
		go func(wo Worker, number int) {
			defer workerWg.Done()
			wo.run(number)
		}(w, i)
	}

	for i := 0; i < count; i++ {
		workerBuffer <- i
	}

	for i := 0; i < worker; i++ {
		workerBuffer <- POISON_PILL
	}

	workerWg.Wait()
}
