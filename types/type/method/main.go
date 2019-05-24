package main

import "reflect"

// 编译器生成2个类型
// main.User,
// *main.User
type User struct {
	name string
	age  int
}

// 编译器生成2个方法
// "".User.Foo
// "".(*User).Foo
func (p User) Foo() {}

// 编译器生成1个方法
// "".(*User).Notify
func (p *User) Notify() {}

// 编译器生成2个类型
// main.Notifier,
// *main.Notifier 它没有研究的意义
// 编译器生成1个方法
// "".Notifier.Notify
type Notifier interface {
	Notify()
}

// 验证
func main() {
	var u User
	// main.User的方法集: Foo
	ut := reflect.TypeOf(u) // 构造eface, {main.User, u变量的值(是调用runtime.convT2E对u做copy出来的)}
	println(ut.String(),
		ut.Kind().String(),
		ut.NumMethod()) // need exported methods in the type's method set.
	for i := 0; i < ut.NumMethod(); i++ {
		m := ut.Method(i)
		println(m.Name)
	}
	println("")

	// *main.User的方法集: Foo, Notify
	rut := reflect.TypeOf(&u) // 构造eface, {*main.User, &u变量的值}
	println(rut.String(), rut.Kind().String(), rut.NumMethod())
	for i := 0; i < rut.NumMethod(); i++ {
		m := rut.Method(i)
		println(m.Name)
	}
	println("")

	// main.Notifier的方法集: Notify
	nt := reflect.TypeOf((*Notifier)(nil)).Elem() // 构造eface, {go.itab.*main.User,main.Notifier, &u变量的值}
	println(nt.String(), nt.Kind().String(), nt.NumMethod())
	for i := 0; i < nt.NumMethod(); i++ {
		m := nt.Method(i)
		println(m.Name)
	}
	println("")

	// *main.Notifier的方法集: 空
	rnt := reflect.TypeOf((*Notifier)(nil)) // 构造eface, {go.itab.*main.User,*main.Notifier, &u变量的值}
	println(rnt.String(), rnt.Kind().String(), rnt.NumMethod())
	for i := 0; i < rnt.NumMethod(); i++ {
		m := rnt.Method(i)
		println(m.Name)
	}

}
