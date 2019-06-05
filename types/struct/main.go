package main

func main() {
	// // struct{}类型落在栈上
	// // a被优化掉了
	// a := struct{}{}
	// println(unsafe.Sizeof(a)) // MOVQ	$0, (SP)

	// // struct{}类型落在堆上
	// // a就是basezero，由runtime保证
	// a := func() *struct{} {
	// 	return &struct{}{}
	// }()
	// println(unsafe.Sizeof(a))

	// struct{}类型落在栈上，但是对它取地址
	// 编译器优化前，a是SP+15
	// 编译器优化后，a是BP
	// 反正编译器就是拐弯抹角的想把a忽略掉
	a := struct{}{}
	println(&a)
}
