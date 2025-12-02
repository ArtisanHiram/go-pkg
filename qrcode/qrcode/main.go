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
	__obf_4e7846922206d84b := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_4f35d582367206a8 := flag.Int("s", 256, "image size (pixel)")
	__obf_d604bbd70bf8d064 := flag.Bool("t", false, "print as text-art on stdout")
	__obf_64990083b12a4c4a := flag.Bool("i", false, "invert black and white")
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
		__obf_8fd953bd8420b5ac(fmt.Errorf("error: no content given"))
	}
	__obf_75cbdbb7871ac4cc := strings.Join(flag.Args(), " ")

	var __obf_bdbfd480b6a9ca38 error
	var __obf_1db07364e9d67bf9 *qrcode.QRCode
	__obf_1db07364e9d67bf9, __obf_bdbfd480b6a9ca38 = qrcode.New(__obf_75cbdbb7871ac4cc, qrcode.Highest)
	__obf_8fd953bd8420b5ac(__obf_bdbfd480b6a9ca38)

	if *Borderless {
		__obf_1db07364e9d67bf9.
			Borderless = true
	}

	if *__obf_d604bbd70bf8d064 {
		__obf_4f1d3498ce49e96f := __obf_1db07364e9d67bf9.ToString(*__obf_64990083b12a4c4a)
		fmt.Println(__obf_4f1d3498ce49e96f)
		return
	}

	if *__obf_64990083b12a4c4a {
		__obf_1db07364e9d67bf9.
			ForegroundColor, __obf_1db07364e9d67bf9.BackgroundColor = __obf_1db07364e9d67bf9.BackgroundColor, __obf_1db07364e9d67bf9.ForegroundColor
	}

	var __obf_d4c2bf955dd21f59 []byte
	__obf_d4c2bf955dd21f59, __obf_bdbfd480b6a9ca38 = __obf_1db07364e9d67bf9.PNG(*__obf_4f35d582367206a8)
	__obf_8fd953bd8420b5ac(__obf_bdbfd480b6a9ca38)

	if *__obf_4e7846922206d84b == "" {
		os.Stdout.Write(__obf_d4c2bf955dd21f59)
	} else {
		var __obf_edb6a87261d8c9c5 *os.File
		__obf_edb6a87261d8c9c5, __obf_bdbfd480b6a9ca38 = os.Create(*__obf_4e7846922206d84b + ".png")
		__obf_8fd953bd8420b5ac(__obf_bdbfd480b6a9ca38)
		defer __obf_edb6a87261d8c9c5.Close()
		__obf_edb6a87261d8c9c5.
			Write(__obf_d4c2bf955dd21f59)
	}
}

func __obf_8fd953bd8420b5ac(__obf_bdbfd480b6a9ca38 error) {
	if __obf_bdbfd480b6a9ca38 != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_bdbfd480b6a9ca38)
		os.Exit(1)
	}
}
