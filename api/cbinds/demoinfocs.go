package main

//#include "preamble.h"
import "C"
import (
	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"os"
	"unsafe"
)

// Keep Go pointers in a map indexed by an integer (handle), and pass this handle
// to C, to comply to pointer passing rules (see Cgo documentation)
var cpointers = PtrProxy()

// Keep demo files in a map keyed by the parser handle. This is necessary, so
// the file pointers aren't garbage collected between function calls.
var demoFiles = make(map[uint]*os.File)

//export GetNewParser
func GetNewParser(c_filePath *C.char) C.ParserHandle {
	filePath := C.GoString(c_filePath)
	demoFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	parser := demoinfocs.NewParser(demoFile)
	// Create a handle for our parser
	handle := cpointers.Ref(unsafe.Pointer(&parser))
	// Keep the demo file open after we leave the function and connect it
	// to our parser
	demoFiles[uint(handle)] = demoFile
	return handle
}

//export GetGameState
func GetGameState(handle C.ParserHandle) *C.GameState {
	parser := getParser(handle)
	gameState := parser.GameState()
	// GameState does change during parsing. So don't keep the Go pointer,
	// but copy the current state to the C heap. The C user has to free
	// the data himself.
	c_gameState := (*C.GameState)(C.malloc(C.sizeof_GameState))
	c_gameState.TotalRoundsPlayed = C.int(gameState.TotalRoundsPlayed())
	return c_gameState
}

//export DestroyGameState
func DestroyGameState(cPtr *C.GameState) {
	C.free(unsafe.Pointer(cPtr))
}

//export ParseHeader
func ParseHeader(handle C.ParserHandle) *C.DemoHeader {
	parser := getParser(handle)
	header, err := parser.ParseHeader()
	if err != nil {
		panic(err)
	}
	// Copy the header struct from Go heap memory to C heap memory, because the Go
	// struct might get garbage collected. (See https://golang.org/cmd/cgo/#hdr-Passing_pointers).
	// This memory has to be freed by the caller.
	c_header := (*C.DemoHeader)(C.malloc(C.sizeof_DemoHeader))
	// CString allocates memory on the C heap by calling C.malloc internally. So the
	// caller must free this at some point.
	c_header.Filestamp = C.CString(header.Filestamp)
	c_header.ServerName = C.CString(header.ServerName)
	c_header.ClientName = C.CString(header.ClientName)
	c_header.MapName = C.CString(header.MapName)
	c_header.GameDirectory = C.CString(header.GameDirectory)
	c_header.PlaybackTime = C.longlong(header.PlaybackTime)
	c_header.Protocol = C.int(header.Protocol)
	c_header.NetworkProtocol = C.int(header.NetworkProtocol)
	c_header.PlaybackTicks = C.int(header.PlaybackTicks)
	c_header.PlaybackFrames = C.int(header.PlaybackFrames)
	c_header.SignonLength = C.int(header.SignonLength)
	return c_header
}

//export ParseToEnd
func ParseToEnd(handle C.ParserHandle) {
	parser := getParser(handle)
	demoFile := demoFiles[uint(handle)]
	delete(demoFiles, uint(handle))
	defer demoFile.Close()
	defer cpointers.Free(handle)
	err := parser.ParseToEnd()
	if err != nil {
		panic(err)
	}
}

func getParser(handle C.uint) demoinfocs.Parser {
	ptr, ok := cpointers.Deref(handle)
	if !ok {
		panic("No parser found")
	}
	//ptr is unsafe.Pointer
	//(*demoinfocs.Parser)(ptr) is *demoinfocs.Parser (interface)
	//(*(*demoinfocs.Parser)(ptr)) is *demoinfocs.parser (actual struct)
	parser := (*(*demoinfocs.Parser)(ptr))
	return parser
}

func main() {}
