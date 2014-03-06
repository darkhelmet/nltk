package filter

/*
#cgo CFLAGS: -Wno-parentheses
#include "double_metaphone.h"
*/
import "C"

import (
	"unsafe"

	"github.com/darkhelmet/nltk"
)

func DoubleMetaphone(in nltk.TokenChan) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		cs := C.CString(tok.String())
		defer C.free(unsafe.Pointer(cs))
		codes := C.double_metaphone(cs)
		primary, secondary := C.GoString(codes.primary), C.GoString(codes.secondary)
		defer C.free_dm_result(codes)
		out <- nltk.Token(primary)
		if primary != secondary {
			out <- nltk.Token(secondary)
		}
	})
}
