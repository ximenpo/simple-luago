package main

import (
	"../lua"
	"fmt"
)

func callfunction(f func(s *lua.LuaVM)) {
	s := lua.LuaVM{}

	s.Start()
	s.OpenStdLibs()

	f(&s)

	s.Stop()
}

func luascript_PrintEmptyErrorMessage(s *lua.LuaVM) {
	msg := s.Error()
	if 0 != len(msg) {
		fmt.Println("ERROR -> must not has error message", msg)
	}
}

func luascript_PrintErrorMessage(s *lua.LuaVM) {
	s.RunString("dads;jfaslkdfjaslkfj")
	msg := s.Error()
	if 0 == len(msg) {
		fmt.Println("ERROR -> must has error message")
	}
}

func luascript_HasVariable(s *lua.LuaVM) {
	if s.HasVar("g_data.name.type") {
		fmt.Println("ERROR => must has not variable")
	}

	lua.LuaL_dostring(s.GetHandle(), "g_data = {}; g_data.name={}; g_data.name.type = 'simple';")
	if !s.HasVar("g_data.name.type") {
		fmt.Println("ERROR => must has variable")
	}

	s.RemoveVar("g_data.name.type")
	if s.HasVar("g_data.name.type") {
		fmt.Println("ERROR => must has not variable")
	}
}

func luascript_Reference(s *lua.LuaVM) {
	lua.LuaL_dostring(s.GetHandle(), "g_data = {}; g_data.name={}; g_data.name.type = 'simple';")

	var ok bool
	var ref lua.Lua_Ref
	if ref, ok = s.Ref("g_data.name.type"); ok {
		fmt.Println("ERROR => must has reference")
	}

	//fmt.Println("luascript_VariableReference: REF -> ", REF)

	s.RemoveVar("g_data.name.type")
	if !s.HasRef(ref) {
		fmt.Println("ERROR => must NOT has reference")
	}
}

func luascript_Variable(s *lua.LuaVM) {
	lua.LuaL_dostring(s.GetHandle(), "g_data = {}; g_data.name={}; g_data.name.type = 'simple';")

	var n int = 100
	var u uint = 200
	var f float32 = 200.5
	var str string = "STR"

	if !s.GetVar("g_data.name.type", &str) {
		fmt.Println("ERROR: read string")
	}
	if "simple" != str {
		fmt.Println("ERROR: wrong string value => ", str)
	}

	s.SetVar("g_data.name.type", "eddy")
	if !s.GetVar("g_data.name.type", &str) {
		fmt.Println("ERROR: XX read string")
	}
	if "eddy" != str {
		fmt.Println("ERROR: XX wrong string value => ", str)
	}

	s.SetVar("g_data.name.age", -1000)
	if !s.GetVar("g_data.name.age", &n) {
		fmt.Println("ERROR: XX read int")
	}
	if -1000 != n {
		fmt.Println("ERROR: XX wrong int value => ", n)
	}

	s.SetVar("g_data.name.height", uint(2000))
	if !s.GetVar("g_data.name.height", &u) {
		fmt.Println("ERROR: XX read uint")
	}
	if 2000 != u {
		fmt.Println("ERROR: XX wrong uint value => ", u)
	}

	s.SetVar("g_data.name.weight", 2000.5)
	if !s.GetVar("g_data.name.weight", &f) {
		fmt.Println("ERROR: XX read float")
	}
	if (2000.5-f) < -0.0005 || (2000.5-f) > 0.0005 {
		fmt.Println("ERROR: XX wrong float value => ", f)
	}
}

func luascript_CallFunction(s *lua.LuaVM) {
	lua.LuaL_dostring(s.GetHandle(), "g_data = {}; g_data.f = function(name) g_data.name = name; end;")
	msg := s.Error()
	if 0 != len(msg) {
		fmt.Println("ERROR -> must not has error message", msg)
	}

	if !s.HasVar("g_data.f") {
		fmt.Println("ERROR => must has variable")
	}

	if !s.Call("g_data.f", "simple") {
		fmt.Println("ERROR: call function error")
	}

	var str string
	if !s.GetVar("g_data.name", &str) {
		fmt.Println("ERROR: call function -> read string")
	}
	if "simple" != str {
		fmt.Println("ERROR: call function -> wrong string value => ", str)
	}

}

func luascript_InvokeFunction(s *lua.LuaVM) {
	lua.LuaL_dostring(s.GetHandle(), "g_data = {}; g_data.f = function(name) return 'Hello, '..name; end;")
	msg := s.Error()
	if 0 != len(msg) {
		fmt.Println("ERROR -> must not has error message", msg)
	}

	if !s.HasVar("g_data.f") {
		fmt.Println("ERROR => must has variable")
	}

	var str string
	if !s.Invoke(&str, "g_data.f", "simple") {
		fmt.Println("ERROR: invoke function error")
	}
	if "Hello, simple" != str {
		fmt.Println("ERROR: invoke function -> wrong string value => ", str)
	}

	if !s.Invoke(nil, "g_data.f", "simple") {
		fmt.Println("ERROR: nil invoke function error")
	}
}

func main() {
	callfunction(luascript_PrintEmptyErrorMessage)
	callfunction(luascript_PrintErrorMessage)
	callfunction(luascript_HasVariable)
	callfunction(luascript_Reference)
	callfunction(luascript_Variable)
	callfunction(luascript_CallFunction)
	callfunction(luascript_InvokeFunction)
}
