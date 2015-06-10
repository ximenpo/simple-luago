package lua

/*
#cgo darwin CFLAGS: -DLUA_USE_MACOSX
#cgo darwin LDFLAGS: -lreadline

#cgo linux  CFLAGS: -DLUA_USE_LINUX
#cgo linux  LDFLAGS: -Wl,-E -ldl -lreadline

*/
import  "C"

import  (
    "unsafe"
)

type    Lua_Alloc       func(ud unsafe.Pointer, ptr unsafe.Pointer, osize C.size_t, nsize C.size_t) unsafe.Pointer
type    Lua_CFunction   func(L unsafe.Pointer) int32

func LuaF_Alloc(fp unsafe.Pointer) (*_swig_fnptr) {
    return  LuaF_AsPtr(uintptr(fp))
}

func LuaF_CFunction(fp unsafe.Pointer) (*_swig_fnptr) {
    return  LuaF_AsPtr(uintptr(fp))
}

func luaF_Reader(fp unsafe.Pointer) (*_swig_fnptr) {
    return  LuaF_AsPtr(uintptr(fp))
}

func luaF_Writer(fp unsafe.Pointer) (*_swig_fnptr) {
    return  LuaF_AsPtr(uintptr(fp))
}

func luaF_Hook(fp unsafe.Pointer) (*_swig_fnptr) {
    return  LuaF_AsPtr(uintptr(fp))
}
