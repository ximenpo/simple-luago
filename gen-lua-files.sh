#!/bin/bash

if [ "$LUA_SRC" == "" ]; then
    echo    please define LUA_SRC first \(eg. export LUA_SRC=/path/to/lua/src\)
    exit    1
fi

# clean
rm -f lua/$.*.h         2> /dev/null
rm -f lua/$.*.c         2> /dev/null
rm -f lua/*.swig        2> /dev/null
rm -f lua/*.swig.*      2> /dev/null

if [ ! -d "lua" ]; then
    mkdir lua
fi

# gen files
(
    CUR_DIR=`pwd`
    cd  $LUA_SRC
    # h
    headers='luaconf.h lua.h lualib.h lauxlib.h lua.hpp '
    for	H   in  $headers
    do
        echo	"#include \"$LUA_SRC/$H\"" > $CUR_DIR/lua/$.$H
    done
    # c
    for C   in  *.c
    do
    	echo	"#include \"$LUA_SRC/$C\"" > $CUR_DIR/lua/$.$C
    done
    cd  $CUR_DIR
    rm	lua/$.lua.c
    rm  lua/$.luac.c
    # swig.i
    gcc -E -P -dD   "$CUR_DIR/gen-lua-swig.h" | grep -E '_MAX|_MIN|_BIT' > "$CUR_DIR/lua/lua.swig.i"
    # swig
    sed -e "s#\${LUA_SRC}#$LUA_SRC#g" gen-lua-swig.tpl >   lua/lua.swig
)
