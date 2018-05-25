package nil

import (
	"fmt"
	"reflect"
	"testing"
)

type Mock interface {
	Hello()
}

type Mock2 interface {
	Hello2()
}

type impl struct {
	Name string
	h    Mock2
}

func (i *impl) Hello() {
	fmt.Println(i.Name, " say hello...")
}

//声明接口类型变量，不赋值，== nil
func Test_should_return_nil_when_empty_inf_value(t *testing.T) {
	var tt Mock

	fmt.Println(reflect.TypeOf(tt))
	fmt.Println(tt == nil)
}

//把具体实现的控制赋值给接口变量，接口变量会存储类型信息
func Test_should_return_nil_when_empty_struct_value(t *testing.T) {
	var tt Mock
	var ll *impl
	tt = ll

	fmt.Println(tt)
	fmt.Println(reflect.TypeOf(tt))
	fmt.Println(tt == nil)

	var ll2 = &impl{}
	fmt.Println(ll2.h == nil)

	//转换为实际的需要的类型， 判断是否为nil
	a, _ := tt.(*impl)
	fmt.Println(a == nil)
}

//声明struct类型变量，不赋值，不能与nil进行比较， 指针变量 == nil
func Test_cannot_compare_nil_when_struct_value(t *testing.T) {
	//cannot
	//var tt1 impl
	//fmt.Println(tt1 == nil)

	var tt *impl
	fmt.Println(tt == nil)
}
