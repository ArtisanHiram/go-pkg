package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	__obf_3f814554f334ee3c string
	__obf_719136bb3474572d string
	__obf_1b241975e6ff3a3f int = 10
)

func init() {
	flag.StringVar(&__obf_3f814554f334ee3c, "u", __obf_3f814554f334ee3c, "M3U8 URL, required")
	flag.IntVar(&__obf_1b241975e6ff3a3f, "c", __obf_1b241975e6ff3a3f, "Maximum number of occurrences")
	flag.StringVar(&__obf_719136bb3474572d, "o", "./", "Output folder, required")
}

func main() {
	flag.Parse()
	defer func() {
		if __obf_0d4408494a9fc731 := recover(); __obf_0d4408494a9fc731 != nil {
			fmt.Println("[error]", __obf_0d4408494a9fc731)
			os.Exit(-1)
		}
	}()
	if __obf_3f814554f334ee3c == "" {
		__obf_9118aad4d5a1e257("u")
	}
	if __obf_719136bb3474572d == "" {
		__obf_9118aad4d5a1e257("o")
	}
	if __obf_1b241975e6ff3a3f <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	__obf_00703952ccc55e21, __obf_8e97e8da92c78189 := NewTask(__obf_719136bb3474572d, __obf_3f814554f334ee3c)
	if __obf_8e97e8da92c78189 != nil {
		panic(__obf_8e97e8da92c78189)
	}
	if __obf_8e97e8da92c78189 := __obf_00703952ccc55e21.Start(__obf_1b241975e6ff3a3f); __obf_8e97e8da92c78189 != nil {
		panic(__obf_8e97e8da92c78189)
	}
	fmt.Println("Done!")
}

func __obf_9118aad4d5a1e257(__obf_4b55bd1eeb60ed8a string) {
	panic("parameter '" + __obf_4b55bd1eeb60ed8a + "' is required")
}
