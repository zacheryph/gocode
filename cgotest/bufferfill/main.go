package main

/*
#include <grp.h>
#include <string.h>
#include "main.h"
*/
import "C"
import (
	"syscall"
	"unsafe"
)

func main() {
	C.cEntryPoint()
}

// align buffer to begin at an address that is modulo n. n must be a
// power of two.
func align(b []byte, n uintptr) []byte {
	off := uintptr(unsafe.Pointer(&b[0])) % n
	if off == 0 {
		return b
	}
	if n-off > uintptr(len(b)) {
		return nil
	}
	b = b[n-off:]
	return b
}

// buffer manages filling a buffer from both ends: strings go in the
// back, pointers go in the front. If the buffer runs out of space,
// err is set, and further operations are ignored.
type buffer struct {
	err syscall.Errno
	b   []byte
}

func newBuffer(b []byte) *buffer {
	buf := &buffer{
		b: align(b, unsafe.Alignof((*C.char)(nil))),
	}
	if buf.b == nil {
		buf.err = syscall.ERANGE
	}
	return buf
}

func (b *buffer) putString(s string) *C.char {
	if b.err != 0 {
		return nil
	}
	need := len(s) + 1
	if len(b.b) < need {
		b.err = syscall.ERANGE
		return nil
	}
	p := (*C.char)(unsafe.Pointer(&b.b[len(b.b)-need]))
	copy(b.b[len(b.b)-need:], s)
	b.b[len(b.b)-1] = 0x00
	b.b = b.b[:len(b.b)-need]
	return p
}

func (b *buffer) putPointer(p *C.char) {
	if b.err != 0 {
		return
	}
	const need = unsafe.Sizeof(p)
	if uintptr(len(b.b)) < need {
		b.err = syscall.ERANGE
		return
	}
	dst := (**C.char)(unsafe.Pointer(&b.b[0]))
	*dst = p
	b.b = b.b[need:]
	return
}

func (b *buffer) ptrToHead() **C.char {
	if b.err != 0 {
		return nil
	}
	return (**C.char)(unsafe.Pointer(&b.b[0]))
}

func fillGroup(dst *C.struct_group, dstbuf []byte, name, passwd string, gid C.__gid_t, groups []string) syscall.Errno {
	buf := newBuffer(dstbuf)
	dst.gr_mem = buf.ptrToHead()
	dst.gr_name = buf.putString(name)
	dst.gr_passwd = buf.putString(passwd)
	dst.gr_gid = gid

	for _, g := range groups {
		p := buf.putString(g)
		buf.putPointer(p)
	}
	buf.putPointer(nil)
	return buf.err
}

//export passingGrpBuffer
func passingGrpBuffer(dst *C.struct_group, dstbuf *C.char, dstbuflen C.int, errnop *C.int) C.int {
	b := C.GoBytes(unsafe.Pointer(dstbuf), dstbuflen)
	err := fillGroup(dst, b, "name", "passwd", 52, []string{
		"one", "two", "three",
	})
	if err != 0 {
		*errnop = C.int(err)
		return -1
	}
	return 0
}
