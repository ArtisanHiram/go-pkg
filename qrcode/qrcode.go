package __obf_8d73621f5856b288

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
func Encode(__obf_6a4173e663037fd7 string, __obf_877aa2e5148b2067 RecoveryLevel, __obf_8ab5ee435c7948e3 int) ([]byte, error) {
	var __obf_70d9a0635566ae8a *QRCode
	__obf_70d9a0635566ae8a, __obf_a4be63d440a7a9e0 := New(__obf_6a4173e663037fd7, __obf_877aa2e5148b2067)

	if __obf_a4be63d440a7a9e0 != nil {
		return nil, __obf_a4be63d440a7a9e0
	}

	return __obf_70d9a0635566ae8a.PNG(__obf_8ab5ee435c7948e3)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_6a4173e663037fd7 string, __obf_877aa2e5148b2067 RecoveryLevel, __obf_8ab5ee435c7948e3 int, __obf_6d8835e55e160def string) error {
	var __obf_70d9a0635566ae8a *QRCode
	__obf_70d9a0635566ae8a, __obf_a4be63d440a7a9e0 := New(__obf_6a4173e663037fd7, __obf_877aa2e5148b2067)

	if __obf_a4be63d440a7a9e0 != nil {
		return __obf_a4be63d440a7a9e0
	}

	return __obf_70d9a0635566ae8a.WriteFile(__obf_8ab5ee435c7948e3,

		// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
		// With WriteColorFile you can also specify the colors you want to use.
		//
		// size is both the image width and height in pixels. If size is too small then
		// a larger image is silently written. Negative values for size cause a variable
		// sized image to be written: See the documentation for Image().
		__obf_6d8835e55e160def)
}

