package stemmer

// #include "porter.h"
import "C"

import (
	"runtime"
	"unsafe"
)

type PorterStemmer struct {
	cstemmer *C.struct_stemmer
}

func freePorterStemmer(ps *PorterStemmer) {
	C.free_stemmer(ps.cstemmer)
}

func (ps *PorterStemmer) Stem(s string) string {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	end := C.stem(ps.cstemmer, cs, C.int(len(s)-1)) + 1
	return s[0:end]
}

func NewPorterStemmer() Stemmer {
	s := &PorterStemmer{cstemmer: C.create_stemmer()}
	runtime.SetFinalizer(s, freePorterStemmer)
	return s
}
