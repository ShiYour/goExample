package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

//通过反射调用函数
func Test_should_return_success_when_invoke_func(t *testing.T) {

	m := make(map[string]interface{})
	m["add"] = add

	newFunc := m["add"]
	t1 := reflect.TypeOf(newFunc)
	fmt.Println(t1)
	fmt.Println(t1.Kind())
	fmt.Println(t1.PkgPath())

	t2 := reflect.ValueOf(newFunc)
	fmt.Println(t2)
	fmt.Println(t2.Kind())
	rs := t2.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(5)})
	fmt.Println(rs[0].Interface().(int))
}

func add(x, y int) int {
	return x + y
}

type myString string

type mockObjA struct {
	Name string
	Age  int
}

type mockObjC struct {
	AA  mockObjA
	AA2 *mockObjA
}

//通过发射设置字段的值
func Test_should_return_success_when_set_field_value(t *testing.T) {
	objA := &mockObjA{Name: "11", Age: 10}
	v := reflect.ValueOf(objA).Elem()
	fmt.Println("field num: ", v.NumField())
	//字段名区分大小写，必须一样才行
	name := v.FieldByName("Name")
	fmt.Println("name valid: ", name.IsValid())
	fmt.Println("name before set: ", name)

	fmt.Println("name type: ", reflect.TypeOf(name.Interface()))
	name.Set(reflect.ValueOf("12"))
	fmt.Println("name after set: ", name)

	fmt.Println("name type == reflect.string", reflect.TypeOf(name.Interface()).Kind() == reflect.String)

	//设置另一个结构体
	objC := &mockObjC{}
	//查看指针类型的默认值是啥
	fmt.Println("ptr type def v: ", objC.AA2, objC.AA2 == nil)
	v2 := reflect.ValueOf(objC).Elem()
	fmt.Println("mockObjC field num: ", v2.NumField())
	nameAA := v2.FieldByName("AA")
	fmt.Println("AA name valid: ", nameAA.IsValid())
	fmt.Println("AA name before set: ", nameAA.Interface().(mockObjA).Name)

	fmt.Println("AA name type: ", reflect.TypeOf(nameAA.Interface()))
	nameAA.Set(reflect.ValueOf(objA).Elem())
	fmt.Println("AA name after set: ", nameAA.Interface().(mockObjA).Name)
	fmt.Println("objC: ", objC.AA.Name, objC.AA.Age)

	//结构体类型的通过函数设置
	_objA1 := &mockObjA{Name: "world111", Age: 200}
	setFieldValue(nameAA, _objA1)
	fmt.Println("objC_objA1: ", objC.AA.Name, objC.AA.Age)

	//指针类型的也可以设置
	nameAA2 := v2.FieldByName("AA2")
	fmt.Println("nameAA2 can set: ", nameAA2.CanSet())
	nameAA2.Set(reflect.ValueOf(objA))

	fmt.Println("AA2 name after set: ", objC.AA.Name)

	//指针类型通过函数后是否可以设置
	_objA := &mockObjA{Name: "world", Age: 20}
	setFieldValue(nameAA2, _objA)
	fmt.Println("objC_objA: ", objC.AA2.Name, objC.AA2.Age)

}

func setFieldValue(field reflect.Value, obj interface{}) {
	if reflect.TypeOf(obj).Kind().String() == "ptr" && reflect.TypeOf(field.Interface()).Kind().String() != "ptr" {
		field.Set(reflect.ValueOf(obj).Elem())
	} else {
		field.Set(reflect.ValueOf(obj))
	}
}

//非指针类型不可以反射设置字段
//reflect.Value.Set using unaddressable value
func Test_should_return_success_when_set_field_value2(t *testing.T) {
	var objA interface{}
	objA = &mockObjA{Name: "11", Age: 10}
	//v := reflect.ValueOf(objA)
	//如下才可以
	v := reflect.ValueOf(objA).Elem()
	fmt.Println("objA type: ", reflect.TypeOf(v.Interface()))

	fmt.Println("v kind: ", reflect.TypeOf(&objA).Kind())
	fmt.Println("field num: ", v.NumField())
	//字段名区分大小写，必须一样才行
	name := v.FieldByName("Name")
	fmt.Println("name valid: ", name.IsValid())
	fmt.Println("name before set: ", name)

	fmt.Println("name type: ", reflect.TypeOf(name.Interface()))
	name.Set(reflect.ValueOf("12"))
	fmt.Println("name after set: ", name)
}

type mockObjB struct {
	Name  string
	Age   int
	Sex   bool
	Age2  uint
	Age3  int8
	Age4  uint8
	Age5  int16
	Age6  uint16
	Age7  int32
	Age8  uint32
	Age9  int64
	Age10 uint64
	Age11 float32
	Age12 float64
}

func Test_should_return_success_when_string2basictype(t *testing.T) {
	objB := mockObjB{}
	bRef := reflect.ValueOf(&objB).Elem()
	oneField := bRef.FieldByName("Age2")
	fmt.Println(typeOf(oneField.Interface()))
}

func typeOf(v interface{}) string {
	switch t := v.(type) {
	case string:
		return "string"
	case uint:
		return "uint"
	case int:
		return "int"
	case float64:
		return "float64"
		//... etc
	default:
		_ = t
		return "unknown"
	}
}

type mockInf interface {
	String() string
}

func (*mockObjB) String() string {
	return "ObjB"
}

type mockObjD struct {
	ObjB mockInf
}

//获取接口的类型
func Test_should_return_inf_type(t *testing.T) {
	d := mockObjD{}
	ty := reflect.TypeOf(d)
	fmt.Println(ty.Field(0).Name)
	field, _ := ty.FieldByName("ObjB")
	fmt.Println(field.Type)

	v := reflect.ValueOf(&d).Elem()
	fmt.Println(v.FieldByName("ObjB"))
	fieldv := v.FieldByName("ObjB")
	fmt.Println(fieldv.CanSet())
	fmt.Println(fieldv.Kind())

	fieldv.Set(reflect.ValueOf(&mockObjB{Age: 10}))

	fmt.Println(d.ObjB.String())
}
