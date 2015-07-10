package lua

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

//
//	LuaScript
//
type LuaScript struct {
	handle Lua_State
}

func (s *LuaScript) GetHandle() Lua_State {
	return s.handle
}

func (s *LuaScript) SetHandle(l Lua_State) {
	s.handle = l
}

func (s *LuaScript) OpenStdLibs() {
	if nil != s.handle {
		LuaL_openlibs(s.handle)
	}
}

func (s *LuaScript) HasRef(ref Lua_Ref) bool {
	if !LuaU_GetRef(s.handle, ref) {
		return false
	}

	R := !Lua_isnoneornil(s.handle, -1)
	Lua_pop(s.handle, 1)
	return R
}

func (s *LuaScript) Ref(var_name string) (ref Lua_Ref, ok bool) {
	ref = Lua_Ref(LUA_NOREF)
	if !LuaU_GetGlobal(s.handle, var_name) {
		return ref, false
	}

	if ok = !Lua_isnoneornil(s.handle, -1); ok {
		ref = Lua_Ref(LuaL_ref(s.handle, LUA_REGISTRYINDEX))
	} else {
		Lua_pop(s.handle, 1)
	}
	return
}

func (s *LuaScript) UnRef(ref Lua_Ref) bool {
	LuaL_unref(s.handle, LUA_REGISTRYINDEX, Lua_CInt(ref))

	return true
}

func (s *LuaScript) LoadRef(ref Lua_Ref) bool {
	return LuaU_GetRef(s.handle, ref)
}

func (s *LuaScript) HasVar(var_name string) bool {
	if !LuaU_GetGlobal(s.handle, var_name) {
		return false
	}

	R := !Lua_isnoneornil(s.handle, -1)
	Lua_pop(s.handle, 1)
	return R
}

func (s *LuaScript) RemoveVar(var_name string) bool {
	Lua_pushnil(s.handle)
	if !LuaU_SetGlobal(s.handle, var_name) {
		Lua_pop(s.handle, 1)
		return false
	}
	return true
}

func (s *LuaScript) GetVar(var_name string, value interface{}) bool {
	return s.GetObject(var_name, value, true)
}

func (s *LuaScript) SetVar(var_name string, value interface{}) bool {
	return s.SetObject(var_name, value, false)
}

func (s *LuaScript) GetObject(var_name string, value interface{}, ignore_nonexistent_field bool) bool {
	r := reflect.ValueOf(value)
	if r.Kind() != reflect.Ptr {
		return false
	}

	v := r.Elem()
	if !v.CanSet() {
		return false
	}

	if !LuaU_GetGlobal(s.handle, var_name) {
		return false
	}

	return LuaU_FetchVar(s.handle, value, ignore_nonexistent_field)
}

func (s *LuaScript) SetObject(var_name string, value interface{}, keep_nonexistent_field bool) bool {
	if !keep_nonexistent_field {
		s.RemoveVar(var_name)
	}

	if !LuaU_PushVar(s.handle, value) {
		return false
	}

	if !LuaU_SetGlobal(s.handle, var_name) {
		Lua_pop(s.handle, 1)
		return false
	}

	return true
}

func (s *LuaScript) Call(func_name string, args ...interface{}) error {
	if !LuaU_GetGlobal(s.handle, func_name) {
		return errors.New("can't find function " + func_name)
	}

	for i := 0; i < len(args); i++ {
		if !LuaU_PushVar(s.handle, args[i]) {
			Lua_pop(s.handle, Lua_CInt(i+1)) // 0, 1, .. i - 1, + LuaU_GetGlobal
			return errors.New(fmt.Sprintf("push param [%d] failed", i))
		}
	}

	var err_msg string
	if !LuaU_InvokeFunc(s.handle, len(args), 0, nil, &err_msg) {
		return errors.New(err_msg)
	}

	return nil
}

func (s *LuaScript) Invoke(ret_value interface{}, func_name string, args ...interface{}) error {
	if !LuaU_GetGlobal(s.handle, func_name) {
		return errors.New("can't find function " + func_name)
	}

	for i := 0; i < len(args); i++ {
		if !LuaU_PushVar(s.handle, args[i]) {
			Lua_pop(s.handle, Lua_CInt(i+1)) // 0, 1, .. i - 1, + LuaU_GetGlobal
			return errors.New(fmt.Sprintf("push param [%d] failed", i))
		}
	}

	retsum := 0
	if nil != ret_value {
		retsum = 1
	}

	var err_msg string
	if !LuaU_InvokeFunc(s.handle, len(args), retsum, nil, &err_msg) {
		return errors.New(err_msg)
	}

	if (nil != ret_value) && !LuaU_FetchVar(s.handle, ret_value, true) {
		return errors.New("fetch function result failed")
	}

	return nil
}

func (s *LuaScript) RunFile(file string) error {
	var err_msg string
	if R := LuaL_loadfile(s.handle, file); LUA_OK != R {
		err_msg = Lua_tostring(s.handle, -1)
		Lua_pop(s.handle, 1)
		return errors.New(err_msg)
	}

	if !LuaU_InvokeFunc(s.handle, 0, int(LUA_MULTRET), nil, &err_msg) {
		return errors.New(err_msg)
	}

	return nil
}

func (s *LuaScript) RunString(code string) error {
	var err_msg string
	if R := LuaL_loadstring(s.handle, code); LUA_OK != R {
		err_msg = Lua_tostring(s.handle, -1)
		Lua_pop(s.handle, 1)
		return errors.New(err_msg)
	}

	if !LuaU_InvokeFunc(s.handle, 0, int(LUA_MULTRET), nil, &err_msg) {
		return errors.New(err_msg)
	}

	return nil
}

func (s *LuaScript) RunBuffer(buffer unsafe.Pointer, size uint) error {
	var err_msg string
	if LUA_OK == LuaL_loadbuffer(s.handle, uintptr(buffer), size, "LuaScript.RunBuffer") {
		if !LuaU_InvokeFunc(s.handle, 0, int(LUA_MULTRET), nil, &err_msg) {
			return errors.New(err_msg)
		}
		return nil
	}

	err_msg = Lua_tostring(s.handle, -1)
	Lua_pop(s.handle, 1)
	return errors.New(err_msg)
}

//
//	LuaVM -> LuaScript
//
type LuaVM struct {
	LuaScript
}

func NewLuaVM() *LuaVM {
	vm := &LuaVM{}
	return vm
}

func (vm *LuaVM) Start() {
	if vm.handle != nil {
		Lua_close(vm.handle)
	}

	vm.handle = LuaL_newstate()
}

func (vm *LuaVM) Stop() {
	if vm.handle != nil {
		Lua_close(vm.handle)
	}
}
