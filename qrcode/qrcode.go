package __obf_63dc7f87a7750084

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
func Encode(__obf_e41866e579426802 string, __obf_47158df92f26b0f4 RecoveryLevel, __obf_80f4ab6d4332abb9 int) ([]byte, error) {
	var __obf_5b0073c32b55ab92 *QRCode
	__obf_5b0073c32b55ab92, __obf_b59cbf05240604a0 := New(__obf_e41866e579426802, __obf_47158df92f26b0f4)

	if __obf_b59cbf05240604a0 != nil {
		return nil, __obf_b59cbf05240604a0
	}

	return __obf_5b0073c32b55ab92.PNG(__obf_80f4ab6d4332abb9)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_e41866e579426802 string, __obf_47158df92f26b0f4 RecoveryLevel, __obf_80f4ab6d4332abb9 int, __obf_52c631b3d40eb87e string) error {
	var __obf_5b0073c32b55ab92 *QRCode
	__obf_5b0073c32b55ab92, __obf_b59cbf05240604a0 := New(__obf_e41866e579426802, __obf_47158df92f26b0f4)

	if __obf_b59cbf05240604a0 != nil {
		return __obf_b59cbf05240604a0
	}

	return __obf_5b0073c32b55ab92.WriteFile(__obf_80f4ab6d4332abb9,

		// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
		// With WriteColorFile you can also specify the colors you want to use.
		//
		// size is both the image width and height in pixels. If size is too small then
		// a larger image is silently written. Negative values for size cause a variable
		// sized image to be written: See the documentation for Image().
		__obf_52c631b3d40eb87e)
}

