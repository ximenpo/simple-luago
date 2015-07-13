package lua

/*

#include	"lua_header.h"
//struct lua_State;
//int     Lua_yield(struct lua_State *L, int nresults);

int	script_WaitFrames(void* l);
int	script_WaitSeconds(void* l);

static	int	script_yield_WaitFrames(void* l){
	return	lua_yield(l, script_WaitFrames(l));
}

static	int	script_yield_WaitSeconds(void* l){
	return	lua_yield(l, script_WaitSeconds(l));
}

static  const void* script_funcs(){
	static	const	struct{void *name, *func;}	lfs[]	= {
		{"WaitFrames",		script_yield_WaitFrames},
		{"WaitSeconds",  	script_yield_WaitSeconds},
		{0,0},
	};
	return  lfs;
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

//	LuaThread	status
const (
	THREAD_NOT_LOADED   = iota // 脚本未载入
	THREAD_LOADED              // 脚本已载入
	THREAD_RUNNING             // 脚本运行中
	THREAD_WAIT_SECONDS        // 脚本等待多少秒
	THREAD_WAIT_FRAMES         // 脚本等待多少帧
	THREAD_ERROR               // 脚本出现错误
	THREAD_DONE                // 脚本执行完毕
)

//	LuaThreadMgr
type LuaThreadMgr struct {
	*LuaVM
	threads map[*LuaThread]bool
}

func NewLuaThreadMgr() *LuaThreadMgr {
	self := &LuaThreadMgr{}
	return self
}

func (self *LuaThreadMgr) Start(vm *LuaVM) error {
	if nil != self.LuaVM {
		return errors.New("vm already exist")
	}

	self.LuaVM = vm
	self.threads = map[*LuaThread]bool{}
	return nil
}

func (self *LuaThreadMgr) Stop() {
	for k := range self.threads {
		delete(self.threads, k)
	}
	self.LuaVM = nil
}

func (self *LuaThreadMgr) CreateThread(bAutoDelete bool) *LuaThread {
	if self.LuaVM == nil || self.Handle == nil {
		return nil
	}

	t := new(LuaThread)
	t.Handle = Lua_newthread(self.Handle)
	t.auto_delete = bAutoDelete
	if nil == t.Handle {
		t.Handle = nil
		return nil
	}

	Lua_pushglobaltable(t.Handle)
	Lua_pushthread(t.Handle)
	Lua_pushlightuserdata(t.Handle, uintptr(unsafe.Pointer(t)))
	Lua_settable(t.Handle, -3)
	Lua_pop(t.Handle, 1)

	self.threads[t] = true
	return t
}

func (self *LuaThreadMgr) DestroyThread(t *LuaThread) {
	if _, ok := self.threads[t]; ok {
		self.threads[t] = false
	}
}

func (self *LuaThreadMgr) IsValidThread(t *LuaThread) bool {
	todel, ok := self.threads[t]
	return ok && todel
}

func (self *LuaThreadMgr) GetThreadSum() int {
	return len(self.threads)
}

func (self *LuaThreadMgr) OpenScriptLib() {
	if nil != self.LuaVM && nil != self.Handle {
		LuaL_newlib(self.Handle, C.script_funcs())
		Lua_setglobal(self.Handle, "script")
	}
}

func (self *LuaThreadMgr) Update(dt float64) {
	var dead []*LuaThread

	for t, v := range self.threads {
		if !v {
			if dead == nil {
				dead = make([]*LuaThread, 0, len(self.threads))
			}
			dead = append(dead, t)
			continue
		}

		t.update(dt)
		switch t.GetStatus() {
		case THREAD_DONE, THREAD_ERROR, THREAD_NOT_LOADED:
			if t.auto_delete {
				if dead == nil {
					dead = make([]*LuaThread, 0, len(self.threads))
				}
				dead = append(dead, t)
			}
		}
	}

	if dead != nil {
		for i := range dead {
			delete(self.threads, dead[i])
		}
	}
}

//	LuaThread
type LuaThread struct {
	LuaScript

	status      int
	auto_delete bool
	mgr         *LuaThreadMgr

	timestamp        float64 // current time
	timestamp_wakeup float64 // time to wake up
	frames_wakeup    int     // number of frames to wait
}

func (self *LuaThread) update(dt float64) (err error) {
	self.timestamp += float64(dt)

	switch self.status {
	case THREAD_WAIT_SECONDS:
		{ // 脚本等待多少秒
			if self.timestamp >= self.timestamp_wakeup {
				err = self.resume(false)
			}
		}
	case THREAD_WAIT_FRAMES:
		{ // 脚本等待多少帧
			self.frames_wakeup--
			if self.frames_wakeup <= 0 {
				err = self.resume(false)
			}
		}
	}

	return
}

func (self *LuaThread) resume(bAbortWait bool) error {
	switch self.status {
	case THREAD_NOT_LOADED:
		{
			return errors.New("thread not loaded")
		}
	case THREAD_ERROR:
		{
			return errors.New("thread was error")
		}
	}

	// we're about to run/resume the thread, so set the global
	self.status = THREAD_RUNNING

	// param is treated as a return value from the function that yielded
	Lua_pushboolean(self.Handle, bAbortWait)

	var err_msg string
	switch Lua_resume(self.Handle, nil, 1) {
	case LUA_OK:
		{
			self.status = THREAD_DONE
			return nil
		}
	case LUA_YIELD:
		{
			return nil
		}
		break
	default:
		{
			self.status = THREAD_ERROR
			err_msg = Lua_tostring(self.Handle, -1)
			Lua_pop(self.Handle, -1)
		}
	}

	return errors.New(err_msg)
}

func (self *LuaThread) GetMgr() *LuaThreadMgr {
	return self.mgr
}

func (self *LuaThread) GetStatus() int {
	return self.status
}

func (self *LuaThread) GetAutoDelete() bool {
	return self.auto_delete
}

func (self *LuaThread) SetAutoDelete(bAutoDelete bool) {
	self.auto_delete = bAutoDelete
}

func (self *LuaThread) RunFile(file string) error {
	self.status = THREAD_NOT_LOADED

	if LUA_OK == LuaL_loadfile(self.Handle, file) {
		self.status = THREAD_LOADED
		return self.resume(false)
	}

	self.status = THREAD_NOT_LOADED
	err_msg := Lua_tostring(self.Handle, -1)
	Lua_pop(self.Handle, 1)

	return errors.New(err_msg)
}

func (self *LuaThread) RunString(code string) error {
	self.status = THREAD_NOT_LOADED

	if LUA_OK != LuaL_loadstring(self.Handle, code) {
		err_msg := Lua_tostring(self.Handle, -1)
		Lua_pop(self.Handle, 1)
		return errors.New(err_msg)
	}

	self.status = THREAD_LOADED

	return self.resume(false)
}

func (self *LuaThread) RunBuffer(buffer unsafe.Pointer, size uint) error {
	self.status = THREAD_NOT_LOADED

	if LUA_OK != LuaL_loadbuffer(self.Handle, uintptr(buffer), size, "LuaThread.RunBuffer") {
		err_msg := Lua_tostring(self.Handle, -1)
		Lua_pop(self.Handle, 1)
		return errors.New(err_msg)
	}

	self.status = THREAD_LOADED

	return self.resume(false)
}

func (self *LuaThread) AbortWait() {
	self.resume(true)
}

//
//	script Lib
//
func script_GetScriptObject(L Lua_State) *LuaThread {
	Lua_pushglobaltable(L)
	Lua_pushthread(L)
	Lua_gettable(L, -2)

	R := (*LuaThread)(unsafe.Pointer(Lua_touserdata(L, -1)))

	Lua_pop(L, 2)
	return R
}

//export script_WaitFrames
func script_WaitFrames(L unsafe.Pointer) Lua_CInt {
	l := Lua_State(L)
	s := script_GetScriptObject(l)

	s.frames_wakeup = int(LuaL_optinteger(l, 1, 1))
	s.status = THREAD_WAIT_FRAMES

	return 0
}

//export script_WaitSeconds
func script_WaitSeconds(L unsafe.Pointer) Lua_CInt {
	l := Lua_State(L)
	s := script_GetScriptObject(l)

	s.timestamp_wakeup = s.timestamp + float64(LuaL_optnumber(l, 1, 1.0))
	s.status = THREAD_WAIT_SECONDS

	return 0
}
