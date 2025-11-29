package __obf_07675b7eb1c7284a

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
func Encode(__obf_64f354ace3e47f6e string, __obf_d70a5da97264c717 RecoveryLevel, __obf_0abbe8d67dfa1a14 int) ([]byte, error) {
	var __obf_02dd10b1f05b344f *QRCode
	__obf_02dd10b1f05b344f, __obf_c88308f5c3b1b66a := New(__obf_64f354ace3e47f6e, __obf_d70a5da97264c717)

	if __obf_c88308f5c3b1b66a != nil {
		return nil, __obf_c88308f5c3b1b66a
	}

	return __obf_02dd10b1f05b344f.PNG(__obf_0abbe8d67dfa1a14)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_64f354ace3e47f6e string, __obf_d70a5da97264c717 RecoveryLevel, __obf_0abbe8d67dfa1a14 int, __obf_cd7bbcd124cf8d97 string) error {
	var __obf_02dd10b1f05b344f *QRCode
	__obf_02dd10b1f05b344f, __obf_c88308f5c3b1b66a := New(__obf_64f354ace3e47f6e, __obf_d70a5da97264c717)

	if __obf_c88308f5c3b1b66a != nil {
		return __obf_c88308f5c3b1b66a
	}

	return __obf_02dd10b1f05b344f.WriteFile(__obf_0abbe8d67dfa1a14,

		// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
		// With WriteColorFile you can also specify the colors you want to use.
		//
		// size is both the image width and height in pixels. If size is too small then
		// a larger image is silently written. Negative values for size cause a variable
		// sized image to be written: See the documentation for Image().
		__obf_cd7bbcd124cf8d97)
}

