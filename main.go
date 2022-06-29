package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/ii64/golinker/_example/internal/native"
)

func main() {

	ret := native.Subr()
	conv := strconv.FormatInt(int64(ret), 16)
	fmt.Println(conv)

	m := make([]byte, 20)
	ret = native.Subr2(&m)
	fmt.Println(m, string(m), ret)

	ret = native.Subr3(&m)
	fmt.Println(m, string(m), ret)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c := make([]byte, 200)

		var ret2, ret3 any
		upper := c[:100]
		lower := c[100:]
		ret2 = native.MemsetSIMD(&c, 'a')
		ret3 = native.MemsetSIMD(&lower, 'b')
		fmt.Println(ret2, ret3)

		fmt.Println(c)
		fmt.Println(upper)
		fmt.Println(lower)
	}()

	wg.Wait()

}
