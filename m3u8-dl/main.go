package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	__obf_cbdb4071e572fbe6 string
	__obf_6050b974c19fc9b9 string
	__obf_449e4cbc4b0f25a9 int = 10
)

func init() {
	flag.StringVar(&__obf_cbdb4071e572fbe6, "u", __obf_cbdb4071e572fbe6, "M3U8 URL, required")
	flag.IntVar(&__obf_449e4cbc4b0f25a9, "c", __obf_449e4cbc4b0f25a9, "Maximum number of occurrences")
	flag.StringVar(&__obf_6050b974c19fc9b9, "o", "./", "Output folder, required")
}

func main() {
	flag.Parse()
	defer func() {
		if __obf_53476594443889ef := recover(); __obf_53476594443889ef != nil {
			fmt.Println("[error]", __obf_53476594443889ef)
			os.Exit(-1)
		}
	}()
	if __obf_cbdb4071e572fbe6 == "" {
		__obf_0937e5a2e790b8fa("u")
	}
	if __obf_6050b974c19fc9b9 == "" {
		__obf_0937e5a2e790b8fa("o")
	}
	if __obf_449e4cbc4b0f25a9 <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	__obf_bf0d9ac08dec47a9, __obf_31a042aea6f8bcd2 := NewTask(__obf_6050b974c19fc9b9, __obf_cbdb4071e572fbe6)
	if __obf_31a042aea6f8bcd2 != nil {
		panic(__obf_31a042aea6f8bcd2)
	}
	if __obf_31a042aea6f8bcd2 := __obf_bf0d9ac08dec47a9.Start(__obf_449e4cbc4b0f25a9); __obf_31a042aea6f8bcd2 != nil {
		panic(__obf_31a042aea6f8bcd2)
	}
	fmt.Println("Done!")
}

func __obf_0937e5a2e790b8fa(__obf_6a11eaa81afb1fdd string) {
	panic("parameter '" + __obf_6a11eaa81afb1fdd + "' is required")
}
