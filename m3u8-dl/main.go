package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	__obf_7179f2bbaa5772e8 string
	__obf_7b98c3dafb8ea7fa string
	__obf_0dacba26c1c15211 int = 10
)

func init() {
	flag.StringVar(&__obf_7179f2bbaa5772e8, "u", __obf_7179f2bbaa5772e8, "M3U8 URL, required")
	flag.IntVar(&__obf_0dacba26c1c15211, "c", __obf_0dacba26c1c15211, "Maximum number of occurrences")
	flag.StringVar(&__obf_7b98c3dafb8ea7fa, "o", "./", "Output folder, required")
}

func main() {
	flag.Parse()
	defer func() {
		if __obf_536c6d3cfc5cc1fa := recover(); __obf_536c6d3cfc5cc1fa != nil {
			fmt.Println("[error]", __obf_536c6d3cfc5cc1fa)
			os.Exit(-1)
		}
	}()
	if __obf_7179f2bbaa5772e8 == "" {
		__obf_f2d455bb9aa90cda("u")
	}
	if __obf_7b98c3dafb8ea7fa == "" {
		__obf_f2d455bb9aa90cda("o")
	}
	if __obf_0dacba26c1c15211 <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	__obf_cc334175f963ac2d, __obf_66cc3fda1d4448dc := NewTask(__obf_7b98c3dafb8ea7fa, __obf_7179f2bbaa5772e8)
	if __obf_66cc3fda1d4448dc != nil {
		panic(__obf_66cc3fda1d4448dc)
	}
	if __obf_66cc3fda1d4448dc := __obf_cc334175f963ac2d.Start(__obf_0dacba26c1c15211); __obf_66cc3fda1d4448dc != nil {
		panic(__obf_66cc3fda1d4448dc)
	}
	fmt.Println("Done!")
}

func __obf_f2d455bb9aa90cda(__obf_80531b6d187ff2ea string) {
	panic("parameter '" + __obf_80531b6d187ff2ea + "' is required")
}
