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
	__obf_e248d8b7f3df44d4 := flag.String("o", "", "out PNG file prefix, empty for stdout")
	__obf_6308ae60c5ccb4ed := flag.Int("s", 256, "image size (pixel)")
	__obf_a90c9b150eb92bde := flag.Bool("t", false, "print as text-art on stdout")
	__obf_db73c8705694a2d7 := flag.Bool("i", false, "invert black and white")
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
		__obf_753de98598b14a49(fmt.Errorf("error: no content given"))
	}
	__obf_663be6e1c437776f := strings.Join(flag.Args(), " ")

	var __obf_6e0f9e653a73e22e error
	var __obf_bbc609cb17d79925 *qrcode.QRCode
	__obf_bbc609cb17d79925, __obf_6e0f9e653a73e22e = qrcode.New(__obf_663be6e1c437776f, qrcode.Highest)
	__obf_753de98598b14a49(__obf_6e0f9e653a73e22e)

	if *Borderless {
		__obf_bbc609cb17d79925.
			Borderless = true
	}

	if *__obf_a90c9b150eb92bde {
		__obf_1c7442f0cc06d716 := __obf_bbc609cb17d79925.ToString(*__obf_db73c8705694a2d7)
		fmt.Println(__obf_1c7442f0cc06d716)
		return
	}

	if *__obf_db73c8705694a2d7 {
		__obf_bbc609cb17d79925.
			ForegroundColor, __obf_bbc609cb17d79925.BackgroundColor = __obf_bbc609cb17d79925.BackgroundColor, __obf_bbc609cb17d79925.ForegroundColor
	}

	var __obf_24a4ccb8f420cbb6 []byte
	__obf_24a4ccb8f420cbb6, __obf_6e0f9e653a73e22e = __obf_bbc609cb17d79925.PNG(*__obf_6308ae60c5ccb4ed)
	__obf_753de98598b14a49(__obf_6e0f9e653a73e22e)

	if *__obf_e248d8b7f3df44d4 == "" {
		os.Stdout.Write(__obf_24a4ccb8f420cbb6)
	} else {
		var __obf_149fc002fc8a1683 *os.File
		__obf_149fc002fc8a1683, __obf_6e0f9e653a73e22e = os.Create(*__obf_e248d8b7f3df44d4 + ".png")
		__obf_753de98598b14a49(__obf_6e0f9e653a73e22e)
		defer __obf_149fc002fc8a1683.Close()
		__obf_149fc002fc8a1683.
			Write(__obf_24a4ccb8f420cbb6)
	}
}

func __obf_753de98598b14a49(__obf_6e0f9e653a73e22e error) {
	if __obf_6e0f9e653a73e22e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", __obf_6e0f9e653a73e22e)
		os.Exit(1)
	}
}