func WriteColorFile(__obf_6a4173e663037fd7 string, __obf_877aa2e5148b2067 RecoveryLevel, __obf_8ab5ee435c7948e3 int, __obf_32ffa4fdade6529d, __obf_8101a45e1d12176a color.Color, __obf_6d8835e55e160def string) error {

	var __obf_70d9a0635566ae8a *QRCode
	__obf_70d9a0635566ae8a, __obf_a4be63d440a7a9e0 := New(__obf_6a4173e663037fd7, __obf_877aa2e5148b2067)
	__obf_70d9a0635566ae8a.
		BackgroundColor = __obf_32ffa4fdade6529d
	__obf_70d9a0635566ae8a.
		ForegroundColor = __obf_8101a45e1d12176a

	if __obf_a4be63d440a7a9e0 != nil {
		return __obf_a4be63d440a7a9e0
	}

	return __obf_70d9a0635566ae8a.WriteFile(__obf_8ab5ee435c7948e3,

		// A QRCode represents a valid encoded QRCode.
		__obf_6d8835e55e160def)
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
	__obf_a0f443901f570b73 *__obf_db79b5bc73d5a614
	__obf_4139f181d582c6ec __obf_c64ed2c3cfb19b80
	__obf_7a0637e203e4ea7a *bitset.Bitset
	__obf_df7d9041d687f9b6 *__obf_df7d9041d687f9b6
	__obf_ed10df3327d493cc int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_6a4173e663037fd7 string, __obf_877aa2e5148b2067 RecoveryLevel) (*QRCode, error) {
	__obf_8f87f0a2ad59e157 := []__obf_a9d4704514b6286c{__obf_80807e9626ee48fa, __obf_9bb2ab43c09df8ad, __obf_7f04905a311519fa}

	var __obf_a0f443901f570b73 *__obf_db79b5bc73d5a614
	var __obf_3efd75da4d4d8559 *bitset.Bitset
	var __obf_64f4c3b68809189b *__obf_c64ed2c3cfb19b80
	var __obf_a4be63d440a7a9e0 error

	for _, __obf_33479cec7dd55141 := range __obf_8f87f0a2ad59e157 {
		__obf_a0f443901f570b73 = __obf_2a2ab161febefc8e(__obf_33479cec7dd55141)
		__obf_3efd75da4d4d8559, __obf_a4be63d440a7a9e0 = __obf_a0f443901f570b73.__obf_35d7b1c8ddee5d52([]byte(__obf_6a4173e663037fd7))

		if __obf_a4be63d440a7a9e0 != nil {
			continue
		}
		__obf_64f4c3b68809189b = __obf_6fffebea060112f9(__obf_877aa2e5148b2067, __obf_a0f443901f570b73, __obf_3efd75da4d4d8559.Len())

		if __obf_64f4c3b68809189b != nil {
			break
		}
	}

	if __obf_a4be63d440a7a9e0 != nil {
		return nil, __obf_a4be63d440a7a9e0
	} else if __obf_64f4c3b68809189b == nil {
		return nil, errors.New("content too long to encode")
	}
	__obf_70d9a0635566ae8a := &QRCode{
		Content: __obf_6a4173e663037fd7,

		Level:         __obf_877aa2e5148b2067,
		VersionNumber: __obf_64f4c3b68809189b.__obf_4139f181d582c6ec,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_a0f443901f570b73: __obf_a0f443901f570b73, __obf_7a0637e203e4ea7a: __obf_3efd75da4d4d8559, __obf_4139f181d582c6ec: *__obf_64f4c3b68809189b,
	}

	return __obf_70d9a0635566ae8a, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_6a4173e663037fd7 string, __obf_4139f181d582c6ec int, __obf_877aa2e5148b2067 RecoveryLevel) (*QRCode, error) {
	var __obf_a0f443901f570b73 *__obf_db79b5bc73d5a614

	switch {
	case __obf_4139f181d582c6ec >= 1 && __obf_4139f181d582c6ec <= 9:
		__obf_a0f443901f570b73 = __obf_2a2ab161febefc8e(__obf_80807e9626ee48fa)
	case __obf_4139f181d582c6ec >= 10 && __obf_4139f181d582c6ec <= 26:
		__obf_a0f443901f570b73 = __obf_2a2ab161febefc8e(__obf_9bb2ab43c09df8ad)
	case __obf_4139f181d582c6ec >= 27 && __obf_4139f181d582c6ec <= 40:
		__obf_a0f443901f570b73 = __obf_2a2ab161febefc8e(__obf_7f04905a311519fa)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_4139f181d582c6ec)
	}

	var __obf_3efd75da4d4d8559 *bitset.Bitset
	__obf_3efd75da4d4d8559, __obf_a4be63d440a7a9e0 := __obf_a0f443901f570b73.__obf_35d7b1c8ddee5d52([]byte(__obf_6a4173e663037fd7))

	if __obf_a4be63d440a7a9e0 != nil {
		return nil, __obf_a4be63d440a7a9e0
	}
	__obf_64f4c3b68809189b := __obf_4af69a6c1a68877b(__obf_877aa2e5148b2067, __obf_4139f181d582c6ec)

	if __obf_64f4c3b68809189b == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_3efd75da4d4d8559.Len() > __obf_64f4c3b68809189b.__obf_fbab2e67c946f055() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)", __obf_4139f181d582c6ec, __obf_3efd75da4d4d8559.
			Len(), __obf_64f4c3b68809189b.__obf_fbab2e67c946f055())
	}
	__obf_70d9a0635566ae8a := &QRCode{
		Content: __obf_6a4173e663037fd7,

		Level:         __obf_877aa2e5148b2067,
		VersionNumber: __obf_64f4c3b68809189b.__obf_4139f181d582c6ec,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_a0f443901f570b73: __obf_a0f443901f570b73, __obf_7a0637e203e4ea7a: __obf_3efd75da4d4d8559, __obf_4139f181d582c6ec: *__obf_64f4c3b68809189b,
	}

	return __obf_70d9a0635566ae8a, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_70d9a0635566ae8a *QRCode) Bitmap() [][]bool {
	__obf_70d9a0635566ae8a.
		// Build QR code.
		__obf_35d7b1c8ddee5d52()

	return __obf_70d9a0635566ae8a.

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
		__obf_df7d9041d687f9b6.__obf_9c007cb7461851ed()
}

