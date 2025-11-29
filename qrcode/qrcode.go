package __obf_540b74a9d13704fa

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
func Encode(__obf_286889314b0bc51c string, __obf_a221f051d10eb503 RecoveryLevel, __obf_5315f29e47197a10 int) ([]byte, error) {
	var __obf_5bc919920862a3b8 *QRCode
	__obf_5bc919920862a3b8, __obf_800077dfc326e0b7 := New(__obf_286889314b0bc51c, __obf_a221f051d10eb503)

	if __obf_800077dfc326e0b7 != nil {
		return nil, __obf_800077dfc326e0b7
	}

	return __obf_5bc919920862a3b8.PNG(__obf_5315f29e47197a10)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_286889314b0bc51c string, __obf_a221f051d10eb503 RecoveryLevel, __obf_5315f29e47197a10 int, __obf_b1306d2c9a355f51 string) error {
	var __obf_5bc919920862a3b8 *QRCode
	__obf_5bc919920862a3b8, __obf_800077dfc326e0b7 := New(__obf_286889314b0bc51c, __obf_a221f051d10eb503)

	if __obf_800077dfc326e0b7 != nil {
		return __obf_800077dfc326e0b7
	}

	return __obf_5bc919920862a3b8.WriteFile(__obf_5315f29e47197a10,

		// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
		// With WriteColorFile you can also specify the colors you want to use.
		//
		// size is both the image width and height in pixels. If size is too small then
		// a larger image is silently written. Negative values for size cause a variable
		// sized image to be written: See the documentation for Image().
		__obf_b1306d2c9a355f51)
}

