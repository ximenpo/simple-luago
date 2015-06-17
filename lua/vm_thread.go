package lua

/*

#include	"lua_header.h"

extern  int	script_WaitFrames(void* l);
extern  int	script_WaitSeconds(void* l);

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
	m := &LuaThreadMgr{}
	return m
}

func (m *LuaThreadMgr) Start(vm *LuaVM) error {
	if nil != m.LuaVM {
		return errors.New("vm already exist")
	}

	m.LuaVM = vm
	m.threads = map[*LuaThread]bool{}
	return nil
}

func (m *LuaThreadMgr) Stop() {
	for k := range m.threads {
		delete(m.threads, k)
	}
	m.LuaVM = nil
}

func (m *LuaThreadMgr) CreateThread(bAutoDelete bool) *LuaThread {
	if m.LuaVM == nil || m.handle == nil {
		return nil
	}

	t := new(LuaThread)
	t.handle = Lua_newthread(m.handle)
	t.auto_delete = bAutoDelete
	if nil == t.handle {
		t.handle = nil
		return nil
	}

	Lua_pushglobaltable(t.handle)
	Lua_pushthread(t.handle)
	Lua_pushlightuserdata(t.handle, uintptr(unsafe.Pointer(t)))
	Lua_settable(t.handle, -3)
	Lua_pop(t.handle, 1)

	m.threads[t] = true
	return t
}

func (m *LuaThreadMgr) DestroyThread(t *LuaThread) {
	if _, ok := m.threads[t]; ok {
		m.threads[t] = false
	}
}

func (m *LuaThreadMgr) IsValidThread(t *LuaThread) bool {
	todel, ok := m.threads[t]
	return ok && todel
}

func (m *LuaThreadMgr) GetThreadSum() int {
	return len(m.threads)
}

func (m *LuaThreadMgr) OpenScriptLib() {
	if nil != m.LuaVM && nil != m.handle {
		LuaL_newlib(m.handle, C.script_funcs())
		Lua_setglobal(m.handle, "script")
	}
}

func (m *LuaThreadMgr) Update(dt float64) {
	var dead []*LuaThread

	for t, v := range m.threads {
		if !v {
			if dead == nil {
				dead = make([]*LuaThread, 0, len(m.threads))
			}
			dead = append(dead, t)
			continue
		}

		t.update(dt)
		switch t.GetStatus() {
		case THREAD_DONE, THREAD_ERROR, THREAD_NOT_LOADED:
			if t.auto_delete {
				if dead == nil {
					dead = make([]*LuaThread, 0, len(m.threads))
				}
				dead = append(dead, t)
			}
		}
	}

	if dead != nil {
		for i := range dead {
			delete(m.threads, dead[i])
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

func (t *LuaThread) update(dt float64) {
	t.timestamp += float64(dt)

	switch t.status {
	case THREAD_WAIT_SECONDS:
		{ // 脚本等待多少秒
			if t.timestamp >= t.timestamp_wakeup {
				t.resume(false)
			}
		}
	case THREAD_WAIT_FRAMES:
		{ // 脚本等待多少帧
			t.frames_wakeup--
			if t.frames_wakeup <= 0 {
				t.resume(false)
			}
		}
	}
}

func (t *LuaThread) resume(bAbortWait bool) bool {
	switch t.status {
	case THREAD_NOT_LOADED:
		{
			return false
		}
	case THREAD_ERROR:
		{
			return false
		}
	}

	// we're about to run/resume the thread, so set the global
	t.status = THREAD_RUNNING

	// param is treated as a return value from the function that yielded
	Lua_pushboolean(t.handle, bAbortWait)

	switch Lua_resume(t.handle, Lua_NilState(0), 1) {
	case LUA_OK:
		{
			t.status = THREAD_DONE
			return true
		}
	case LUA_YIELD:
		{
			return true
		}
		break
	default:
		{
			t.status = THREAD_ERROR
			t.err_msg = Lua_tostring(t.handle, -1)
			Lua_pop(t.handle, -1)
		}
	}

	return false
}

func (t *LuaThread) GetMgr() *LuaThreadMgr {
	return t.mgr
}

func (t *LuaThread) GetStatus() int {
	return t.status
}

func (t *LuaThread) GetAutoDelete() bool {
	return t.auto_delete
}

func (t *LuaThread) SetAutoDelete(bAutoDelete bool) {
	t.auto_delete = bAutoDelete
}

func (t *LuaThread) RunFile(file string) bool {
	t.status = THREAD_NOT_LOADED

	if LUA_OK == LuaL_loadfile(t.handle, file) {
		t.status = THREAD_LOADED
		return t.resume(false)
	}

	t.status = THREAD_NOT_LOADED
	t.err_msg = Lua_tostring(t.handle, -1)
	Lua_pop(t.handle, 1)

	return false
}

func (t *LuaThread) RunString(code string) bool {
	t.status = THREAD_NOT_LOADED

	if LUA_OK != LuaL_loadstring(t.handle, code) {
		t.err_msg = Lua_tostring(t.handle, -1)
		Lua_pop(t.handle, 1)
		return false
	}

	t.status = THREAD_LOADED

	return t.resume(false)
}

func (t *LuaThread) RunBuffer(buffer unsafe.Pointer, size uint) bool {
	t.status = THREAD_NOT_LOADED

	if LUA_OK != LuaL_loadbuffer(t.handle, uintptr(buffer), size, "LuaThread.RunBuffer") {
		t.err_msg = Lua_tostring(t.handle, -1)
		Lua_pop(t.handle, 1)
		return false
	}

	t.status = THREAD_LOADED

	return t.resume(false)
}

func (t *LuaThread) AbortWait() {
	t.resume(true)
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
	l := LuaF_State(L)
	s := script_GetScriptObject(l)

	s.frames_wakeup = int(LuaL_optinteger(l, 1, 1))
	s.status = THREAD_WAIT_FRAMES

	return 0
}

//export script_WaitSeconds
func script_WaitSeconds(L unsafe.Pointer) Lua_CInt {
	l := LuaF_State(L)
	s := script_GetScriptObject(l)

	s.timestamp_wakeup = s.timestamp + float64(LuaL_optnumber(l, 1, 1.0))
	s.status = THREAD_WAIT_SECONDS

	return 0
}
