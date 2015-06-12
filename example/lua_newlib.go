package main

import (
	. "../lua"
	. "fmt"
	"unsafe"
)

/*
extern  int	myGoCFunc(void* l);
extern  int	anotherGoCFunc(void* l);

static  const void* lfuncs(){
	static	const	struct{void *name, *func;}	lfs[]	= {
        {"myGoCFunc",       myGoCFunc},
        {"anotherGoCFunc",  anotherGoCFunc},
        {0,                 0},
    };
    return  lfs;
}

*/
import "C"

//export myGoCFunc
func myGoCFunc(L unsafe.Pointer) int32 {
	l := LuaF_Handle(L)
	Println("Hello, " + LuaL_checkstring(l, 1) + "! call from GO->myGoCFunc")
	return 0
}

//export anotherGoCFunc
func anotherGoCFunc(L unsafe.Pointer) int32 {
	Println("Aha, call from GO->anotherGoCFunc")
	return 0
}

func main() {
	l := LuaL_newstate()
	LuaL_openlibs(l)

	LuaL_newlib(l, C.lfuncs())
	Lua_setglobal(l, "GoLib")

	LuaL_dostring(l, `GoLib.myGoCFunc('world')`)
	LuaL_dostring(l, `GoLib.anotherGoCFunc()`)

	Lua_close(l)
}