func WriteColorFile(__obf_286889314b0bc51c string, __obf_a221f051d10eb503 RecoveryLevel, __obf_5315f29e47197a10 int, __obf_c419bb317a60f238, __obf_ef05128670b0e783 color.Color, __obf_b1306d2c9a355f51 string) error {

	var __obf_5bc919920862a3b8 *QRCode
	__obf_5bc919920862a3b8, __obf_800077dfc326e0b7 := New(__obf_286889314b0bc51c, __obf_a221f051d10eb503)
	__obf_5bc919920862a3b8.
		BackgroundColor = __obf_c419bb317a60f238
	__obf_5bc919920862a3b8.
		ForegroundColor = __obf_ef05128670b0e783

	if __obf_800077dfc326e0b7 != nil {
		return __obf_800077dfc326e0b7
	}

	return __obf_5bc919920862a3b8.WriteFile(__obf_5315f29e47197a10,

		// A QRCode represents a valid encoded QRCode.
		__obf_b1306d2c9a355f51)
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
	__obf_2646c2f27fac28d6 *__obf_5aac2ee6cac292a2
	__obf_7bf1f2404ecb0c0a __obf_52ef8fe7f69f3df7
	__obf_3bcdb4f6c983dc7c *bitset.Bitset
	__obf_ebb7a995f536d449 *__obf_ebb7a995f536d449
	__obf_a1713e992775f27b int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_286889314b0bc51c string, __obf_a221f051d10eb503 RecoveryLevel) (*QRCode, error) {
	__obf_7afb1a3c69134082 := []__obf_ede9586c20e29cd6{__obf_36e336d7eab6b5c8, __obf_ce4133b003340f88, __obf_d9027dc4f9acd3b0}

	var __obf_2646c2f27fac28d6 *__obf_5aac2ee6cac292a2
	var __obf_cb404255cbbc8689 *bitset.Bitset
	var __obf_2d8435b35e31964c *__obf_52ef8fe7f69f3df7
	var __obf_800077dfc326e0b7 error

	for _, __obf_56c65516aa2ae2b8 := range __obf_7afb1a3c69134082 {
		__obf_2646c2f27fac28d6 = __obf_46cf60d6f5b35f0a(__obf_56c65516aa2ae2b8)
		__obf_cb404255cbbc8689, __obf_800077dfc326e0b7 = __obf_2646c2f27fac28d6.__obf_c7821a6383b7b718([]byte(__obf_286889314b0bc51c))

		if __obf_800077dfc326e0b7 != nil {
			continue
		}
		__obf_2d8435b35e31964c = __obf_846df71d822a28c1(__obf_a221f051d10eb503, __obf_2646c2f27fac28d6, __obf_cb404255cbbc8689.Len())

		if __obf_2d8435b35e31964c != nil {
			break
		}
	}

	if __obf_800077dfc326e0b7 != nil {
		return nil, __obf_800077dfc326e0b7
	} else if __obf_2d8435b35e31964c == nil {
		return nil, errors.New("content too long to encode")
	}
	__obf_5bc919920862a3b8 := &QRCode{
		Content: __obf_286889314b0bc51c,

		Level:         __obf_a221f051d10eb503,
		VersionNumber: __obf_2d8435b35e31964c.__obf_7bf1f2404ecb0c0a,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_2646c2f27fac28d6: __obf_2646c2f27fac28d6, __obf_3bcdb4f6c983dc7c: __obf_cb404255cbbc8689, __obf_7bf1f2404ecb0c0a: *__obf_2d8435b35e31964c,
	}

	return __obf_5bc919920862a3b8, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_286889314b0bc51c string, __obf_7bf1f2404ecb0c0a int, __obf_a221f051d10eb503 RecoveryLevel) (*QRCode, error) {
	var __obf_2646c2f27fac28d6 *__obf_5aac2ee6cac292a2

	switch {
	case __obf_7bf1f2404ecb0c0a >= 1 && __obf_7bf1f2404ecb0c0a <= 9:
		__obf_2646c2f27fac28d6 = __obf_46cf60d6f5b35f0a(__obf_36e336d7eab6b5c8)
	case __obf_7bf1f2404ecb0c0a >= 10 && __obf_7bf1f2404ecb0c0a <= 26:
		__obf_2646c2f27fac28d6 = __obf_46cf60d6f5b35f0a(__obf_ce4133b003340f88)
	case __obf_7bf1f2404ecb0c0a >= 27 && __obf_7bf1f2404ecb0c0a <= 40:
		__obf_2646c2f27fac28d6 = __obf_46cf60d6f5b35f0a(__obf_d9027dc4f9acd3b0)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_7bf1f2404ecb0c0a)
	}

	var __obf_cb404255cbbc8689 *bitset.Bitset
	__obf_cb404255cbbc8689, __obf_800077dfc326e0b7 := __obf_2646c2f27fac28d6.__obf_c7821a6383b7b718([]byte(__obf_286889314b0bc51c))

	if __obf_800077dfc326e0b7 != nil {
		return nil, __obf_800077dfc326e0b7
	}
	__obf_2d8435b35e31964c := __obf_23004ab726110f63(__obf_a221f051d10eb503, __obf_7bf1f2404ecb0c0a)

	if __obf_2d8435b35e31964c == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_cb404255cbbc8689.Len() > __obf_2d8435b35e31964c.__obf_6ca7a7ab38ace36f() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)", __obf_7bf1f2404ecb0c0a, __obf_cb404255cbbc8689.
			Len(), __obf_2d8435b35e31964c.__obf_6ca7a7ab38ace36f())
	}
	__obf_5bc919920862a3b8 := &QRCode{
		Content: __obf_286889314b0bc51c,

		Level:         __obf_a221f051d10eb503,
		VersionNumber: __obf_2d8435b35e31964c.__obf_7bf1f2404ecb0c0a,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_2646c2f27fac28d6: __obf_2646c2f27fac28d6, __obf_3bcdb4f6c983dc7c: __obf_cb404255cbbc8689, __obf_7bf1f2404ecb0c0a: *__obf_2d8435b35e31964c,
	}

	return __obf_5bc919920862a3b8, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_5bc919920862a3b8 *QRCode) Bitmap() [][]bool {
	__obf_5bc919920862a3b8.
		// Build QR code.
		__obf_c7821a6383b7b718()

	return __obf_5bc919920862a3b8.

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
		__obf_ebb7a995f536d449.__obf_0f6cf8adbc21238e()
}

