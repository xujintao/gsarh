package main

func main() {
	// foo()
	bar()
}

func foo() {
	var ebp30 bool

	func() {
		// 注册异常处理1
		defer func() {
			recover()
		}()
		// 注册异常处理2
		defer func() {
			recover()
			ebp30 = true
		}()
		panic("int3")
	}()

	// 结果
	if ebp30 {
		println("right")
	} else {
		println("wrong")
	}
}

func bar() {
	var ebp30 bool

	// 注册异常处理1
	defer func() {
		recover()
	}()
	// 注册异常处理2
	defer func() {
		recover()
		ebp30 = true
	}()
	panic("int3")

	// 结果
	if ebp30 {
		println("right")
	} else {
		println("wrong")
	}
}
