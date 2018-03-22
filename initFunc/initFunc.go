package initFunc

import "fmt"

// 定义多个init函数
func init() {
	fmt.Println("init 1 ...")
}

func init() {
	fmt.Println("init 2 ...")
}

func Hello() {
	fmt.Println("initFunc hello")
}
