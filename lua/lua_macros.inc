#include "lua_header.h"

//
// lua.h    lua_?
//
#ifdef  lua_upvalueindex
int Lua_upvalueindex(int i){
    return  lua_upvalueindex(i);
}
#endif

#ifdef  lua_call
void    Lua_call(lua_State *L, int nargs, int nresults){
    lua_call(L, nargs, nresults);
}
#endif

#ifdef  lua_pcall
int     Lua_pcall(lua_State *L, int nargs, int nresults, int errfunc){
    return  lua_pcall(L, nargs, nresults, errfunc);
}
#endif

#ifdef  lua_yield
int     Lua_yield(lua_State *L, int nresults){
    return  lua_yield(L, nresults);
}
#endif

#ifdef  lua_getextraspace
void*   Lua_getextraspace(lua_State *L){
    return  lua_getextraspace(L);
}
#endif

#ifdef  lua_tonumber
lua_Number  Lua_tonumber(lua_State *L, int i){
    return  lua_tonumber(L, i);
}
#endif

#ifdef  lua_tointeger
lua_Integer Lua_tointeger(lua_State *L, int i){
    return  lua_tointeger(L, i);
}
#endif

#ifdef  lua_pop
void Lua_pop(lua_State *L, int n) {
    lua_pop(L,n);
}
#endif

#ifdef  lua_newtable
void Lua_newtable(lua_State *L) {
    lua_newtable(L);
}
#endif

#ifdef  lua_register
void Lua_register(lua_State *L, const char* n, lua_CFunction f) {
    lua_register(L, n, f);
}
#endif

#ifdef  lua_pushcfunction
void Lua_pushcfunction(lua_State *L, lua_CFunction f) {
    lua_pushcfunction(L, f);
}
#endif

//  lua_funcs: lua_isfunction(L,n)
int X_Lua_isfunction(lua_State *L, int n) {
    return  lua_isfunction(L,n) ? 1 : 0;
}

//  lua_funcs: lua_istable(L,n)
int X_Lua_istable(lua_State *L, int n) {
    return  lua_istable(L,n) ? 1 : 0;
}

//  lua_funcs: lua_islightuserdata(L,n)
int X_Lua_islightuserdata(lua_State *L, int n) {
    return  lua_islightuserdata(L,n) ? 1 : 0;
}

//  lua_funcs: lua_isnil(L,n)
int X_Lua_isnil(lua_State *L, int n) {
    return  lua_isnil(L,n) ? 1 : 0;
}

//  lua_funcs: lua_isboolean(L,n)
int X_Lua_isboolean(lua_State *L, int n) {
    return  lua_isboolean(L,n) ? 1 : 0;
}

//  lua_funcs: lua_isthread(L,n)
int X_Lua_isthread(lua_State *L, int n) {
    return  lua_isthread(L,n) ? 1 : 0;
}

//  lua_funcs: lua_isnone(L,n)
int X_Lua_isnone(lua_State *L, int n) {
    return  lua_isnone(L,n) ? 1 : 0;
}

//  lua_funcs: lua_isnoneornil(L,n)
int X_Lua_isnoneornil(lua_State *L, int n) {
    return  lua_isnoneornil(L,n) ? 1 : 0;
}

#ifdef  lua_pushliteral
// TODO: lua_pushliteral
#endif

#ifdef  lua_pushglobaltable
int Lua_pushglobaltable(lua_State *L){
    return  lua_pushglobaltable(L);
}
#endif

#ifdef  lua_tostring
const char* Lua_tostring(lua_State *L, int i){
    return  lua_tostring(L, i);
}
#endif

#ifdef  lua_insert
void Lua_insert(lua_State *L, int idx) {
    lua_insert(L, idx);
}
#endif

#ifdef  lua_remove
void Lua_remove(lua_State *L, int idx) {
    lua_remove(L, idx);
}
#endif

#ifdef  lua_replace
void Lua_replace(lua_State *L, int idx) {
    lua_replace(L, idx);
}
#endif

//
// lauxlib.h    lua_?
//
#ifdef  lua_writestring
void    Lua_writestring(const char* s, size_t l){
    lua_writestring(s, l);
}
#endif

#ifdef  lua_writeline
void    Lua_writeline(){
    lua_writeline();
}
#endif

#ifdef  lua_writestringerror
void    Lua_writestringerror(const char* fmt, const char* msg){
    lua_writestringerror(fmt, msg);
}
#endif

//
// lauxlib.h    luaL_?
//
#ifdef  luaL_checkversion
void    LuaL_checkversion(lua_State *L) {
    luaL_checkversion(L);
}
#endif

#ifdef  luaL_loadfile
int     LuaL_loadfile(lua_State *L, const char *f){
    return  luaL_loadfile(L, f);
}
#endif

#ifdef  luaL_newlibtable
// lua_funcs: luaL_newlibtable(L, l)
#endif

#ifdef  luaL_newlib
// lua_funcs: luaL_newlib(L, l)
#endif

#ifdef  luaL_argcheck
void    LuaL_argcheck(lua_State *L, int cond, int arg, const char* extramsg) {
    luaL_argcheck(L, cond, arg, extramsg);
}
#endif

#ifdef  luaL_checkstring
const char* LuaL_checkstring(lua_State *L, int n) {
    return  luaL_checkstring(L, n);
}
#endif

#ifdef  luaL_optstring
const char* LuaL_optstring(lua_State *L, int n, const char* d) {
    return  luaL_optstring(L, n, d);
}
#endif

#ifdef  luaL_typename
const char* LuaL_typename(lua_State *L, int i){
    return  luaL_typename(L, i);
}
#endif

#ifdef  luaL_dofile
int     LuaL_dofile(lua_State *L, const char* fn){
    return  luaL_dofile(L, fn);
}
#endif

#ifdef  luaL_dostring
int     LuaL_dostring(lua_State *L, const char* s){
    return  luaL_dostring(L, s);
}
#endif

#ifdef  luaL_getmetatable
void    LuaL_getmetatable(lua_State *L, const char* n) {
    luaL_getmetatable(L, n);
}
#endif

#ifdef  luaL_opt
//  TODO: support luaL_opt
#endif

#ifdef  luaL_loadbuffer
int     LuaL_loadbuffer(lua_State *L, const void *s, size_t sz, const char *n){
    return  luaL_loadbuffer(L,(const char *)s,sz,n);
}
#endif

#ifdef  luaL_addchar
void    LuaL_addchar(luaL_Buffer *B, char c) {
    luaL_addchar(B, c);
}
#endif

#ifdef  luaL_addsize
void    LuaL_addsize(luaL_Buffer *B, size_t s) {
    luaL_addsize(B, s);
}
#endif

#ifdef  luaL_prepbuffer
void    LuaL_prepbuffer(luaL_Buffer *B) {
    luaL_prepbuffer(B);
}
#endif
