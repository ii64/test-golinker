// +build !noasm !appengine
// Code generated by golinker, DO NOT EDIT.
// Command: golinker -fallback-rawbytes-x86 -rawbytes-x86 -extsymstub -stub=./internal/native/stub.go -out=./internal/native -entryname=__native_entry__ ./native/libnative-amd64.a
package native

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
	__subr_subr = __native_entry__() + 16
	__subr_subr2 = __native_entry__() + 32
	__subr_subr3 = __native_entry__() + 112
	__subr_subrSimd = __native_entry__() + 208
)

const (
	__stack_subr = 0
	__stack_subr2 = 0
	__stack_subr3 = 0
	__stack_subrSimd = 0
)

const (
	_ = __stack_subr
	_ = __stack_subr2
	_ = __stack_subr3
	_ = __stack_subrSimd
)
var (
	_ = __subr_subr
	_ = __subr_subr2
	_ = __subr_subr3
	_ = __subr_subrSimd
)