func WriteColorFile(__obf_e41866e579426802 string, __obf_47158df92f26b0f4 RecoveryLevel, __obf_80f4ab6d4332abb9 int, __obf_835d47f4d838864b, __obf_c267ce7d0e11b5fd color.Color, __obf_52c631b3d40eb87e string) error {

	var __obf_5b0073c32b55ab92 *QRCode
	__obf_5b0073c32b55ab92, __obf_b59cbf05240604a0 := New(__obf_e41866e579426802, __obf_47158df92f26b0f4)
	__obf_5b0073c32b55ab92.
		BackgroundColor = __obf_835d47f4d838864b
	__obf_5b0073c32b55ab92.
		ForegroundColor = __obf_c267ce7d0e11b5fd

	if __obf_b59cbf05240604a0 != nil {
		return __obf_b59cbf05240604a0
	}

	return __obf_5b0073c32b55ab92.WriteFile(__obf_80f4ab6d4332abb9,

		// A QRCode represents a valid encoded QRCode.
		__obf_52c631b3d40eb87e)
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
	__obf_2753e87a957b69f8 *__obf_3ee8d058dd75e4b0
	__obf_2066e640e4e50f14 __obf_987afc77278a1676
	__obf_eb8b3f015415859a *bitset.Bitset
	__obf_a9c08ca051c7e6c3 *__obf_a9c08ca051c7e6c3
	__obf_3a54d62911a3bcfb int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_e41866e579426802 string, __obf_47158df92f26b0f4 RecoveryLevel) (*QRCode, error) {
	__obf_be2613891f4b1b32 := []__obf_40700b751bbe7fa9{__obf_4ef346b16a795892, __obf_e891267bd1bcc38a, __obf_054184439c33471a}

	var __obf_2753e87a957b69f8 *__obf_3ee8d058dd75e4b0
	var __obf_175ed6e946c14a9b *bitset.Bitset
	var __obf_1d6b91b9673acb44 *__obf_987afc77278a1676
	var __obf_b59cbf05240604a0 error

	for _, __obf_494cef4b0d092224 := range __obf_be2613891f4b1b32 {
		__obf_2753e87a957b69f8 = __obf_d8bb8ebef294bcff(__obf_494cef4b0d092224)
		__obf_175ed6e946c14a9b, __obf_b59cbf05240604a0 = __obf_2753e87a957b69f8.__obf_395c063b7326115c([]byte(__obf_e41866e579426802))

		if __obf_b59cbf05240604a0 != nil {
			continue
		}
		__obf_1d6b91b9673acb44 = __obf_361e6bfddb7983f8(__obf_47158df92f26b0f4, __obf_2753e87a957b69f8, __obf_175ed6e946c14a9b.Len())

		if __obf_1d6b91b9673acb44 != nil {
			break
		}
	}

	if __obf_b59cbf05240604a0 != nil {
		return nil, __obf_b59cbf05240604a0
	} else if __obf_1d6b91b9673acb44 == nil {
		return nil, errors.New("content too long to encode")
	}
	__obf_5b0073c32b55ab92 := &QRCode{
		Content: __obf_e41866e579426802,

		Level:         __obf_47158df92f26b0f4,
		VersionNumber: __obf_1d6b91b9673acb44.__obf_2066e640e4e50f14,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_2753e87a957b69f8: __obf_2753e87a957b69f8, __obf_eb8b3f015415859a: __obf_175ed6e946c14a9b, __obf_2066e640e4e50f14: *__obf_1d6b91b9673acb44,
	}

	return __obf_5b0073c32b55ab92, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_e41866e579426802 string, __obf_2066e640e4e50f14 int, __obf_47158df92f26b0f4 RecoveryLevel) (*QRCode, error) {
	var __obf_2753e87a957b69f8 *__obf_3ee8d058dd75e4b0

	switch {
	case __obf_2066e640e4e50f14 >= 1 && __obf_2066e640e4e50f14 <= 9:
		__obf_2753e87a957b69f8 = __obf_d8bb8ebef294bcff(__obf_4ef346b16a795892)
	case __obf_2066e640e4e50f14 >= 10 && __obf_2066e640e4e50f14 <= 26:
		__obf_2753e87a957b69f8 = __obf_d8bb8ebef294bcff(__obf_e891267bd1bcc38a)
	case __obf_2066e640e4e50f14 >= 27 && __obf_2066e640e4e50f14 <= 40:
		__obf_2753e87a957b69f8 = __obf_d8bb8ebef294bcff(__obf_054184439c33471a)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_2066e640e4e50f14)
	}

	var __obf_175ed6e946c14a9b *bitset.Bitset
	__obf_175ed6e946c14a9b, __obf_b59cbf05240604a0 := __obf_2753e87a957b69f8.__obf_395c063b7326115c([]byte(__obf_e41866e579426802))

	if __obf_b59cbf05240604a0 != nil {
		return nil, __obf_b59cbf05240604a0
	}
	__obf_1d6b91b9673acb44 := __obf_7704d615d7f75575(__obf_47158df92f26b0f4, __obf_2066e640e4e50f14)

	if __obf_1d6b91b9673acb44 == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_175ed6e946c14a9b.Len() > __obf_1d6b91b9673acb44.__obf_6ca1f14a38153ca1() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)", __obf_2066e640e4e50f14, __obf_175ed6e946c14a9b.
			Len(), __obf_1d6b91b9673acb44.__obf_6ca1f14a38153ca1())
	}
	__obf_5b0073c32b55ab92 := &QRCode{
		Content: __obf_e41866e579426802,

		Level:         __obf_47158df92f26b0f4,
		VersionNumber: __obf_1d6b91b9673acb44.__obf_2066e640e4e50f14,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_2753e87a957b69f8: __obf_2753e87a957b69f8, __obf_eb8b3f015415859a: __obf_175ed6e946c14a9b, __obf_2066e640e4e50f14: *__obf_1d6b91b9673acb44,
	}

	return __obf_5b0073c32b55ab92, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_5b0073c32b55ab92 *QRCode) Bitmap() [][]bool {
	__obf_5b0073c32b55ab92.
		// Build QR code.
		__obf_395c063b7326115c()

	return __obf_5b0073c32b55ab92.

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
		__obf_a9c08ca051c7e6c3.__obf_94902da22ea0120e()
}

