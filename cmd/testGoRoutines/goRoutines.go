package main

import "sync"

func ProcessEvents(id int, events <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for event := range events {
		println("goroutine", id, "received", event)
	}
}
