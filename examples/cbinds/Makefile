LIBDIR="../../api/cbinds"

all: examples

examples: libdemoinfocs example1

example1: libdemoinfocs
	gcc -Wall example1.c -L. -ldemoinfocs -o $@ 

libdemoinfocs:
	cd $(LIBDIR) && $(MAKE) c-shared
	cp $(LIBDIR)/build/* .
	cp $(LIBDIR)/preamble.h .

run-example1: example1
	LD_LIBRARY_PATH=. ./example1

clean:
	rm -f example1
	rm -f libdemoinfocs.so
	rm -f libdemoinfocs.a
	rm -f libdemoinfocs.h
	rm -f preamble.h
