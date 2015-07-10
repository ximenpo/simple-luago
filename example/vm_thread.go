package main

import (
	"../lua"
	"fmt"
)

func callfunction(f func(s *lua.LuaThreadMgr)) {

	s := lua.LuaVM{}
	s.Start()
	defer s.Stop()

	mgr := lua.LuaThreadMgr{}
	mgr.Start(&s)
	defer mgr.Stop()

	mgr.OpenStdLibs()
	mgr.OpenScriptLib()
	if 0 != lua.Lua_gettop(mgr.GetHandle()) {
		fmt.Println("STACK: ", lua.Lua_gettop(mgr.GetHandle()))
	}

	f(&mgr)

	mgr.Stop()
}

func luascript_Test(m *lua.LuaThreadMgr) {
	//fmt.Println("--------------------------")
	//fmt.Println("--------------------------")
}

func luascript_Demo(m *lua.LuaThreadMgr) {
	s := m.CreateThread(true)

	if 0 != lua.Lua_gettop(s.GetHandle()) {
		fmt.Println("STACK: ", lua.Lua_gettop(s.GetHandle()))
	}
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum: ", m.GetThreadSum(), "!=", 1)
	}

	if err := s.RunString("print('hello, world')"); err != nil {
		fmt.Println(err)
	}
	if 0 != lua.Lua_gettop(s.GetHandle()) {
		fmt.Println("STACK: ", lua.Lua_gettop(s.GetHandle()))
	}

	m.Update(0.0)
	if 0 != m.GetThreadSum() {
		fmt.Println("#?thread sum: ", m.GetThreadSum(), "!=", 0, " ?", s.GetStatus())
	}
}

func luascript_WaitFrames(m *lua.LuaThreadMgr) {
	s := m.CreateThread(true)

	if 0 != lua.Lua_gettop(s.GetHandle()) {
		fmt.Println("STACK 1: ", lua.Lua_gettop(s.GetHandle()))
	}
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum: ", m.GetThreadSum(), "!=", 1)
	}

	s.RunString("print('#1');script.WaitFrames(3);print('#2');")
	if 0 != lua.Lua_gettop(s.GetHandle()) {
		fmt.Println("STACK 2: ", lua.Lua_gettop(s.GetHandle()))
	}
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum: ", m.GetThreadSum(), "!=", 1)
	}

	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 0: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(0.0)
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 1: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(0.0)
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 2: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(0.0)
	if 0 != m.GetThreadSum() {
		fmt.Println("#thread sum 3: ", m.GetThreadSum(), "!=", 0)
	}
}

func luascript_WaitSeconds(m *lua.LuaThreadMgr) {
	s := m.CreateThread(true)

	if 0 != lua.Lua_gettop(s.GetHandle()) {
		fmt.Println("STACK 1: ", lua.Lua_gettop(s.GetHandle()))
	}
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum: ", m.GetThreadSum(), "!=", 1)
	}

	s.RunString("print('#1');script.WaitSeconds(3.1);print('#2');")
	if 0 != lua.Lua_gettop(s.GetHandle()) {
		fmt.Println("STACK 2: ", lua.Lua_gettop(s.GetHandle()))
	}
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum: ", m.GetThreadSum(), "!=", 1)
	}

	//fmt.Println(" ?", s.GetStatus())
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 0: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(1.0)
	//fmt.Println(" ?", s.GetStatus())
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 1: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(1.0)
	//fmt.Println(" ?", s.GetStatus())
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 2: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(1.0)
	//fmt.Println(" ?", s.GetStatus())
	if 1 != m.GetThreadSum() {
		fmt.Println("#thread sum 3: ", m.GetThreadSum(), "!=", 1)
	}
	m.Update(1.0)
	//fmt.Println(" ?", s.GetStatus())
	if 0 != m.GetThreadSum() {
		fmt.Println("#thread sum 3: ", m.GetThreadSum(), "!=", 0)
	}
}

func main() {
	callfunction(luascript_Test)
	callfunction(luascript_Demo)
	callfunction(luascript_WaitFrames)
	callfunction(luascript_WaitSeconds)
}
