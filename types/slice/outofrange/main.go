package main

// slice越界检查
func main() {
	a := make([]int, 1, 2)
	a[2] = 0
}

// "".main STEXT nosplit size=19 args=0x0 locals=0x8
// 	0x0000 00000 (main.go:4)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $8-0
// 	0x0000 00000 (main.go:4)	SUBQ	$8, SP
// 	0x0004 00004 (main.go:4)	MOVQ	BP, (SP)
// 	0x0008 00008 (main.go:4)	LEAQ	(SP), BP
// 	0x000c 00012 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x000c 00012 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x000c 00012 (main.go:4)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x000c 00012 (main.go:33)	PCDATA	$2, $0
// 	0x000c 00012 (main.go:33)	PCDATA	$0, $0
// 	0x000c 00012 (main.go:33)	CALL	runtime.panicindex(SB)
// 	0x0011 00017 (main.go:33)	UNDEF
