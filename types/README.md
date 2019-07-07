## golang type system

golang类型系统有点难, 但是只有掌握了类型系统才能游刃有余的进行oop, 才能理解mm  
先打个广告，对golang类型系统有什么想法的可以加企鹅群95786324找阿皮，这一点很重要。
### 一，类型与kind  
kind是类型的分类
```golang
// reflect/type.go
var kindNames = []string{
	Invalid:       "invalid",
	Bool:          "bool",
	Int:           "int",
	Int8:          "int8",
	Int16:         "int16",
	Int32:         "int32",
	Int64:         "int64",
	Uint:          "uint",
	Uint8:         "uint8",
	Uint16:        "uint16",
	Uint32:        "uint32",
	Uint64:        "uint64",
	Uintptr:       "uintptr",
	Float32:       "float32",
	Float64:       "float64",
	Complex64:     "complex64",
	Complex128:    "complex128",
	Array:         "array",
	Chan:          "chan",
	Func:          "func",
	Interface:     "interface",
	Map:           "map",
	Ptr:           "ptr",
	Slice:         "slice",
	String:        "string",
	Struct:        "struct",
	UnsafePointer: "unsafe.Pointer",
}
```
速度可以有前进1，前进2，前进3  
同样golang也有类型1，类型2，类型3  

#### 类型1：要落到bool, numeric, string这3种kind中  
由于这3种类型已经被编译器穷举了，所以类型的identifier和种类的名称是一样的  
```
// 比如
func main() {
	println(reflect.TypeOf(true).String(),
		reflect.TypeOf(true).Kind().String()) // bool bool
	println(reflect.TypeOf(uint8(1)).String(),
		reflect.TypeOf(uint8(1)).Kind().String()) // uint8 uint8
	println(reflect.TypeOf("abc").String(),
		reflect.TypeOf("abc").Kind().String()) // string string
}
```
类型1对应的指针类型：*bool, *numeric, *string要落到ptr中  
```
// 比如*bool
func main() {
	println(reflect.TypeOf((*bool)(nil)).String(),
		reflect.TypeOf((*bool)(nil)).Kind().String()) // *bool ptr
}
```

#### 类型2：由用户穷举，然后要落到array, slice, map, chan, func, struct, interface这7种kind中
注意：interface是有类型的，而空接口interface{}没有类型。  
《香水》这部电影描述的是除了男主没有气味，其他人都有气味，然后男主就开始收集气味。
同样美剧《权力的游戏》里除了瓦里斯没有蛋蛋其他人都有蛋蛋，有蛋蛋的为自己的蛋蛋活着，没有蛋蛋的为所有人活着。
所以，建议认真看完这两部作品以加深对空接口的理解。
```
func main() {
	println(reflect.TypeOf([2]int{}).String(),
		reflect.TypeOf([2]int{}).Kind().String()) // [2]int array
	println(reflect.TypeOf([]int{}).String(),
		reflect.TypeOf([]int{}).Kind().String()) // []int slice
	println(reflect.TypeOf(&[]int{}).String(),
		reflect.TypeOf(&[]int{}).Kind().String()) // *[]int ptr
	println(reflect.TypeOf(map[string]int{}).String(),
		reflect.TypeOf(map[string]int{}).Kind().String()) // map[string]int map
	println(reflect.TypeOf((*chan int)(nil)).Elem().String(),
		reflect.TypeOf((*chan int)(nil)).Elem().Kind().String()) //chan int chan
	println(reflect.TypeOf(func(int) {}).String(),
		reflect.TypeOf(func(int) {}).Kind().String()) // func(int) func
	println(reflect.TypeOf(struct{ int }{}).String(),
		reflect.TypeOf(struct{ int }{}).Kind().String()) // struct {} struct
	println(reflect.TypeOf((*interface{ foo() })(nil)).Elem().String(),
		reflect.TypeOf((*interface{ foo() })(nil)).Elem().Kind().String())
	// interface { main.foo() } interface
}
```
同样，类型2对应的指针类型要落到ptr中  
```
// 比如*[]int
func main() {
	println(reflect.TypeOf(&[]int{}).String(),
		reflect.TypeOf(&[]int{}).Kind().String()) // *[]int ptr
}
```

