package lua

/*
#include	<stdlib.h>
#include	"lua_header.h"
#include	"lua_utils.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

// utils.

func LuaU_GetRef(L Lua_State, refid Lua_Ref) bool {
	return C.luaU_GetRef(LuaF_StateCPtr(L), C.int(refid)) != 0
}

func LuaU_GetGlobal(L Lua_State, varname string) bool {
	s := C.CString(varname)
	defer C.free(unsafe.Pointer(s))

	r := C.luaU_GetGlobal(LuaF_StateCPtr(L), s)
	return (0 != r)
}

func LuaU_SetGlobal(L Lua_State, varname string) bool {
	s := C.CString(varname)
	defer C.free(unsafe.Pointer(s))

	r := C.luaU_SetGlobal(LuaF_StateCPtr(L), s)
	return (0 != r)
}

func LuaU_InvokeFunc(L Lua_State, nargs int, nresults int, err_code *int, err_msg *string) bool {
	var ret int
	if nil != err_msg {
		*err_msg = ""
	}
	if nil == err_code {
		err_code = &ret
	}

	switch *err_code = int(Lua_pcall(L, Lua_CInt(nargs), Lua_CInt(nresults), 0)); Lua_CInt(*err_code) {
	case LUA_OK, LUA_YIELD:
		{
			return true
		}
	default:
		{
			if nil != err_msg {
				*err_msg = Lua_tostring(L, -1)
			}
			Lua_pop(L, 1)
		}
	}
	return false
}

func LuaU_PushVar(L Lua_State, value interface{}) bool {
	r := reflect.ValueOf(value)
	if r.Kind() == reflect.Ptr {
		v := r.Elem()
		if !v.CanSet() {
			return false
		}

		return LuaU_PushValue(L, &v)
	} else {
		return LuaU_PushValue(L, &r)
	}
	return false
}

func LuaU_FetchVar(L Lua_State, value interface{}, ignore_nonexistent_field bool) bool {
	r := reflect.ValueOf(value)
	if r.Kind() != reflect.Ptr {
		Lua_pop(L, 1)
		return false
	}

	v := r.Elem()
	if !v.CanSet() {
		Lua_pop(L, 1)
		return false
	}

	return LuaU_FetchValue(L, &v, ignore_nonexistent_field)
}

func LuaU_FetchValue(L Lua_State, value *reflect.Value, ignore_nonexistent_field bool) bool {
	if !value.CanSet() {
		Lua_pop(L, 1)
		return false
	}

	if Lua_isnil(L, -1) {
		Lua_pop(L, 1)
		return false
	}

	bValid := false
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bValid = luaU_FetchInt(L, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bValid = luaU_FetchUint(L, value)
	case reflect.Float32, reflect.Float64:
		bValid = luaU_FetchFloat(L, value)
	case reflect.Bool:
		bValid = luaU_FetchBool(L, value)
	case reflect.String:
		bValid = luaU_FetchString(L, value)
	case reflect.Ptr:
		bValid = luaU_FetchPointer(L, value)
	case reflect.Struct:
		if !Lua_istable(L, -1) {
			Lua_pop(L, 1)
			break
		}
		bValid = true
		t := value.Type()
		for i := 0; i < value.NumField(); i++ {
			item := value.Field(i)
			if !item.CanSet() {
				continue
			}
			Lua_pushstring(L, t.Field(i).Name)
			Lua_gettable(L, -2)
			if !LuaU_FetchValue(L, &item, ignore_nonexistent_field) {
				if !ignore_nonexistent_field {
					bValid = false
					break //for
				}
			}
		}
		Lua_pop(L, 1)
	case reflect.Slice:
		if !Lua_istable(L, -1) {
			Lua_pop(L, 1)
			break
		}
		bValid = true
		l := int(Lua_rawlen(L, -1))
		tempRfV := reflect.MakeSlice(value.Type(), 0, l) //*value
		for i := 0; i < l; i++ {
			item := reflect.New(tempRfV.Type().Elem()).Elem()
			Lua_rawgeti(L, -1, Lua_Integer(i+1))
			if !LuaU_FetchValue(L, &item, ignore_nonexistent_field) {
				if !ignore_nonexistent_field {
					bValid = false
					break //for
				}
			}
			tempRfV = reflect.Append(tempRfV, item)
		}
		Lua_pop(L, 1)
		value.Set(tempRfV)
	case reflect.Array:
		if !Lua_istable(L, -1) {
			Lua_pop(L, 1)
			break
		}
		bValid = true
		l := value.Len()
		for i := 0; i < l; i++ {
			Lua_rawgeti(L, -1, Lua_Integer(i+1))
			item := value.Index(i)
			if !LuaU_FetchValue(L, &item, ignore_nonexistent_field) {
				if !ignore_nonexistent_field {
					bValid = false
					break //for
				}
			}
		}
		Lua_pop(L, 1)
	case reflect.Map:
		if !Lua_istable(L, -1) {
			Lua_pop(L, 1)
			break
		}
		bValid = true

		if value.IsNil() {
			value.Set(reflect.MakeMap(value.Type()))
		}

		Lua_pushnil(L)
		for 0 != Lua_next(L, -2) {
			Lua_pushvalue(L, -2) // table key value key

			k := reflect.New(value.Type().Key()).Elem()
			if !LuaU_FetchValue(L, &k, ignore_nonexistent_field) {
				if !ignore_nonexistent_field {
					Lua_pop(L, 2)
					bValid = false
					break //for
				}
				Lua_pop(L, 1)
				continue
			}
			v := reflect.New(value.Type().Elem()).Elem()
			if !LuaU_FetchValue(L, &v, ignore_nonexistent_field) {
				if !ignore_nonexistent_field {
					Lua_pop(L, 1)
					bValid = false
					break //for
				}
				continue
			}
			value.SetMapIndex(k, v)
		}
		Lua_pop(L, 1)
	default:
		bValid = luaU_FetchFailed(L, value)
	}
	return bValid
}

func LuaU_PushValue(L Lua_State, value *reflect.Value) bool {
	if !value.IsValid() {
		return false
	}

	ok := true
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if r, ok := value.Interface().(Lua_Ref); ok {
			LuaU_GetRef(L, r)
		} else {
			Lua_pushinteger(L, Lua_Integer(value.Int()))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		Lua_pushinteger(L, Lua_Integer(value.Uint()))
	case reflect.Float32, reflect.Float64:
		Lua_pushnumber(L, Lua_Number(value.Float()))
	case reflect.Bool:
		Lua_pushboolean(L, value.Bool())
	case reflect.String:
		Lua_pushstring(L, value.String())
	case reflect.Struct:
		Lua_newtable(L)
		t := value.Type()
		for i := 0; i < value.NumField(); i++ {
			item := value.Field(i)
			if !item.CanSet() {
				continue
			}
			Lua_pushstring(L, t.Field(i).Name)
			if !LuaU_PushValue(L, &item) {
				Lua_pop(L, 2)
				ok = false
				break
			}
			Lua_settable(L, -3)
		}
	case reflect.Slice, reflect.Array:
		Lua_newtable(L)
		l := value.Len()
		for i := 0; i < l; i++ {
			v := value.Index(i)
			if !LuaU_PushValue(L, &v) {
				Lua_pop(L, 1)
				ok = false
				break
			}
			Lua_rawseti(L, -2, Lua_Integer(i+1)) // go begin with 0, lua table is 1
		}
	case reflect.Map:
		Lua_newtable(L)
		keys := value.MapKeys()
		l := len(keys)
		for i := 0; i < l; i++ {
			k := keys[i]
			v := value.MapIndex(k)
			if !LuaU_PushValue(L, &k) {
				Lua_pop(L, 1)
				ok = false
				break
			}
			if !LuaU_PushValue(L, &v) {
				Lua_pop(L, 2)
				ok = false
				break
			}
			Lua_settable(L, -3)
		}
	default:
		{
			v := value.Interface()
			switch v.(type) {
			case Lua_State:
				Lua_pushthread(L)
			case Lua_CFunction:
				Lua_pushcfunction(L, v.(Lua_CFunction))
			default:
				ok = false
			}
		}
	}
	return ok
}

/////////////////////////////////////////////////////////////////
//	internal impls.
/////////////////////////////////////////////////////////////////

func luaU_FetchInt(l Lua_State, v *reflect.Value) (ok bool) {
	if ok = Lua_isnumber(l, -1); ok {
		v.SetInt(int64(Lua_tointeger(l, -1)))
	}
	Lua_pop(l, 1)
	return
}

func luaU_FetchUint(l Lua_State, v *reflect.Value) (ok bool) {
	if ok = Lua_isnumber(l, -1); ok {
		v.SetUint(uint64(Lua_tointegerx(l, -1, nil)))
	}
	Lua_pop(l, 1)
	return
}

func luaU_FetchFloat(l Lua_State, v *reflect.Value) (ok bool) {
	if ok = Lua_isnumber(l, -1); ok {
		v.SetFloat(float64(Lua_tonumber(l, -1)))
	}
	Lua_pop(l, 1)
	return
}

func luaU_FetchBool(l Lua_State, v *reflect.Value) (ok bool) {
	if ok = Lua_isboolean(l, -1); ok {
		v.SetBool(Lua_toboolean(l, -1))
	}
	Lua_pop(l, 1)
	return
}

func luaU_FetchString(l Lua_State, v *reflect.Value) (ok bool) {
	if ok = Lua_isstring(l, -1); ok {
		v.SetString(Lua_tostring(l, -1))
	}
	Lua_pop(l, 1)
	return
}

func luaU_FetchPointer(l Lua_State, v *reflect.Value) (ok bool) {
	if ok = Lua_isthread(l, -1); ok {
		v.SetPointer(unsafe.Pointer(Lua_tothread(l, -1).Swigcptr()))
	}
	Lua_pop(l, 1)
	return
}

func luaU_FetchFailed(l Lua_State, v *reflect.Value) (ok bool) {
	ok = false
	Lua_pop(l, 1)
	return
}

func luaU_FetchParams(L Lua_State, ignore_nonexistent_field bool, args ...interface{}) (ok bool) {
	ok = true
	for i := len(args) - 1; i >= 0; i-- {
		if !LuaU_FetchVar(L, args[i], ignore_nonexistent_field) {
			ok = false
			break
		}
	}
	return
}
