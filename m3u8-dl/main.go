package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	__obf_fba5742c8939d4d5 string
	__obf_75dc0fa65ea3ee4a string
	__obf_11b2063fd33fc8a2 int = 10
)

func init() {
	flag.StringVar(&__obf_fba5742c8939d4d5, "u", __obf_fba5742c8939d4d5, "M3U8 URL, required")
	flag.IntVar(&__obf_11b2063fd33fc8a2, "c", __obf_11b2063fd33fc8a2, "Maximum number of occurrences")
	flag.StringVar(&__obf_75dc0fa65ea3ee4a, "o", "./", "Output folder, required")
}

func main() {
	flag.Parse()
	defer func() {
		if __obf_e2258fa59931527f := recover(); __obf_e2258fa59931527f != nil {
			fmt.Println("[error]", __obf_e2258fa59931527f)
			os.Exit(-1)
		}
	}()
	if __obf_fba5742c8939d4d5 == "" {
		__obf_9bd24d2c36608d92("u")
	}
	if __obf_75dc0fa65ea3ee4a == "" {
		__obf_9bd24d2c36608d92("o")
	}
	if __obf_11b2063fd33fc8a2 <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	__obf_b8a072f370649f15, __obf_c3746095cbc77276 := NewTask(__obf_75dc0fa65ea3ee4a, __obf_fba5742c8939d4d5)
	if __obf_c3746095cbc77276 != nil {
		panic(__obf_c3746095cbc77276)
	}
	if __obf_c3746095cbc77276 := __obf_b8a072f370649f15.Start(__obf_11b2063fd33fc8a2); __obf_c3746095cbc77276 != nil {
		panic(__obf_c3746095cbc77276)
	}
	fmt.Println("Done!")
}

func __obf_9bd24d2c36608d92(__obf_d8423d39b6f4eea9 string) {
	panic("parameter '" + __obf_d8423d39b6f4eea9 + "' is required")
}
