package main

import (
	. "../lua"
	. "fmt"
	"unsafe"
)

/*
extern  int	myGoCFunc(void* l);
extern  int	anotherGoCFunc(void* l);

typedef int	(*lua_CFunction)(void* l);
typedef struct luaL_Reg {
  const char *name;
  lua_CFunction func;
} luaL_Reg;

static  const void* lfuncs(){
    static  const   luaL_Reg   lfs[]  = {
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

	Lua_createtable(l, 0, 2)
	LuaL_setfuncs(l, C.lfuncs(), 0)
	Lua_setglobal(l, "GoLib")

	LuaL_dostring(l, `GoLib.myGoCFunc('world')`)
	LuaL_dostring(l, `GoLib.anotherGoCFunc()`)

	Lua_close(l)
}