#### 类型3：使用type definition自定义出的类型  
哦，对了，自定义类型的identify要带名字空间所以会带包名  
```
type T1 int                // main.T1
type T2 struct{ a int }    // main.T2
type T3 interface{ foo() } // main.T3
func main() {
	println(reflect.TypeOf(T1(1)).String(),
		reflect.TypeOf(T1(1)).Kind().String()) // main.T1 int
	println(reflect.TypeOf(T2{}).String(),
		reflect.TypeOf(T2{}).Kind().String()) // main.T2 struct
	println(reflect.TypeOf((*T3)(nil)).Elem().String(),
		reflect.TypeOf((*T3)(nil)).Elem().Kind().String()) // main.T3 interface
}
```
一样的，类型3对应的指针类型要落到ptr中  
```
// 例如*main.T2
func main() {
	// println(reflect.TypeOf(&T2{1}).String(),
	// 	reflect.TypeOf(&T2{1}).Kind().String()) // *main.T2 ptr
	println(reflect.PtrTo(reflect.TypeOf(T2{1})).String(),
		reflect.PtrTo(reflect.TypeOf(T2{1})).Kind().String()) // *main.T2 ptr
}
```

##### 类型可以千变万化，但最终还是要落到kind框架中  
##### 一个值得注意的规律是，类型与其对应的指针类型要么都不出现，要么总是成对出现，它们是双子  
##### 另一个特色就是不同类型不能互相赋值  
```
buf := make([]byte, 1024) // type byte = uint8 at src/builtin/builtin.go

// ok
buf[0] = uint8(1)

// cannot use false (type bool) as type byte in assignment
buf[0] = false

// cannot use struct {} literal (type struct {}) as type byte in assignment
buf[0] = struct{}{}

// cannot use &buf[0] (type *byte) as type uintptr in assignment
var addr1 uintptr = &buf[0]

// cannot use unsafe.Pointer(&buf[0]) (type unsafe.Pointer) as type uintptr in assignment
var addr2 uintptr = unsafe.Pointer(&buf[0])

println(addr1, addr2)
```
那么问题来了，类型是作用于值还是作用于变量?  

### 二，类型的进一步探索  
编译器会为每种类型自动生成一个类型常量  
类型1对应的是_type类型  
类型2和3需要对_type类型进行拓展  

```
// runtime/type.go
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}

type arraytype struct {
	typ   _type
	elem  *_type
	slice *_type
	len   uintptr
}

type chantype struct {
	typ  _type
	elem *_type
	dir  uintptr
}

type functype struct {
	typ      _type
	inCount  uint16
	outCount uint16
}

type imethod struct {
	name nameOff
	ityp typeOff
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type maptype struct {
	typ        _type
	key        *_type
	elem       *_type
	bucket     *_type // internal type representing a hash bucket
	keysize    uint8  // size of key slot
	valuesize  uint8  // size of value slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

type ptrtype struct {
	typ  _type
	elem *_type
}

type slicetype struct {
	typ  _type
	elem *_type
}

type structfield struct {
	name       name
	typ        *_type
	offsetAnon uintptr
}

func (f *structfield) offset() uintptr {
	return f.offsetAnon >> 1
}

type structtype struct {
	typ     _type
	pkgPath name
	fields  []structfield
}
```

##### 可以看到我们平时使用某个类型申明变量(常量)以后，变量(常量)的abi落到了内存，变量(常量)的类型也以常量(read only)的形式落到了内存。
##### 变量(常量)的abi很重要。
abi信息应该能从编译器代码那边找到，但是我找半天没找到，只能照着runtime代码手写几个重要的来释义。
```
// string类型
type string struct {
	base unsafe.Pointer
	len int
}

// slice类型
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// map类型(hmap的定义在runtime/map.go中)
type map *hmap

// chan类型(hchan的定义在runtime/chan.go中)
type chan *hchan

// func类型(funcval的定义在runtime/runtime2.go中)
type func *funcval

// interface类型(itab的定义在runtime/runtime2.go中)
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

// interface{}
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

### 三，从编译器角度理解类型(留给老冯写)  


### 四，自定义类型的方法集
```
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
```

### 五，类型继承  
