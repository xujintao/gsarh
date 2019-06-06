package main

// 1, hashmap的桶数组
// 2, kafka的partition数组
// 3, nginx的upstream数组
// 4, mysql分库分表
// 5, tcp多进程监听器/udp多进程
// 6, golang的通道数组
// 7, golang的定时器桶数组
// 对这些数组做负载均衡就是在做hash

func foo() map[string]int {
	// // 方式1
	// m := make(map[string]int) // runtime.makemap_small
	// m["a"] = 'a'              // runtime.mapassign_faststr
	// m["b"] = 'b'              // runtime.mapassign_faststr
	// m["c"] = 'c'              // runtime.mapassign_faststr
	// m["d"] = 'd'              // runtime.mapassign_faststr

	// 方式2
	m := make(map[string]int, 10) // runtime.makemap, 因为hint大于8了，所以会先分配一个桶
	m["a"] = 'a'                  // runtime.mapassign_faststr
	m["b"] = 'b'                  // runtime.mapassign_faststr
	m["c"] = 'c'                  // runtime.mapassign_faststr
	m["d"] = 'd'                  // runtime.mapassign_faststr

	// // 方式3 与 方式1 一样
	// m := map[string]int{"a": 'a', "b": 'b', "c": 'c', "d": 'd'}
	return m
}

func main() {
	m := foo()

	// // 查方式1
	// println(m["d"]) // runtime.mapaccess1_faststr

	// // 查方式2
	// if v, ok := m["d"]; ok { // runtime.mapaccess2_faststr
	// 	println(v)
	// }

	// 遍历
	for k, v := range m { // runtime.mapiterinit, runtime.mapiternext
		println(k, v)
	}
	// println(len(m))
}
