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

##### 一个值得注意的规律是，类型与其对应的指针类型要么都不出现，要么总是成对出现，它们是双子  

### 二，类型的进一步探索  
然而，任何golang类型都是一个_type类型的常量  
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
```
我们可以查看每个类型的_type结构
```
type _type struct {
	Size       uintptr `json:"size"`
	Ptrdata    uintptr `json:"ptrdata"`
	Hash       uint32  `json:"hash"`
	Tflag      uint8   `json:"tflag"`
	Align      uint8   `json:"align"`
	Fieldalign uint8   `json:"fieldalign"`
	Kind       uint8   `json:"kind"`
	Alg        uintptr `json:"alg"`
	Gcdata     *byte   `json:"gcdata"`
	Str        int32   `json:"str"`
	PtrToThis  int32   `json:"ptrToThis"`
}

func printType(t reflect.Type) {
	at := (*[2]unsafe.Pointer)(unsafe.Pointer(&t)) // 把reflect.Type转成指针数组
	res, _ := json.Marshal((*_type)(at[1])) // 取data域并序列化出来
	println(string(res))
}

type T1 int                // main.T1
type T2 struct{ a int }    // main.T2
type T3 interface{ foo() } // main.T3
func main() {
	printType(reflect.TypeOf("abc"))
	printType(reflect.PtrTo(reflect.TypeOf("abc")))
	printType(reflect.TypeOf((*chan int)(nil)).Elem())
	printType(reflect.TypeOf((*chan int)(nil)))
	printType(reflect.TypeOf((*interface{ foo() })(nil)).Elem())
	printType(reflect.TypeOf((*interface{ foo() })(nil)))
	printType(reflect.TypeOf(T1(1)))
	printType(reflect.PtrTo(reflect.TypeOf(T1(1))))
	printType(reflect.TypeOf(T2{}))
	printType(reflect.PtrTo(reflect.TypeOf(T2{})))
}
```

类型1
```
看看常量type.string和常量type.*string的几种表达方式

// 使用golang语言表达
const type.string = _type{}
const type.*string = _type{}
// 把内容序列化出来
"type.string": {
	"size": 16,
	"ptrdata": 8,
	"hash": 3774831796,
	"tflag": 7,
	"align": 8,
	"fieldAlign": 8,
	"kind": 24, // 24&kindMask = 24，是string
	"alg": 5954192,
	"gcdata": 1,
	"str": 6768,
	"ptrToThis": 63200
}
"type.*string": {
	"size": 8,
	"ptrdata": 8,
	"hash": 1511480045,
	"tflag": 0,
	"align": 8,
	"fieldAlign": 8,
	"kind": 54, // 54&kindMask = 54-32 = 22，是ptr
	"alg": 5954160,
	"gcdata": 1,
	"str": 6768,
	"ptrToThis": 0
}

// 汇编表达暂时没法看
```

类型2
```
看看常量"type.chan int"和常量"type.*chan int"的几种表达方式

// 使用golang语言表达
const type.chan_int = _type{}
const type.*chan_int = _type{}
// 把内容序列化出来
"type.chan int": {
	"size": 8,
	"ptrdata": 8,
	"hash": 1909151121,
	"tflag": 2,
	"align": 8,
	"fieldAlign": 8,
	"kind": 50, // 50&kindMask = 50-32 = 18，是chan
	"alg": 5954160,
	"gcdata": 1,
	"str": 10888,
	"ptrToThis": 51040	
}
"type.*chan int": {
	"size": 8,
	"ptrdata": 8,
	"hash": 1005419501,
	"tflag": 0,
	"align": 8,
	"fieldAlign": 8,
	"kind": 54, // 54&kindMask = 54-32 = 22，是ptr
	"alg": 5954160,
	"gcdata": 1,
	"str": 10888,
	"ptrToThis": 0	
}

// 使用golang汇编表达
type.chan int SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 91 55 cb 71 02 08 08 32 00 00 00 00 00 00 00 00  .U.q...2........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 44+4 t=6 type.*chan int+0
	rel 48+8 t=1 type.int+0
