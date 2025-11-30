// go-qrcode
// Copyright 2014 Tom Harwood

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	qrcode "github.com/ArtisanHiram/go-pkg/qrcode"
)

func main() {
	__obf_d39440cd9076798d := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_88276e4abc0e9177 := flag.Int("s", 256, "image size (pixel)")
	__obf_94489eb496c4aced := flag.Bool("t", false, "print as text-art on stdout")
	__obf_7ef4f966cab02adc := flag.Bool("i", false, "invert black and white")
	Borderless := flag.Bool("b", false, "QR Code borderless or not")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `qrcode -- QR Code encoder in Go https://github.com/carmel/go-qrcode
Flags:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Usage:
  1. Arguments except for flags are joined by " " and used to generate QR code.
     Default output is STDOUT, pipe to imagemagick command "display" to display
     on any X server.

       qrcode hello word | display

  2. Save to file if "display" not available:

       qrcode "homepage: https://github.com/carmel/go-qrcode" > out.png

`)
	}
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		__obf_da321f4d28f3b957(fmt.Errorf("error: no content given"))
	}
	__obf_780237ab561642c4 := strings.Join(flag.Args(), " ")

	var __obf_b681ef89a0f35050 error
	var __obf_d22b310a000124ce *qrcode.QRCode
	__obf_d22b310a000124ce, __obf_b681ef89a0f35050 = qrcode.New(__obf_780237ab561642c4, qrcode.Highest)
	__obf_da321f4d28f3b957(__obf_b681ef89a0f35050)

	if *Borderless {
		__obf_d22b310a000124ce.
			Borderless = true
	}

	if *__obf_94489eb496c4aced {
		__obf_ac2ee32a9ebab12f := __obf_d22b310a000124ce.ToString(*__obf_7ef4f966cab02adc)
		fmt.Println(__obf_ac2ee32a9ebab12f)
		return
	}

	if *__obf_7ef4f966cab02adc {
		__obf_d22b310a000124ce.
			ForegroundColor, __obf_d22b310a000124ce.BackgroundColor = __obf_d22b310a000124ce.BackgroundColor, __obf_d22b310a000124ce.ForegroundColor
	}

	var __obf_bdd85c90d8f1b94e []byte
	__obf_bdd85c90d8f1b94e, __obf_b681ef89a0f35050 = __obf_d22b310a000124ce.PNG(*__obf_88276e4abc0e9177)
	__obf_da321f4d28f3b957(__obf_b681ef89a0f35050)

	if *__obf_d39440cd9076798d == "" {
		os.Stdout.Write(__obf_bdd85c90d8f1b94e)
	} else {
		var __obf_74b96bbacd24208f *os.File
		__obf_74b96bbacd24208f, __obf_b681ef89a0f35050 = os.Create(*__obf_d39440cd9076798d + ".png")
		__obf_da321f4d28f3b957(__obf_b681ef89a0f35050)
		defer __obf_74b96bbacd24208f.Close()
		__obf_74b96bbacd24208f.
			Write(__obf_bdd85c90d8f1b94e)
	}
}

func __obf_da321f4d28f3b957(__obf_b681ef89a0f35050 error) {
	if __obf_b681ef89a0f35050 != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_b681ef89a0f35050)
		os.Exit(1)
	}
}
