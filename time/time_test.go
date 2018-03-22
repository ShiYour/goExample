package time

import (
	"fmt"
	"testing"
	"time"
)

func Test_time_ticker(t *testing.T) {

	tickChan := time.NewTicker(2 * time.Second).C
	for {
		select {
		case <-tickChan:
			fmt.Println("time reached ...")
		}
	}

}

func Test_time_timer(t *testing.T) {

	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("time reached ...")
			return
		}
	}

}

func Test_time_duration(t *testing.T) {

	fmt.Println(time.Duration(6) * time.Second)

}
