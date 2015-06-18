@PUSHD  %~dp0
@setlocal enabledelayedexpansion

@IF "%LUA_SRC%" == "" @(
    ECHO	please define LUA_SRC first, unix seperator, no quotes.
	ECHO	eg. SET LUA_SRC=E:\path\to\lua\src
    EXIT    /B  1
)

@SET LUA_SRC=%LUA_SRC:\=/%

:: clean
@IF EXIST lua\$.*.h		DEL /Q /F lua\$.*.h		2> nul
@IF EXIST lua\$.*.c		DEL /Q /F lua\$.*.c		2> nul
@IF EXIST lua\*.swig	DEL /Q /F lua\*.swig	2> nul
@IF EXIST lua\*.swig.*	DEL /Q /F lua\*.swig.*	2> nul

@IF NOT EXIST lua MD lua

:: gen files
@SET	CUR_DIR=%CD%
@CD		"%LUA_SRC:/=\%"

:: h
@SET	HEADERS=luaconf.h lua.h lualib.h lauxlib.h lua.hpp
@FOR	%%F IN (%HEADERS%)	DO	@(
	@ECHO   #include "%LUA_SRC%/%%F"	>	"%CUR_DIR%\lua\$.%%F"
)

:: c
@FOR %%F in (*.c) do @(
	@ECHO	#include "%LUA_SRC%/%%F"	>	"%CUR_DIR%\lua\$.%%F"
)
@CD  %CUR_DIR%
@DEL /Q /F lua\$.lua.c
@DEL /Q /F lua\$.luac.c

:: swig.i
@gcc -E -P -dD   "%CUR_DIR%/gen-lua-swig.h" | findstr "_MAX _MIN _BIT" > "%CUR_DIR%/lua/lua.swig.i"

:: swig
@for /F "delims=" %%i in (gen-lua-swig.tpl) do @(
	@set var=%%i
	@echo !var:${LUA_SRC}=%LUA_SRC%!	>>	lua/lua.swig
)

@POPD
