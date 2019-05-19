package main

func main() {}

// select写通道
// compiler implements
//
//	select {
//	case c <- v:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if runtime.selectnbsend(c, v) { // runtime/chan.go
//		... foo
//	} else {
//		... bar
//	}
//

// select读通道
// compiler implements
//
//	select {
//	case v = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if runtime.selectnbrecv(&v, c) { // runtime/chan.go
//		... foo
//	} else {
//		... bar
//	}
//

// select读通道(带ok)
// compiler implements
//
//	select {
//	case v, ok = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if c != nil && runtime.selectnbrecv2(&v, &ok, c) { // runtime/chan.go
//		... foo
//	} else {
//		... bar
//	}
//

// select多通道
// runtime.selectgo // runtime/select.go

// select{}
// runtime.block // runtime/select.go
