package __obf_48b3d686b985a20b

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
func Encode(__obf_42ae8f69ca3d94e4 string, __obf_28ff6edf33e2b2ce RecoveryLevel, __obf_311d29b1744e83e0 int) ([]byte, error) {
	var __obf_9c281bbbde209bea *QRCode
	__obf_9c281bbbde209bea, __obf_3f84ccce09b63a68 := New(__obf_42ae8f69ca3d94e4, __obf_28ff6edf33e2b2ce)

	if __obf_3f84ccce09b63a68 != nil {
		return nil, __obf_3f84ccce09b63a68
	}

	return __obf_9c281bbbde209bea.PNG(__obf_311d29b1744e83e0)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_42ae8f69ca3d94e4 string, __obf_28ff6edf33e2b2ce RecoveryLevel, __obf_311d29b1744e83e0 int, __obf_ffac0de5194f7875 string) error {
	var __obf_9c281bbbde209bea *QRCode
	__obf_9c281bbbde209bea, __obf_3f84ccce09b63a68 := New(__obf_42ae8f69ca3d94e4, __obf_28ff6edf33e2b2ce)

	if __obf_3f84ccce09b63a68 != nil {
		return __obf_3f84ccce09b63a68
	}

	return __obf_9c281bbbde209bea.WriteFile(__obf_311d29b1744e83e0,

		// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
		// With WriteColorFile you can also specify the colors you want to use.
		//
		// size is both the image width and height in pixels. If size is too small then
		// a larger image is silently written. Negative values for size cause a variable
		// sized image to be written: See the documentation for Image().
		__obf_ffac0de5194f7875)
}

