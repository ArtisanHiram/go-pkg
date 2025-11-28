package __obf_d8b53a99900c2ed7

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
	reedsolomon "github.com/ArtisanHiram/go-pkg/qrcode/reedsolomon"
)

// Encode a QR Code and return a raw PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
//
// To serve over HTTP, remember to send a Content-Type: image/png header.
func Encode(__obf_89ac1aaea9af4356 string, __obf_14fc221ac535188d RecoveryLevel, __obf_658024c8e4ec9df7 int) ([]byte, error) {
	var __obf_c75d8da5927c1067 *QRCode

	__obf_c75d8da5927c1067, __obf_f3cf37a038d12d80 := New(__obf_89ac1aaea9af4356, __obf_14fc221ac535188d)

	if __obf_f3cf37a038d12d80 != nil {
		return nil, __obf_f3cf37a038d12d80
	}

	return __obf_c75d8da5927c1067.PNG(__obf_658024c8e4ec9df7)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_89ac1aaea9af4356 string, __obf_14fc221ac535188d RecoveryLevel, __obf_658024c8e4ec9df7 int, __obf_6f439a57b7efb1a7 string) error {
	var __obf_c75d8da5927c1067 *QRCode

	__obf_c75d8da5927c1067, __obf_f3cf37a038d12d80 := New(__obf_89ac1aaea9af4356, __obf_14fc221ac535188d)

	if __obf_f3cf37a038d12d80 != nil {
		return __obf_f3cf37a038d12d80
	}

	return __obf_c75d8da5927c1067.WriteFile(__obf_658024c8e4ec9df7, __obf_6f439a57b7efb1a7)
}

// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
// With WriteColorFile you can also specify the colors you want to use.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteColorFile(__obf_89ac1aaea9af4356 string, __obf_14fc221ac535188d RecoveryLevel, __obf_658024c8e4ec9df7 int, __obf_fb7fb7beb312276b,
	__obf_25bf6e6c802629f6 color.Color, __obf_6f439a57b7efb1a7 string) error {

	var __obf_c75d8da5927c1067 *QRCode

	__obf_c75d8da5927c1067, __obf_f3cf37a038d12d80 := New(__obf_89ac1aaea9af4356, __obf_14fc221ac535188d)

	__obf_c75d8da5927c1067.BackgroundColor = __obf_fb7fb7beb312276b
	__obf_c75d8da5927c1067.ForegroundColor = __obf_25bf6e6c802629f6

	if __obf_f3cf37a038d12d80 != nil {
		return __obf_f3cf37a038d12d80
	}

	return __obf_c75d8da5927c1067.WriteFile(__obf_658024c8e4ec9df7, __obf_6f439a57b7efb1a7)
}

