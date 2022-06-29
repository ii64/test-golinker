package native

//go:nosplit
//go:noescape
func subr() (ret int64)

//go:nosplit
//go:noescape
func subr2(b *[]byte) (ret int64)

//go:nosplit
//go:noescape
func subr3(b *[]byte) (ret int64)

//go:nosplit
//go:noescape
func subrSimd(b *[]byte, v byte) (ret int32)

func Subr() int64 {
	return subr()
}

func Subr2(b *[]byte) int64 {
	return subr2(b)
}

func Subr3(b *[]byte) int64 {
	return subr3(b)
}

func MemsetSIMD(b *[]byte, v byte) int32 {
	return subrSimd(b, v)
}
