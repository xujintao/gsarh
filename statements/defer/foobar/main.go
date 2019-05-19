package main

func bar() {
	defer func() {
		println("exit bar")
	}()
	defer func() {
		recover()
	}()

	println("entry bar")
	panic("bar")
}

func foo() {
	defer func() {
		println("exit foo")
		// recover()
	}()

	println("entry foo")
	bar()
	println("foo")
}

func main() {
	defer func() {
		println("exit main")
	}()

	println("entry main")
	foo()
	println("main")
}
