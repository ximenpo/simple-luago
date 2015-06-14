#ifndef	__SIMPLE_LUAGO_LUA_UTILS_H__
#define	__SIMPLE_LUAGO_LUA_UTILS_H__

struct	lua_State;

int		luaU_GetRef(struct lua_State* L, int nVariableReference);

int		luaU_GetGlobal(struct lua_State* L, const char* sVariableName);
int		luaU_SetGlobal(struct lua_State* L, const char* sVariableName);

#endif