// A QRCode represents a valid encoded QRCode.
type QRCode struct {
	// Original content encoded.
	Content string

	// QR Code type.
	Level         RecoveryLevel
	VersionNumber int

	// User settable drawing options.
	ForegroundColor color.Color
	BackgroundColor color.Color

	// Disable the QR Code border.
	Borderless bool

	__obf_e2ebde233760b639 *__obf_c5fc77e1fc362b78
	__obf_710138c0bddc8a72 __obf_76b5b454b361aca4

	__obf_597c6046cd04cc8f *bitset.Bitset
	__obf_dbf6f541061eeeb5 *__obf_dbf6f541061eeeb5
	__obf_5f3b1446be4be5d2 int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_89ac1aaea9af4356 string, __obf_14fc221ac535188d RecoveryLevel) (*QRCode, error) {
	__obf_0ce521758dc385a0 := []__obf_80ac4aaf2f637970{__obf_c531a6f523406e39, __obf_c0bbd372ece2268e,
		__obf_14b09e7535ac520c}

	var __obf_e2ebde233760b639 *__obf_c5fc77e1fc362b78
	var __obf_db73de0d3577b11c *bitset.Bitset
	var __obf_d50eff52bcf98db2 *__obf_76b5b454b361aca4
	var __obf_f3cf37a038d12d80 error

	for _, __obf_c40c4d1c8e73185e := range __obf_0ce521758dc385a0 {
		__obf_e2ebde233760b639 = __obf_0a11f28667b2f0e3(__obf_c40c4d1c8e73185e)
		__obf_db73de0d3577b11c, __obf_f3cf37a038d12d80 = __obf_e2ebde233760b639.__obf_d69ee0563ed868f4([]byte(__obf_89ac1aaea9af4356))

		if __obf_f3cf37a038d12d80 != nil {
			continue
		}

		__obf_d50eff52bcf98db2 = __obf_4312a51a81d74153(__obf_14fc221ac535188d, __obf_e2ebde233760b639, __obf_db73de0d3577b11c.Len())

		if __obf_d50eff52bcf98db2 != nil {
			break
		}
	}

	if __obf_f3cf37a038d12d80 != nil {
		return nil, __obf_f3cf37a038d12d80
	} else if __obf_d50eff52bcf98db2 == nil {
		return nil, errors.New("content too long to encode")
	}

	__obf_c75d8da5927c1067 := &QRCode{
		Content: __obf_89ac1aaea9af4356,

		Level:         __obf_14fc221ac535188d,
		VersionNumber: __obf_d50eff52bcf98db2.__obf_710138c0bddc8a72,

		ForegroundColor: color.Black,
		BackgroundColor: color.White,

		__obf_e2ebde233760b639: __obf_e2ebde233760b639,
		__obf_597c6046cd04cc8f: __obf_db73de0d3577b11c,
		__obf_710138c0bddc8a72: *__obf_d50eff52bcf98db2,
	}

	return __obf_c75d8da5927c1067, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_89ac1aaea9af4356 string, __obf_710138c0bddc8a72 int, __obf_14fc221ac535188d RecoveryLevel) (*QRCode, error) {
	var __obf_e2ebde233760b639 *__obf_c5fc77e1fc362b78

	switch {
	case __obf_710138c0bddc8a72 >= 1 && __obf_710138c0bddc8a72 <= 9:
		__obf_e2ebde233760b639 = __obf_0a11f28667b2f0e3(__obf_c531a6f523406e39)
	case __obf_710138c0bddc8a72 >= 10 && __obf_710138c0bddc8a72 <= 26:
		__obf_e2ebde233760b639 = __obf_0a11f28667b2f0e3(__obf_c0bbd372ece2268e)
	case __obf_710138c0bddc8a72 >= 27 && __obf_710138c0bddc8a72 <= 40:
		__obf_e2ebde233760b639 = __obf_0a11f28667b2f0e3(__obf_14b09e7535ac520c)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_710138c0bddc8a72)
	}

	var __obf_db73de0d3577b11c *bitset.Bitset
	__obf_db73de0d3577b11c, __obf_f3cf37a038d12d80 := __obf_e2ebde233760b639.__obf_d69ee0563ed868f4([]byte(__obf_89ac1aaea9af4356))

	if __obf_f3cf37a038d12d80 != nil {
		return nil, __obf_f3cf37a038d12d80
	}

	__obf_d50eff52bcf98db2 := __obf_edc406ccd0ae2ee8(__obf_14fc221ac535188d, __obf_710138c0bddc8a72)

	if __obf_d50eff52bcf98db2 == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_db73de0d3577b11c.Len() > __obf_d50eff52bcf98db2.__obf_c353137244584482() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)",
			__obf_710138c0bddc8a72,
			__obf_db73de0d3577b11c.Len(),
			__obf_d50eff52bcf98db2.__obf_c353137244584482())
	}

	__obf_c75d8da5927c1067 := &QRCode{
		Content: __obf_89ac1aaea9af4356,

		Level:         __obf_14fc221ac535188d,
		VersionNumber: __obf_d50eff52bcf98db2.__obf_710138c0bddc8a72,

		ForegroundColor: color.Black,
		BackgroundColor: color.White,

		__obf_e2ebde233760b639: __obf_e2ebde233760b639,
		__obf_597c6046cd04cc8f: __obf_db73de0d3577b11c,
		__obf_710138c0bddc8a72: *__obf_d50eff52bcf98db2,
	}

	return __obf_c75d8da5927c1067, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_c75d8da5927c1067 *QRCode) Bitmap() [][]bool {
	// Build QR code.
	__obf_c75d8da5927c1067.__obf_d69ee0563ed868f4()

	return __obf_c75d8da5927c1067.__obf_dbf6f541061eeeb5.__obf_86e6f11fca3fe468()
}

