package main

import "C"
import (
	"bytes"
	"reflect"
	"unsafe"
)

// Using Reflect
// does not work
func passingGrpBufferWithReflect(dst *C.struct_group, dstbuf *C.char, dstbuflen C.int, errnop *C.int) {
	ourGid := 52
	byteArr := C.GoBytes(unsafe.Pointer(buf), buflen)

	byteBuf := bytes.NewBuffer(byteArr[:0])
	for _, s := range []string{"name", "passwd", "one", "two", "three"} {
		byteBuf.WriteString(s)
		byteBuf.WriteByte(0)
	}

	var p *C.char

	memPtr := []*C.char{
		(*C.char)(unsafe.Pointer(&byteArr[12])),
		(*C.char)(unsafe.Pointer(&byteArr[16])),
		(*C.char)(unsafe.Pointer(&byteArr[20])),
		(*C.char)(nil),
	}

	lenCap := len(memPtr) * int(unsafe.Sizeof(p))
	refSlice := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&memPtr[0])),
		Len:  lenCap, Cap: lenCap}

	memBytes := *(*[]byte)(unsafe.Pointer(refSlice))
	byteBuf.Write(memBytes)

	grp.gr_name = (*C.char)(unsafe.Pointer(&byteArr[0]))
	grp.gr_passwd = (*C.char)(unsafe.Pointer(&byteArr[5]))
	grp.gr_gid = C.gid_t(ourGid)
	grp.gr_mem = (**C.char)(unsafe.Pointer(&byteArr[26]))
}
