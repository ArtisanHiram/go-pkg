package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	__obf_1d184b1c1810901d string
	__obf_f153ab1806df0bbb string
	__obf_b7c0978bba4a885e int = 10
)

func init() {
	flag.StringVar(&__obf_1d184b1c1810901d, "u", __obf_1d184b1c1810901d, "M3U8 URL, required")
	flag.IntVar(&__obf_b7c0978bba4a885e, "c", __obf_b7c0978bba4a885e, "Maximum number of occurrences")
	flag.StringVar(&__obf_f153ab1806df0bbb, "o", "./", "Output folder, required")
}

func main() {
	flag.Parse()
	defer func() {
		if __obf_a2f63a45f1ea7757 := recover(); __obf_a2f63a45f1ea7757 != nil {
			fmt.Println("[error]", __obf_a2f63a45f1ea7757)
			os.Exit(-1)
		}
	}()
	if __obf_1d184b1c1810901d == "" {
		__obf_6428b2a4d2e0797c("u")
	}
	if __obf_f153ab1806df0bbb == "" {
		__obf_6428b2a4d2e0797c("o")
	}
	if __obf_b7c0978bba4a885e <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	__obf_45e984273d82e5c6, __obf_dc3122c5b8033a43 := NewTask(__obf_f153ab1806df0bbb, __obf_1d184b1c1810901d)
	if __obf_dc3122c5b8033a43 != nil {
		panic(__obf_dc3122c5b8033a43)
	}
	if __obf_dc3122c5b8033a43 := __obf_45e984273d82e5c6.Start(__obf_b7c0978bba4a885e); __obf_dc3122c5b8033a43 != nil {
		panic(__obf_dc3122c5b8033a43)
	}
	fmt.Println("Done!")
}

func __obf_6428b2a4d2e0797c(__obf_39275d6955320b2f string) {
	panic("parameter '" + __obf_39275d6955320b2f + "' is required")
}
