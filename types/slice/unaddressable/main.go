package main

import (
	"crypto/md5"
	"strconv"
	"unsafe"
)

func addressable() {
	buf := make([]byte, 1024)
	for i := range buf {
		buf[i] = 0
	}
	println("slice addr:", &buf)
	slicep := (*[3]uint)(unsafe.Pointer(&buf))
	println("slice element:", strconv.FormatUint(uint64(slicep[0]), 16), slicep[1], slicep[2])
	println("buf[0] addr:", &buf[0])
	// println(&buf[0] % 16)
	// println(slicep[0] % 16)
}

func unaddressable() {
	// // invalid operation [3]int literal[:] (slice of unaddressable value)
	// arr := ([3]int{1, 2, 3})[:]
	// println(arr[2])

	// // invalid operation md5.Sum(([]byte)("123"))[:] (slice of unaddressable value)
	// arrsum := md5.Sum([]byte("123"))[:]
	// println(string(arrsum))

	// ok
	arr := [3]int{1, 2, 3}
	println(arr[:])

	// ok
	arrsum := md5.Sum([]byte("123"))
	println(string(arrsum[:]))

	// github.com/golang/go/test/complit1.go
	// 想构造slice，底层数组必须可寻址
}

func main() {
	addressable()
	unaddressable()
}