type.*chan int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ed 7b ed 3b 00 08 08 36 00 00 00 00 00 00 00 00  .{.;...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 48+8 t=1 type.chan int+0

还要看看常量type.interface{ main.foo()}和常量type.*interface{ main.foo() }的几种表达方式

// 使用golang语言表达
const type.interface{ main.foo()} = _type{}
const type.*interface{ main.foo()} = _type{}
// 把内容序列化出来
"type.interface{ main.foo()}": {
	"size": 16,
	"ptrdata": 16,
	"hash": 1827387388,
	"tflag": 2,
	"align": 8,
	"fieldAlign": 8,
	"kind": 20, // 20&kindMask = 20，是interface
	"alg": 5954208,
	"gcdata": 2,
	"str": 39757,
	"ptrToThis": 52480
}
"type.*interface{ main.foo()}": {
	"size": 8,
	"ptrdata": 8,
	"hash": 3471523764,
	"tflag": 0,
	"align": 8,
	"fieldAlign": 8,
	"kind": 54, // 54&kindMask = 54-32 = 22，是ptr
	"alg": 5954160,
	"gcdata": 1,
	"str": 39757,
	"ptrToThis": 0
}

// 使用golang汇编表达
type.interface { main.foo() } SRODATA dupok size=88
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 fc b7 eb 6c 02 08 08 14 00 00 00 00 00 00 00 00  ...l............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+128
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface { main.foo() }-+0
	rel 44+4 t=6 type.*interface { main.foo() }+0
	rel 56+8 t=1 type.interface { main.foo() }+80
	rel 80+4 t=5 type..namedata."".0+0
	rel 84+4 t=5 type.func()+0
type.*interface { main.foo() } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b4 3f eb ce 00 08 08 36 00 00 00 00 00 00 00 00  .?.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface { main.foo() }-+0
	rel 48+8 t=1 type.interface { main.foo() }+0
```

类型3
```
看看常量type.main.T1和常量type.*main.T1的几种表达方式

// 使用golang语言表达
const type.main.T1 = _type{}
const type.*main.T1 = _type{}
// 把内容序列化出来
"type.main.T1": {
	"size": 8,
	"ptrdata": 0,
	"hash": 2079407416,
	"tflag": 7,
	"align": 8,
	"fieldAlign": 8,
	"kind": 130, // 130&kindMask = 130-128 = 2，是int
	"alg": 5954160,
	"gcdata": 1,
	"str": 8841,
	"ptrToThis": 52704
}
"type.*main.T1": {
	"size": 8,
	"ptrdata": 8,
	"hash": 1599159816,
	"tflag": 0,
	"align": 8,
	"fieldAlign": 8,
	"kind": 54, // 54&kindMask = 54-32 = 22，是ptr
	"alg": 5954160,
	"gcdata": 1,
	"str": 8841,
	"ptrToThis": 0
}

// 汇编表达
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.T1. SRODATA dupok size=11
	0x0000 01 00 08 2a 6d 61 69 6e 2e 54 31                 ...*main.T1
type.*"main".T1 SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 08 3e 51 5f 00 08 08 36 00 00 00 00 00 00 00 00  .>Q_...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.T1.+0
	rel 48+8 t=1 type."".T1+0
runtime.gcbits. SRODATA dupok size=0
type."main".T1 SRODATA size=64
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 38 3d f1 7b 07 08 08 82 00 00 00 00 00 00 00 00  8=.{............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.T1.+0
	rel 44+4 t=5 type.*"".T1+0
	rel 48+4 t=5 type..importpath."".+0

然后再看看常量type.main.T2和常量type.*main.T2的几种表达方式

// golang表达
const type.main.T2 = _type{}
const type.*main.T2 = _type{}
// 把内容序列化出来
"type.main.T2": {
	"size": 8,
	"ptrdata": 0,
	"hash": 2038138561,
	"tflag": 7,
	"align": 8,
	"fieldAlign": 8,
	"kind": 153, // 153&kindMask = 153-128 = 25, 是struct
	"alg": 5954160,
	"gcdata": 1,
	"str": 8852,
	"ptrToThis": 52768
}
"type.*main.T2": {
	"size": 8,
	"ptrdata": 8,
	"hash": 4091785602,
	"tflag": 0,
	"align": 8,
	"fieldAlign": 8,
	"kind": 54, // 54&kindMask = 54-32 = 22, 是ptr
	"alg": 5954160,
	"gcdata": 1,
	"str": 8852,
	"ptrToThis": 0
}

// 汇编表达
type..namedata.*main.T2. SRODATA dupok size=11
	0x0000 01 00 08 2a 6d 61 69 6e 2e 54 32                 ...*main.T2
type.*"".T2 SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 82 b1 e3 f3 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.T2.+0
	rel 48+8 t=1 type."".T2+0
type."".T2 SRODATA size=120
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 c1 86 7b 79 07 08 08 99 00 00 00 00 00 00 00 00  ..{y............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.T2.+0
	rel 44+4 t=5 type.*"".T2+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".T2+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.a-+0
	rel 104+8 t=1 type.int+0
```

##### 可以看到我们平时使用某个类型申明变量以后，变量落到了内存，变量的类型也以常量(read only)的形式落到了内存。
##### 类型常量很有用，mm的gcmalloc需要它，范型interface{}也需要它。
##### _type.size域表达了变量占用字节数，具体含义我会在目录树中逐个详细分析。

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
