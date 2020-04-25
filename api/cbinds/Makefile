gofiles = demoinfocs.go common.go events.go ptrproxy.go

all: c-shared c-archive

c-shared: build/libdemoinfocs.so

c-archive: build/libdemoinfocs.a

build/:
	mkdir build

build/libdemoinfocs.so: $(gofiles)
	go build -buildmode=c-shared -o $@

build/libdemoinfocs.a: $(gofiles)
	go build -buildmode=c-archive -o $@

clean:
	rm -rf build/
