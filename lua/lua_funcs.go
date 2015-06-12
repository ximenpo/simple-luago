package lua

/*
#cgo darwin CFLAGS: -DLUA_USE_MACOSX
#cgo darwin LDFLAGS: -lreadline

#cgo linux  CFLAGS: -DLUA_USE_LINUX
#cgo linux  LDFLAGS: -Wl,-E -ldl -lreadline

#include    "$.lua.h"
#include    "$.lualib.h"
#include    "$.lauxlib.h"

*/
import "C"

import (
	"unsafe"
)

func LuaF_Handle(lp unsafe.Pointer) (ret SwigcptrStruct_SS_lua_State) {
	ret = SwigcptrStruct_SS_lua_State(uintptr(lp))
	return
}

//
//  Lua callback types.
//

//extern    void*   myGoAlloc(void *ud, void *ptr, size_t osize, size_t nsize);
//func              myGoAlloc(ud unsafe.Pointer, ptr unsafe.Pointer, osize uintptr, nsize uintptr) unsafe.Pointer
func LuaF_Alloc(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//extern    int	    myGoCFunc(void* l);
//tfunc             myGoCFunc(L unsafe.Pointer) int32
func LuaF_CFunction(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//extern    const char* myGoReader(void *L, void *ud, size_t *sz);
//func                  myGoReader(L unsafe.Pointer, ud unsafe.Pointer, sz uintptr)unsafe.Pointer
func LuaF_Reader(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//extern    int     myGoWriter(void *L, void* p, size_t sz, void* ud);
//func              myGoWriter(L unsafe.Pointer, p unsafe.Pointer, sz uint, ud unsafe.Pointer)int32
func LuaF_Writer(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}

//extern    void    myGoHook(lua_State *L, lua_Debug *ar);
//func              myGoHook(L unsafe.Pointer, ar unsafe.Pointer)
func LuaF_Hook(fp unsafe.Pointer) *_swig_fnptr {
	return LuaF_Ptr(uintptr(fp))
}
