package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mail uint8
	var lock sync.Mutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(&lock)

	send := func(id, idx int) {
		lock.Lock()
		for mail == 1 {
			sendCond.Wait()
		}
		//log.Printf("sender [%d-%d]: the mailbox is empty.", id, idx)

		mail = 1
		//log.Printf("sender [%d-%d]: the letter has been sent.", id, idx)

		lock.Unlock()
		recvCond.Broadcast()
	}

	recv := func(id, idx int) {
		lock.Lock()
		for mail == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]: the mailbox is full.",
			id, idx)

		mail = 0
		log.Printf("receiver [%d-%d]: the letter has been received.",
			id, idx)

		lock.Unlock()
		sendCond.Signal()
	}
	max := 6
	sign := make(chan struct{}, 3)

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(500 * time.Millisecond)
			send(id, i)
		}
	}(0, max)

	recvFunc := func(id, idx int) {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= idx; j++ {
			time.Sleep(200 * time.Millisecond)
			recv(id, j)
		}
	}

	go recvFunc(1, max/2)
	go recvFunc(2, max/2)

	<-sign
	<-sign
	<-sign

}
