package main

import . "../lua"

import "fmt"

func main() {
	l := LuaL_newstate()

	LuaL_openlibs(l)
	LuaL_dostring(l, `print 'Hello World!\n'`)

	Lua_pushstring(l, "simple")
	fmt.Println(LuaL_ref(l, LUA_REGISTRYINDEX))

	Lua_close(l)
}
