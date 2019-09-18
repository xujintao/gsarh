package main

import "fmt"

func main() {
	// string2slice
	func() {
		// 或者直接 s := []byte("123")
		str := "123"
		s := []byte(str) // runtime.stringtoslicebyte，new and copy
		println("func1", s)
	}()

	// slice2string
	func() {
		slice := []byte{49, 50, 51} // 或者直接 s := []byte("123")
		str := string(slice)        // runtime.slicebytetostring(SB), new and memmove
		println("func2", str)
	}()

	// string2array
	func() {
		// var array [10]byte = "123" // c++允许这样初始化吗？
		var array [10]byte
		l := copy(array[:], "123")
		// println(array) // println不能打印数组，改用fmt
		println("func3", l) // println不能打印数组，改用fmt
	}()

	// array2string
	func() {
		array := [...]byte{49, 50, 51}
		str := string(array[:])
		println("func4", str)
	}()

	// 变长参数，func5
	func(format string, a ...interface{}) {
		fmt.Printf(format, a...)
	}("func5, %d.%d\n", 1, 2)

	func() { // func6
		array := [...]byte{49, 50, 51}
		fmt.Printf("func6, %s\n", array)
		fmt.Printf("func6, %s\n", array[:])
		fmt.Printf("func6, %s\n", string(array[:]))
	}()

	// func7
	func() {
		var s []byte
		str := "123"
		s = append(s, str...) // 这是什么奇技淫巧？
		println("func7", s)
	}()

	// func8
	func() {
		s := []byte("123")
		println("func8", s)
	}()

	// func9
	func() {
		array := [...]byte{0x59, 0x6F, 0x75, 0, 0}
		str := string(array[:])
		println("func9", str, len(str))
	}()
}
