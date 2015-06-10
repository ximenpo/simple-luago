package main

import(
    "unsafe"
    . "../lua"
)

/*
#include <stdlib.h>

extern  void*	allocator(void *ud, void *ptr, size_t osize, size_t nsize);

*/
import  "C"

//export allocator
func allocator(ud unsafe.Pointer, ptr unsafe.Pointer, osize C.size_t, nsize C.size_t) unsafe.Pointer {
	var ret unsafe.Pointer = nil
	if nsize == 0 {
		C.free(ptr)
	} else {
		ret = C.realloc(ptr, nsize)
	}

	return ret
}

func main() {
	l := Lua_newstate(LuaF_Alloc(C.allocator), uintptr(0))

	LuaL_openlibs(l)
	LuaL_dostring(l, `print'Hello World!\n'`)

	Lua_close(l)
}
