package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan uint64)
	go choose(ch)

	var i uint64
	for {
		ch <- i
		i++
		time.Sleep(time.Second * 10)
	}
}

func choose(ch chan uint64) {
	for {
		select {
		case n := <-ch:
			fmt.Println(n)

		}
	}
}
