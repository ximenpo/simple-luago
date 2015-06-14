//
// lua.swig
//

#define     LUA_USE_MACOSX
#define     LUA_32BITS

#define     LUA_COMPAT_APIINTCASTS

%{
#include    "lua_header.h"
#include    "lua_macros.inc"
%}

struct      lua_State;
struct      lua_Debug;
struct      luaL_Buffer;
struct      luaL_Stream;

%ignore     lua_pushvfstring;
%ignore     lua_pushfstring;
%ignore     luaL_Reg;
%ignore     luaL_checkoption;
%ignore     luaL_setfuncs;
%ignore     lua_isnumber;
%ignore     lua_isstring;
%ignore     lua_iscfunction;
%ignore     lua_isinteger;
%ignore     lua_isuserdata;
%ignore     lua_toboolean;
%ignore     lua_pushboolean;

%typemap(gotype)    (lua_Number)    "Lua_Number"
%typemap(gotype)    (lua_Integer)   "Lua_Integer"
%typemap(gotype)    (lua_Unsigned)  "Lua_Unsigned"

%include    "${LUA_SRC}/luaconf.h"
%include    "${LUA_SRC}/lua.h"
%include    "${LUA_SRC}/lualib.h"
%include    "${LUA_SRC}/lauxlib.h"
%include    "lua_macros.inc"
