package main

import  .   "../lua"

func main() {
    l := LuaL_newstate()

    LuaL_openlibs(l)
    LuaL_dostring(l, `print 'Hello World!\n'`)

    Lua_close(l)
}
