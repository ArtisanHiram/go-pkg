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
	__obf_abf9f67500079f7a := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_a2c1bded25dcd1d3 := flag.Int("s", 256, "image size (pixel)")
	__obf_319dec1a134e5c56 := flag.Bool("t", false, "print as text-art on stdout")
	__obf_ebddd617ff5632c1 := flag.Bool("i", false, "invert black and white")
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
		__obf_ba452ab3ca18bde4(fmt.Errorf("error: no content given"))
	}
	__obf_d675106cfc2927e1 := strings.Join(flag.Args(), " ")

	var __obf_716299f548770a5a error
	var __obf_f92063f9a2908c03 *qrcode.QRCode
	__obf_f92063f9a2908c03, __obf_716299f548770a5a = qrcode.New(__obf_d675106cfc2927e1, qrcode.Highest)
	__obf_ba452ab3ca18bde4(__obf_716299f548770a5a)

	if *Borderless {
		__obf_f92063f9a2908c03.
			Borderless = true
	}

	if *__obf_319dec1a134e5c56 {
		__obf_728ed67b96d6045c := __obf_f92063f9a2908c03.ToString(*__obf_ebddd617ff5632c1)
		fmt.Println(__obf_728ed67b96d6045c)
		return
	}

	if *__obf_ebddd617ff5632c1 {
		__obf_f92063f9a2908c03.
			ForegroundColor, __obf_f92063f9a2908c03.BackgroundColor = __obf_f92063f9a2908c03.BackgroundColor, __obf_f92063f9a2908c03.ForegroundColor
	}

	var __obf_56dbe61f7197d293 []byte
	__obf_56dbe61f7197d293, __obf_716299f548770a5a = __obf_f92063f9a2908c03.PNG(*__obf_a2c1bded25dcd1d3)
	__obf_ba452ab3ca18bde4(__obf_716299f548770a5a)

	if *__obf_abf9f67500079f7a == "" {
		os.Stdout.Write(__obf_56dbe61f7197d293)
	} else {
		var __obf_81de63fbf6cc9551 *os.File
		__obf_81de63fbf6cc9551, __obf_716299f548770a5a = os.Create(*__obf_abf9f67500079f7a + ".png")
		__obf_ba452ab3ca18bde4(__obf_716299f548770a5a)
		defer __obf_81de63fbf6cc9551.Close()
		__obf_81de63fbf6cc9551.
			Write(__obf_56dbe61f7197d293)
	}
}

func __obf_ba452ab3ca18bde4(__obf_716299f548770a5a error) {
	if __obf_716299f548770a5a != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_716299f548770a5a)
		os.Exit(1)
	}
}
