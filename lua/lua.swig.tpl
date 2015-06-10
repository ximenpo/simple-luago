//
// lua.swig
//

#define     LUA_USE_MACOSX
#define     LUA_32BITS

%{
#include    "$.lua.h"
#include    "$.lualib.h"
#include    "$.lauxlib.h"
#include    "lua_mfuncs.inc"
%}

struct      lua_Debug;

%ignore     lua_pushvfstring;
%ignore     lua_pushfstring;

%include    "${LUA_SRC}/luaconf.h"
%include    "${LUA_SRC}/lua.h"
%include    "${LUA_SRC}/lualib.h"
%include    "${LUA_SRC}/lauxlib.h"
%include    "lua_mfuncs.inc"