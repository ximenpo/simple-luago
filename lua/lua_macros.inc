#include "lua_header.h"

//
// lua.h    lua_?
//
int Lua_upvalueindex(int i){
    return  lua_upvalueindex(i);
}

void    Lua_call(lua_State *L, int nargs, int nresults){
    lua_call(L, nargs, nresults);
}

int     Lua_pcall(lua_State *L, int nargs, int nresults, int errfunc){
    return  lua_pcall(L, nargs, nresults, errfunc);
}

int     Lua_yield(lua_State *L, int nresults){
    return  lua_yield(L, nresults);
}

#ifdef  lua_yield
#endif

void*   Lua_getextraspace(lua_State *L){
    return  lua_getextraspace(L);
}

lua_Number  Lua_tonumber(lua_State *L, int i){
    return  lua_tonumber(L, i);
}

lua_Integer Lua_tointeger(lua_State *L, int i){
    return  lua_tointeger(L, i);
}

void Lua_pop(lua_State *L, int n) {
    lua_pop(L,n);
}

void Lua_newtable(lua_State *L) {
    Lua_newtable(L);
}

void Lua_register(lua_State *L, const char* n, lua_CFunction f) {
    lua_register(L, n, f);
}

void Lua_pushcfunction(lua_State *L, lua_CFunction f) {
    lua_pushcfunction(L, f);
}

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


// TODO: lua_pushliteral

int Lua_pushglobaltable(lua_State *L){
    return  lua_pushglobaltable(L);
}

const char* Lua_tostring(lua_State *L, int i){
    return  lua_tostring(L, i);
}

void Lua_insert(lua_State *L, int idx) {
    lua_insert(L, idx);
}

void Lua_remove(lua_State *L, int idx) {
    lua_remove(L, idx);
}

void Lua_replace(lua_State *L, int idx) {
    lua_replace(L, idx);
}

//
// lauxlib.h    lua_?
//
void    Lua_writestring(const char* s, size_t l){
    lua_writestring(s, l);
}

void    Lua_writeline(){
    lua_writeline();
}

void    Lua_writestringerror(const char* fmt, const char* msg){
    lua_writestringerror(fmt, msg);
}

//
// lauxlib.h    luaL_?
//
void    LuaL_checkversion(lua_State *L) {
    luaL_checkversion(L);
}

int     LuaL_loadfile(lua_State *L, const char *f){
    return  luaL_loadfile(L, f);
}

// lua_funcs: luaL_newlibtable(L, l)
// lua_funcs: LuaL_newlib(L, l)

void    LuaL_argcheck(lua_State *L, int cond, int arg, const char* extramsg) {
    luaL_argcheck(L, cond, arg, extramsg);
}

const char* LuaL_checkstring(lua_State *L, int n) {
    return  luaL_checkstring(L, n);
}

const char* LuaL_optstring(lua_State *L, int n, const char* d) {
    return  luaL_optstring(L, n, d);
}

const char* LuaL_typename(lua_State *L, int i){
    return  LuaL_typename(L, i);
}

int     LuaL_dofile(lua_State *L, const char* fn){
    return  luaL_dofile(L, fn);
}

int     LuaL_dostring(lua_State *L, const char* s){
    return  luaL_dostring(L, s);
}

void    LuaL_getmetatable(lua_State *L, const char* n) {
    luaL_getmetatable(L, n);
}

//  TODO: support luaL_opt

int     LuaL_loadbuffer(lua_State *L, const void *s, size_t sz, const char *n){
    return  luaL_loadbuffer(L,(const char *)s,sz,n);
}

void    LuaL_addchar(luaL_Buffer *B, char c) {
    luaL_addchar(B, c);
}

void    LuaL_addsize(luaL_Buffer *B, size_t s) {
    LuaL_addsize(B, s);
}

void    LuaL_prepbuffer(luaL_Buffer *B) {
    luaL_prepbuffer(B);
}
