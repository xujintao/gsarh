package main

// var (
// 	a, b int
// )

// func main() {
// 	a = b + 1
// 	b = 0
// 	println(a, b)
// }

// 闭包
// func foo() func() int {
// 	var a int // 情况2，会把a以参数的形式传递给子协程
// 	a += 2    // ADDQ	$2, (AX)
// 	return func() int {
// 		a++ // INCQ	(AX)
// 		println(a)
// 		return 1
// 	}
// }

// func main(){
// 	f := foo()
// 	f()
// }
// 结论，foo的返回值其实是个结构体指针

// var a int // 情况1

// func main() {
// 	var a int // 情况2，会把a以参数的形式传递给子协程
// 	a += 2    // ADDQ	$2, (AX)
// 	go func() {
// 		a++ // INCQ	(AX)
// 		println(a)
// 	}()
// 	// runtime.Gosched()
// 	println(a)
// }

// 结论，现在累加语句都是编译成单个指令

//go:nosplit
//go:noinline
func Load64(ptr *uint64) uint64 {
	return *ptr
}

func Store64(ptr *uint64, val uint64) {
	*ptr = val
}

var a uint64 // 情况1
func main() {
	b := Load64(&a)
	assert(b == 0)
}
