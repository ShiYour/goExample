package main

import (
	"fmt"
	_ "goExample/initFunc"
)

type aa struct {
	R []*bb
	S map[string]string
	T mockInf
}

type mockInf interface {
	String1() string
}

type bb struct {
	Age int
}

func (*bb) String1() string {
	return "ddsa"
}

func main() {
	fmt.Println("main...")

	fmt.Println((3+63)/64)
}
