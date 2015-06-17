# simple-luago

simple-luago is a simple wrapper for lua(5.3.0+) in go(1.4+) using swig(3.0.5+) and cgo. 

### supported platforms

* macosx
* linux(centos)
* windows(TODO)

### install

#### prepare

- go1.4.0+
- swig3.0.5+
- git
- gcc
- lua5.3.0+

#### setup

Macosx:
```
export GOPATH=`pwd`
go get github.com/ximenpo/simple-lua
export LUA_SRC=/path/to/lua/src
go generate github.com/ximenpo/simple-lua
go install github.com/ximenpo/simple-lua/lua
```

#### test & examples

check the example diretory.

```
go run src/github.com/ximenpo/simple-luago/example/lua_???.go
```

------

# simple-luago

simple-luago项目专注于lua在go语言中的应用，使用swig和cgo将lua虚拟机嵌入到go语言中，目前使用的相关版本为：

* lua   5.3.0+
* go    1.4.0+
* swig  3.0.5+

### 支持的平台

* macosx
* linux(centos)
* windows(TODO)

### 安装

#### 准备

- go1.4.0+
- swig3.0.5+
- git
- gcc
- lua5.3.0+

#### 安装

Macosx:
```
export GOPATH=`pwd`
go get github.com/ximenpo/simple-lua
export LUA_SRC=/path/to/lua/src
go generate github.com/ximenpo/simple-lua
go install github.com/ximenpo/simple-lua/lua
```

#### 测试&例子

请查看example目录下的例子，并用以下命令运行（替换掉???的内容）。

```
go run src/github.com/ximenpo/simple-luago/example/lua_???.go
```
