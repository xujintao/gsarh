package main

import (
	"fmt"
	"reflect"
)

type Barer interface {
	bar()
}

type base struct {
	Barer
}

func (b *base) Foo() {
	b.bar()
}

type service1 struct {
	base
}

func (s *service1) bar() {
	println("service1 bar")
}

type service2 struct {
	base
}

func (s *service2) bar() {
	println("service2 bar")
}

// 配置
var services = [...]Barer{
	&service1{}, // what about "service1": https://stackoverflow.com/questions/23030884/is-there-a-way-to-create-an-instance-of-a-struct-from-a-string
	&service2{},
}

func main() {
	// construct
	for _, s := range services {
		v := reflect.ValueOf(s).Elem().FieldByName("base").FieldByName("Barer")
		if v.Kind() == reflect.Invalid {
			panic(fmt.Sprintln(reflect.TypeOf(s).String(), "should inherit base"))
		}
		v.Set(reflect.ValueOf(s))
	}

	// call
	for _, s := range services {
		// v := reflect.ValueOf(s).Elem()
		v := reflect.ValueOf(s)
		v.MethodByName("Foo").Call(nil)
		// s.bar()
	}
}
