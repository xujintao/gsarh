package main

import "math/rand"

func foo() func(int) {
	i := rand.Intn(10)
	return func(a int) {
		// i++ //会导致引用i
		print(i)
		print(a)
	}
}

func bar(f func(int)) {
	var fs []func(int)
	fs = append(fs, f)
	for _, v := range fs {
		v(2)
	}
}

func main() {
	bar(foo())
}

// // 匿名函数在不同场景的参数传递问题
// func main() {
// 	i := rand.Intn(10)
// 	// type funcval struct {
// 	// 	fn  uintptr
// 	// 	data *int
// 	// }
// 	// tmp := funcval{
// 	// 	fn: &main.func1,
// 	// 	data: &i,
// 	// }
// 	// bar(&tmp)

// 	// 1, 匿名函数作为参数(i在堆上，a在栈上)
// 	// bar(func(a int) {
// 	// 	print(i)
// 	// 	print(a)
// 	// })
// 	// 2, 匿名函数同步调用(没意义)
// 	// func() {
// 	// 	print(i)
// 	// }()
// 	// 3, 匿名函数异步调用(i和a都在栈上)
// 	// go func(a int) {
// 	// 	print(i)
// 	// 	print(a)
// 	// }(2)
// 	// 4, 匿名函数延迟调用(与异步调用保持一致)
// 	defer func(a int) {
// 		print(i)
// 		print(a)
// 	}(2)
// 	i++ //会导致引用i
// 	print(i)
// }
