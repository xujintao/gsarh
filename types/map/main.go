package main

// 1, hashmap的桶数组
// 2, kafka的partition数组
// 3, nginx的upstream数组
// 4, mysql分库分表
// 5, golang的通道数组
// 6, golang的定时器桶数组
// 对这些数组做负载均衡就是在做hash
func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		println(k, v)
	}

	// 查
	if v, ok := m["a"]; ok {
		println(v)
	}
	if v, ok := m["d"]; ok {
		println(v)
	}
}
