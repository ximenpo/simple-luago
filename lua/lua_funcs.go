package lua

/*
#cgo darwin CFLAGS: -DLUA_USE_MACOSX
#cgo darwin LDFLAGS: -lreadline

#cgo linux  CFLAGS: -DLUA_USE_LINUX
#cgo linux  LDFLAGS: -Wl,-E -ldl -lreadline

#include    "$.lua.h"
#include    "$.lualib.h"
#include    "$.lauxlib.h"

static	int	_LuaF_RegLen(void* p){
	int	len	= 0;
	if(p != 0){
		luaL_Reg* r	= (luaL_Reg*)p;
		while(r->name && r->func){
			len++;
			r++;
		}
	}
	return	len;
}

*/
import "C"

import (
	"unsafe"
)

func LuaL_setfuncs(L Struct_SS_lua_State, l unsafe.Pointer, nup int) {
	C.luaL_setfuncs((*C.lua_State)(unsafe.Pointer(L.Swigcptr())), (*C.luaL_Reg)(l), C.int(nup))
}

func LuaL_newlibtable(L Struct_SS_lua_State, l unsafe.Pointer) {
	C.lua_createtable((*C.lua_State)(unsafe.Pointer(L.Swigcptr())), 0, LuaF_RegLen(l))
}

func LuaL_newlib(L Struct_SS_lua_State, l unsafe.Pointer) {
	LuaL_checkversion(L)
	LuaL_newlibtable(L, l)
	LuaL_setfuncs(L, l, 0)
}

//
// convert unsafe.Pointer to Lua_State
//
func LuaF_Handle(lp unsafe.Pointer) (ret SwigcptrStruct_SS_lua_State) {
	ret = SwigcptrStruct_SS_lua_State(uintptr(lp))
	return
}

func LuaF_RegLen(p unsafe.Pointer) C.int  {
	return	C._LuaF_RegLen(p)
}

//
//  Lua callback types.
//

//
// convert unsafe.Pointer to Lua_Alloc
//
//extern    void*   myGoAlloc(void *ud, void *ptr, size_t osize, size_t nsize);
//func              myGoAlloc(ud unsafe.Pointer, ptr unsafe.Pointer, osize uintptr, nsize uintptr) unsafe.Pointer
func LuaF_Alloc(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//
// convert unsafe.Pointer to Lua_CFunction
//
//extern    int	    myGoCFunc(void* l);
//tfunc             myGoCFunc(L unsafe.Pointer) int32
func LuaF_CFunction(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//
// convert unsafe.Pointer to Lua_Reader
//
//extern    const char* myGoReader(void *L, void *ud, size_t *sz);
//func                  myGoReader(L unsafe.Pointer, ud unsafe.Pointer, sz uintptr)unsafe.Pointer
func LuaF_Reader(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//
// convert unsafe.Pointer to Lua_Writer
//
//extern    int     myGoWriter(void *L, void* p, size_t sz, void* ud);
//func              myGoWriter(L unsafe.Pointer, p unsafe.Pointer, sz uintptr, ud unsafe.Pointer)int32
func LuaF_Writer(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//
// convert unsafe.Pointer to Lua_Hook
//
//extern    void    myGoHook(void *L, void *ar);
//func              myGoHook(L unsafe.Pointer, ar unsafe.Pointer)
func LuaF_Hook(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}
