//
// lua.swig
//

#define     LUA_USE_MACOSX
#define     LUA_32BITS

%{
#include    "$.lua.h"
#include    "$.lualib.h"
#include    "$.lauxlib.h"
#include    "lua_macros.inc"
%}

struct      lua_Debug;

%ignore     lua_pushvfstring;
%ignore     lua_pushfstring;
%ignore     luaL_checkoption;
%ignore     luaL_setfuncs;

%include    "${LUA_SRC}/luaconf.h"
%include    "${LUA_SRC}/lua.h"
%include    "${LUA_SRC}/lualib.h"
%include    "${LUA_SRC}/lauxlib.h"
%include    "lua_macros.inc"
