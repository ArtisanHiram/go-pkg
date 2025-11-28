// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_29ab288f1d4949da

import (
	"flag"
	"fmt"
	"os"
	"strings"

	qrcode "github.com/ArtisanHiram/go-pkg/qrcode"
)

func __obf_29ab288f1d4949da() {
	__obf_ba9078ce53f3e5c5 := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_655692963246709a := flag.Int("s", 256, "image size (pixel)")
	__obf_f90156b212ef8155 := flag.Bool("t", false, "print as text-art on stdout")
	__obf_6431291e51f58b55 := flag.Bool("i", false, "invert black and white")
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
		__obf_40b6038c7e5ef702(fmt.Errorf("error: no content given"))
	}

	__obf_f105366890fb3380 := strings.Join(flag.Args(), " ")

	var __obf_77be39223e3c6bc4 error
	var __obf_edc215702ba45948 *qrcode.QRCode
	__obf_edc215702ba45948, __obf_77be39223e3c6bc4 = qrcode.New(__obf_f105366890fb3380, qrcode.Highest)
	__obf_40b6038c7e5ef702(__obf_77be39223e3c6bc4)

	if *Borderless {
		__obf_edc215702ba45948.Borderless = true
	}

	if *__obf_f90156b212ef8155 {
		__obf_d036656a0c81ac6c := __obf_edc215702ba45948.ToString(*__obf_6431291e51f58b55)
		fmt.Println(__obf_d036656a0c81ac6c)
		return
	}

	if *__obf_6431291e51f58b55 {
		__obf_edc215702ba45948.ForegroundColor, __obf_edc215702ba45948.BackgroundColor = __obf_edc215702ba45948.BackgroundColor, __obf_edc215702ba45948.ForegroundColor
	}

	var __obf_67c0d633cc606c29 []byte
	__obf_67c0d633cc606c29, __obf_77be39223e3c6bc4 = __obf_edc215702ba45948.PNG(*__obf_655692963246709a)
	__obf_40b6038c7e5ef702(__obf_77be39223e3c6bc4)

	if *__obf_ba9078ce53f3e5c5 == "" {
		os.Stdout.Write(__obf_67c0d633cc606c29)
	} else {
		var __obf_b5eefff29b3c6e4b *os.File
		__obf_b5eefff29b3c6e4b, __obf_77be39223e3c6bc4 = os.Create(*__obf_ba9078ce53f3e5c5 + ".png")
		__obf_40b6038c7e5ef702(__obf_77be39223e3c6bc4)
		defer __obf_b5eefff29b3c6e4b.Close()
		__obf_b5eefff29b3c6e4b.Write(__obf_67c0d633cc606c29)
	}
}

func __obf_40b6038c7e5ef702(__obf_77be39223e3c6bc4 error) {
	if __obf_77be39223e3c6bc4 != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_77be39223e3c6bc4)
		os.Exit(1)
	}
}
