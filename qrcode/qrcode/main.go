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
	__obf_5c1bc8fbc5e5afae := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_aecfe90633f5437e := flag.Int("s", 256, "image size (pixel)")
	__obf_3a9ba3e38c53031d := flag.Bool("t", false, "print as text-art on stdout")
	__obf_4c9a1373e504117b := flag.Bool("i", false, "invert black and white")
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
		__obf_28b3da82ec4535ec(fmt.Errorf("error: no content given"))
	}
	__obf_2bc679c14d9af6fe := strings.Join(flag.Args(), " ")

	var __obf_5c5e106c45896e2a error
	var __obf_9d18f6200f59978a *qrcode.QRCode
	__obf_9d18f6200f59978a, __obf_5c5e106c45896e2a = qrcode.New(__obf_2bc679c14d9af6fe, qrcode.Highest)
	__obf_28b3da82ec4535ec(__obf_5c5e106c45896e2a)

	if *Borderless {
		__obf_9d18f6200f59978a.
			Borderless = true
	}

	if *__obf_3a9ba3e38c53031d {
		__obf_98909d8f08186bdf := __obf_9d18f6200f59978a.ToString(*__obf_4c9a1373e504117b)
		fmt.Println(__obf_98909d8f08186bdf)
		return
	}

	if *__obf_4c9a1373e504117b {
		__obf_9d18f6200f59978a.
			ForegroundColor, __obf_9d18f6200f59978a.BackgroundColor = __obf_9d18f6200f59978a.BackgroundColor, __obf_9d18f6200f59978a.ForegroundColor
	}

	var __obf_5bc9c647ef16a2e7 []byte
	__obf_5bc9c647ef16a2e7, __obf_5c5e106c45896e2a = __obf_9d18f6200f59978a.PNG(*__obf_aecfe90633f5437e)
	__obf_28b3da82ec4535ec(__obf_5c5e106c45896e2a)

	if *__obf_5c1bc8fbc5e5afae == "" {
		os.Stdout.Write(__obf_5bc9c647ef16a2e7)
	} else {
		var __obf_1668a6fcf6ff72e8 *os.File
		__obf_1668a6fcf6ff72e8, __obf_5c5e106c45896e2a = os.Create(*__obf_5c1bc8fbc5e5afae + ".png")
		__obf_28b3da82ec4535ec(__obf_5c5e106c45896e2a)
		defer __obf_1668a6fcf6ff72e8.Close()
		__obf_1668a6fcf6ff72e8.
			Write(__obf_5bc9c647ef16a2e7)
	}
}

func __obf_28b3da82ec4535ec(__obf_5c5e106c45896e2a error) {
	if __obf_5c5e106c45896e2a != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_5c5e106c45896e2a)
		os.Exit(1)
	}
}
