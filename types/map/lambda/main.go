package main

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		println(k, v)
	}

	// 匿名函数共享m，但是不会修改m变量，因为m本身是个引用
	// 等效于func(m map[string]int)
	func() {
		m["d"] = 4
	}()

	// 查
	if v, ok := m["a"]; ok {
		println(v)
	}
	if v, ok := m["d"]; ok {
		println(v)
	}
}
