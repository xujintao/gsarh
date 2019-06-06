package main

// func main() {
// 	m := make(map[string]int) // buckets域直接放栈上了
// 	// m := make(map[string]int, 10) //会调用runtime.makemap
// 	m["abc"] = 2
// 	println(m)
// }

func foo() map[string]int {
	m := make(map[string]int) // runtime.makemap_small
	// m := make(map[string]int, 10) // runtime.makemap
	m["abc"] = 2
	return m
}

func main() {
	m := foo()
	println(m)
}
