package main

type barer interface {
	bar()
}

type fooer interface {
	foo()
}

type base struct {
	barer
}

func (b *base) foo() {
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

func main() {
	var services []interface{}
	// services[0] = &service1{base{&service1{}}}
	// services[1] = &service2{base{&service2{}}}
	s1 := &service1{}
	s1.base.barer = s1
	services = append(services, s1)

	s2 := &service2{}
	s2.base = base{s2}
	services = append(services, s2)

	for _, s := range services {
		s.(fooer).foo()
	}
}
