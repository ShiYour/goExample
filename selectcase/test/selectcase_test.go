package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_should_return_success_when_readChan(t *testing.T) {
	ch := make(chan uint64)
	go choose(ch)

	timer := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-timer.C:
			fmt.Println("time out")
		case n := <-ch:
			fmt.Println("n: ", n)
		}
	}
}

func choose(ch chan uint64) {
	var i uint64
	for {
		time.Sleep(time.Second * 1)
		ch <- i
		close(ch)
		time.Sleep(time.Second * 4)
		i++
	}
}
