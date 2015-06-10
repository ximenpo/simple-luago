package lua

import(
    "testing"
)

func Test_lua_macros(t *testing.T) {
	if len(LUA_COPYRIGHT) <= 0 {
        t.Error("LUA_COPYRIGHT shoudn't be empty")
    }
}

func Test_lua_simple_usage(t *testing.T) {
    l := LuaL_newstate()

    LuaL_openlibs(l)
    LuaL_dostring(l, `print 'Hello World!\n'`)

    Lua_close(l)
}
