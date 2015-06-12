package main

import (
	. "../lua"
	"unsafe"
)

/*
#include <stdlib.h>

extern  void*	allocator(void *ud, void *ptr, size_t osize, size_t nsize);

*/
import "C"

//export allocator
func allocator(ud unsafe.Pointer, ptr unsafe.Pointer, osize uintptr, nsize uintptr) (ret unsafe.Pointer) {
	ret = nil
	if nsize == 0 {
		C.free(ptr)
	} else {
		ret = C.realloc(ptr, C.size_t(nsize))
	}
	return
}

func main() {
	l := Lua_newstate(LuaF_Alloc(C.allocator), uintptr(0))

	LuaL_openlibs(l)
	LuaL_dostring(l, `print'Hello World!\n'`)

	Lua_close(l)
}
