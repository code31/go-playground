package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mut = &sync.Mutex{}


type semaphore struct {
	c chan struct{}
	messages chan string
}

func NewSemaphore(size int) *semaphore {
	return &semaphore{
		c: make(chan struct{}, size),
		messages: make(chan string, 100),
	}
}

func (s *semaphore) start() {
	s.c <- struct{}{}
}

func (s *semaphore) done() {
	<- s.c
}

func worker(handlerId int, workerId int, s *semaphore, wg *sync.WaitGroup) {
	defer s.done()
	defer wg.Done()
	mut.Lock()
	counter++
	mut.Unlock()
	time.Sleep(time.Second * 10)
	s.messages <- fmt.Sprintf("executing worker: handlerId: %d, workerId: %d, channelLength: %d\n", handlerId, workerId, len(s.c))
}

func handler(id int, s *semaphore) {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		s.start()
		go worker(id, i, s, &wg)
	}

	wg.Wait()
}

func main() {
	s := NewSemaphore(10)

	for h := 0; h < 10; h++ {
		go handler(h, s)
	}

	for m := range s.messages {
		fmt.Println(m)
	}

	fmt.Println("Total Requests:", counter)

}
