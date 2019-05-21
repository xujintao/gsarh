package main

// 初始化
func main() {
	// // 方式1
	// a := []int{1, 2} // a在栈上, a的底层数组也放在栈上
	// // println这里被编译成runtime.printslice, 参数单一至少不会影响到a
	// println(a)

	// // 方式2
	// a := []int{1, 2} // 调用newobject(&type.[2]int(SB)), 将底层数组放在堆上
	// // fmt.Print方法族的参数是... interface{}也就是[]interface{}
	// // a作为参数调用runtime.convTslice对a做一份copy到堆上
	// // interface{}的data域指向该副本
	// // []interface{}的底层数组在栈上
	// // 那么为什么这样布局?都放栈上不行?
	// fmt.Println(a)

	// // 方式3
	// a := make([]int, 1, 2) // a在栈上, a的底层数组也放在栈上
	// println(a)

	// // 方式4
	// a := make([]int, 1, 2) // 调用makeslice(&type.int(SB),1,2), 将底层数组放在堆上
	// // 其他与方式2一致
	// fmt.Println(a)

	// a := make([]int, 1, 2)
	// addr := (*[3]*int)(unsafe.Pointer(&a))
	// println(&a)
	// println(unsafe.Pointer(addr[0]))
	// println(addr[1])
	// println(addr[2])
}
