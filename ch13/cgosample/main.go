package main

/*
#include <stdlib.h>
#include <stdio.h>

// Cのコードはimport "C"の真上に書く
void hello(char *str)
{
  printf("hello, %s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	greetingTo := C.CString("cgo")
	// CStringで得たstring型は必ず開放する
	defer C.free(unsafe.Pointer(greetingTo))
	C.hello(greetingTo)
}
