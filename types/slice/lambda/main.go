package main

func main() {
	s := []int{1, 2}
	for _, v := range s {
		println(v) // 1, 2
	}

	// 匿名函数共享s，但是会修改s的值
	// 等效于func(s *[]int)
	func() {
		s = append(s, 3)
	}()

	for _, v := range s {
		println(v) // 1, 2, 3
	}
}