func WriteColorFile(__obf_42ae8f69ca3d94e4 string, __obf_28ff6edf33e2b2ce RecoveryLevel, __obf_311d29b1744e83e0 int, __obf_ed1a26991796b7ee, __obf_6d9997cb21d52d96 color.Color, __obf_ffac0de5194f7875 string) error {

	var __obf_9c281bbbde209bea *QRCode
	__obf_9c281bbbde209bea, __obf_3f84ccce09b63a68 := New(__obf_42ae8f69ca3d94e4, __obf_28ff6edf33e2b2ce)
	__obf_9c281bbbde209bea.
		BackgroundColor = __obf_ed1a26991796b7ee
	__obf_9c281bbbde209bea.
		ForegroundColor = __obf_6d9997cb21d52d96

	if __obf_3f84ccce09b63a68 != nil {
		return __obf_3f84ccce09b63a68
	}

	return __obf_9c281bbbde209bea.WriteFile(__obf_311d29b1744e83e0,

		// A QRCode represents a valid encoded QRCode.
		__obf_ffac0de5194f7875)
}

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
	Borderless             bool
	__obf_af5b88a292691c3c *__obf_5a535fd3d1104862
	__obf_4f23b6c0c8a29b07 __obf_166f0ff41181b0b8
	__obf_e4df198c6616985f *bitset.Bitset
	__obf_14306cbbf0004950 *__obf_14306cbbf0004950
	__obf_0bb74c1e8aa906af int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_42ae8f69ca3d94e4 string, __obf_28ff6edf33e2b2ce RecoveryLevel) (*QRCode, error) {
	__obf_82b8685d9fc05f5c := []__obf_2818b04b6d379c6a{__obf_2ddbe5e0d50f3fdb, __obf_d4c11dce5ccd0394, __obf_e21e9a063b84f8a4}

	var __obf_af5b88a292691c3c *__obf_5a535fd3d1104862
	var __obf_ba2373eb7d967534 *bitset.Bitset
	var __obf_fb47b7063cd80a24 *__obf_166f0ff41181b0b8
	var __obf_3f84ccce09b63a68 error

	for _, __obf_05af1e63275cf6cc := range __obf_82b8685d9fc05f5c {
		__obf_af5b88a292691c3c = __obf_fccacc671db5d5f3(__obf_05af1e63275cf6cc)
		__obf_ba2373eb7d967534, __obf_3f84ccce09b63a68 = __obf_af5b88a292691c3c.__obf_5fa94477cdc64b8f([]byte(__obf_42ae8f69ca3d94e4))

		if __obf_3f84ccce09b63a68 != nil {
			continue
		}
		__obf_fb47b7063cd80a24 = __obf_d50edf8706752281(__obf_28ff6edf33e2b2ce, __obf_af5b88a292691c3c, __obf_ba2373eb7d967534.Len())

		if __obf_fb47b7063cd80a24 != nil {
			break
		}
	}

	if __obf_3f84ccce09b63a68 != nil {
		return nil, __obf_3f84ccce09b63a68
	} else if __obf_fb47b7063cd80a24 == nil {
		return nil, errors.New("content too long to encode")
	}
	__obf_9c281bbbde209bea := &QRCode{
		Content: __obf_42ae8f69ca3d94e4,

		Level:         __obf_28ff6edf33e2b2ce,
		VersionNumber: __obf_fb47b7063cd80a24.__obf_4f23b6c0c8a29b07,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_af5b88a292691c3c: __obf_af5b88a292691c3c, __obf_e4df198c6616985f: __obf_ba2373eb7d967534, __obf_4f23b6c0c8a29b07: *__obf_fb47b7063cd80a24,
	}

	return __obf_9c281bbbde209bea, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_42ae8f69ca3d94e4 string, __obf_4f23b6c0c8a29b07 int, __obf_28ff6edf33e2b2ce RecoveryLevel) (*QRCode, error) {
	var __obf_af5b88a292691c3c *__obf_5a535fd3d1104862

	switch {
	case __obf_4f23b6c0c8a29b07 >= 1 && __obf_4f23b6c0c8a29b07 <= 9:
		__obf_af5b88a292691c3c = __obf_fccacc671db5d5f3(__obf_2ddbe5e0d50f3fdb)
	case __obf_4f23b6c0c8a29b07 >= 10 && __obf_4f23b6c0c8a29b07 <= 26:
		__obf_af5b88a292691c3c = __obf_fccacc671db5d5f3(__obf_d4c11dce5ccd0394)
	case __obf_4f23b6c0c8a29b07 >= 27 && __obf_4f23b6c0c8a29b07 <= 40:
		__obf_af5b88a292691c3c = __obf_fccacc671db5d5f3(__obf_e21e9a063b84f8a4)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_4f23b6c0c8a29b07)
	}

	var __obf_ba2373eb7d967534 *bitset.Bitset
	__obf_ba2373eb7d967534, __obf_3f84ccce09b63a68 := __obf_af5b88a292691c3c.__obf_5fa94477cdc64b8f([]byte(__obf_42ae8f69ca3d94e4))

	if __obf_3f84ccce09b63a68 != nil {
		return nil, __obf_3f84ccce09b63a68
	}
	__obf_fb47b7063cd80a24 := __obf_d2e0f70b27ca0e9a(__obf_28ff6edf33e2b2ce, __obf_4f23b6c0c8a29b07)

	if __obf_fb47b7063cd80a24 == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_ba2373eb7d967534.Len() > __obf_fb47b7063cd80a24.__obf_17d316c470f598b7() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)", __obf_4f23b6c0c8a29b07, __obf_ba2373eb7d967534.
			Len(), __obf_fb47b7063cd80a24.__obf_17d316c470f598b7())
	}
	__obf_9c281bbbde209bea := &QRCode{
		Content: __obf_42ae8f69ca3d94e4,

		Level:         __obf_28ff6edf33e2b2ce,
		VersionNumber: __obf_fb47b7063cd80a24.__obf_4f23b6c0c8a29b07,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_af5b88a292691c3c: __obf_af5b88a292691c3c, __obf_e4df198c6616985f: __obf_ba2373eb7d967534, __obf_4f23b6c0c8a29b07: *__obf_fb47b7063cd80a24,
	}

	return __obf_9c281bbbde209bea, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_9c281bbbde209bea *QRCode) Bitmap() [][]bool {
	__obf_9c281bbbde209bea.
		// Build QR code.
		__obf_5fa94477cdc64b8f()

	return __obf_9c281bbbde209bea.

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
		__obf_14306cbbf0004950.__obf_d7ceef4be5069a61()
}

