package main

func main() {
	b1 := []int{1, 2, 3}
	b2 := make([]*int, 0, 3)
	for k, v := range b1 {
		println(k, v)
		// b2 = append(b2, &v) // 值拷贝
		b2 = append(b2, &b1[k])
	}

	for _, v := range b2 {
		println(*v)
	}
}
