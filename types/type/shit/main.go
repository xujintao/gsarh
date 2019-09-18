package main

type student struct {
	age int
}

func (s student) foo() {
	s.age = 20
}

func (s *student) bar() {
	s.age = 21
}

func main() {

	func() {
		var ps = &student{} // "".(*student)类型变量
		ps.foo()            // 我觉得应该是"".(*student).foo(ps)，结果却是"".student.foo(*ps)
	}()

	func() {
		var s student // "".student类型变量
		s.bar()       // 我觉得应该是"".student.bar(s)，觉得编译会报错，结果却是"".(*student).bar(&s)
	}()

	// 总结，golang类型的方法集设计得像一陀屎，违背用户意志，充满了不确定性
	// 你要么就学习c++，只有指针类型有方法集，只传递this指针
}
