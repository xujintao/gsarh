package main

import "runtime"

// 不逃逸，hmap的48个字节放在栈上作为临时变量，连hamp.buckets也放在栈上
// 调用runtime.makemap
// func main() {
// 	m := make(map[string]int) // buckets域直接放栈上了
// 	// m := make(map[string]int, 10) //会调用runtime.makemap
// 	m["abc"] = 2
// 	println(m)
// }

// 逃逸
// 调用runtime.makemap
// func foo() map[string]int {
// 	m := make(map[string]int, 2)
// 	m["abc"] = 2
// 	return m
// }

// 逃逸，
// 字面量赋值效果和make调用一致
// func foo() map[string]int {
// 	m := map[string]int{
// 		"abc": 1,
// 		"edf": 2,
// 	}
// 	return m
// }

// func main() {
// 	m := foo()
// 	println(m)
// }

// -------------------------------------------------
// 逃逸
// 调用runtime.makeslice初始化底层数组
// func foo() []int {
// 	s := make([]int, 0, 1)
// 	return s
// }

// 逃逸
// new切片，效果和nil切片一样
// func foo() *[]int {
// 	ps := new([]int)
// 	*ps = append(*ps, 2)
// 	return ps
// }

// 逃逸
// nil切片
// func foo() []int {
// 	var s []int
// 	s = append(s, 2)
// 	return s
// }

// 逃逸
// 空切片，调用runtime.makeslice初始化底层数组，只是cap为0
// func foo() []int {
// 	s := make([]int, 0)
// 	s = append(s, 2)
// 	return s
// }

// func main() {
// 	s := foo()
// 	println(s)
// }

// 总结
// 1. new只是把内存new出来了，但是并没有初始化
// 2. 一般而言切片不make，直接用

// ----------------------------------------------------
// 测试结构类型的new和字面量初始化
// type person struct {
// 	name string
// 	age  int
// }

// 情况一，栈
// func main() {
// 	p := person{
// 		name: "xujintao",
// 		age:  30,
// 	}
// 	println(&p)
// }

// 情况二，栈，和情况一保持一致
// 这种情况很有迷惑性
// func main() {
// 	p := new(person)
// 	p.name = "xujintao"
// 	p.age = 30
// 	println(p)
// }

// 情况三，堆
// func foo() *person {
// 	p := &person{
// 		name: "xujintao",
// 		age:  30,
// 	}
// 	return p
// }

// 情况四，堆，和情况三保持一致
// 这种写法很有迷惑性，这种写法在c++里面算错误
// func foo() *person {
// 	p := person{
// 		name: "xujintao",
// 		age:  30,
// 	}
// 	return &p
// }

// func foo() *person {
// 	var p person
// 	p.name = "xujintao"
// 	p.age = 30
// 	return &p
// }

// 情况五，堆，和情况三保持一致
// func foo() *person {
// 	p := new(person)
// 	p.name = "xujintao"
// 	p.age = 30
// 	return p
// }

// func main() {
// 	p := foo()
// 	println(p)
// }

// 结论
// 当gc想把变量分配在栈上，不管什么形式的初始化都阻止不了它。
// 当gc想把变量逃逸到堆上，十头牛都拉不回来。
// 这一点很不好，gc太智能了。就是明摆着让用户别管堆和栈了。
// 既然用户不管堆栈了，那new关键字干脆去掉好了，大家都用字面量来初始化。

// ----------------------------------------------------
// 不逃逸，因为子协程只读i
func main() {
	i := 1
	go func() {
		println(i)
	}()
	runtime.Gosched()
	println(i)
}

// 逃逸，因为子协程会写i
// func main() {
// 	i := 1 // i := new(int)
// 	go func() {
// 		i = 2      // *i = 2
// 		println(i) // println(*i)
// 	}()
// 	runtime.Gosched()
// 	println(i) // println(*i)
// }
