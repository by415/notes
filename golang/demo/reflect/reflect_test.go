package reflect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Study interface {
	Study()
}

type Student struct {
	Name string "名字"
	Age  int    `a:"111"b:"3333"`
}

type Humen struct {
	Name string "名字"
	Age  int    `a:"111"b:"3333"`
}

func (h Humen) Study() {
	panic("implement me")
}

func (s Student) Study() {
	fmt.Println("student-study")
}

func TestReflect(t *testing.T) {
	s := Student{"name---", 33}
	rt := reflect.TypeOf(s)
	fmt.Println("rt.Name() = ", rt.Name())
	fmt.Println("rt.PkgPath() = ", rt.PkgPath())
	fmt.Println("rt.NumField() = ", rt.NumField())

	for i := 0; i < rt.NumField(); i++ {
		fmt.Println("Field = ", rt.Field(i).Name)
	}

	fieldName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println(fieldName.Tag)
		fmt.Println(fieldName.Name)
		fmt.Println(fieldName.PkgPath)
		fmt.Println(fieldName.Type)
	}

	x, _ := json.Marshal(fieldName)
	fmt.Println(string(x))

	fieldAge, ok := rt.FieldByName("Age")
	if ok {
		fmt.Println(fieldAge.Tag)
	}

	x, _ = json.Marshal(fieldAge)
	fmt.Println(string(x))

}
func TypeOf(student Study) {
	fmt.Println(reflect.TypeOf(student))
}

// 实例到Type TypeOf
func TestTypeOf(t *testing.T) {
	TypeOf(Student{})
	TypeOf(Humen{})
}

// 实例到value ValueOf
func TestValueOf(t *testing.T) {
	s := Student{"sfa", 123}
	fmt.Println(reflect.ValueOf(s))
}

// Value到Type  Type
func TestType(t *testing.T) {
	s := Student{"sfa", 123}
	fmt.Println(reflect.ValueOf(s).Type())
}

// Type到Value  New, Zero, NewAt
func TestNew(t *testing.T) {
	s := Student{"sfa", 123}
	fmt.Println(reflect.New(reflect.ValueOf(s).Type()))
	fmt.Println(reflect.Zero(reflect.ValueOf(s).Type()))
	fmt.Println(reflect.NewAt(reflect.ValueOf(s).Type(), (unsafe.Pointer(&s))))
}

//value 到实例
func TestValueToInstance(t *testing.T) {
	s := Student{"sfa", 123}
	fmt.Println(reflect.ValueOf(s).Interface())
	fmt.Println(reflect.TypeOf(reflect.ValueOf(s).Interface()))

	var c int = 333

	fmt.Println(reflect.ValueOf(c).Interface())
	fmt.Println(reflect.TypeOf(reflect.ValueOf(c).Interface()))
}

//point 到value
func TestPointToValue(t *testing.T) {
	s := &Student{"sfa", 123}
	fmt.Println(reflect.ValueOf(s).Elem())
	//fmt.Println(reflect.ValueOf(*s).Elem())		// panic
	fmt.Println(reflect.TypeOf(reflect.ValueOf(s).Interface()))
}

// type 指针与值的相互转化
//必须是 容器类型
func TestPointValue(t *testing.T) {
	/// value to point
	s := []int{3, 4, 5}
	fmt.Println(reflect.PtrTo(reflect.TypeOf(s)))

	///  point  to value
	fmt.Println(reflect.TypeOf(&s).Elem())
	///  point  to value
	fmt.Println(reflect.TypeOf(s).Elem())
}
