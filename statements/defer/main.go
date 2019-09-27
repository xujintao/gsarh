package main

func main() {
	// 1，起始逻辑

	// 2，注册异常处理
	defer func() {
		// recover() // 我信心满满
		// 对异常的处理，无非是做一些资源的释放
	}()

	// 3，可能会发生异常的函数调用(这里模拟一个会产生异常的函数调用)
	func() {
		panic("int3")
	}()

	// 4，后续逻辑
}

// golang的defer在设计上和SEH的__try/__finally很类似，然而区别也很大
// 1，异常产生后应当遍历整个异常处理链然后退出，排除很有信心的recover情况后续逻辑不应该继续执行，这一点上SEH从设计上就错了
// 2，相比SEH又是下划线又是大括号的，defer在代码风格上更加优雅
