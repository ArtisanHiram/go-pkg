package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	__obf_c0fb6d8692497149 string
	__obf_10b0d572c2650427 string
	__obf_0fa0c7d88068e077 int = 10
)

func init() {
	flag.StringVar(&__obf_c0fb6d8692497149, "u", __obf_c0fb6d8692497149, "M3U8 URL, required")
	flag.IntVar(&__obf_0fa0c7d88068e077, "c", __obf_0fa0c7d88068e077, "Maximum number of occurrences")
	flag.StringVar(&__obf_10b0d572c2650427, "o", "./", "Output folder, required")
}

func main() {
	flag.Parse()
	defer func() {
		if __obf_25ab5ff4e3e4b1a7 := recover(); __obf_25ab5ff4e3e4b1a7 != nil {
			fmt.Println("[error]", __obf_25ab5ff4e3e4b1a7)
			os.Exit(-1)
		}
	}()
	if __obf_c0fb6d8692497149 == "" {
		__obf_d59bbad45785e582("u")
	}
	if __obf_10b0d572c2650427 == "" {
		__obf_d59bbad45785e582("o")
	}
	if __obf_0fa0c7d88068e077 <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	__obf_ef0a49037d9aa2bd, __obf_e4142ef6a9f97155 := NewTask(__obf_10b0d572c2650427, __obf_c0fb6d8692497149)
	if __obf_e4142ef6a9f97155 != nil {
		panic(__obf_e4142ef6a9f97155)
	}
	if __obf_e4142ef6a9f97155 := __obf_ef0a49037d9aa2bd.Start(__obf_0fa0c7d88068e077); __obf_e4142ef6a9f97155 != nil {
		panic(__obf_e4142ef6a9f97155)
	}
	fmt.Println("Done!")
}

func __obf_d59bbad45785e582(__obf_1538f46b7e270d9c string) {
	panic("parameter '" + __obf_1538f46b7e270d9c + "' is required")
}
