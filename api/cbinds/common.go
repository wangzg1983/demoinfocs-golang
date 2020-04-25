package main

//#include "preamble.h"
import "C"
import (
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"unsafe"
)

//export DestroyDemoHeader
func DestroyDemoHeader(c_ptr *C.DemoHeader) {
	C.free(unsafe.Pointer(c_ptr.Filestamp))
	C.free(unsafe.Pointer(c_ptr.ServerName))
	C.free(unsafe.Pointer(c_ptr.ClientName))
	C.free(unsafe.Pointer(c_ptr.MapName))
	C.free(unsafe.Pointer(c_ptr.GameDirectory))
	C.free(unsafe.Pointer(c_ptr))
}

//export DestroyEquipment
func DestroyEquipment(c_ptr *C.Equipment) {
	C.free(unsafe.Pointer(c_ptr.Type))
	C.free(unsafe.Pointer(c_ptr))
}

//export DestroyPlayer
func DestroyPlayer(c_ptr *C.Player) {
	C.free(unsafe.Pointer(c_ptr.Name))
	C.free(unsafe.Pointer(c_ptr))
}

// Allocates memory on the C heap for a `struct equipment`, copies data from
// `equip` to that struct and returns a pointer to that struct
func createEquipment(equip *common.Equipment) *C.Equipment {
	c_ptr := (*C.Equipment)(C.malloc(C.sizeof_Equipment))
	c_ptr.Type = C.CString(equip.Type.String())
	return c_ptr
}

func createPlayer(player *common.Player) *C.Player {
	c_ptr := (*C.Player)(C.malloc(C.sizeof_Player))
	c_ptr.Name = C.CString(player.Name)
	return c_ptr
}
