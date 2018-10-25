package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

const (
	POISON_PILL = -1
)

func now() int64 {
	return time.Now().UTC().UnixNano()
}

var START_TIME = now()

func tick() int64 {
	return now() - START_TIME
}

func workerRoutine(wg *sync.WaitGroup, i int, timeout time.Duration, c chan int) {
	defer wg.Done()
	for v := range c {
		if v == POISON_PILL {
			break
		} else {
			fmt.Printf("[GR %d] ; %012d ; receive ; %d\n", i, tick(), v)
			time.Sleep(timeout)
		}
	}
}

func main() {
	var worker int
	var timeout time.Duration
	var messages int
	var buffer int
	flag.IntVar(&worker, "worker", 4, "worker ")
	flag.IntVar(&buffer, "buffer", 3, "buffer ")
	flag.IntVar(&messages, "messages", 10, "messages ")
	flag.DurationVar(&timeout, "timeout", 1*time.Second, "timeout")
	flag.Parse()

	c := make(chan int, buffer)
	defer close(c)

	var workerWg sync.WaitGroup
	workerWg.Add(worker)

	for i := 0; i < worker; i++ {
		go workerRoutine(&workerWg, i, timeout, c)
	}

	for i := 0; i < messages; i++ {
		fmt.Printf("[MAIN] ; %012d ; send ; %d\n", tick(), i)
		c <- i
	}

	for i := 0; i < worker; i++ {
		c <- POISON_PILL
	}

	workerWg.Wait()

}
