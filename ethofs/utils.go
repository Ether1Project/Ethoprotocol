package ethofs

import (
	//"bytes"
	"unsafe"
)

func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
