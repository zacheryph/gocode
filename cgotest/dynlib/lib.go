// #cgo CFLAGS: -shared

// Package dynlib is our dynamic library to be loaded into our C code
package main

import "C"
import "fmt"

func main() {}

//export helloWorld
func helloWorld(name *C.char) *C.char {
	full := fmt.Sprintf("Hello %s!\n", C.GoString(name))
	fmt.Println("Hello World Here!")
	return C.CString(full)
}