func (__obf_70d9a0635566ae8a *QRCode) Image(__obf_8ab5ee435c7948e3 int) image.Image {
	__obf_70d9a0635566ae8a.
		// Build QR code.
		__obf_35d7b1c8ddee5d52()
	__obf_bbb83bd9c85778d6 := // Minimum pixels (both width and height) required.
		__obf_70d9a0635566ae8a.

			// Variable size support.
			__obf_df7d9041d687f9b6.__obf_8ab5ee435c7948e3

	if __obf_8ab5ee435c7948e3 < 0 {
		__obf_8ab5ee435c7948e3 = __obf_8ab5ee435c7948e3 * -1 * __obf_bbb83bd9c85778d6
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_8ab5ee435c7948e3 < __obf_bbb83bd9c85778d6 {
		__obf_8ab5ee435c7948e3 = __obf_bbb83bd9c85778d6
	}
	__obf_b5f1760c9fb9c46a := // Output image.
		image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_8ab5ee435c7948e3,

			// Saves a few bytes to have them in this order
			__obf_8ab5ee435c7948e3}}
	__obf_d0434534fe9dc910 := color.Palette([]color.Color{__obf_70d9a0635566ae8a.BackgroundColor, __obf_70d9a0635566ae8a.ForegroundColor})
	__obf_fb03fc244c726c3d := image.NewPaletted(__obf_b5f1760c9fb9c46a, __obf_d0434534fe9dc910)
	__obf_ce0cc98e9fa1c611 := uint8(__obf_fb03fc244c726c3d.Palette.Index(__obf_70d9a0635566ae8a.ForegroundColor))
	__obf_9c007cb7461851ed := // QR code bitmap.
		__obf_70d9a0635566ae8a.

			// Map each image pixel to the nearest QR code module.
			__obf_df7d9041d687f9b6.__obf_9c007cb7461851ed()
	__obf_9d96f87f2c83d79f := float64(__obf_bbb83bd9c85778d6) / float64(__obf_8ab5ee435c7948e3)
	for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < __obf_8ab5ee435c7948e3; __obf_d70c25f3394ee124++ {
		__obf_d9ccd0b1cf1b8d3b := int(float64(__obf_d70c25f3394ee124) * __obf_9d96f87f2c83d79f)
		for __obf_2db4671bf7cd078f := 0; __obf_2db4671bf7cd078f < __obf_8ab5ee435c7948e3; __obf_2db4671bf7cd078f++ {
			__obf_7445f9ba793b3a9c := int(float64(__obf_2db4671bf7cd078f) * __obf_9d96f87f2c83d79f)
			__obf_1b6e726e11cd9194 := __obf_9c007cb7461851ed[__obf_d9ccd0b1cf1b8d3b][__obf_7445f9ba793b3a9c]

			if __obf_1b6e726e11cd9194 {
				__obf_e8c7f965a80ac9ad := __obf_fb03fc244c726c3d.PixOffset(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124)
				__obf_fb03fc244c726c3d.
					Pix[__obf_e8c7f965a80ac9ad] = __obf_ce0cc98e9fa1c611
			}
		}
	}

	return __obf_fb03fc244c726c3d
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_70d9a0635566ae8a *QRCode) PNG(__obf_8ab5ee435c7948e3 int) ([]byte, error) {
	__obf_fb03fc244c726c3d := __obf_70d9a0635566ae8a.Image(__obf_8ab5ee435c7948e3)
	__obf_a0f443901f570b73 := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_22b7af99cfb59bff bytes.Buffer
	__obf_a4be63d440a7a9e0 := __obf_a0f443901f570b73.Encode(&__obf_22b7af99cfb59bff, __obf_fb03fc244c726c3d)

	if __obf_a4be63d440a7a9e0 != nil {
		return nil, __obf_a4be63d440a7a9e0
	}

	return __obf_22b7af99cfb59bff.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_70d9a0635566ae8a *QRCode) Write(__obf_8ab5ee435c7948e3 int, __obf_11d41c6754731a53 io.Writer) error {
	var png []byte

	png, __obf_a4be63d440a7a9e0 := __obf_70d9a0635566ae8a.PNG(__obf_8ab5ee435c7948e3)

	if __obf_a4be63d440a7a9e0 != nil {
		return __obf_a4be63d440a7a9e0
	}
	_, __obf_a4be63d440a7a9e0 = __obf_11d41c6754731a53.Write(png)
	return __obf_a4be63d440a7a9e0
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_70d9a0635566ae8a *QRCode) WriteFile(__obf_8ab5ee435c7948e3 int, __obf_6d8835e55e160def string) error {
	var png []byte

	png, __obf_a4be63d440a7a9e0 := __obf_70d9a0635566ae8a.PNG(__obf_8ab5ee435c7948e3)

	if __obf_a4be63d440a7a9e0 != nil {
		return __obf_a4be63d440a7a9e0
	}

	return os.WriteFile(__obf_6d8835e55e160def, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_70d9a0635566ae8a *QRCode) __obf_35d7b1c8ddee5d52() {
	__obf_90c19e0e3bc06fff := __obf_70d9a0635566ae8a.__obf_4139f181d582c6ec.__obf_f5b9e51105940fc3(__obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Len())
	__obf_70d9a0635566ae8a.__obf_411ebdb49ce3804f(__obf_90c19e0e3bc06fff)
	__obf_70d9a0635566ae8a.__obf_dd083f10df92af8b()
	__obf_3efd75da4d4d8559 := __obf_70d9a0635566ae8a.__obf_53b4adaf44d106a3()

	const __obf_d1a2a4bb57a362a8 int = 8
	__obf_cca1d7447027f3ba := 0

	for __obf_ed10df3327d493cc := 0; __obf_ed10df3327d493cc < __obf_d1a2a4bb57a362a8; __obf_ed10df3327d493cc++ {
		var __obf_d4fa85a87405a625 *__obf_df7d9041d687f9b6
		var __obf_a4be63d440a7a9e0 error
		__obf_d4fa85a87405a625, __obf_a4be63d440a7a9e0 = __obf_03446a0753219873(__obf_70d9a0635566ae8a.__obf_4139f181d582c6ec, __obf_ed10df3327d493cc, __obf_3efd75da4d4d8559, !__obf_70d9a0635566ae8a.Borderless)

		if __obf_a4be63d440a7a9e0 != nil {
			log.Panic(__obf_a4be63d440a7a9e0.Error())
		}
		__obf_a2c9171a634d8aa4 := __obf_d4fa85a87405a625.__obf_a2c9171a634d8aa4()
		if __obf_a2c9171a634d8aa4 != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)", __obf_a2c9171a634d8aa4, __obf_70d9a0635566ae8a.
				VersionNumber)
		}
		__obf_d0434534fe9dc910 := __obf_d4fa85a87405a625.

			//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())
			__obf_5ff464158ac6e38e()

		if __obf_70d9a0635566ae8a.__obf_df7d9041d687f9b6 == nil || __obf_d0434534fe9dc910 < __obf_cca1d7447027f3ba {
			__obf_70d9a0635566ae8a.__obf_df7d9041d687f9b6 = __obf_d4fa85a87405a625
			__obf_70d9a0635566ae8a.__obf_ed10df3327d493cc = __obf_ed10df3327d493cc
			__obf_cca1d7447027f3ba = __obf_d0434534fe9dc910
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_70d9a0635566ae8a *QRCode) __obf_411ebdb49ce3804f(__obf_90c19e0e3bc06fff int) {
	__obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.
		AppendNumBools(__obf_90c19e0e3bc06fff, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_70d9a0635566ae8a *QRCode) __obf_53b4adaf44d106a3() *bitset.Bitset {
	// Split into blocks.
	type __obf_76357b567394877c struct {
		__obf_7a0637e203e4ea7a *bitset.Bitset
		__obf_86c0337cfb0f16d6 int
	}
	__obf_d8d8b9663688ea90 := make([]__obf_76357b567394877c, __obf_70d9a0635566ae8a.__obf_4139f181d582c6ec.__obf_cb93cbffbdb0156c())
	__obf_bb49f3bbe0778fb0 := 0
	__obf_8bb960799586cabd := 0
	__obf_a4b1424dbe52c852 := 0

	for _, __obf_22b7af99cfb59bff := range __obf_70d9a0635566ae8a.__obf_4139f181d582c6ec.__obf_d8d8b9663688ea90 {
		for __obf_7c0eb8656083f52b := 0; __obf_7c0eb8656083f52b < __obf_22b7af99cfb59bff.__obf_cb93cbffbdb0156c; __obf_7c0eb8656083f52b++ {
			__obf_bb49f3bbe0778fb0 = __obf_8bb960799586cabd
			__obf_8bb960799586cabd = __obf_bb49f3bbe0778fb0 + __obf_22b7af99cfb59bff.__obf_5a1d504b8645501c*8
			__obf_8224d715fd98c28a := // Apply error correction to each block.
				__obf_22b7af99cfb59bff.__obf_431c275946223cee - __obf_22b7af99cfb59bff.__obf_5a1d504b8645501c
			__obf_d8d8b9663688ea90[__obf_a4b1424dbe52c852].__obf_7a0637e203e4ea7a = reedsolomon.Encode(__obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Substr(__obf_bb49f3bbe0778fb0, __obf_8bb960799586cabd), __obf_8224d715fd98c28a)
			__obf_d8d8b9663688ea90[__obf_a4b1424dbe52c852].__obf_86c0337cfb0f16d6 = __obf_8bb960799586cabd - __obf_bb49f3bbe0778fb0
			__obf_a4b1424dbe52c852++
		}
	}
	__obf_e81697a14e2164ea := // Interleave the blocks.

		bitset.New()
	__obf_15a764f443669e7a := // Combine data blocks.
		true
	for __obf_015b1aa3060900e0 := 0; __obf_15a764f443669e7a; __obf_015b1aa3060900e0 += 8 {
		__obf_15a764f443669e7a = false

		for __obf_7c0eb8656083f52b, __obf_22b7af99cfb59bff := range __obf_d8d8b9663688ea90 {
			if __obf_015b1aa3060900e0 >= __obf_d8d8b9663688ea90[__obf_7c0eb8656083f52b].__obf_86c0337cfb0f16d6 {
				continue
			}
			__obf_e81697a14e2164ea.
				Append(__obf_22b7af99cfb59bff.__obf_7a0637e203e4ea7a.Substr(__obf_015b1aa3060900e0, __obf_015b1aa3060900e0+8))
			__obf_15a764f443669e7a = true
		}
	}
	__obf_15a764f443669e7a = // Combine error correction blocks.
		true
	for __obf_015b1aa3060900e0 := 0; __obf_15a764f443669e7a; __obf_015b1aa3060900e0 += 8 {
		__obf_15a764f443669e7a = false

		for __obf_7c0eb8656083f52b, __obf_22b7af99cfb59bff := range __obf_d8d8b9663688ea90 {
			__obf_5590cf26b2fe9eec := __obf_015b1aa3060900e0 + __obf_d8d8b9663688ea90[__obf_7c0eb8656083f52b].__obf_86c0337cfb0f16d6
			if __obf_5590cf26b2fe9eec >= __obf_d8d8b9663688ea90[__obf_7c0eb8656083f52b].__obf_7a0637e203e4ea7a.Len() {
				continue
			}
			__obf_e81697a14e2164ea.
				Append(__obf_22b7af99cfb59bff.__obf_7a0637e203e4ea7a.Substr(__obf_5590cf26b2fe9eec, __obf_5590cf26b2fe9eec+8))
			__obf_15a764f443669e7a = true
		}
	}
	__obf_e81697a14e2164ea.

		// Append remainder bits.
		AppendNumBools(__obf_70d9a0635566ae8a.__obf_4139f181d582c6ec.__obf_297bf7a24bf7fde6, false)

	return __obf_e81697a14e2164ea
}

// addPadding pads the encoded data upto the full length required.
func (__obf_70d9a0635566ae8a *QRCode) __obf_dd083f10df92af8b() {
	__obf_fbab2e67c946f055 := __obf_70d9a0635566ae8a.__obf_4139f181d582c6ec.__obf_fbab2e67c946f055()

	if __obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Len() == __obf_fbab2e67c946f055 {
		return
	}
	__obf_70d9a0635566ae8a.

		// Pad to the nearest codeword boundary.
		__obf_7a0637e203e4ea7a.
		AppendNumBools(__obf_70d9a0635566ae8a.__obf_4139f181d582c6ec.__obf_d2760344a3799976(__obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Len()), false)
	__obf_4d8f8256edcf0b16 := // Pad codewords 0b11101100 and 0b00010001.
		[2]*bitset.Bitset{
			bitset.New(true, true, true, false, true, true, false, false),
			bitset.New(false, false, false, true, false, false, false, true),
		}
	__obf_015b1aa3060900e0 := // Insert pad codewords alternately.
		0
	for __obf_fbab2e67c946f055-__obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Len() >= 8 {
		__obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.
			Append(__obf_4d8f8256edcf0b16[__obf_015b1aa3060900e0])
		__obf_015b1aa3060900e0 = 1 - __obf_015b1aa3060900e0 // Alternate between 0 and 1.
	}

	if __obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Len() != __obf_fbab2e67c946f055 {
		log.Panicf("BUG: got len %d, expected %d", __obf_70d9a0635566ae8a.__obf_7a0637e203e4ea7a.Len(), __obf_fbab2e67c946f055)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_70d9a0635566ae8a *QRCode) ToString(__obf_4095cb445c6d354d bool) string {
	__obf_14b5cf89e78cdf62 := __obf_70d9a0635566ae8a.Bitmap()
	var __obf_de5e715db568da7e bytes.Buffer
	for __obf_d70c25f3394ee124 := range __obf_14b5cf89e78cdf62 {
		for __obf_2db4671bf7cd078f := range __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124] {
			if __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124][__obf_2db4671bf7cd078f] != __obf_4095cb445c6d354d {
				__obf_de5e715db568da7e.
					WriteString("  ")
			} else {
				__obf_de5e715db568da7e.
					WriteString("██")
			}
		}
		__obf_de5e715db568da7e.
			WriteString("\n")
	}
	return __obf_de5e715db568da7e.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_70d9a0635566ae8a *QRCode) ToSmallString(__obf_4095cb445c6d354d bool) string {
	__obf_14b5cf89e78cdf62 := __obf_70d9a0635566ae8a.Bitmap()
	var __obf_de5e715db568da7e bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < len(__obf_14b5cf89e78cdf62)-1; __obf_d70c25f3394ee124 += 2 {
		for __obf_2db4671bf7cd078f := range __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124] {
			if __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124][__obf_2db4671bf7cd078f] == __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124+1][__obf_2db4671bf7cd078f] {
				if __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124][__obf_2db4671bf7cd078f] != __obf_4095cb445c6d354d {
					__obf_de5e715db568da7e.
						WriteString(" ")
				} else {
					__obf_de5e715db568da7e.
						WriteString("█")
				}
			} else {
				if __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124][__obf_2db4671bf7cd078f] != __obf_4095cb445c6d354d {
					__obf_de5e715db568da7e.
						WriteString("▄")
				} else {
					__obf_de5e715db568da7e.
						WriteString("▀")
				}
			}
		}
		__obf_de5e715db568da7e.
			WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_14b5cf89e78cdf62)%2 == 1 {
		__obf_d70c25f3394ee124 := len(__obf_14b5cf89e78cdf62) - 1
		for __obf_2db4671bf7cd078f := range __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124] {
			if __obf_14b5cf89e78cdf62[__obf_d70c25f3394ee124][__obf_2db4671bf7cd078f] != __obf_4095cb445c6d354d {
				__obf_de5e715db568da7e.
					WriteString(" ")
			} else {
				__obf_de5e715db568da7e.
					WriteString("▀")
			}
		}
		__obf_de5e715db568da7e.
			WriteString("\n")
	}
	return __obf_de5e715db568da7e.String()
}
