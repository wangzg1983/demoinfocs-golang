# C API for demoinfocs-golang

Use *demoinfocs-golang* as a library in your program written in C.

Currently, there is only a working prototype, which proofs that the library can be called from C. So **a lot of events and other data structs are missing.**

## Example

```C
#include <stdio.h>
#include "libdemoinfocs.h"

void printKills(ParserHandle parser, KillEvent *data) {
    printf("%s killed %s\n", data->Killer->Name, data->Victim->Name);
    DestroyKillEvent(data);
}

int main() {
    ParserHandle parser = GetNewParser("demo.dem");
    RegisterEventHandler(parser, KILL_EVENT, (EventHandler *)printKills);
    ParseToEnd(parser);
    return 0;
}
```

This simple example creates a parser, registers an event handler `printKills()` and executes the parser. Every time a `demoinfocs.events.Kill` is fired, the Go library calls `printKills()`.

To see an example in action:

```
$ cd examples/cbinds
$ make run-example1
```

This will build a shared library `libdemoinfocs.so`, compile `example1.c` against it and run the resulting binary. You have to provide a `demo.dem` in the same directory.

## How it works

`RegisterEventHandler()` or `GetGameState()`, and others, are function wrappers written in Go. They wrap the equivalent library functions `demoinfocs.Parser.RegisterEventHandler()` or `demoinfocs.Parser.GameState()`. These wrappers are exported to C (see [Cgo documentation](https://golang.org/cmd/cgo/#hdr-C_references_to_Go)), so they can be called from there.

Because C is not a garbage collecting programming environment, we export some more functions, like `DestroyKillEvent()`. The reason is, all data passed to C is allocated on the C heap with `malloc()`. So `RegisterEventHandler(p, KILL_EVENT, (EventHandler *)e)`, for example, allocates memory on the C heap for the `KillEvent` struct, which can then be used in the event handler. At some point you have to free (or destroy) this data, to not leak memory.

The parser itself (`demoinfocs.parser` struct) has to live in the Go memory and is there for not passed to the C environment (see [Pointer passing rules in Cgo](https://golang.org/cmd/cgo/#hdr-Passing_pointers)). But to be able to reference the parser we use an indirection, called `ParserHandle`. So every time you want to use a function on a specfic parser (created by `GetNewParser()`), you need to pass the handle for this specific parser with the function. 

The Cgo concepts in this project are based on [this awesome work](https://github.com/shazow/gohttplib). Also read [this blog post](https://blog.heroku.com/see_python_see_python_go_go_python_go) to gain further insight in to the concepts.

## Build the library

To build the shared library:

```
$ cd api/cbinds
$ make c-shared
```

To build an archive:

```
$ cd api/cbinds
$ make c-archive
```
