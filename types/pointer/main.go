package main

import "unsafe"

func main() {
	val := 4
	valp := &val
	valpnext := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(valp)) + unsafe.Sizeof(val)))
	println(valp, valpnext)
}

// c处理数组一般使用指针，而golang处理数组一般使用切片(对指针的一种封装)
// 所以golang虽然支持指针运算，但一般没必要使用
