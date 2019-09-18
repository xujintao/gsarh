package main

func main() {
	y := make([]int, 10)
	println(len(y))
	func(x interface{}) {
		y1 := x.(*[]int)
		(*y1) = (*y1)[1:]
		println(len(*y1))
	}(&y)
	println(len(y))
	foo()
}
