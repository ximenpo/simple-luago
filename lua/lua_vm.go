package lua

import (
	"reflect"
	"unsafe"
)

//
//	LuaVM
//
type LuaVM struct {
	handle  Lua_State
	err_msg string
}

func NewLuaVM() *LuaVM {
	vm := &LuaVM{}
	return vm
}

func (vm *LuaVM) GetHandle() Lua_State {
	return vm.handle
}

func (vm *LuaVM) SetHandle(l Lua_State) {
	vm.handle = l
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

// called must after dostring / dofile
func (vm *LuaVM) Error() string {
	return vm.err_msg
}

func (vm *LuaVM) OpenStdLibs() {
	if nil != vm.handle {
		LuaL_openlibs(vm.handle)
	}
}

func (vm *LuaVM) HasRef(ref Lua_Ref) bool {
	if !LuaU_GetRef(vm.handle, ref) {
		return false
	}

	R := !Lua_isnoneornil(vm.handle, -1)
	Lua_pop(vm.handle, 1)
	return R
}

func (vm *LuaVM) Ref(var_name string) (ref Lua_Ref, ok bool) {
	ref	= Lua_Ref(LUA_NOREF)
	if !LuaU_GetGlobal(vm.handle, var_name) {
		return ref, false
	}

	if ok = !Lua_isnoneornil(vm.handle, -1); ok {
		ref = Lua_Ref(LuaL_ref(vm.handle, LUA_REGISTRYINDEX))
	} else {
		Lua_pop(vm.handle, 1)
	}
	return
}

func (vm *LuaVM) UnRef(ref Lua_Ref) bool {
	LuaL_unref(vm.handle, LUA_REGISTRYINDEX, int(ref))

	return true
}

func (vm *LuaVM) LoadRef(ref Lua_Ref) bool {
	return LuaU_GetRef(vm.handle, ref)
}

func (vm *LuaVM) HasVar(var_name string) bool {
	if !LuaU_GetGlobal(vm.handle, var_name) {
		return false
	}

	R := !Lua_isnoneornil(vm.handle, -1)
	Lua_pop(vm.handle, 1)
	return R
}

func (vm *LuaVM) RemoveVar(var_name string) {
	Lua_pushnil(vm.handle)
	if !LuaU_SetGlobal(vm.handle, var_name) {
		Lua_pop(vm.handle, 1)
	}
}

func (vm *LuaVM) GetVar(var_name string, value interface{}) bool {
	return vm.GetObject(var_name, value, true)
}

func (vm *LuaVM) SetVar(var_name string, value interface{}) bool {
	return vm.SetObject(var_name, value, false)
}

func (vm *LuaVM) GetObject(var_name string, value interface{}, ignore_nonexistent_field bool) bool {
	r := reflect.ValueOf(value)
	if r.Kind() != reflect.Ptr {
		return false
	}

	v := r.Elem()
	if !v.CanSet() {
		return false
	}

	if !LuaU_GetGlobal(vm.handle, var_name) {
		return false
	}

	return LuaU_FetchVar(vm.handle, value, ignore_nonexistent_field)
}

func (vm *LuaVM) SetObject(var_name string, value interface{}, keep_nonexistent_field bool) bool {
	if !keep_nonexistent_field {
		vm.RemoveVar(var_name)
	}

	if !LuaU_PushVar(vm.handle, value) {
		return false
	}

	if !LuaU_SetGlobal(vm.handle, var_name) {
		Lua_pop(vm.handle, 1)
		return false
	}

	return true
}

func (vm *LuaVM) Call(func_name string, args ...interface{}) bool {
	if !LuaU_GetGlobal(vm.handle, func_name) {
		return false
	}

	for i := 0; i < len(args); i++ {
		if !LuaU_PushVar(vm.handle, args[i]) {
			Lua_pop(vm.handle, i+1) // 0, 1, .. i - 1, + LuaU_GetGlobal
			return false
		}
	}

	return LuaU_InvokeFunc(vm.handle, len(args), 0, nil, &vm.err_msg)
}

func (vm *LuaVM) Invoke(ret_value interface{}, func_name string, args ...interface{}) bool {
	if !LuaU_GetGlobal(vm.handle, func_name) {
		return false
	}

	for i := 0; i < len(args); i++ {
		if !LuaU_PushVar(vm.handle, args[i]) {
			Lua_pop(vm.handle, i+1) // 0, 1, .. i - 1, + LuaU_GetGlobal
			return false
		}
	}

	retsum := 0
	if nil != ret_value {
		retsum = 1
	}

	if !LuaU_InvokeFunc(vm.handle, len(args), retsum, nil, &vm.err_msg) {
		return false
	}

	if nil != ret_value {
		return LuaU_FetchVar(vm.handle, ret_value, true)
	}

	return true
}

func (vm *LuaVM) RunFile(file string) bool {
	if R := LuaL_loadfile(vm.handle, file); LUA_OK != R {
		vm.err_msg = Lua_tostring(vm.handle, -1)
		Lua_pop(vm.handle, 1)
		return false
	}
	return LuaU_InvokeFunc(vm.handle, 0, LUA_MULTRET, nil, &vm.err_msg)
}

func (vm *LuaVM) RunString(code string) bool {
	if R := LuaL_loadstring(vm.handle, code); LUA_OK != R {
		vm.err_msg = Lua_tostring(vm.handle, -1)
		Lua_pop(vm.handle, 1)
		return false
	}
	return LuaU_InvokeFunc(vm.handle, 0, LUA_MULTRET, nil, &vm.err_msg)
}

func (vm *LuaVM) RunBuffer(buffer unsafe.Pointer, size uint) bool {
	if LUA_OK == LuaL_loadbuffer(vm.handle, uintptr(buffer), size, "LuaVM.RunBuffer") {
		return LuaU_InvokeFunc(vm.handle, 0, LUA_MULTRET, nil, &vm.err_msg)
	}
	vm.err_msg = Lua_tostring(vm.handle, -1)
	Lua_pop(vm.handle, 1)
	return false
}