// Image returns the QR Code as an image.Image.
//
// A positive size sets a fixed image width and height (e.g. 256 yields an
// 256x256px image).
//
// Depending on the amount of data encoded, fixed size images can have different
// amounts of padding (white space around the QR Code). As an alternative, a
// variable sized image can be generated instead:
//
// A negative size causes a variable sized image to be returned. The image
// returned is the minimum size required for the QR Code. Choose a larger
// negative number to increase the scale of the image. e.g. a size of -5 causes
// each module (QR Code "pixel") to be 5px in size.
func (__obf_c75d8da5927c1067 *QRCode) Image(__obf_658024c8e4ec9df7 int) image.Image {
	// Build QR code.
	__obf_c75d8da5927c1067.__obf_d69ee0563ed868f4()

	// Minimum pixels (both width and height) required.
	__obf_2f57903f822bed40 := __obf_c75d8da5927c1067.__obf_dbf6f541061eeeb5.__obf_658024c8e4ec9df7

	// Variable size support.
	if __obf_658024c8e4ec9df7 < 0 {
		__obf_658024c8e4ec9df7 = __obf_658024c8e4ec9df7 * -1 * __obf_2f57903f822bed40
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_658024c8e4ec9df7 < __obf_2f57903f822bed40 {
		__obf_658024c8e4ec9df7 = __obf_2f57903f822bed40
	}

	// Output image.
	__obf_c0fc79f5659efe1f := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_658024c8e4ec9df7, __obf_658024c8e4ec9df7}}

	// Saves a few bytes to have them in this order
	__obf_8183023ece8d232e := color.Palette([]color.Color{__obf_c75d8da5927c1067.BackgroundColor, __obf_c75d8da5927c1067.ForegroundColor})
	__obf_440964e6cf6618a5 := image.NewPaletted(__obf_c0fc79f5659efe1f, __obf_8183023ece8d232e)
	__obf_6cc3245f897d8e68 := uint8(__obf_440964e6cf6618a5.Palette.Index(__obf_c75d8da5927c1067.ForegroundColor))

	// QR code bitmap.
	__obf_86e6f11fca3fe468 := __obf_c75d8da5927c1067.__obf_dbf6f541061eeeb5.__obf_86e6f11fca3fe468()

	// Map each image pixel to the nearest QR code module.
	__obf_b2e27c025092907a := float64(__obf_2f57903f822bed40) / float64(__obf_658024c8e4ec9df7)
	for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < __obf_658024c8e4ec9df7; __obf_349392c72c5c7113++ {
		__obf_b0ce45b43efb5a07 := int(float64(__obf_349392c72c5c7113) * __obf_b2e27c025092907a)
		for __obf_2f6be407475f551f := 0; __obf_2f6be407475f551f < __obf_658024c8e4ec9df7; __obf_2f6be407475f551f++ {
			__obf_e0488ee550dac6f8 := int(float64(__obf_2f6be407475f551f) * __obf_b2e27c025092907a)

			__obf_4650eaa332d33944 := __obf_86e6f11fca3fe468[__obf_b0ce45b43efb5a07][__obf_e0488ee550dac6f8]

			if __obf_4650eaa332d33944 {
				__obf_20616a4614a83ed1 := __obf_440964e6cf6618a5.PixOffset(__obf_2f6be407475f551f, __obf_349392c72c5c7113)
				__obf_440964e6cf6618a5.Pix[__obf_20616a4614a83ed1] = __obf_6cc3245f897d8e68
			}
		}
	}

	return __obf_440964e6cf6618a5
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_c75d8da5927c1067 *QRCode) PNG(__obf_658024c8e4ec9df7 int) ([]byte, error) {
	__obf_440964e6cf6618a5 := __obf_c75d8da5927c1067.Image(__obf_658024c8e4ec9df7)

	__obf_e2ebde233760b639 := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_86534ca54b6c1d7e bytes.Buffer
	__obf_f3cf37a038d12d80 := __obf_e2ebde233760b639.Encode(&__obf_86534ca54b6c1d7e, __obf_440964e6cf6618a5)

	if __obf_f3cf37a038d12d80 != nil {
		return nil, __obf_f3cf37a038d12d80
	}

	return __obf_86534ca54b6c1d7e.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_c75d8da5927c1067 *QRCode) Write(__obf_658024c8e4ec9df7 int, __obf_4d3b014f5312ae55 io.Writer) error {
	var png []byte

	png, __obf_f3cf37a038d12d80 := __obf_c75d8da5927c1067.PNG(__obf_658024c8e4ec9df7)

	if __obf_f3cf37a038d12d80 != nil {
		return __obf_f3cf37a038d12d80
	}
	_, __obf_f3cf37a038d12d80 = __obf_4d3b014f5312ae55.Write(png)
	return __obf_f3cf37a038d12d80
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_c75d8da5927c1067 *QRCode) WriteFile(__obf_658024c8e4ec9df7 int, __obf_6f439a57b7efb1a7 string) error {
	var png []byte

	png, __obf_f3cf37a038d12d80 := __obf_c75d8da5927c1067.PNG(__obf_658024c8e4ec9df7)

	if __obf_f3cf37a038d12d80 != nil {
		return __obf_f3cf37a038d12d80
	}

	return os.WriteFile(__obf_6f439a57b7efb1a7, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_c75d8da5927c1067 *QRCode) __obf_d69ee0563ed868f4() {
	__obf_9865e6bb0b1ac45d := __obf_c75d8da5927c1067.__obf_710138c0bddc8a72.__obf_7b333c1169728dbe(__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Len())

	__obf_c75d8da5927c1067.__obf_7e32d8dbcbb4f098(__obf_9865e6bb0b1ac45d)
	__obf_c75d8da5927c1067.__obf_36845ad3dc487070()

	__obf_db73de0d3577b11c := __obf_c75d8da5927c1067.__obf_472d82495118feee()

	const __obf_5aea21259fa2e577 int = 8
	__obf_e751bf876b38498b := 0

	for __obf_5f3b1446be4be5d2 := 0; __obf_5f3b1446be4be5d2 < __obf_5aea21259fa2e577; __obf_5f3b1446be4be5d2++ {
		var __obf_da19f8f4aed522cb *__obf_dbf6f541061eeeb5
		var __obf_f3cf37a038d12d80 error

		__obf_da19f8f4aed522cb, __obf_f3cf37a038d12d80 = __obf_2f378f59e7aa349a(__obf_c75d8da5927c1067.__obf_710138c0bddc8a72, __obf_5f3b1446be4be5d2, __obf_db73de0d3577b11c, !__obf_c75d8da5927c1067.Borderless)

		if __obf_f3cf37a038d12d80 != nil {
			log.Panic(__obf_f3cf37a038d12d80.Error())
		}

		__obf_1362019ac7bdc111 := __obf_da19f8f4aed522cb.__obf_1362019ac7bdc111()
		if __obf_1362019ac7bdc111 != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)",
				__obf_1362019ac7bdc111, __obf_c75d8da5927c1067.VersionNumber)
		}

		__obf_8183023ece8d232e := __obf_da19f8f4aed522cb.__obf_2ce595e50aac84df()

		//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())

		if __obf_c75d8da5927c1067.__obf_dbf6f541061eeeb5 == nil || __obf_8183023ece8d232e < __obf_e751bf876b38498b {
			__obf_c75d8da5927c1067.__obf_dbf6f541061eeeb5 = __obf_da19f8f4aed522cb
			__obf_c75d8da5927c1067.__obf_5f3b1446be4be5d2 = __obf_5f3b1446be4be5d2
			__obf_e751bf876b38498b = __obf_8183023ece8d232e
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_c75d8da5927c1067 *QRCode) __obf_7e32d8dbcbb4f098(__obf_9865e6bb0b1ac45d int) {
	__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.AppendNumBools(__obf_9865e6bb0b1ac45d, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_c75d8da5927c1067 *QRCode) __obf_472d82495118feee() *bitset.Bitset {
	// Split into blocks.
	type __obf_956b00fe0df2c201 struct {
		__obf_597c6046cd04cc8f *bitset.Bitset
		__obf_af3ebcc20b2e0ed6 int
	}

	__obf_25d864324bfece84 := make([]__obf_956b00fe0df2c201, __obf_c75d8da5927c1067.__obf_710138c0bddc8a72.__obf_9ce91964fbf227fa())

	__obf_c6edaa3d3a18688d := 0
	__obf_8ba3e53b76b5dd4a := 0
	__obf_5e394f85a15a8d6f := 0

	for _, __obf_86534ca54b6c1d7e := range __obf_c75d8da5927c1067.__obf_710138c0bddc8a72.__obf_25d864324bfece84 {
		for __obf_b09a06cc5f53e0ed := 0; __obf_b09a06cc5f53e0ed < __obf_86534ca54b6c1d7e.__obf_9ce91964fbf227fa; __obf_b09a06cc5f53e0ed++ {
			__obf_c6edaa3d3a18688d = __obf_8ba3e53b76b5dd4a
			__obf_8ba3e53b76b5dd4a = __obf_c6edaa3d3a18688d + __obf_86534ca54b6c1d7e.__obf_235ba13aede89443*8

			// Apply error correction to each block.
			__obf_c4a3390e7b5b4ecb := __obf_86534ca54b6c1d7e.__obf_5fec32634d8d8a1d - __obf_86534ca54b6c1d7e.__obf_235ba13aede89443
			__obf_25d864324bfece84[__obf_5e394f85a15a8d6f].__obf_597c6046cd04cc8f = reedsolomon.Encode(__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Substr(__obf_c6edaa3d3a18688d, __obf_8ba3e53b76b5dd4a), __obf_c4a3390e7b5b4ecb)
			__obf_25d864324bfece84[__obf_5e394f85a15a8d6f].__obf_af3ebcc20b2e0ed6 = __obf_8ba3e53b76b5dd4a - __obf_c6edaa3d3a18688d

			__obf_5e394f85a15a8d6f++
		}
	}

	// Interleave the blocks.

	__obf_9afe29b453625a5f := bitset.New()

	// Combine data blocks.
	__obf_a438039674bde83d := true
	for __obf_6ffda43d9ecd66ed := 0; __obf_a438039674bde83d; __obf_6ffda43d9ecd66ed += 8 {
		__obf_a438039674bde83d = false

		for __obf_b09a06cc5f53e0ed, __obf_86534ca54b6c1d7e := range __obf_25d864324bfece84 {
			if __obf_6ffda43d9ecd66ed >= __obf_25d864324bfece84[__obf_b09a06cc5f53e0ed].__obf_af3ebcc20b2e0ed6 {
				continue
			}

			__obf_9afe29b453625a5f.Append(__obf_86534ca54b6c1d7e.__obf_597c6046cd04cc8f.Substr(__obf_6ffda43d9ecd66ed, __obf_6ffda43d9ecd66ed+8))

			__obf_a438039674bde83d = true
		}
	}

	// Combine error correction blocks.
	__obf_a438039674bde83d = true
	for __obf_6ffda43d9ecd66ed := 0; __obf_a438039674bde83d; __obf_6ffda43d9ecd66ed += 8 {
		__obf_a438039674bde83d = false

		for __obf_b09a06cc5f53e0ed, __obf_86534ca54b6c1d7e := range __obf_25d864324bfece84 {
			__obf_3d6c564e8bf580b9 := __obf_6ffda43d9ecd66ed + __obf_25d864324bfece84[__obf_b09a06cc5f53e0ed].__obf_af3ebcc20b2e0ed6
			if __obf_3d6c564e8bf580b9 >= __obf_25d864324bfece84[__obf_b09a06cc5f53e0ed].__obf_597c6046cd04cc8f.Len() {
				continue
			}

			__obf_9afe29b453625a5f.Append(__obf_86534ca54b6c1d7e.__obf_597c6046cd04cc8f.Substr(__obf_3d6c564e8bf580b9, __obf_3d6c564e8bf580b9+8))

			__obf_a438039674bde83d = true
		}
	}

	// Append remainder bits.
	__obf_9afe29b453625a5f.AppendNumBools(__obf_c75d8da5927c1067.__obf_710138c0bddc8a72.__obf_c2659eba064f3784, false)

	return __obf_9afe29b453625a5f
}

// addPadding pads the encoded data upto the full length required.
func (__obf_c75d8da5927c1067 *QRCode) __obf_36845ad3dc487070() {
	__obf_c353137244584482 := __obf_c75d8da5927c1067.__obf_710138c0bddc8a72.__obf_c353137244584482()

	if __obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Len() == __obf_c353137244584482 {
		return
	}

	// Pad to the nearest codeword boundary.
	__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.AppendNumBools(__obf_c75d8da5927c1067.__obf_710138c0bddc8a72.__obf_fca930c165cbee29(__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Len()), false)

	// Pad codewords 0b11101100 and 0b00010001.
	__obf_a34b5d4d29b646cd := [2]*bitset.Bitset{
		bitset.New(true, true, true, false, true, true, false, false),
		bitset.New(false, false, false, true, false, false, false, true),
	}

	// Insert pad codewords alternately.
	__obf_6ffda43d9ecd66ed := 0
	for __obf_c353137244584482-__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Len() >= 8 {
		__obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Append(__obf_a34b5d4d29b646cd[__obf_6ffda43d9ecd66ed])

		__obf_6ffda43d9ecd66ed = 1 - __obf_6ffda43d9ecd66ed // Alternate between 0 and 1.
	}

	if __obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Len() != __obf_c353137244584482 {
		log.Panicf("BUG: got len %d, expected %d", __obf_c75d8da5927c1067.__obf_597c6046cd04cc8f.Len(), __obf_c353137244584482)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_c75d8da5927c1067 *QRCode) ToString(__obf_f7f046c3b32e3870 bool) string {
	__obf_5e143d3cdb7f0128 := __obf_c75d8da5927c1067.Bitmap()
	var __obf_f0e134e3438b87ed bytes.Buffer
	for __obf_349392c72c5c7113 := range __obf_5e143d3cdb7f0128 {
		for __obf_2f6be407475f551f := range __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113] {
			if __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113][__obf_2f6be407475f551f] != __obf_f7f046c3b32e3870 {
				__obf_f0e134e3438b87ed.WriteString("  ")
			} else {
				__obf_f0e134e3438b87ed.WriteString("██")
			}
		}
		__obf_f0e134e3438b87ed.WriteString("\n")
	}
	return __obf_f0e134e3438b87ed.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_c75d8da5927c1067 *QRCode) ToSmallString(__obf_f7f046c3b32e3870 bool) string {
	__obf_5e143d3cdb7f0128 := __obf_c75d8da5927c1067.Bitmap()
	var __obf_f0e134e3438b87ed bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < len(__obf_5e143d3cdb7f0128)-1; __obf_349392c72c5c7113 += 2 {
		for __obf_2f6be407475f551f := range __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113] {
			if __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113][__obf_2f6be407475f551f] == __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113+1][__obf_2f6be407475f551f] {
				if __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113][__obf_2f6be407475f551f] != __obf_f7f046c3b32e3870 {
					__obf_f0e134e3438b87ed.WriteString(" ")
				} else {
					__obf_f0e134e3438b87ed.WriteString("█")
				}
			} else {
				if __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113][__obf_2f6be407475f551f] != __obf_f7f046c3b32e3870 {
					__obf_f0e134e3438b87ed.WriteString("▄")
				} else {
					__obf_f0e134e3438b87ed.WriteString("▀")
				}
			}
		}
		__obf_f0e134e3438b87ed.WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_5e143d3cdb7f0128)%2 == 1 {
		__obf_349392c72c5c7113 := len(__obf_5e143d3cdb7f0128) - 1
		for __obf_2f6be407475f551f := range __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113] {
			if __obf_5e143d3cdb7f0128[__obf_349392c72c5c7113][__obf_2f6be407475f551f] != __obf_f7f046c3b32e3870 {
				__obf_f0e134e3438b87ed.WriteString(" ")
			} else {
				__obf_f0e134e3438b87ed.WriteString("▀")
			}
		}
		__obf_f0e134e3438b87ed.WriteString("\n")
	}
	return __obf_f0e134e3438b87ed.String()
}
