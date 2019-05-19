package main

func main() {

	// string转[]byte
	// s := []byte("abc") // 直接构造出byte切片
	// rawstr := "abc"
	// s := []byte(rawstr) // runtime.stringtoslicebyte
	// print(s)

	// []byte转string
	// str := string(s) // runtime.slicebytetostring
	// print(str)

	// // string转[]rune
	s := []rune("你好a") // 直接构造rune切片
	// rawstr := "你好a"
	// s := []rune(rawstr) // runtime.stringtoslicerune
	println(s)

	// // []rune转string
	// str := string(s[:1]) // runtime.slicerunetostring
	// println(str)
}
