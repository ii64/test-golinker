default: build-debug

GLINKER := $(GLINKER)
GLINKER += -fallback-rawbytes-x86
GLINKER += -rawbytes-x86 # temporary
GLINKER += -extsymstub

build-debug:
	go build -x -gcflags="all=-N -l" main.go

build-native:
	make -C native
	golinker $(GLINKER) \
		-stub=./internal/native/stub.go \
		-out=./internal/native \
		-entryname=__native_entry__ \
		./native/libnative-amd64.a

build:
	go build main.go


dmp: build-native build 
	objdump -d main > dmp
	ld --relocatable --whole-archive native/libnative-amd64.a -o /tmp/a.o
	objdump -d /tmp/a.o > dmp2

clean:
	make -C native clean
	rm -rf main dmp dmp2