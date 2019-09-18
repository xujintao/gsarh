package main

import "net/http"

func main() {
	// num1 := 1
	// num1, num2 := 2, 2
	// num2, num3 := 1, 1
	// println(num1, num2, num3)
	err := foo()
	println(err)
}

func foo() (err error) {
	resp, err := http.Get("...")
	if err != nil {
		return
	}
	resp.Body.Close()
	return nil
}
