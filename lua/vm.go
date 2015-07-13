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
	Handle Lua_State
}

func (self *LuaScript) OpenStdLibs() {
	if nil != self.Handle {
		LuaL_openlibs(self.Handle)
	}
}

func (self *LuaScript) HasRef(ref Lua_Ref) bool {
	if !LuaU_GetRef(self.Handle, ref) {
		return false
	}

	R := !Lua_isnoneornil(self.Handle, -1)
	Lua_pop(self.Handle, 1)
	return R
}

func (self *LuaScript) Ref(var_name string) (ref Lua_Ref, ok bool) {
	ref = Lua_Ref(LUA_NOREF)
	if !LuaU_GetGlobal(self.Handle, var_name) {
		return ref, false
	}

	if ok = !Lua_isnoneornil(self.Handle, -1); ok {
		ref = Lua_Ref(LuaL_ref(self.Handle, LUA_REGISTRYINDEX))
	} else {
		Lua_pop(self.Handle, 1)
	}
	return
}

func (self *LuaScript) UnRef(ref Lua_Ref) bool {
	LuaL_unref(self.Handle, LUA_REGISTRYINDEX, Lua_CInt(ref))

	return true
}

func (self *LuaScript) LoadRef(ref Lua_Ref) bool {
	return LuaU_GetRef(self.Handle, ref)
}

func (self *LuaScript) HasVar(var_name string) bool {
	if !LuaU_GetGlobal(self.Handle, var_name) {
		return false
	}

	R := !Lua_isnoneornil(self.Handle, -1)
	Lua_pop(self.Handle, 1)
	return R
}

func (self *LuaScript) RemoveVar(var_name string) bool {
	Lua_pushnil(self.Handle)
	if !LuaU_SetGlobal(self.Handle, var_name) {
		Lua_pop(self.Handle, 1)
		return false
	}
	return true
}

func (self *LuaScript) GetVar(var_name string, value interface{}) bool {
	return self.GetObject(var_name, value, true)
}

func (self *LuaScript) SetVar(var_name string, value interface{}) bool {
	return self.SetObject(var_name, value, false)
}

func (self *LuaScript) GetObject(var_name string, value interface{}, ignore_nonexistent_field bool) bool {
	r := reflect.ValueOf(value)
	if r.Kind() != reflect.Ptr {
		return false
	}

	v := r.Elem()
	if !v.CanSet() {
		return false
	}

	if !LuaU_GetGlobal(self.Handle, var_name) {
		return false
	}

	return LuaU_FetchVar(self.Handle, value, ignore_nonexistent_field)
}

func (self *LuaScript) SetObject(var_name string, value interface{}, keep_nonexistent_field bool) bool {
	if !keep_nonexistent_field {
		self.RemoveVar(var_name)
	}

	if !LuaU_PushVar(self.Handle, value) {
		return false
	}

	if !LuaU_SetGlobal(self.Handle, var_name) {
		Lua_pop(self.Handle, 1)
		return false
	}

	return true
}

func (self *LuaScript) Call(func_name string, args ...interface{}) (err error) {
	if !LuaU_GetGlobal(self.Handle, func_name) {
		return errors.New("can't find function " + func_name)
	}

	for i := 0; i < len(args); i++ {
		if !LuaU_PushVar(self.Handle, args[i]) {
			Lua_pop(self.Handle, Lua_CInt(i+1)) // 0, 1, .. i - 1, + LuaU_GetGlobal
			return errors.New(fmt.Sprintf("push param [%d] failed", i))
		}
	}

	_, err = LuaU_InvokeFunc(self.Handle, len(args), 0)
	return
}

func (self *LuaScript) Invoke(ret_value interface{}, func_name string, args ...interface{}) (err error) {
	if !LuaU_GetGlobal(self.Handle, func_name) {
		return errors.New("can't find function " + func_name)
	}

	for i := 0; i < len(args); i++ {
		if !LuaU_PushVar(self.Handle, args[i]) {
			Lua_pop(self.Handle, Lua_CInt(i+1)) // 0, 1, .. i - 1, + LuaU_GetGlobal
			return errors.New(fmt.Sprintf("push param [%d] failed", i))
		}
	}

	retsum := 0
	if nil != ret_value {
		retsum = 1
	}

	if _, err = LuaU_InvokeFunc(self.Handle, len(args), retsum); err != nil {
		return
	}

	if (nil != ret_value) && !LuaU_FetchVar(self.Handle, ret_value, true) {
		return errors.New("fetch function result failed")
	}

	return nil
}

func (self *LuaScript) RunFile(file string) (err error) {
	if R := LuaL_loadfile(self.Handle, file); LUA_OK != R {
		err = errors.New(Lua_tostring(self.Handle, -1))
		Lua_pop(self.Handle, 1)
		return
	}

	_, err = LuaU_InvokeFunc(self.Handle, 0, int(LUA_MULTRET))
	return
}

func (self *LuaScript) RunString(code string) (err error) {
	if R := LuaL_loadstring(self.Handle, code); LUA_OK != R {
		err = errors.New(Lua_tostring(self.Handle, -1))
		Lua_pop(self.Handle, 1)
		return
	}

	_, err = LuaU_InvokeFunc(self.Handle, 0, int(LUA_MULTRET))
	return
}

func (self *LuaScript) RunBuffer(buffer unsafe.Pointer, size uint) (err error) {
	if LUA_OK == LuaL_loadbuffer(self.Handle, uintptr(buffer), size, "LuaScript.RunBuffer") {
		_, err = LuaU_InvokeFunc(self.Handle, 0, int(LUA_MULTRET))
		return
	}

	err = errors.New(Lua_tostring(self.Handle, -1))
	Lua_pop(self.Handle, 1)
	return
}

//
//	LuaVM -> LuaScript
//
type LuaVM struct {
	LuaScript
}

func NewLuaVM() *LuaVM {
	self := &LuaVM{}
	return self
}

func (self *LuaVM) Start() {
	if self.Handle != nil {
		Lua_close(self.Handle)
	}

	self.Handle = LuaL_newstate()
}

func (self *LuaVM) Stop() {
	if self.Handle != nil {
		Lua_close(self.Handle)
	}
}