func (__obf_5b0073c32b55ab92 *QRCode) Image(__obf_80f4ab6d4332abb9 int) image.Image {
	__obf_5b0073c32b55ab92.
		// Build QR code.
		__obf_395c063b7326115c()
	__obf_e3d393e0900b18ef := // Minimum pixels (both width and height) required.
		__obf_5b0073c32b55ab92.

			// Variable size support.
			__obf_a9c08ca051c7e6c3.__obf_80f4ab6d4332abb9

	if __obf_80f4ab6d4332abb9 < 0 {
		__obf_80f4ab6d4332abb9 = __obf_80f4ab6d4332abb9 * -1 * __obf_e3d393e0900b18ef
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_80f4ab6d4332abb9 < __obf_e3d393e0900b18ef {
		__obf_80f4ab6d4332abb9 = __obf_e3d393e0900b18ef
	}
	__obf_081a6c25fe653dd2 := // Output image.
		image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_80f4ab6d4332abb9,

			// Saves a few bytes to have them in this order
			__obf_80f4ab6d4332abb9}}
	__obf_012d6cea24c783cf := color.Palette([]color.Color{__obf_5b0073c32b55ab92.BackgroundColor, __obf_5b0073c32b55ab92.ForegroundColor})
	__obf_73ee74cf74079256 := image.NewPaletted(__obf_081a6c25fe653dd2, __obf_012d6cea24c783cf)
	__obf_1940ca6767f73b49 := uint8(__obf_73ee74cf74079256.Palette.Index(__obf_5b0073c32b55ab92.ForegroundColor))
	__obf_94902da22ea0120e := // QR code bitmap.
		__obf_5b0073c32b55ab92.

			// Map each image pixel to the nearest QR code module.
			__obf_a9c08ca051c7e6c3.__obf_94902da22ea0120e()
	__obf_4d3c5e1b247f5406 := float64(__obf_e3d393e0900b18ef) / float64(__obf_80f4ab6d4332abb9)
	for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < __obf_80f4ab6d4332abb9; __obf_16bfdfd36c335382++ {
		__obf_08b0d7a65e9bf07f := int(float64(__obf_16bfdfd36c335382) * __obf_4d3c5e1b247f5406)
		for __obf_7529943f67ee44cb := 0; __obf_7529943f67ee44cb < __obf_80f4ab6d4332abb9; __obf_7529943f67ee44cb++ {
			__obf_dfeeae06784a6756 := int(float64(__obf_7529943f67ee44cb) * __obf_4d3c5e1b247f5406)
			__obf_d3308b8b9ec8593a := __obf_94902da22ea0120e[__obf_08b0d7a65e9bf07f][__obf_dfeeae06784a6756]

			if __obf_d3308b8b9ec8593a {
				__obf_f9234edaec51b37e := __obf_73ee74cf74079256.PixOffset(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382)
				__obf_73ee74cf74079256.
					Pix[__obf_f9234edaec51b37e] = __obf_1940ca6767f73b49
			}
		}
	}

	return __obf_73ee74cf74079256
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_5b0073c32b55ab92 *QRCode) PNG(__obf_80f4ab6d4332abb9 int) ([]byte, error) {
	__obf_73ee74cf74079256 := __obf_5b0073c32b55ab92.Image(__obf_80f4ab6d4332abb9)
	__obf_2753e87a957b69f8 := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_0de44d307506fdeb bytes.Buffer
	__obf_b59cbf05240604a0 := __obf_2753e87a957b69f8.Encode(&__obf_0de44d307506fdeb, __obf_73ee74cf74079256)

	if __obf_b59cbf05240604a0 != nil {
		return nil, __obf_b59cbf05240604a0
	}

	return __obf_0de44d307506fdeb.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_5b0073c32b55ab92 *QRCode) Write(__obf_80f4ab6d4332abb9 int, __obf_f603e57a51325f8b io.Writer) error {
	var png []byte

	png, __obf_b59cbf05240604a0 := __obf_5b0073c32b55ab92.PNG(__obf_80f4ab6d4332abb9)

	if __obf_b59cbf05240604a0 != nil {
		return __obf_b59cbf05240604a0
	}
	_, __obf_b59cbf05240604a0 = __obf_f603e57a51325f8b.Write(png)
	return __obf_b59cbf05240604a0
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_5b0073c32b55ab92 *QRCode) WriteFile(__obf_80f4ab6d4332abb9 int, __obf_52c631b3d40eb87e string) error {
	var png []byte

	png, __obf_b59cbf05240604a0 := __obf_5b0073c32b55ab92.PNG(__obf_80f4ab6d4332abb9)

	if __obf_b59cbf05240604a0 != nil {
		return __obf_b59cbf05240604a0
	}

	return os.WriteFile(__obf_52c631b3d40eb87e, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_5b0073c32b55ab92 *QRCode) __obf_395c063b7326115c() {
	__obf_1d07d10e4aa75fd1 := __obf_5b0073c32b55ab92.__obf_2066e640e4e50f14.__obf_3955f9cbbeb36b7e(__obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Len())
	__obf_5b0073c32b55ab92.__obf_e2cb4eb9375cceea(__obf_1d07d10e4aa75fd1)
	__obf_5b0073c32b55ab92.__obf_53d37d46577398d2()
	__obf_175ed6e946c14a9b := __obf_5b0073c32b55ab92.__obf_0ec45dbc5e632011()

	const __obf_09b2c99454666af6 int = 8
	__obf_a1a2e98604fc16f7 := 0

	for __obf_3a54d62911a3bcfb := 0; __obf_3a54d62911a3bcfb < __obf_09b2c99454666af6; __obf_3a54d62911a3bcfb++ {
		var __obf_ab8d068af1cb47d6 *__obf_a9c08ca051c7e6c3
		var __obf_b59cbf05240604a0 error
		__obf_ab8d068af1cb47d6, __obf_b59cbf05240604a0 = __obf_3e8c0b7890083010(__obf_5b0073c32b55ab92.__obf_2066e640e4e50f14, __obf_3a54d62911a3bcfb, __obf_175ed6e946c14a9b, !__obf_5b0073c32b55ab92.Borderless)

		if __obf_b59cbf05240604a0 != nil {
			log.Panic(__obf_b59cbf05240604a0.Error())
		}
		__obf_955c9aaf43109184 := __obf_ab8d068af1cb47d6.__obf_955c9aaf43109184()
		if __obf_955c9aaf43109184 != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)", __obf_955c9aaf43109184, __obf_5b0073c32b55ab92.
				VersionNumber)
		}
		__obf_012d6cea24c783cf := __obf_ab8d068af1cb47d6.

			//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())
			__obf_9dab437c9fcdbe57()

		if __obf_5b0073c32b55ab92.__obf_a9c08ca051c7e6c3 == nil || __obf_012d6cea24c783cf < __obf_a1a2e98604fc16f7 {
			__obf_5b0073c32b55ab92.__obf_a9c08ca051c7e6c3 = __obf_ab8d068af1cb47d6
			__obf_5b0073c32b55ab92.__obf_3a54d62911a3bcfb = __obf_3a54d62911a3bcfb
			__obf_a1a2e98604fc16f7 = __obf_012d6cea24c783cf
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_5b0073c32b55ab92 *QRCode) __obf_e2cb4eb9375cceea(__obf_1d07d10e4aa75fd1 int) {
	__obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.
		AppendNumBools(__obf_1d07d10e4aa75fd1, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_5b0073c32b55ab92 *QRCode) __obf_0ec45dbc5e632011() *bitset.Bitset {
	// Split into blocks.
	type __obf_95a0aeaac032c66b struct {
		__obf_eb8b3f015415859a *bitset.Bitset
		__obf_4ec9daef443b36d0 int
	}
	__obf_497bb62bbd9ebbf5 := make([]__obf_95a0aeaac032c66b, __obf_5b0073c32b55ab92.__obf_2066e640e4e50f14.__obf_211dc13f80ab3495())
	__obf_3a15471c26131738 := 0
	__obf_a9837c7e909f1665 := 0
	__obf_a0846be02695f75b := 0

	for _, __obf_0de44d307506fdeb := range __obf_5b0073c32b55ab92.__obf_2066e640e4e50f14.__obf_497bb62bbd9ebbf5 {
		for __obf_0d2b5442049f8295 := 0; __obf_0d2b5442049f8295 < __obf_0de44d307506fdeb.__obf_211dc13f80ab3495; __obf_0d2b5442049f8295++ {
			__obf_3a15471c26131738 = __obf_a9837c7e909f1665
			__obf_a9837c7e909f1665 = __obf_3a15471c26131738 + __obf_0de44d307506fdeb.__obf_d1b26b395ae776d9*8
			__obf_c881148aff49931d := // Apply error correction to each block.
				__obf_0de44d307506fdeb.__obf_cd96a3f45948faa6 - __obf_0de44d307506fdeb.__obf_d1b26b395ae776d9
			__obf_497bb62bbd9ebbf5[__obf_a0846be02695f75b].__obf_eb8b3f015415859a = reedsolomon.Encode(__obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Substr(__obf_3a15471c26131738, __obf_a9837c7e909f1665), __obf_c881148aff49931d)
			__obf_497bb62bbd9ebbf5[__obf_a0846be02695f75b].__obf_4ec9daef443b36d0 = __obf_a9837c7e909f1665 - __obf_3a15471c26131738
			__obf_a0846be02695f75b++
		}
	}
	__obf_5f9887ab3541002d := // Interleave the blocks.

		bitset.New()
	__obf_b35f8d156abe5ba0 := // Combine data blocks.
		true
	for __obf_b39cd0fd88eaf53c := 0; __obf_b35f8d156abe5ba0; __obf_b39cd0fd88eaf53c += 8 {
		__obf_b35f8d156abe5ba0 = false

		for __obf_0d2b5442049f8295, __obf_0de44d307506fdeb := range __obf_497bb62bbd9ebbf5 {
			if __obf_b39cd0fd88eaf53c >= __obf_497bb62bbd9ebbf5[__obf_0d2b5442049f8295].__obf_4ec9daef443b36d0 {
				continue
			}
			__obf_5f9887ab3541002d.
				Append(__obf_0de44d307506fdeb.__obf_eb8b3f015415859a.Substr(__obf_b39cd0fd88eaf53c, __obf_b39cd0fd88eaf53c+8))
			__obf_b35f8d156abe5ba0 = true
		}
	}
	__obf_b35f8d156abe5ba0 = // Combine error correction blocks.
		true
	for __obf_b39cd0fd88eaf53c := 0; __obf_b35f8d156abe5ba0; __obf_b39cd0fd88eaf53c += 8 {
		__obf_b35f8d156abe5ba0 = false

		for __obf_0d2b5442049f8295, __obf_0de44d307506fdeb := range __obf_497bb62bbd9ebbf5 {
			__obf_8d7e72d2b9ea39cd := __obf_b39cd0fd88eaf53c + __obf_497bb62bbd9ebbf5[__obf_0d2b5442049f8295].__obf_4ec9daef443b36d0
			if __obf_8d7e72d2b9ea39cd >= __obf_497bb62bbd9ebbf5[__obf_0d2b5442049f8295].__obf_eb8b3f015415859a.Len() {
				continue
			}
			__obf_5f9887ab3541002d.
				Append(__obf_0de44d307506fdeb.__obf_eb8b3f015415859a.Substr(__obf_8d7e72d2b9ea39cd, __obf_8d7e72d2b9ea39cd+8))
			__obf_b35f8d156abe5ba0 = true
		}
	}
	__obf_5f9887ab3541002d.

		// Append remainder bits.
		AppendNumBools(__obf_5b0073c32b55ab92.__obf_2066e640e4e50f14.__obf_4bd5e5d599d53519, false)

	return __obf_5f9887ab3541002d
}

// addPadding pads the encoded data upto the full length required.
func (__obf_5b0073c32b55ab92 *QRCode) __obf_53d37d46577398d2() {
	__obf_6ca1f14a38153ca1 := __obf_5b0073c32b55ab92.__obf_2066e640e4e50f14.__obf_6ca1f14a38153ca1()

	if __obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Len() == __obf_6ca1f14a38153ca1 {
		return
	}
	__obf_5b0073c32b55ab92.

		// Pad to the nearest codeword boundary.
		__obf_eb8b3f015415859a.
		AppendNumBools(__obf_5b0073c32b55ab92.__obf_2066e640e4e50f14.__obf_8f66da0487eb5a9c(__obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Len()), false)
	__obf_3ae6fc7621d929fa := // Pad codewords 0b11101100 and 0b00010001.
		[2]*bitset.Bitset{
			bitset.New(true, true, true, false, true, true, false, false),
			bitset.New(false, false, false, true, false, false, false, true),
		}
	__obf_b39cd0fd88eaf53c := // Insert pad codewords alternately.
		0
	for __obf_6ca1f14a38153ca1-__obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Len() >= 8 {
		__obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.
			Append(__obf_3ae6fc7621d929fa[__obf_b39cd0fd88eaf53c])
		__obf_b39cd0fd88eaf53c = 1 - __obf_b39cd0fd88eaf53c // Alternate between 0 and 1.
	}

	if __obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Len() != __obf_6ca1f14a38153ca1 {
		log.Panicf("BUG: got len %d, expected %d", __obf_5b0073c32b55ab92.__obf_eb8b3f015415859a.Len(), __obf_6ca1f14a38153ca1)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_5b0073c32b55ab92 *QRCode) ToString(__obf_04e5e601afeee63b bool) string {
	__obf_09f55376f1035355 := __obf_5b0073c32b55ab92.Bitmap()
	var __obf_6ff0ba134ee1a062 bytes.Buffer
	for __obf_16bfdfd36c335382 := range __obf_09f55376f1035355 {
		for __obf_7529943f67ee44cb := range __obf_09f55376f1035355[__obf_16bfdfd36c335382] {
			if __obf_09f55376f1035355[__obf_16bfdfd36c335382][__obf_7529943f67ee44cb] != __obf_04e5e601afeee63b {
				__obf_6ff0ba134ee1a062.
					WriteString("  ")
			} else {
				__obf_6ff0ba134ee1a062.
					WriteString("██")
			}
		}
		__obf_6ff0ba134ee1a062.
			WriteString("\n")
	}
	return __obf_6ff0ba134ee1a062.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_5b0073c32b55ab92 *QRCode) ToSmallString(__obf_04e5e601afeee63b bool) string {
	__obf_09f55376f1035355 := __obf_5b0073c32b55ab92.Bitmap()
	var __obf_6ff0ba134ee1a062 bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < len(__obf_09f55376f1035355)-1; __obf_16bfdfd36c335382 += 2 {
		for __obf_7529943f67ee44cb := range __obf_09f55376f1035355[__obf_16bfdfd36c335382] {
			if __obf_09f55376f1035355[__obf_16bfdfd36c335382][__obf_7529943f67ee44cb] == __obf_09f55376f1035355[__obf_16bfdfd36c335382+1][__obf_7529943f67ee44cb] {
				if __obf_09f55376f1035355[__obf_16bfdfd36c335382][__obf_7529943f67ee44cb] != __obf_04e5e601afeee63b {
					__obf_6ff0ba134ee1a062.
						WriteString(" ")
				} else {
					__obf_6ff0ba134ee1a062.
						WriteString("█")
				}
			} else {
				if __obf_09f55376f1035355[__obf_16bfdfd36c335382][__obf_7529943f67ee44cb] != __obf_04e5e601afeee63b {
					__obf_6ff0ba134ee1a062.
						WriteString("▄")
				} else {
					__obf_6ff0ba134ee1a062.
						WriteString("▀")
				}
			}
		}
		__obf_6ff0ba134ee1a062.
			WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_09f55376f1035355)%2 == 1 {
		__obf_16bfdfd36c335382 := len(__obf_09f55376f1035355) - 1
		for __obf_7529943f67ee44cb := range __obf_09f55376f1035355[__obf_16bfdfd36c335382] {
			if __obf_09f55376f1035355[__obf_16bfdfd36c335382][__obf_7529943f67ee44cb] != __obf_04e5e601afeee63b {
				__obf_6ff0ba134ee1a062.
					WriteString(" ")
			} else {
				__obf_6ff0ba134ee1a062.
					WriteString("▀")
			}
		}
		__obf_6ff0ba134ee1a062.
			WriteString("\n")
	}
	return __obf_6ff0ba134ee1a062.String()
}
