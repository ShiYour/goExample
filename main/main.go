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

	//m2 := map[string]string{
	//	"a": "a",
	//	"b": "b",
	//}
	//
	//aaObj2 := &aa{}
	//aaObj2.R = append(aaObj2.R, &bb{Age: 1})
	//aaObj2.S = m2
	//aaObj2.T = &bb{Age: 12}
	//
	//a, _ := json.Marshal(aaObj2)
	//fmt.Println(aaObj2)
	//fmt.Println(a)
	//
	//aaObj := &aa{}
	//err := json.Unmarshal(a, aaObj)
	//fmt.Println(err)
	//fmt.Println(aaObj)
}
