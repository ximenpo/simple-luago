package main

import(
    . "fmt"
    "unsafe"
    . "../lua"
)

/*
extern  int	myGoCFunc(void* l);
extern  int	anotherGoCFunc(void* l);

*/
import  "C"

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

	Lua_register(l, "GoFunc", LuaF_CFunction(C.myGoCFunc))
	LuaL_dostring(l, `GoFunc('world')`)

	Lua_register(l, "GoFunc2", LuaF_CFunction(C.anotherGoCFunc))
	LuaL_dostring(l, `GoFunc2()`)

	Lua_close(l)
}