func (__obf_9c281bbbde209bea *QRCode) Image(__obf_311d29b1744e83e0 int) image.Image {
	__obf_9c281bbbde209bea.
		// Build QR code.
		__obf_5fa94477cdc64b8f()
	__obf_15a556eb60c6115e := // Minimum pixels (both width and height) required.
		__obf_9c281bbbde209bea.

			// Variable size support.
			__obf_14306cbbf0004950.__obf_311d29b1744e83e0

	if __obf_311d29b1744e83e0 < 0 {
		__obf_311d29b1744e83e0 = __obf_311d29b1744e83e0 * -1 * __obf_15a556eb60c6115e
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_311d29b1744e83e0 < __obf_15a556eb60c6115e {
		__obf_311d29b1744e83e0 = __obf_15a556eb60c6115e
	}
	__obf_0ea6ba11e00dc609 := // Output image.
		image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_311d29b1744e83e0,

			// Saves a few bytes to have them in this order
			__obf_311d29b1744e83e0}}
	__obf_8af79729b34fbe30 := color.Palette([]color.Color{__obf_9c281bbbde209bea.BackgroundColor, __obf_9c281bbbde209bea.ForegroundColor})
	__obf_1f15c3284581f61d := image.NewPaletted(__obf_0ea6ba11e00dc609, __obf_8af79729b34fbe30)
	__obf_3d78340bf032213a := uint8(__obf_1f15c3284581f61d.Palette.Index(__obf_9c281bbbde209bea.ForegroundColor))
	__obf_d7ceef4be5069a61 := // QR code bitmap.
		__obf_9c281bbbde209bea.

			// Map each image pixel to the nearest QR code module.
			__obf_14306cbbf0004950.__obf_d7ceef4be5069a61()
	__obf_ade5f2c1e365f7b8 := float64(__obf_15a556eb60c6115e) / float64(__obf_311d29b1744e83e0)
	for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < __obf_311d29b1744e83e0; __obf_8650f6fe673c98a1++ {
		__obf_7dd81ce83bf63772 := int(float64(__obf_8650f6fe673c98a1) * __obf_ade5f2c1e365f7b8)
		for __obf_cf967a75716d8446 := 0; __obf_cf967a75716d8446 < __obf_311d29b1744e83e0; __obf_cf967a75716d8446++ {
			__obf_ff6736d3df8961f5 := int(float64(__obf_cf967a75716d8446) * __obf_ade5f2c1e365f7b8)
			__obf_f85dcc6c0aea860b := __obf_d7ceef4be5069a61[__obf_7dd81ce83bf63772][__obf_ff6736d3df8961f5]

			if __obf_f85dcc6c0aea860b {
				__obf_33ab3a0f563da5d7 := __obf_1f15c3284581f61d.PixOffset(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1)
				__obf_1f15c3284581f61d.
					Pix[__obf_33ab3a0f563da5d7] = __obf_3d78340bf032213a
			}
		}
	}

	return __obf_1f15c3284581f61d
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_9c281bbbde209bea *QRCode) PNG(__obf_311d29b1744e83e0 int) ([]byte, error) {
	__obf_1f15c3284581f61d := __obf_9c281bbbde209bea.Image(__obf_311d29b1744e83e0)
	__obf_af5b88a292691c3c := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_21b947e29586a23f bytes.Buffer
	__obf_3f84ccce09b63a68 := __obf_af5b88a292691c3c.Encode(&__obf_21b947e29586a23f, __obf_1f15c3284581f61d)

	if __obf_3f84ccce09b63a68 != nil {
		return nil, __obf_3f84ccce09b63a68
	}

	return __obf_21b947e29586a23f.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_9c281bbbde209bea *QRCode) Write(__obf_311d29b1744e83e0 int, __obf_0bdb8847c4477e94 io.Writer) error {
	var png []byte

	png, __obf_3f84ccce09b63a68 := __obf_9c281bbbde209bea.PNG(__obf_311d29b1744e83e0)

	if __obf_3f84ccce09b63a68 != nil {
		return __obf_3f84ccce09b63a68
	}
	_, __obf_3f84ccce09b63a68 = __obf_0bdb8847c4477e94.Write(png)
	return __obf_3f84ccce09b63a68
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_9c281bbbde209bea *QRCode) WriteFile(__obf_311d29b1744e83e0 int, __obf_ffac0de5194f7875 string) error {
	var png []byte

	png, __obf_3f84ccce09b63a68 := __obf_9c281bbbde209bea.PNG(__obf_311d29b1744e83e0)

	if __obf_3f84ccce09b63a68 != nil {
		return __obf_3f84ccce09b63a68
	}

	return os.WriteFile(__obf_ffac0de5194f7875, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_9c281bbbde209bea *QRCode) __obf_5fa94477cdc64b8f() {
	__obf_e1aceb419eb8b230 := __obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07.__obf_7ea55d54ea75607a(__obf_9c281bbbde209bea.__obf_e4df198c6616985f.Len())
	__obf_9c281bbbde209bea.__obf_4f04a66c29623c8b(__obf_e1aceb419eb8b230)
	__obf_9c281bbbde209bea.__obf_9e8f828ad9c392e5()
	__obf_ba2373eb7d967534 := __obf_9c281bbbde209bea.__obf_9bccde5759463161()

	const __obf_82a4f6a200c355ef int = 8
	__obf_d5cbfb4fd6fc4841 := 0

	for __obf_0bb74c1e8aa906af := 0; __obf_0bb74c1e8aa906af < __obf_82a4f6a200c355ef; __obf_0bb74c1e8aa906af++ {
		var __obf_d73c540e0bdff21b *__obf_14306cbbf0004950
		var __obf_3f84ccce09b63a68 error
		__obf_d73c540e0bdff21b, __obf_3f84ccce09b63a68 = __obf_38396e56554d9d40(__obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07, __obf_0bb74c1e8aa906af, __obf_ba2373eb7d967534, !__obf_9c281bbbde209bea.Borderless)

		if __obf_3f84ccce09b63a68 != nil {
			log.Panic(__obf_3f84ccce09b63a68.Error())
		}
		__obf_2a27c02f01243dcc := __obf_d73c540e0bdff21b.__obf_2a27c02f01243dcc()
		if __obf_2a27c02f01243dcc != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)", __obf_2a27c02f01243dcc, __obf_9c281bbbde209bea.
				VersionNumber)
		}
		__obf_8af79729b34fbe30 := __obf_d73c540e0bdff21b.

			//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())
			__obf_0a829429859f1a12()

		if __obf_9c281bbbde209bea.__obf_14306cbbf0004950 == nil || __obf_8af79729b34fbe30 < __obf_d5cbfb4fd6fc4841 {
			__obf_9c281bbbde209bea.__obf_14306cbbf0004950 = __obf_d73c540e0bdff21b
			__obf_9c281bbbde209bea.__obf_0bb74c1e8aa906af = __obf_0bb74c1e8aa906af
			__obf_d5cbfb4fd6fc4841 = __obf_8af79729b34fbe30
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_9c281bbbde209bea *QRCode) __obf_4f04a66c29623c8b(__obf_e1aceb419eb8b230 int) {
	__obf_9c281bbbde209bea.__obf_e4df198c6616985f.
		AppendNumBools(__obf_e1aceb419eb8b230, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_9c281bbbde209bea *QRCode) __obf_9bccde5759463161() *bitset.Bitset {
	// Split into blocks.
	type __obf_af3ec74d9b6a805d struct {
		__obf_e4df198c6616985f *bitset.Bitset
		__obf_33bc61986d34301f int
	}
	__obf_f51213722f02a78f := make([]__obf_af3ec74d9b6a805d, __obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07.__obf_86dcf07e23f37f95())
	__obf_b1d9b56dd18d11cd := 0
	__obf_ced3f69a482bda05 := 0
	__obf_9944eaf319105499 := 0

	for _, __obf_21b947e29586a23f := range __obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07.__obf_f51213722f02a78f {
		for __obf_02acc1a50ace36fe := 0; __obf_02acc1a50ace36fe < __obf_21b947e29586a23f.__obf_86dcf07e23f37f95; __obf_02acc1a50ace36fe++ {
			__obf_b1d9b56dd18d11cd = __obf_ced3f69a482bda05
			__obf_ced3f69a482bda05 = __obf_b1d9b56dd18d11cd + __obf_21b947e29586a23f.__obf_b3e075c2526905ef*8
			__obf_6dece183843f8a54 := // Apply error correction to each block.
				__obf_21b947e29586a23f.__obf_e27ceaf2974832d2 - __obf_21b947e29586a23f.__obf_b3e075c2526905ef
			__obf_f51213722f02a78f[__obf_9944eaf319105499].__obf_e4df198c6616985f = reedsolomon.Encode(__obf_9c281bbbde209bea.__obf_e4df198c6616985f.Substr(__obf_b1d9b56dd18d11cd, __obf_ced3f69a482bda05), __obf_6dece183843f8a54)
			__obf_f51213722f02a78f[__obf_9944eaf319105499].__obf_33bc61986d34301f = __obf_ced3f69a482bda05 - __obf_b1d9b56dd18d11cd
			__obf_9944eaf319105499++
		}
	}
	__obf_a425e7f63b42864b := // Interleave the blocks.

		bitset.New()
	__obf_be901b0ce3b1f6cf := // Combine data blocks.
		true
	for __obf_8715c27aa2bf2586 := 0; __obf_be901b0ce3b1f6cf; __obf_8715c27aa2bf2586 += 8 {
		__obf_be901b0ce3b1f6cf = false

		for __obf_02acc1a50ace36fe, __obf_21b947e29586a23f := range __obf_f51213722f02a78f {
			if __obf_8715c27aa2bf2586 >= __obf_f51213722f02a78f[__obf_02acc1a50ace36fe].__obf_33bc61986d34301f {
				continue
			}
			__obf_a425e7f63b42864b.
				Append(__obf_21b947e29586a23f.__obf_e4df198c6616985f.Substr(__obf_8715c27aa2bf2586, __obf_8715c27aa2bf2586+8))
			__obf_be901b0ce3b1f6cf = true
		}
	}
	__obf_be901b0ce3b1f6cf = // Combine error correction blocks.
		true
	for __obf_8715c27aa2bf2586 := 0; __obf_be901b0ce3b1f6cf; __obf_8715c27aa2bf2586 += 8 {
		__obf_be901b0ce3b1f6cf = false

		for __obf_02acc1a50ace36fe, __obf_21b947e29586a23f := range __obf_f51213722f02a78f {
			__obf_7fa183b6bdaac3f6 := __obf_8715c27aa2bf2586 + __obf_f51213722f02a78f[__obf_02acc1a50ace36fe].__obf_33bc61986d34301f
			if __obf_7fa183b6bdaac3f6 >= __obf_f51213722f02a78f[__obf_02acc1a50ace36fe].__obf_e4df198c6616985f.Len() {
				continue
			}
			__obf_a425e7f63b42864b.
				Append(__obf_21b947e29586a23f.__obf_e4df198c6616985f.Substr(__obf_7fa183b6bdaac3f6, __obf_7fa183b6bdaac3f6+8))
			__obf_be901b0ce3b1f6cf = true
		}
	}
	__obf_a425e7f63b42864b.

		// Append remainder bits.
		AppendNumBools(__obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07.__obf_1818e415f9ac87f1, false)

	return __obf_a425e7f63b42864b
}

// addPadding pads the encoded data upto the full length required.
func (__obf_9c281bbbde209bea *QRCode) __obf_9e8f828ad9c392e5() {
	__obf_17d316c470f598b7 := __obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07.__obf_17d316c470f598b7()

	if __obf_9c281bbbde209bea.__obf_e4df198c6616985f.Len() == __obf_17d316c470f598b7 {
		return
	}
	__obf_9c281bbbde209bea.

		// Pad to the nearest codeword boundary.
		__obf_e4df198c6616985f.
		AppendNumBools(__obf_9c281bbbde209bea.__obf_4f23b6c0c8a29b07.__obf_4ec92ab3f8e0bcb4(__obf_9c281bbbde209bea.__obf_e4df198c6616985f.Len()), false)
	__obf_0f5b22a68e2e6497 := // Pad codewords 0b11101100 and 0b00010001.
		[2]*bitset.Bitset{
			bitset.New(true, true, true, false, true, true, false, false),
			bitset.New(false, false, false, true, false, false, false, true),
		}
	__obf_8715c27aa2bf2586 := // Insert pad codewords alternately.
		0
	for __obf_17d316c470f598b7-__obf_9c281bbbde209bea.__obf_e4df198c6616985f.Len() >= 8 {
		__obf_9c281bbbde209bea.__obf_e4df198c6616985f.
			Append(__obf_0f5b22a68e2e6497[__obf_8715c27aa2bf2586])
		__obf_8715c27aa2bf2586 = 1 - __obf_8715c27aa2bf2586 // Alternate between 0 and 1.
	}

	if __obf_9c281bbbde209bea.__obf_e4df198c6616985f.Len() != __obf_17d316c470f598b7 {
		log.Panicf("BUG: got len %d, expected %d", __obf_9c281bbbde209bea.__obf_e4df198c6616985f.Len(), __obf_17d316c470f598b7)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_9c281bbbde209bea *QRCode) ToString(__obf_ff4d721f04307bde bool) string {
	__obf_6a9faca8b54d2a16 := __obf_9c281bbbde209bea.Bitmap()
	var __obf_e08b0a12f63f2f60 bytes.Buffer
	for __obf_8650f6fe673c98a1 := range __obf_6a9faca8b54d2a16 {
		for __obf_cf967a75716d8446 := range __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1] {
			if __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1][__obf_cf967a75716d8446] != __obf_ff4d721f04307bde {
				__obf_e08b0a12f63f2f60.
					WriteString("  ")
			} else {
				__obf_e08b0a12f63f2f60.
					WriteString("██")
			}
		}
		__obf_e08b0a12f63f2f60.
			WriteString("\n")
	}
	return __obf_e08b0a12f63f2f60.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_9c281bbbde209bea *QRCode) ToSmallString(__obf_ff4d721f04307bde bool) string {
	__obf_6a9faca8b54d2a16 := __obf_9c281bbbde209bea.Bitmap()
	var __obf_e08b0a12f63f2f60 bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < len(__obf_6a9faca8b54d2a16)-1; __obf_8650f6fe673c98a1 += 2 {
		for __obf_cf967a75716d8446 := range __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1] {
			if __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1][__obf_cf967a75716d8446] == __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1+1][__obf_cf967a75716d8446] {
				if __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1][__obf_cf967a75716d8446] != __obf_ff4d721f04307bde {
					__obf_e08b0a12f63f2f60.
						WriteString(" ")
				} else {
					__obf_e08b0a12f63f2f60.
						WriteString("█")
				}
			} else {
				if __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1][__obf_cf967a75716d8446] != __obf_ff4d721f04307bde {
					__obf_e08b0a12f63f2f60.
						WriteString("▄")
				} else {
					__obf_e08b0a12f63f2f60.
						WriteString("▀")
				}
			}
		}
		__obf_e08b0a12f63f2f60.
			WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_6a9faca8b54d2a16)%2 == 1 {
		__obf_8650f6fe673c98a1 := len(__obf_6a9faca8b54d2a16) - 1
		for __obf_cf967a75716d8446 := range __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1] {
			if __obf_6a9faca8b54d2a16[__obf_8650f6fe673c98a1][__obf_cf967a75716d8446] != __obf_ff4d721f04307bde {
				__obf_e08b0a12f63f2f60.
					WriteString(" ")
			} else {
				__obf_e08b0a12f63f2f60.
					WriteString("▀")
			}
		}
		__obf_e08b0a12f63f2f60.
			WriteString("\n")
	}
	return __obf_e08b0a12f63f2f60.String()
}