func WriteColorFile(__obf_64f354ace3e47f6e string, __obf_d70a5da97264c717 RecoveryLevel, __obf_0abbe8d67dfa1a14 int, __obf_51504a44f3c2fb20, __obf_9dc1479d789173b7 color.Color, __obf_cd7bbcd124cf8d97 string) error {

	var __obf_02dd10b1f05b344f *QRCode
	__obf_02dd10b1f05b344f, __obf_c88308f5c3b1b66a := New(__obf_64f354ace3e47f6e, __obf_d70a5da97264c717)
	__obf_02dd10b1f05b344f.
		BackgroundColor = __obf_51504a44f3c2fb20
	__obf_02dd10b1f05b344f.
		ForegroundColor = __obf_9dc1479d789173b7

	if __obf_c88308f5c3b1b66a != nil {
		return __obf_c88308f5c3b1b66a
	}

	return __obf_02dd10b1f05b344f.WriteFile(__obf_0abbe8d67dfa1a14,

		// A QRCode represents a valid encoded QRCode.
		__obf_cd7bbcd124cf8d97)
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
	__obf_14a8535206b233c8 *__obf_a1f99035c479a58e
	__obf_499888902ab9ce20 __obf_07b921eb65a6057a
	__obf_b08d86cf0c048279 *bitset.Bitset
	__obf_e4275ce80c964a2d *__obf_e4275ce80c964a2d
	__obf_09052c1dd0468dff int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_64f354ace3e47f6e string, __obf_d70a5da97264c717 RecoveryLevel) (*QRCode, error) {
	__obf_306a3612c7672c0a := []__obf_d9ac8a0ee0d0f27f{__obf_96e9dfef6bbb06ba, __obf_bd631aba8d4005cf, __obf_5da2f2a34edb06bc}

	var __obf_14a8535206b233c8 *__obf_a1f99035c479a58e
	var __obf_9b1dcfbd17fe9618 *bitset.Bitset
	var __obf_1d8a883769dad5ad *__obf_07b921eb65a6057a
	var __obf_c88308f5c3b1b66a error

	for _, __obf_066728c6553a619e := range __obf_306a3612c7672c0a {
		__obf_14a8535206b233c8 = __obf_11835c8cedb230b3(__obf_066728c6553a619e)
		__obf_9b1dcfbd17fe9618, __obf_c88308f5c3b1b66a = __obf_14a8535206b233c8.__obf_7626f31ec1adfcf5([]byte(__obf_64f354ace3e47f6e))

		if __obf_c88308f5c3b1b66a != nil {
			continue
		}
		__obf_1d8a883769dad5ad = __obf_2dd42a191f3f8abb(__obf_d70a5da97264c717, __obf_14a8535206b233c8, __obf_9b1dcfbd17fe9618.Len())

		if __obf_1d8a883769dad5ad != nil {
			break
		}
	}

	if __obf_c88308f5c3b1b66a != nil {
		return nil, __obf_c88308f5c3b1b66a
	} else if __obf_1d8a883769dad5ad == nil {
		return nil, errors.New("content too long to encode")
	}
	__obf_02dd10b1f05b344f := &QRCode{
		Content: __obf_64f354ace3e47f6e,

		Level:         __obf_d70a5da97264c717,
		VersionNumber: __obf_1d8a883769dad5ad.__obf_499888902ab9ce20,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_14a8535206b233c8: __obf_14a8535206b233c8, __obf_b08d86cf0c048279: __obf_9b1dcfbd17fe9618, __obf_499888902ab9ce20: *__obf_1d8a883769dad5ad,
	}

	return __obf_02dd10b1f05b344f, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_64f354ace3e47f6e string, __obf_499888902ab9ce20 int, __obf_d70a5da97264c717 RecoveryLevel) (*QRCode, error) {
	var __obf_14a8535206b233c8 *__obf_a1f99035c479a58e

	switch {
	case __obf_499888902ab9ce20 >= 1 && __obf_499888902ab9ce20 <= 9:
		__obf_14a8535206b233c8 = __obf_11835c8cedb230b3(__obf_96e9dfef6bbb06ba)
	case __obf_499888902ab9ce20 >= 10 && __obf_499888902ab9ce20 <= 26:
		__obf_14a8535206b233c8 = __obf_11835c8cedb230b3(__obf_bd631aba8d4005cf)
	case __obf_499888902ab9ce20 >= 27 && __obf_499888902ab9ce20 <= 40:
		__obf_14a8535206b233c8 = __obf_11835c8cedb230b3(__obf_5da2f2a34edb06bc)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_499888902ab9ce20)
	}

	var __obf_9b1dcfbd17fe9618 *bitset.Bitset
	__obf_9b1dcfbd17fe9618, __obf_c88308f5c3b1b66a := __obf_14a8535206b233c8.__obf_7626f31ec1adfcf5([]byte(__obf_64f354ace3e47f6e))

	if __obf_c88308f5c3b1b66a != nil {
		return nil, __obf_c88308f5c3b1b66a
	}
	__obf_1d8a883769dad5ad := __obf_d933820d3995e993(__obf_d70a5da97264c717, __obf_499888902ab9ce20)

	if __obf_1d8a883769dad5ad == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_9b1dcfbd17fe9618.Len() > __obf_1d8a883769dad5ad.__obf_7e2c981a689bf0cc() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)", __obf_499888902ab9ce20, __obf_9b1dcfbd17fe9618.
			Len(), __obf_1d8a883769dad5ad.__obf_7e2c981a689bf0cc())
	}
	__obf_02dd10b1f05b344f := &QRCode{
		Content: __obf_64f354ace3e47f6e,

		Level:         __obf_d70a5da97264c717,
		VersionNumber: __obf_1d8a883769dad5ad.__obf_499888902ab9ce20,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_14a8535206b233c8: __obf_14a8535206b233c8, __obf_b08d86cf0c048279: __obf_9b1dcfbd17fe9618, __obf_499888902ab9ce20: *__obf_1d8a883769dad5ad,
	}

	return __obf_02dd10b1f05b344f, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_02dd10b1f05b344f *QRCode) Bitmap() [][]bool {
	__obf_02dd10b1f05b344f.
		// Build QR code.
		__obf_7626f31ec1adfcf5()

	return __obf_02dd10b1f05b344f.

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
		__obf_e4275ce80c964a2d.__obf_e6043f786f7c3cb5()
}

