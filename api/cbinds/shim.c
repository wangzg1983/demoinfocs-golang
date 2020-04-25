#include "_cgo_export.h"

void CallEventHandler(ParserHandle p, EventPtr e, EventHandler *fn) {
    return fn(p, e);
} 
