//
// lua.swig
//

struct      lua_Debug;
struct      luaL_Buffer;
struct      luaL_Stream;

//
//  impls manully
//
%ignore     lua_ident;
%ignore     lua_pushvfstring;
%ignore     lua_pushfstring;
%ignore     lua_isnumber;
%ignore     lua_isstring;
%ignore     lua_iscfunction;
%ignore     lua_isinteger;
%ignore     lua_isuserdata;
%ignore     lua_toboolean;
%ignore     lua_pushboolean;
%ignore     luaL_Reg;
%ignore     luaL_checkoption;
%ignore     luaL_setfuncs;
%ignore     luaL_loadbufferx;

%rename(fun)    func;

//
//  types
//
%typemap(gotype)    (lua_State*)    "Lua_State"
%typemap(gotype)    (lua_Number)    "Lua_Number"
%typemap(gotype)    (lua_Integer)   "Lua_Integer"
%typemap(gotype)    (lua_Unsigned)  "Lua_Unsigned"
%typemap(gotype)    (lua_KContext)  "Lua_KContext"
%typemap(gotype)    (int)           "Lua_CInt"
%typemap(gotype)    (unsigned int)  "Lua_CUint"
%typemap(gotype)    (FILE*)         "Lua_CFile"
%typemap(gotype)    (size_t)        "uint"
%typemap(gotype)    (void*)         "uintptr"
%typemap(gotype)    (CallInfo*)     "uintptr"

//
// interfaces
//
%import     "lua.swig.i"
%include    "${LUA_SRC}/luaconf.h"
%include    "${LUA_SRC}/lua.h"
%include    "${LUA_SRC}/lualib.h"
%include    "${LUA_SRC}/lauxlib.h"
%include    "lua_macros.inc"

//
// c wrapper code
//
%insert(header) %{
#include    "lua_header.h"
#include    "lua_macros.inc"
%}

//
// go wrapper code
//
#if     SWIG_VERSION < 0x030006
%insert(go_begin) %{
/*
#include "lua_header.h"
*/
import "C"
%}
#endif