func (__obf_02dd10b1f05b344f *QRCode) Image(__obf_0abbe8d67dfa1a14 int) image.Image {
	__obf_02dd10b1f05b344f.
		// Build QR code.
		__obf_7626f31ec1adfcf5()
	__obf_b1b454c7a178865b := // Minimum pixels (both width and height) required.
		__obf_02dd10b1f05b344f.

			// Variable size support.
			__obf_e4275ce80c964a2d.__obf_0abbe8d67dfa1a14

	if __obf_0abbe8d67dfa1a14 < 0 {
		__obf_0abbe8d67dfa1a14 = __obf_0abbe8d67dfa1a14 * -1 * __obf_b1b454c7a178865b
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_0abbe8d67dfa1a14 < __obf_b1b454c7a178865b {
		__obf_0abbe8d67dfa1a14 = __obf_b1b454c7a178865b
	}
	__obf_db384bb3f8cdd85f := // Output image.
		image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_0abbe8d67dfa1a14,

			// Saves a few bytes to have them in this order
			__obf_0abbe8d67dfa1a14}}
	__obf_a787eb999ce91781 := color.Palette([]color.Color{__obf_02dd10b1f05b344f.BackgroundColor, __obf_02dd10b1f05b344f.ForegroundColor})
	__obf_1cc946816c06cc8e := image.NewPaletted(__obf_db384bb3f8cdd85f, __obf_a787eb999ce91781)
	__obf_7614e55f725dfa9e := uint8(__obf_1cc946816c06cc8e.Palette.Index(__obf_02dd10b1f05b344f.ForegroundColor))
	__obf_e6043f786f7c3cb5 := // QR code bitmap.
		__obf_02dd10b1f05b344f.

			// Map each image pixel to the nearest QR code module.
			__obf_e4275ce80c964a2d.__obf_e6043f786f7c3cb5()
	__obf_73daaa941b4f506f := float64(__obf_b1b454c7a178865b) / float64(__obf_0abbe8d67dfa1a14)
	for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < __obf_0abbe8d67dfa1a14; __obf_21e5ea661cba6209++ {
		__obf_bdc1af469fce1977 := int(float64(__obf_21e5ea661cba6209) * __obf_73daaa941b4f506f)
		for __obf_7107b684b2ee3954 := 0; __obf_7107b684b2ee3954 < __obf_0abbe8d67dfa1a14; __obf_7107b684b2ee3954++ {
			__obf_afe002de887f8da6 := int(float64(__obf_7107b684b2ee3954) * __obf_73daaa941b4f506f)
			__obf_283aa709fd2621d8 := __obf_e6043f786f7c3cb5[__obf_bdc1af469fce1977][__obf_afe002de887f8da6]

			if __obf_283aa709fd2621d8 {
				__obf_a9820374fc324dd5 := __obf_1cc946816c06cc8e.PixOffset(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209)
				__obf_1cc946816c06cc8e.
					Pix[__obf_a9820374fc324dd5] = __obf_7614e55f725dfa9e
			}
		}
	}

	return __obf_1cc946816c06cc8e
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_02dd10b1f05b344f *QRCode) PNG(__obf_0abbe8d67dfa1a14 int) ([]byte, error) {
	__obf_1cc946816c06cc8e := __obf_02dd10b1f05b344f.Image(__obf_0abbe8d67dfa1a14)
	__obf_14a8535206b233c8 := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_32f895d189b42b64 bytes.Buffer
	__obf_c88308f5c3b1b66a := __obf_14a8535206b233c8.Encode(&__obf_32f895d189b42b64, __obf_1cc946816c06cc8e)

	if __obf_c88308f5c3b1b66a != nil {
		return nil, __obf_c88308f5c3b1b66a
	}

	return __obf_32f895d189b42b64.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_02dd10b1f05b344f *QRCode) Write(__obf_0abbe8d67dfa1a14 int, __obf_ea16583d54fc8e00 io.Writer) error {
	var png []byte

	png, __obf_c88308f5c3b1b66a := __obf_02dd10b1f05b344f.PNG(__obf_0abbe8d67dfa1a14)

	if __obf_c88308f5c3b1b66a != nil {
		return __obf_c88308f5c3b1b66a
	}
	_, __obf_c88308f5c3b1b66a = __obf_ea16583d54fc8e00.Write(png)
	return __obf_c88308f5c3b1b66a
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_02dd10b1f05b344f *QRCode) WriteFile(__obf_0abbe8d67dfa1a14 int, __obf_cd7bbcd124cf8d97 string) error {
	var png []byte

	png, __obf_c88308f5c3b1b66a := __obf_02dd10b1f05b344f.PNG(__obf_0abbe8d67dfa1a14)

	if __obf_c88308f5c3b1b66a != nil {
		return __obf_c88308f5c3b1b66a
	}

	return os.WriteFile(__obf_cd7bbcd124cf8d97, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_02dd10b1f05b344f *QRCode) __obf_7626f31ec1adfcf5() {
	__obf_6eb9abeddc219f06 := __obf_02dd10b1f05b344f.__obf_499888902ab9ce20.__obf_0498600f3f5752f7(__obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Len())
	__obf_02dd10b1f05b344f.__obf_e514d12c24576bd7(__obf_6eb9abeddc219f06)
	__obf_02dd10b1f05b344f.__obf_ecae963c8624bd60()
	__obf_9b1dcfbd17fe9618 := __obf_02dd10b1f05b344f.__obf_b2d525ebf7e163a7()

	const __obf_1bb8f7f61bf12b39 int = 8
	__obf_a633b45f4e29f865 := 0

	for __obf_09052c1dd0468dff := 0; __obf_09052c1dd0468dff < __obf_1bb8f7f61bf12b39; __obf_09052c1dd0468dff++ {
		var __obf_e988dab6de39cff8 *__obf_e4275ce80c964a2d
		var __obf_c88308f5c3b1b66a error
		__obf_e988dab6de39cff8, __obf_c88308f5c3b1b66a = __obf_4c41d2af47304d6b(__obf_02dd10b1f05b344f.__obf_499888902ab9ce20, __obf_09052c1dd0468dff, __obf_9b1dcfbd17fe9618, !__obf_02dd10b1f05b344f.Borderless)

		if __obf_c88308f5c3b1b66a != nil {
			log.Panic(__obf_c88308f5c3b1b66a.Error())
		}
		__obf_8933e3589fae1c15 := __obf_e988dab6de39cff8.__obf_8933e3589fae1c15()
		if __obf_8933e3589fae1c15 != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)", __obf_8933e3589fae1c15, __obf_02dd10b1f05b344f.
				VersionNumber)
		}
		__obf_a787eb999ce91781 := __obf_e988dab6de39cff8.

			//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())
			__obf_fcbea8a1aa2a09d3()

		if __obf_02dd10b1f05b344f.__obf_e4275ce80c964a2d == nil || __obf_a787eb999ce91781 < __obf_a633b45f4e29f865 {
			__obf_02dd10b1f05b344f.__obf_e4275ce80c964a2d = __obf_e988dab6de39cff8
			__obf_02dd10b1f05b344f.__obf_09052c1dd0468dff = __obf_09052c1dd0468dff
			__obf_a633b45f4e29f865 = __obf_a787eb999ce91781
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_02dd10b1f05b344f *QRCode) __obf_e514d12c24576bd7(__obf_6eb9abeddc219f06 int) {
	__obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.
		AppendNumBools(__obf_6eb9abeddc219f06, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_02dd10b1f05b344f *QRCode) __obf_b2d525ebf7e163a7() *bitset.Bitset {
	// Split into blocks.
	type __obf_4bb7789b8129b49b struct {
		__obf_b08d86cf0c048279 *bitset.Bitset
		__obf_36c81a99c712c086 int
	}
	__obf_58f56c8799376640 := make([]__obf_4bb7789b8129b49b, __obf_02dd10b1f05b344f.__obf_499888902ab9ce20.__obf_bad9b23e64008899())
	__obf_6a6a77e65664e51f := 0
	__obf_539abc224fbaa37f := 0
	__obf_f9c8c623edf9a2c1 := 0

	for _, __obf_32f895d189b42b64 := range __obf_02dd10b1f05b344f.__obf_499888902ab9ce20.__obf_58f56c8799376640 {
		for __obf_95b5b5fef8a19647 := 0; __obf_95b5b5fef8a19647 < __obf_32f895d189b42b64.__obf_bad9b23e64008899; __obf_95b5b5fef8a19647++ {
			__obf_6a6a77e65664e51f = __obf_539abc224fbaa37f
			__obf_539abc224fbaa37f = __obf_6a6a77e65664e51f + __obf_32f895d189b42b64.__obf_7baefc526a23613a*8
			__obf_f65aad0d139787e4 := // Apply error correction to each block.
				__obf_32f895d189b42b64.__obf_5c1b4d20dab7a903 - __obf_32f895d189b42b64.__obf_7baefc526a23613a
			__obf_58f56c8799376640[__obf_f9c8c623edf9a2c1].__obf_b08d86cf0c048279 = reedsolomon.Encode(__obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Substr(__obf_6a6a77e65664e51f, __obf_539abc224fbaa37f), __obf_f65aad0d139787e4)
			__obf_58f56c8799376640[__obf_f9c8c623edf9a2c1].__obf_36c81a99c712c086 = __obf_539abc224fbaa37f - __obf_6a6a77e65664e51f
			__obf_f9c8c623edf9a2c1++
		}
	}
	__obf_34f15b8f5f4ba96d := // Interleave the blocks.

		bitset.New()
	__obf_bb319c58a129cbd7 := // Combine data blocks.
		true
	for __obf_2e4170383a5ddec7 := 0; __obf_bb319c58a129cbd7; __obf_2e4170383a5ddec7 += 8 {
		__obf_bb319c58a129cbd7 = false

		for __obf_95b5b5fef8a19647, __obf_32f895d189b42b64 := range __obf_58f56c8799376640 {
			if __obf_2e4170383a5ddec7 >= __obf_58f56c8799376640[__obf_95b5b5fef8a19647].__obf_36c81a99c712c086 {
				continue
			}
			__obf_34f15b8f5f4ba96d.
				Append(__obf_32f895d189b42b64.__obf_b08d86cf0c048279.Substr(__obf_2e4170383a5ddec7, __obf_2e4170383a5ddec7+8))
			__obf_bb319c58a129cbd7 = true
		}
	}
	__obf_bb319c58a129cbd7 = // Combine error correction blocks.
		true
	for __obf_2e4170383a5ddec7 := 0; __obf_bb319c58a129cbd7; __obf_2e4170383a5ddec7 += 8 {
		__obf_bb319c58a129cbd7 = false

		for __obf_95b5b5fef8a19647, __obf_32f895d189b42b64 := range __obf_58f56c8799376640 {
			__obf_8f9689595b4940c6 := __obf_2e4170383a5ddec7 + __obf_58f56c8799376640[__obf_95b5b5fef8a19647].__obf_36c81a99c712c086
			if __obf_8f9689595b4940c6 >= __obf_58f56c8799376640[__obf_95b5b5fef8a19647].__obf_b08d86cf0c048279.Len() {
				continue
			}
			__obf_34f15b8f5f4ba96d.
				Append(__obf_32f895d189b42b64.__obf_b08d86cf0c048279.Substr(__obf_8f9689595b4940c6, __obf_8f9689595b4940c6+8))
			__obf_bb319c58a129cbd7 = true
		}
	}
	__obf_34f15b8f5f4ba96d.

		// Append remainder bits.
		AppendNumBools(__obf_02dd10b1f05b344f.__obf_499888902ab9ce20.__obf_a3022c36220e944d, false)

	return __obf_34f15b8f5f4ba96d
}

// addPadding pads the encoded data upto the full length required.
func (__obf_02dd10b1f05b344f *QRCode) __obf_ecae963c8624bd60() {
	__obf_7e2c981a689bf0cc := __obf_02dd10b1f05b344f.__obf_499888902ab9ce20.__obf_7e2c981a689bf0cc()

	if __obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Len() == __obf_7e2c981a689bf0cc {
		return
	}
	__obf_02dd10b1f05b344f.

		// Pad to the nearest codeword boundary.
		__obf_b08d86cf0c048279.
		AppendNumBools(__obf_02dd10b1f05b344f.__obf_499888902ab9ce20.__obf_5e4a04f55cd12858(__obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Len()), false)
	__obf_897968ed72c73a38 := // Pad codewords 0b11101100 and 0b00010001.
		[2]*bitset.Bitset{
			bitset.New(true, true, true, false, true, true, false, false),
			bitset.New(false, false, false, true, false, false, false, true),
		}
	__obf_2e4170383a5ddec7 := // Insert pad codewords alternately.
		0
	for __obf_7e2c981a689bf0cc-__obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Len() >= 8 {
		__obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.
			Append(__obf_897968ed72c73a38[__obf_2e4170383a5ddec7])
		__obf_2e4170383a5ddec7 = 1 - __obf_2e4170383a5ddec7 // Alternate between 0 and 1.
	}

	if __obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Len() != __obf_7e2c981a689bf0cc {
		log.Panicf("BUG: got len %d, expected %d", __obf_02dd10b1f05b344f.__obf_b08d86cf0c048279.Len(), __obf_7e2c981a689bf0cc)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_02dd10b1f05b344f *QRCode) ToString(__obf_f3fec1dc09e4a3e0 bool) string {
	__obf_516b7bae007fd0cd := __obf_02dd10b1f05b344f.Bitmap()
	var __obf_a26e371daad8bd99 bytes.Buffer
	for __obf_21e5ea661cba6209 := range __obf_516b7bae007fd0cd {
		for __obf_7107b684b2ee3954 := range __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209] {
			if __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209][__obf_7107b684b2ee3954] != __obf_f3fec1dc09e4a3e0 {
				__obf_a26e371daad8bd99.
					WriteString("  ")
			} else {
				__obf_a26e371daad8bd99.
					WriteString("██")
			}
		}
		__obf_a26e371daad8bd99.
			WriteString("\n")
	}
	return __obf_a26e371daad8bd99.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_02dd10b1f05b344f *QRCode) ToSmallString(__obf_f3fec1dc09e4a3e0 bool) string {
	__obf_516b7bae007fd0cd := __obf_02dd10b1f05b344f.Bitmap()
	var __obf_a26e371daad8bd99 bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < len(__obf_516b7bae007fd0cd)-1; __obf_21e5ea661cba6209 += 2 {
		for __obf_7107b684b2ee3954 := range __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209] {
			if __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209][__obf_7107b684b2ee3954] == __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209+1][__obf_7107b684b2ee3954] {
				if __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209][__obf_7107b684b2ee3954] != __obf_f3fec1dc09e4a3e0 {
					__obf_a26e371daad8bd99.
						WriteString(" ")
				} else {
					__obf_a26e371daad8bd99.
						WriteString("█")
				}
			} else {
				if __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209][__obf_7107b684b2ee3954] != __obf_f3fec1dc09e4a3e0 {
					__obf_a26e371daad8bd99.
						WriteString("▄")
				} else {
					__obf_a26e371daad8bd99.
						WriteString("▀")
				}
			}
		}
		__obf_a26e371daad8bd99.
			WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_516b7bae007fd0cd)%2 == 1 {
		__obf_21e5ea661cba6209 := len(__obf_516b7bae007fd0cd) - 1
		for __obf_7107b684b2ee3954 := range __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209] {
			if __obf_516b7bae007fd0cd[__obf_21e5ea661cba6209][__obf_7107b684b2ee3954] != __obf_f3fec1dc09e4a3e0 {
				__obf_a26e371daad8bd99.
					WriteString(" ")
			} else {
				__obf_a26e371daad8bd99.
					WriteString("▀")
			}
		}
		__obf_a26e371daad8bd99.
			WriteString("\n")
	}
	return __obf_a26e371daad8bd99.String()
}
