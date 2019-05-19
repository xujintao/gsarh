package main

// func main() {
// 	m := make(map[string]int) // buckets域直接放栈上了
// 	// m := make(map[string]int, 10) //会调用runtime.makemap
// 	m["abc"] = 2
// 	println(m)
// }

// func foo() map[string]int {
// 	m := make(map[string]int, 10)
// 	m["abc"] = 2
// 	return m
// }

// func main() {
// 	m := foo()
// 	println(m)
// }

func main() {
	m := make(map[int]int) //runtime.makemap_small
	go func() {
		for {
			m[1] = 1 //runtime.mapassign_fast64
		}
	}()
	go func() {
		for {
			_ = m[1] //runtime.mapaccess1_fast64
		}
	}()
	select {}
}
