package main

import (
	"os"
	"unsafe"
)

// runtime/type.go
type _type struct {
	size    uintptr
	ptrdata uintptr // size of memory prefix holding all pointers
	hash    uint32
	// tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	// alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata *byte
	// str       nameOff
	// ptrToThis typeOff
}

type eface struct {
	typ  *_type
	data unsafe.Pointer
}

//go:inline build-in
func assert(v eface, typ *_type) (unsafe.Pointer, bool) {
	if v.typ == typ {
		return v.data, true
	}
	return nil, false
}

func foo(w interface{}) {
	// if w, ok := assert(w, &type.*os.File)
	if w, ok := w.(*os.File); ok {
		w.Name()
	}
}

func main() {

}
