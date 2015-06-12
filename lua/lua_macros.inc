
//
//  function helper
//
typedef void(*_LuaF_FuncPtr)();
static  _LuaF_FuncPtr   LuaF_Ptr(void* fp){
    return  (_LuaF_FuncPtr)fp;
}

//
// lua.h    lua_?
//
static  int Lua_upvalueindex(int i){
    return  lua_upvalueindex(i);
}

static  void    Lua_call(lua_State *L, int nargs, int nresults){
    lua_call(L, nargs, nresults);
}

static  int     Lua_pcall(lua_State *L, int nargs, int nresults, int errfunc){
    return  lua_pcall(L, nargs, nresults, errfunc);
}

static  int     Lua_yield(lua_State *L, int nresults){
    return  lua_yield(L, nresults);
}

static  void*   Lua_getextraspace(lua_State *L){
    return  lua_getextraspace(L);
}

static  lua_Number  Lua_tonumber(lua_State *L, int i){
    return  lua_tonumber(L, i);
}

static  lua_Integer Lua_tointeger(lua_State *L, int i){
    return  lua_tointeger(L, i);
}

static  void Lua_pop(lua_State *L, int n) {
    lua_pop(L,n);
}

static  void Lua_newtable(lua_State *L) {
    Lua_newtable(L);
}

static  void Lua_register(lua_State *L, const char* n, lua_CFunction f) {
    lua_register(L, n, f);
}

static  void Lua_pushcfunction(lua_State *L, lua_CFunction f) {
    lua_pushcfunction(L, f);
}

static  int Lua_isfunction(lua_State *L, int n) {
    return  lua_isfunction(L,n) ? 1 : 0;
}

static  int Lua_istable(lua_State *L, int n) {
    return  lua_istable(L,n) ? 1 : 0;
}

static  int Lua_islightuserdata(lua_State *L, int n) {
    return  lua_islightuserdata(L,n) ? 1 : 0;
}

static  int Lua_isnil(lua_State *L, int n) {
    return  lua_isnil(L,n) ? 1 : 0;
}

static  int Lua_isboolean(lua_State *L, int n) {
    return  lua_isboolean(L,n) ? 1 : 0;
}

static  int Lua_isthread(lua_State *L, int n) {
    return  lua_isthread(L,n) ? 1 : 0;
}

static  int Lua_isnone(lua_State *L, int n) {
    return  lua_isnone(L,n) ? 1 : 0;
}

static  int Lua_isnoneornil(lua_State *L, int n) {
    return  lua_isnoneornil(L,n) ? 1 : 0;
}

// TODO: lua_pushliteral

static  int Lua_pushglobaltable(lua_State *L){
    return  lua_pushglobaltable(L);
}

static  const char* Lua_tostring(lua_State *L, int i){
    return  lua_tostring(L, i);
}

static  void Lua_insert(lua_State *L, int idx) {
    lua_insert(L, idx);
}

static  void Lua_remove(lua_State *L, int idx) {
    lua_remove(L, idx);
}

static  void Lua_replace(lua_State *L, int idx) {
    lua_replace(L, idx);
}

//
// lauxlib.h    lua_?
//
static  void    Lua_writestring(const char* s, size_t l){
    lua_writestring(s, l);
}

static  void    Lua_writeline(){
    lua_writeline();
}

static  void    Lua_writestringerror(const char* fmt, const char* msg){
    lua_writestringerror(fmt, msg);
}

//
// lauxlib.h    luaL_?
//
static  void    LuaL_checkversion(lua_State *L) {
    luaL_checkversion(L);
}

static  int     LuaL_loadfile(lua_State *L, const char *f){
    return  luaL_loadfile(L, f);
}

// TODO: luaL_newlibtable(L, l)
static  void    LuaL_newlibtable(lua_State *L, luaL_Reg* l, size_t sz) {
    //luaL_newlibtable(L, l);
    lua_createtable(L, 0, sz);
}

// TODO: LuaL_newlib(L, l)
static  void    LuaL_newlib(lua_State *L, luaL_Reg* l, size_t sz) {
    //luaL_newlib(L, l);
    luaL_checkversion(L);
    lua_createtable(L,0,sz);
    luaL_setfuncs(L,l,0);
}

static  void    LuaL_argcheck(lua_State *L, int cond, int arg, const char* extramsg) {
    luaL_argcheck(L, cond, arg, extramsg);
}

static  const char* LuaL_checkstring(lua_State *L, int n) {
    return  luaL_checkstring(L, n);
}

static  const char* LuaL_optstring(lua_State *L, int n, const char* d) {
    return  luaL_optstring(L, n, d);
}

static  const char* LuaL_typename(lua_State *L, int i){
    return  LuaL_typename(L, i);
}

static  int     LuaL_dofile(lua_State *L, const char* fn){
    return  luaL_dofile(L, fn);
}

static  int     LuaL_dostring(lua_State *L, const char* s){
    return  luaL_dostring(L, s);
}

static  void    LuaL_getmetatable(lua_State *L, const char* n) {
    luaL_getmetatable(L, n);
}

//  TODO: support luaL_opt

static  int     LuaL_loadbuffer(lua_State *L, const char *s, size_t sz, const char *n){
    return  luaL_loadbuffer(L,s,sz,n);
}

static  void    LuaL_addchar(luaL_Buffer *B, char c) {
    luaL_addchar(B, c);
}

static  void    LuaL_addsize(luaL_Buffer *B, size_t s) {
    LuaL_addsize(B, s);
}

static  void    LuaL_prepbuffer(luaL_Buffer *B) {
    luaL_prepbuffer(B);
}
