package main

/*
#include <grp.h>
#include <string.h>
#include "main.h"
*/
import "C"
import (
	"bytes"
	"unsafe"
)

type group C.struct_group

func main() {
	C.cEntryPoint()
}

//export passingGrpBuffer
func passingGrpBuffer(grp *group, buf *C.char, buflen C.int) {
	ourGid := 52
	byteArr := C.GoBytes(unsafe.Pointer(buf), buflen)

	// this parts easy
	grp.gr_gid = C.gid_t(ourGid)

	// using copy
	// copy(byteArr[0:], []byte("name"))
	// copy(byteArr[5:], []byte("passwd"))

	byteBuf := bytes.NewBuffer(byteArr[:0])
	for _, s := range []string{"name", "passwd", "one", "two", "three"} {
		byteBuf.WriteString(s)
		byteBuf.WriteByte(0)
	}

	// memPtr := []*C.char{
	// 	(*C.char)(unsafe.Pointer(&byteArr[12])),
	// 	(*C.char)(unsafe.Pointer(&byteArr[16])),
	// 	(*C.char)(unsafe.Pointer(&byteArr[20])),
	// }
	//
	// TODO: the troubling part. writing back to the buffer
	// This dont work obviously.
	// byteBuf.Write([]byte(memPtr))
	// byteBuf.WriteByte(0)

	grp.gr_name = (*C.char)(unsafe.Pointer(&byteArr[0]))
	grp.gr_passwd = (*C.char)(unsafe.Pointer(&byteArr[5]))
	grp.gr_gid = C.gid_t(ourGid)
	grp.gr_mem = (**C.char)(unsafe.Pointer(&byteArr[26]))
}
