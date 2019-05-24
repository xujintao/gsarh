package main

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

// func main() {
// 	println(reflect.TypeOf(true).String(),
// 		reflect.TypeOf(true).Kind().String()) // bool bool
// 	println(reflect.TypeOf(uint8(1)).String(),
// 		reflect.TypeOf(uint8(1)).Kind().String()) // uint8 uint8
// 	println(reflect.TypeOf("abc").String(),
// 		reflect.TypeOf("abc").Kind().String()) // string string
// }

// func main() {
// 	println(reflect.TypeOf((*bool)(nil)).String(),
// 		reflect.TypeOf((*bool)(nil)).Kind().String()) // *bool ptr
// }

// func main() {
// 	println(reflect.TypeOf([2]int{}).String(),
// 		reflect.TypeOf([2]int{}).Kind().String()) // [2]int array
// 	println(reflect.TypeOf([]int{}).String(),
// 		reflect.TypeOf([]int{}).Kind().String()) // []int slice
// 	println(reflect.TypeOf(&[]int{}).String(),
// 		reflect.TypeOf(&[]int{}).Kind().String()) // *[]int ptr
// 	println(reflect.TypeOf(map[string]int{}).String(),
// 		reflect.TypeOf(map[string]int{}).Kind().String()) // map[string]int map
// 	println(reflect.TypeOf((*chan int)(nil)).Elem().String(),
// 		reflect.TypeOf((*chan int)(nil)).Elem().Kind().String()) //chan int chan
// 	println(reflect.TypeOf(func(int) {}).String(),
// 		reflect.TypeOf(func(int) {}).Kind().String()) // func(int) func
// 	println(reflect.TypeOf(struct{ int }{}).String(),
// 		reflect.TypeOf(struct{ int }{}).Kind().String()) // struct {} struct
// 	println(reflect.TypeOf((*interface{ foo() })(nil)).Elem().String(),
// 		reflect.TypeOf((*interface{ foo() })(nil)).Elem().Kind().String())
// 	// interface { main.foo() } interface
// }

// func main() {
// 	println(reflect.TypeOf(&[]int{}).String(),
// 		reflect.TypeOf(&[]int{}).Kind().String()) // *[]int ptr
// }

// type T1 int32              // main.T1
// type T2 struct{ a int }    // main.T2
// type T3 interface{ foo() } // main.T3
// func main() {
// 	println(reflect.TypeOf(T1(1)).String(),
// 		reflect.TypeOf(T1(1)).Kind().String()) // main.T1 int32
// 	println(reflect.TypeOf(T2{}).String(),
// 		reflect.TypeOf(T2{}).Kind().String()) // main.T2 struct
// 	println(reflect.TypeOf(&T2{1}).String(),
// 		reflect.TypeOf(&T2{1}).Kind().String()) // *main.T2 ptr
// 	println(reflect.TypeOf((*T3)(nil)).Elem().String(),
// 		reflect.TypeOf((*T3)(nil)).Elem().Kind().String()) // main.T3 interface
// 	// println(reflect.TypeOf(t3).String()) // panic: runtime error: invalid memory address or nil pointer dereference
// }

// type T1 int32              // main.T1
// type T2 struct{ a int }    // main.T2
// type T3 interface{ foo() } // main.T3
// func main() {
// 	// println(reflect.TypeOf(&T2{1}).String(),
// 	// 	reflect.TypeOf(&T2{1}).Kind().String()) // *main.T2 ptr
// 	println(reflect.PtrTo(reflect.TypeOf(T2{1})).String(),
// 		reflect.PtrTo(reflect.TypeOf(T2{1})).Kind().String()) // *main.T2 ptr
// }

// -----
type _type struct {
	Size       uintptr `json:"size"`
	Ptrdata    uintptr `json:"ptrdata"`
	Hash       uint32  `json:"hash"`
	Tflag      uint8   `json:"tflag"`
	Align      uint8   `json:"align"`
	Fieldalign uint8   `json:"fieldalign"`
	Kind       uint8   `json:"kind"`
	Alg        uintptr `json:"alg"`
	Gcdata     *byte   `json:"gcdata"`
	Str        int32   `json:"str"`
	PtrToThis  int32   `json:"ptrToThis"`
}

func printType(t reflect.Type) {
	at := (*[2]unsafe.Pointer)(unsafe.Pointer(&t))
	res, _ := json.Marshal((*_type)(at[1]))
	println(string(res))
}

type T1 int                // main.T1
type T2 struct{ a int }    // main.T2
type T3 interface{ foo() } // main.T3
func main() {
	printType(reflect.TypeOf("abc"))
	printType(reflect.PtrTo(reflect.TypeOf("abc")))
	printType(reflect.TypeOf((*chan int)(nil)).Elem())
	printType(reflect.TypeOf((*chan int)(nil)))
	printType(reflect.TypeOf((*interface{ foo() })(nil)).Elem())
	printType(reflect.TypeOf((*interface{ foo() })(nil)))
	printType(reflect.TypeOf(T1(1)))
	printType(reflect.PtrTo(reflect.TypeOf(T1(1))))
	printType(reflect.TypeOf(T2{}))
	printType(reflect.PtrTo(reflect.TypeOf(T2{})))
}
