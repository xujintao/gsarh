package main

func inc() (v int) {
	defer func() { v++ }()
	return 42
}

func main() {
	print(inc()) // 43
}
