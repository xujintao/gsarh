package main

// 清零
func main() {
	a := make([]int, 2)
	a = a[:0]
	println(a)
}
