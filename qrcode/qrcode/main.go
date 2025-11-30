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
	__obf_fc3c6d7147e08a64 := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_cf8af38c4e70a950 := flag.Int("s", 256, "image size (pixel)")
	__obf_5758dbeebca500f0 := flag.Bool("t", false, "print as text-art on stdout")
	__obf_c82620279b8c9763 := flag.Bool("i", false, "invert black and white")
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
		__obf_3bd8a318e71d9fac(fmt.Errorf("error: no content given"))
	}
	__obf_172c386b56d52ed2 := strings.Join(flag.Args(), " ")

	var __obf_6926d398f9bd6dbf error
	var __obf_065dadba607edcc4 *qrcode.QRCode
	__obf_065dadba607edcc4, __obf_6926d398f9bd6dbf = qrcode.New(__obf_172c386b56d52ed2, qrcode.Highest)
	__obf_3bd8a318e71d9fac(__obf_6926d398f9bd6dbf)

	if *Borderless {
		__obf_065dadba607edcc4.
			Borderless = true
	}

	if *__obf_5758dbeebca500f0 {
		__obf_925cc8598e232373 := __obf_065dadba607edcc4.ToString(*__obf_c82620279b8c9763)
		fmt.Println(__obf_925cc8598e232373)
		return
	}

	if *__obf_c82620279b8c9763 {
		__obf_065dadba607edcc4.
			ForegroundColor, __obf_065dadba607edcc4.BackgroundColor = __obf_065dadba607edcc4.BackgroundColor, __obf_065dadba607edcc4.ForegroundColor
	}

	var __obf_525c70d4e2ab1dfe []byte
	__obf_525c70d4e2ab1dfe, __obf_6926d398f9bd6dbf = __obf_065dadba607edcc4.PNG(*__obf_cf8af38c4e70a950)
	__obf_3bd8a318e71d9fac(__obf_6926d398f9bd6dbf)

	if *__obf_fc3c6d7147e08a64 == "" {
		os.Stdout.Write(__obf_525c70d4e2ab1dfe)
	} else {
		var __obf_df7caa77904ffd74 *os.File
		__obf_df7caa77904ffd74, __obf_6926d398f9bd6dbf = os.Create(*__obf_fc3c6d7147e08a64 + ".png")
		__obf_3bd8a318e71d9fac(__obf_6926d398f9bd6dbf)
		defer __obf_df7caa77904ffd74.Close()
		__obf_df7caa77904ffd74.
			Write(__obf_525c70d4e2ab1dfe)
	}
}

func __obf_3bd8a318e71d9fac(__obf_6926d398f9bd6dbf error) {
	if __obf_6926d398f9bd6dbf != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_6926d398f9bd6dbf)
		os.Exit(1)
	}
}
