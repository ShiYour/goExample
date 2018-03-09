package pointer

import "testing"

type ptrA struct {
}

func Test_should_return_pointer_type_when_create_instance(t *testing.T) {
	//没有*interface写法
	//var a *interface{} = &ptrA{}
	//map1 := map[string]*interface{}{
	//	"1": a,
	//}
}
