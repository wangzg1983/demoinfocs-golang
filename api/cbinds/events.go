package main

//#include "preamble.h"
import "C"
import (
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	"unsafe"
)

//export RegisterEventHandler
func RegisterEventHandler(handle C.ParserHandle, event C.Event, c_func *C.EventHandler) {
	switch event {
	case C.KILL_EVENT:
		registerEventHandler_Kill(handle, c_func)
	case C.ROUND_END_EVENT:
		registerEventHandler_RoundEnd(handle, c_func)
	case C.ROUND_START_EVENT:
		registerEventHandler_RoundStart(handle, c_func)
	}
}

func registerEventHandler_Kill(handle C.ParserHandle, c_func *C.EventHandler) {
	parser := getParser(handle)
	parser.RegisterEventHandler(func(e events.Kill) {
		c_weapon := createEquipment(e.Weapon)
		c_victim := createPlayer(e.Victim)
		c_killer := createPlayer(e.Killer)
		c_ptr := (*C.KillEvent)(C.malloc(C.sizeof_KillEvent))
		c_ptr.Weapon = c_weapon
		c_ptr.Victim = c_victim
		c_ptr.Killer = c_killer
		c_ptr.PenetratedObjects = C.int(e.PenetratedObjects)
		c_ptr.IsHeadshot = C.bool(e.IsHeadshot)
		C.CallEventHandler(handle, C.EventPtr(c_ptr), c_func)
	})
}

//export DestroyKillEvent
func DestroyKillEvent(c_ptr *C.KillEvent) {
	DestroyEquipment(c_ptr.Weapon)
	DestroyPlayer(c_ptr.Victim)
	DestroyPlayer(c_ptr.Killer)
	C.free(unsafe.Pointer(c_ptr))
}

func registerEventHandler_RoundEnd(handle C.ParserHandle, c_func *C.EventHandler) {
	parser := getParser(handle)
	parser.RegisterEventHandler(func(e events.RoundEnd) {
		c_ptr := (*C.RoundEndEvent)(C.malloc(C.sizeof_RoundEndEvent))
		c_ptr.Message = C.CString(e.Message)
		C.CallEventHandler(handle, C.EventPtr(c_ptr), c_func)
	})
}

//export DestroyRoundEndEvent
func DestroyRoundEndEvent(c_ptr *C.RoundEndEvent) {
	C.free(unsafe.Pointer(c_ptr.Message))
	C.free(unsafe.Pointer(c_ptr))
}

func registerEventHandler_RoundStart(handle C.ParserHandle, c_func *C.EventHandler) {
	parser := getParser(handle)
	parser.RegisterEventHandler(func(e events.RoundStart) {
		c_ptr := (*C.RoundStartEvent)(C.malloc(C.sizeof_RoundStartEvent))
		c_ptr.Objective = C.CString(e.Objective)
		C.CallEventHandler(handle, C.EventPtr(c_ptr), c_func)
	})
}

//export DestroyRoundStartEvent
func DestroyRoundStartEvent(c_ptr *C.RoundStartEvent) {
	C.free(unsafe.Pointer(c_ptr.Objective))
	C.free(unsafe.Pointer(c_ptr))
}