func (__obf_5bc919920862a3b8 *QRCode) Image(__obf_5315f29e47197a10 int) image.Image {
	__obf_5bc919920862a3b8.
		// Build QR code.
		__obf_c7821a6383b7b718()
	__obf_c85be7c8168a0bb0 := // Minimum pixels (both width and height) required.
		__obf_5bc919920862a3b8.

			// Variable size support.
			__obf_ebb7a995f536d449.__obf_5315f29e47197a10

	if __obf_5315f29e47197a10 < 0 {
		__obf_5315f29e47197a10 = __obf_5315f29e47197a10 * -1 * __obf_c85be7c8168a0bb0
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_5315f29e47197a10 < __obf_c85be7c8168a0bb0 {
		__obf_5315f29e47197a10 = __obf_c85be7c8168a0bb0
	}
	__obf_d6c54c53ec63a878 := // Output image.
		image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_5315f29e47197a10,

			// Saves a few bytes to have them in this order
			__obf_5315f29e47197a10}}
	__obf_e341ead915610610 := color.Palette([]color.Color{__obf_5bc919920862a3b8.BackgroundColor, __obf_5bc919920862a3b8.ForegroundColor})
	__obf_94138721ee1f7bdf := image.NewPaletted(__obf_d6c54c53ec63a878, __obf_e341ead915610610)
	__obf_b6b334964db24c9d := uint8(__obf_94138721ee1f7bdf.Palette.Index(__obf_5bc919920862a3b8.ForegroundColor))
	__obf_0f6cf8adbc21238e := // QR code bitmap.
		__obf_5bc919920862a3b8.

			// Map each image pixel to the nearest QR code module.
			__obf_ebb7a995f536d449.__obf_0f6cf8adbc21238e()
	__obf_84255904ff809450 := float64(__obf_c85be7c8168a0bb0) / float64(__obf_5315f29e47197a10)
	for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < __obf_5315f29e47197a10; __obf_9637ef17b76683ca++ {
		__obf_ae902644e212a47c := int(float64(__obf_9637ef17b76683ca) * __obf_84255904ff809450)
		for __obf_a7fd5f72cb6fabea := 0; __obf_a7fd5f72cb6fabea < __obf_5315f29e47197a10; __obf_a7fd5f72cb6fabea++ {
			__obf_fac970a6155cd4e3 := int(float64(__obf_a7fd5f72cb6fabea) * __obf_84255904ff809450)
			__obf_75eba50799fccf42 := __obf_0f6cf8adbc21238e[__obf_ae902644e212a47c][__obf_fac970a6155cd4e3]

			if __obf_75eba50799fccf42 {
				__obf_871b18e0f4f602d5 := __obf_94138721ee1f7bdf.PixOffset(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca)
				__obf_94138721ee1f7bdf.
					Pix[__obf_871b18e0f4f602d5] = __obf_b6b334964db24c9d
			}
		}
	}

	return __obf_94138721ee1f7bdf
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_5bc919920862a3b8 *QRCode) PNG(__obf_5315f29e47197a10 int) ([]byte, error) {
	__obf_94138721ee1f7bdf := __obf_5bc919920862a3b8.Image(__obf_5315f29e47197a10)
	__obf_2646c2f27fac28d6 := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_aa760948f0d0122d bytes.Buffer
	__obf_800077dfc326e0b7 := __obf_2646c2f27fac28d6.Encode(&__obf_aa760948f0d0122d, __obf_94138721ee1f7bdf)

	if __obf_800077dfc326e0b7 != nil {
		return nil, __obf_800077dfc326e0b7
	}

	return __obf_aa760948f0d0122d.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_5bc919920862a3b8 *QRCode) Write(__obf_5315f29e47197a10 int, __obf_e321ae222fa0efa4 io.Writer) error {
	var png []byte

	png, __obf_800077dfc326e0b7 := __obf_5bc919920862a3b8.PNG(__obf_5315f29e47197a10)

	if __obf_800077dfc326e0b7 != nil {
		return __obf_800077dfc326e0b7
	}
	_, __obf_800077dfc326e0b7 = __obf_e321ae222fa0efa4.Write(png)
	return __obf_800077dfc326e0b7
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_5bc919920862a3b8 *QRCode) WriteFile(__obf_5315f29e47197a10 int, __obf_b1306d2c9a355f51 string) error {
	var png []byte

	png, __obf_800077dfc326e0b7 := __obf_5bc919920862a3b8.PNG(__obf_5315f29e47197a10)

	if __obf_800077dfc326e0b7 != nil {
		return __obf_800077dfc326e0b7
	}

	return os.WriteFile(__obf_b1306d2c9a355f51, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_5bc919920862a3b8 *QRCode) __obf_c7821a6383b7b718() {
	__obf_14757c29e2bdea6e := __obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a.__obf_bdf7c8fa55e0922e(__obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Len())
	__obf_5bc919920862a3b8.__obf_4c60ba94aa4b706c(__obf_14757c29e2bdea6e)
	__obf_5bc919920862a3b8.__obf_6765bb58335d8266()
	__obf_cb404255cbbc8689 := __obf_5bc919920862a3b8.__obf_e7f8caf8e7027c18()

	const __obf_55b14566a8b075dc int = 8
	__obf_17e9e44595f51024 := 0

	for __obf_a1713e992775f27b := 0; __obf_a1713e992775f27b < __obf_55b14566a8b075dc; __obf_a1713e992775f27b++ {
		var __obf_6900aee466231dfd *__obf_ebb7a995f536d449
		var __obf_800077dfc326e0b7 error
		__obf_6900aee466231dfd, __obf_800077dfc326e0b7 = __obf_c86a54d7f0dfd343(__obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a, __obf_a1713e992775f27b, __obf_cb404255cbbc8689, !__obf_5bc919920862a3b8.Borderless)

		if __obf_800077dfc326e0b7 != nil {
			log.Panic(__obf_800077dfc326e0b7.Error())
		}
		__obf_730a1c84bbe28f0b := __obf_6900aee466231dfd.__obf_730a1c84bbe28f0b()
		if __obf_730a1c84bbe28f0b != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)", __obf_730a1c84bbe28f0b, __obf_5bc919920862a3b8.
				VersionNumber)
		}
		__obf_e341ead915610610 := __obf_6900aee466231dfd.

			//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())
			__obf_60fe70efba5f9777()

		if __obf_5bc919920862a3b8.__obf_ebb7a995f536d449 == nil || __obf_e341ead915610610 < __obf_17e9e44595f51024 {
			__obf_5bc919920862a3b8.__obf_ebb7a995f536d449 = __obf_6900aee466231dfd
			__obf_5bc919920862a3b8.__obf_a1713e992775f27b = __obf_a1713e992775f27b
			__obf_17e9e44595f51024 = __obf_e341ead915610610
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_5bc919920862a3b8 *QRCode) __obf_4c60ba94aa4b706c(__obf_14757c29e2bdea6e int) {
	__obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.
		AppendNumBools(__obf_14757c29e2bdea6e, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_5bc919920862a3b8 *QRCode) __obf_e7f8caf8e7027c18() *bitset.Bitset {
	// Split into blocks.
	type __obf_f0bd13559f35ec53 struct {
		__obf_3bcdb4f6c983dc7c *bitset.Bitset
		__obf_c6c9232845eb2bd4 int
	}
	__obf_ef7d737cbd163114 := make([]__obf_f0bd13559f35ec53, __obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a.__obf_bbfd4dbabec7ff0b())
	__obf_711a399c8c9ecede := 0
	__obf_7d758e370d49db11 := 0
	__obf_a9f36a4f0eb8df2e := 0

	for _, __obf_aa760948f0d0122d := range __obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a.__obf_ef7d737cbd163114 {
		for __obf_cf60660918b25aa8 := 0; __obf_cf60660918b25aa8 < __obf_aa760948f0d0122d.__obf_bbfd4dbabec7ff0b; __obf_cf60660918b25aa8++ {
			__obf_711a399c8c9ecede = __obf_7d758e370d49db11
			__obf_7d758e370d49db11 = __obf_711a399c8c9ecede + __obf_aa760948f0d0122d.__obf_ebdbba63d7655685*8
			__obf_404eed09b75cdabd := // Apply error correction to each block.
				__obf_aa760948f0d0122d.__obf_b9de796f50df339e - __obf_aa760948f0d0122d.__obf_ebdbba63d7655685
			__obf_ef7d737cbd163114[__obf_a9f36a4f0eb8df2e].__obf_3bcdb4f6c983dc7c = reedsolomon.Encode(__obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Substr(__obf_711a399c8c9ecede, __obf_7d758e370d49db11), __obf_404eed09b75cdabd)
			__obf_ef7d737cbd163114[__obf_a9f36a4f0eb8df2e].__obf_c6c9232845eb2bd4 = __obf_7d758e370d49db11 - __obf_711a399c8c9ecede
			__obf_a9f36a4f0eb8df2e++
		}
	}
	__obf_3ceb8341caf311f5 := // Interleave the blocks.

		bitset.New()
	__obf_48a8d80c46ac5755 := // Combine data blocks.
		true
	for __obf_2e2482997826bcc5 := 0; __obf_48a8d80c46ac5755; __obf_2e2482997826bcc5 += 8 {
		__obf_48a8d80c46ac5755 = false

		for __obf_cf60660918b25aa8, __obf_aa760948f0d0122d := range __obf_ef7d737cbd163114 {
			if __obf_2e2482997826bcc5 >= __obf_ef7d737cbd163114[__obf_cf60660918b25aa8].__obf_c6c9232845eb2bd4 {
				continue
			}
			__obf_3ceb8341caf311f5.
				Append(__obf_aa760948f0d0122d.__obf_3bcdb4f6c983dc7c.Substr(__obf_2e2482997826bcc5, __obf_2e2482997826bcc5+8))
			__obf_48a8d80c46ac5755 = true
		}
	}
	__obf_48a8d80c46ac5755 = // Combine error correction blocks.
		true
	for __obf_2e2482997826bcc5 := 0; __obf_48a8d80c46ac5755; __obf_2e2482997826bcc5 += 8 {
		__obf_48a8d80c46ac5755 = false

		for __obf_cf60660918b25aa8, __obf_aa760948f0d0122d := range __obf_ef7d737cbd163114 {
			__obf_0c5c3413cb95f3c7 := __obf_2e2482997826bcc5 + __obf_ef7d737cbd163114[__obf_cf60660918b25aa8].__obf_c6c9232845eb2bd4
			if __obf_0c5c3413cb95f3c7 >= __obf_ef7d737cbd163114[__obf_cf60660918b25aa8].__obf_3bcdb4f6c983dc7c.Len() {
				continue
			}
			__obf_3ceb8341caf311f5.
				Append(__obf_aa760948f0d0122d.__obf_3bcdb4f6c983dc7c.Substr(__obf_0c5c3413cb95f3c7, __obf_0c5c3413cb95f3c7+8))
			__obf_48a8d80c46ac5755 = true
		}
	}
	__obf_3ceb8341caf311f5.

		// Append remainder bits.
		AppendNumBools(__obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a.__obf_58489f07bc720472, false)

	return __obf_3ceb8341caf311f5
}

// addPadding pads the encoded data upto the full length required.
func (__obf_5bc919920862a3b8 *QRCode) __obf_6765bb58335d8266() {
	__obf_6ca7a7ab38ace36f := __obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a.__obf_6ca7a7ab38ace36f()

	if __obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Len() == __obf_6ca7a7ab38ace36f {
		return
	}
	__obf_5bc919920862a3b8.

		// Pad to the nearest codeword boundary.
		__obf_3bcdb4f6c983dc7c.
		AppendNumBools(__obf_5bc919920862a3b8.__obf_7bf1f2404ecb0c0a.__obf_187f84545b0e27f4(__obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Len()), false)
	__obf_a9b68917158d68db := // Pad codewords 0b11101100 and 0b00010001.
		[2]*bitset.Bitset{
			bitset.New(true, true, true, false, true, true, false, false),
			bitset.New(false, false, false, true, false, false, false, true),
		}
	__obf_2e2482997826bcc5 := // Insert pad codewords alternately.
		0
	for __obf_6ca7a7ab38ace36f-__obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Len() >= 8 {
		__obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.
			Append(__obf_a9b68917158d68db[__obf_2e2482997826bcc5])
		__obf_2e2482997826bcc5 = 1 - __obf_2e2482997826bcc5 // Alternate between 0 and 1.
	}

	if __obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Len() != __obf_6ca7a7ab38ace36f {
		log.Panicf("BUG: got len %d, expected %d", __obf_5bc919920862a3b8.__obf_3bcdb4f6c983dc7c.Len(), __obf_6ca7a7ab38ace36f)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_5bc919920862a3b8 *QRCode) ToString(__obf_ba2cc914a0566cff bool) string {
	__obf_99abec224bac8344 := __obf_5bc919920862a3b8.Bitmap()
	var __obf_394a04a82694b061 bytes.Buffer
	for __obf_9637ef17b76683ca := range __obf_99abec224bac8344 {
		for __obf_a7fd5f72cb6fabea := range __obf_99abec224bac8344[__obf_9637ef17b76683ca] {
			if __obf_99abec224bac8344[__obf_9637ef17b76683ca][__obf_a7fd5f72cb6fabea] != __obf_ba2cc914a0566cff {
				__obf_394a04a82694b061.
					WriteString("  ")
			} else {
				__obf_394a04a82694b061.
					WriteString("██")
			}
		}
		__obf_394a04a82694b061.
			WriteString("\n")
	}
	return __obf_394a04a82694b061.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_5bc919920862a3b8 *QRCode) ToSmallString(__obf_ba2cc914a0566cff bool) string {
	__obf_99abec224bac8344 := __obf_5bc919920862a3b8.Bitmap()
	var __obf_394a04a82694b061 bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < len(__obf_99abec224bac8344)-1; __obf_9637ef17b76683ca += 2 {
		for __obf_a7fd5f72cb6fabea := range __obf_99abec224bac8344[__obf_9637ef17b76683ca] {
			if __obf_99abec224bac8344[__obf_9637ef17b76683ca][__obf_a7fd5f72cb6fabea] == __obf_99abec224bac8344[__obf_9637ef17b76683ca+1][__obf_a7fd5f72cb6fabea] {
				if __obf_99abec224bac8344[__obf_9637ef17b76683ca][__obf_a7fd5f72cb6fabea] != __obf_ba2cc914a0566cff {
					__obf_394a04a82694b061.
						WriteString(" ")
				} else {
					__obf_394a04a82694b061.
						WriteString("█")
				}
			} else {
				if __obf_99abec224bac8344[__obf_9637ef17b76683ca][__obf_a7fd5f72cb6fabea] != __obf_ba2cc914a0566cff {
					__obf_394a04a82694b061.
						WriteString("▄")
				} else {
					__obf_394a04a82694b061.
						WriteString("▀")
				}
			}
		}
		__obf_394a04a82694b061.
			WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_99abec224bac8344)%2 == 1 {
		__obf_9637ef17b76683ca := len(__obf_99abec224bac8344) - 1
		for __obf_a7fd5f72cb6fabea := range __obf_99abec224bac8344[__obf_9637ef17b76683ca] {
			if __obf_99abec224bac8344[__obf_9637ef17b76683ca][__obf_a7fd5f72cb6fabea] != __obf_ba2cc914a0566cff {
				__obf_394a04a82694b061.
					WriteString(" ")
			} else {
				__obf_394a04a82694b061.
					WriteString("▀")
			}
		}
		__obf_394a04a82694b061.
			WriteString("\n")
	}
	return __obf_394a04a82694b061.String()
}
