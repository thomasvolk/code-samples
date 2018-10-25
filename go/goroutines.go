package main

import (
	"bytes"
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

type Worker struct {
	waittime       time.Duration
	workerBuffer   chan int
	receiveHistory chan int64
}

func (w Worker) run(i int) {
	for v := range w.workerBuffer {
		if v == POISON_PILL {
			break
		} else {
			t := tick()
			w.receiveHistory <- t
			time.Sleep(w.waittime)
		}
	}
}

type History struct {
	timeframe      time.Duration
	sendHistory    chan int64
	receiveHistory chan int64
	messagesTotal  int
	buffer         chan int
}

type StringBuffer struct {
	buffer bytes.Buffer
}

func (sb *StringBuffer) String() string {
	return sb.buffer.String()
}

func (sb *StringBuffer) MultiAdd(item string, l int) {
	for i := 0; i < l; i++ {
		sb.Add(item)
	}
}

func (sb *StringBuffer) Add(item string) {
	sb.buffer.WriteString(item)
}

func (h History) run() {
	receivedTotal := 0
	send := 0
	received := 0
	active := true
	for active {
		select {
		case <-h.sendHistory:
			send++
		case <-h.receiveHistory:
			received++
			receivedTotal++
		case <-time.After(h.timeframe):
			var b StringBuffer
			b.Add("|")
			b.MultiAdd(".", len(h.buffer))
			b.MultiAdd("S", send)
			b.MultiAdd("R", received)
			fmt.Println(b.String())
			send = 0
			received = 0
			if receivedTotal >= h.messagesTotal {
				active = false
			}
		}
	}
}

func main() {
	var worker int
	var waittime time.Duration
	var timeframe time.Duration
	var messages int
	var buffer int
	flag.IntVar(&worker, "worker", 4, "worker ")
	flag.IntVar(&buffer, "buffer", 3, "buffer ")
	flag.IntVar(&messages, "messages", 10, "messages ")
	flag.DurationVar(&waittime, "waittime", 1*time.Second, "waittime")
	flag.DurationVar(&timeframe, "timeframe", 100*time.Millisecond, "timeframe")
	flag.Parse()

	workerBuffer := make(chan int, buffer)
	defer close(workerBuffer)
	sendHistory := make(chan int64, messages)
	defer close(sendHistory)
	receiveHistory := make(chan int64, messages)
	defer close(receiveHistory)

	var workerWg sync.WaitGroup
	workerWg.Add(worker)

	var historyWg sync.WaitGroup
	historyWg.Add(1)
	go func(h History) {
		defer historyWg.Done()
		h.run()
	}(History{timeframe, sendHistory, receiveHistory, messages, workerBuffer})

	for i := 0; i < worker; i++ {
		w := Worker{waittime, workerBuffer, receiveHistory}
		go func(w Worker) {
			defer workerWg.Done()
			w.run(i)
		}(w)
	}

	for i := 0; i < messages; i++ {
		t := tick()
		sendHistory <- t
		workerBuffer <- i
	}

	for i := 0; i < worker; i++ {
		workerBuffer <- POISON_PILL
	}

	workerWg.Wait()
	historyWg.Wait()
}
