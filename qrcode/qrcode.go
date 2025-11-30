package __obf_380fc15a74e6e942

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
func Encode(__obf_6863f49ead0d68c9 string, __obf_fc3c629b710f5957 RecoveryLevel, __obf_5077244fac9670a3 int) ([]byte, error) {
	var __obf_a486326c9a0a156b *QRCode
	__obf_a486326c9a0a156b, __obf_643ee744984d6f9b := New(__obf_6863f49ead0d68c9, __obf_fc3c629b710f5957)

	if __obf_643ee744984d6f9b != nil {
		return nil, __obf_643ee744984d6f9b
	}

	return __obf_a486326c9a0a156b.PNG(__obf_5077244fac9670a3)
}

// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func WriteFile(__obf_6863f49ead0d68c9 string, __obf_fc3c629b710f5957 RecoveryLevel, __obf_5077244fac9670a3 int, __obf_db953d3270cbb769 string) error {
	var __obf_a486326c9a0a156b *QRCode
	__obf_a486326c9a0a156b, __obf_643ee744984d6f9b := New(__obf_6863f49ead0d68c9, __obf_fc3c629b710f5957)

	if __obf_643ee744984d6f9b != nil {
		return __obf_643ee744984d6f9b
	}

	return __obf_a486326c9a0a156b.WriteFile(__obf_5077244fac9670a3,

		// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
		// With WriteColorFile you can also specify the colors you want to use.
		//
		// size is both the image width and height in pixels. If size is too small then
		// a larger image is silently written. Negative values for size cause a variable
		// sized image to be written: See the documentation for Image().
		__obf_db953d3270cbb769)
}

func WriteColorFile(__obf_6863f49ead0d68c9 string, __obf_fc3c629b710f5957 RecoveryLevel, __obf_5077244fac9670a3 int, __obf_b39b603f880dbb9e, __obf_d6d8d5e70d691b0d color.Color, __obf_db953d3270cbb769 string) error {

	var __obf_a486326c9a0a156b *QRCode
	__obf_a486326c9a0a156b, __obf_643ee744984d6f9b := New(__obf_6863f49ead0d68c9, __obf_fc3c629b710f5957)
	__obf_a486326c9a0a156b.
		BackgroundColor = __obf_b39b603f880dbb9e
	__obf_a486326c9a0a156b.
		ForegroundColor = __obf_d6d8d5e70d691b0d

	if __obf_643ee744984d6f9b != nil {
		return __obf_643ee744984d6f9b
	}

	return __obf_a486326c9a0a156b.WriteFile(__obf_5077244fac9670a3,

		// A QRCode represents a valid encoded QRCode.
		__obf_db953d3270cbb769)
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
	__obf_c275e7e500011398 *__obf_8a2e265fe487f1a2
	__obf_264d1de845406472 __obf_870ce853fd3b5622
	__obf_75065e802bd2136f *bitset.Bitset
	__obf_87d3969cdacfe5ff *__obf_87d3969cdacfe5ff
	__obf_8833ff9f1beb17df int
}

// New constructs a QRCode.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.New("my content", qrcode.Medium)
//
// An error occurs if the content is too long.
func New(__obf_6863f49ead0d68c9 string, __obf_fc3c629b710f5957 RecoveryLevel) (*QRCode, error) {
	__obf_24d80c2aaf84d6c1 := []__obf_d05a9d69f4a51132{__obf_f0e68bac73c2a21c, __obf_b270254bf247f115, __obf_602ce8ec28abacf4}

	var __obf_c275e7e500011398 *__obf_8a2e265fe487f1a2
	var __obf_fe987f8387318f16 *bitset.Bitset
	var __obf_67cb565eb7f8949d *__obf_870ce853fd3b5622
	var __obf_643ee744984d6f9b error

	for _, __obf_bfc4cdf552630747 := range __obf_24d80c2aaf84d6c1 {
		__obf_c275e7e500011398 = __obf_6bdf8930191977ce(__obf_bfc4cdf552630747)
		__obf_fe987f8387318f16, __obf_643ee744984d6f9b = __obf_c275e7e500011398.__obf_d827976a41e20008([]byte(__obf_6863f49ead0d68c9))

		if __obf_643ee744984d6f9b != nil {
			continue
		}
		__obf_67cb565eb7f8949d = __obf_945c524e956c9051(__obf_fc3c629b710f5957, __obf_c275e7e500011398, __obf_fe987f8387318f16.Len())

		if __obf_67cb565eb7f8949d != nil {
			break
		}
	}

	if __obf_643ee744984d6f9b != nil {
		return nil, __obf_643ee744984d6f9b
	} else if __obf_67cb565eb7f8949d == nil {
		return nil, errors.New("content too long to encode")
	}
	__obf_a486326c9a0a156b := &QRCode{
		Content: __obf_6863f49ead0d68c9,

		Level:         __obf_fc3c629b710f5957,
		VersionNumber: __obf_67cb565eb7f8949d.__obf_264d1de845406472,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_c275e7e500011398: __obf_c275e7e500011398, __obf_75065e802bd2136f: __obf_fe987f8387318f16, __obf_264d1de845406472: *__obf_67cb565eb7f8949d,
	}

	return __obf_a486326c9a0a156b, nil
}

// NewWithForcedVersion constructs a QRCode of a specific version.
//
//	var q *qrcode.QRCode
//	q, err := qrcode.NewWithForcedVersion("my content", 25, qrcode.Medium)
//
// An error occurs in case of invalid version.
func NewWithForcedVersion(__obf_6863f49ead0d68c9 string, __obf_264d1de845406472 int, __obf_fc3c629b710f5957 RecoveryLevel) (*QRCode, error) {
	var __obf_c275e7e500011398 *__obf_8a2e265fe487f1a2

	switch {
	case __obf_264d1de845406472 >= 1 && __obf_264d1de845406472 <= 9:
		__obf_c275e7e500011398 = __obf_6bdf8930191977ce(__obf_f0e68bac73c2a21c)
	case __obf_264d1de845406472 >= 10 && __obf_264d1de845406472 <= 26:
		__obf_c275e7e500011398 = __obf_6bdf8930191977ce(__obf_b270254bf247f115)
	case __obf_264d1de845406472 >= 27 && __obf_264d1de845406472 <= 40:
		__obf_c275e7e500011398 = __obf_6bdf8930191977ce(__obf_602ce8ec28abacf4)
	default:
		return nil, fmt.Errorf("invalid version %d (expected 1-40 inclusive)", __obf_264d1de845406472)
	}

	var __obf_fe987f8387318f16 *bitset.Bitset
	__obf_fe987f8387318f16, __obf_643ee744984d6f9b := __obf_c275e7e500011398.__obf_d827976a41e20008([]byte(__obf_6863f49ead0d68c9))

	if __obf_643ee744984d6f9b != nil {
		return nil, __obf_643ee744984d6f9b
	}
	__obf_67cb565eb7f8949d := __obf_c5839fc8a96e682a(__obf_fc3c629b710f5957, __obf_264d1de845406472)

	if __obf_67cb565eb7f8949d == nil {
		return nil, errors.New("cannot find QR Code version")
	}

	if __obf_fe987f8387318f16.Len() > __obf_67cb565eb7f8949d.__obf_35f302a87ce4a7f1() {
		return nil, fmt.Errorf("cannot encode QR code: content too large for fixed size QR Code version %d (encoded length is %d bits, maximum length is %d bits)", __obf_264d1de845406472, __obf_fe987f8387318f16.
			Len(), __obf_67cb565eb7f8949d.__obf_35f302a87ce4a7f1())
	}
	__obf_a486326c9a0a156b := &QRCode{
		Content: __obf_6863f49ead0d68c9,

		Level:         __obf_fc3c629b710f5957,
		VersionNumber: __obf_67cb565eb7f8949d.__obf_264d1de845406472,

		ForegroundColor: color.Black,
		BackgroundColor: color.White, __obf_c275e7e500011398: __obf_c275e7e500011398, __obf_75065e802bd2136f: __obf_fe987f8387318f16, __obf_264d1de845406472: *__obf_67cb565eb7f8949d,
	}

	return __obf_a486326c9a0a156b, nil
}

// Bitmap returns the QR Code as a 2D array of 1-bit pixels.
//
// bitmap[y][x] is true if the pixel at (x, y) is set.
//
// The bitmap includes the required "quiet zone" around the QR Code to aid
// decoding.
func (__obf_a486326c9a0a156b *QRCode) Bitmap() [][]bool {
	__obf_a486326c9a0a156b.
		// Build QR code.
		__obf_d827976a41e20008()

	return __obf_a486326c9a0a156b.

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
		__obf_87d3969cdacfe5ff.__obf_0d0a40eb3c5b3d00()
}

func (__obf_a486326c9a0a156b *QRCode) Image(__obf_5077244fac9670a3 int) image.Image {
	__obf_a486326c9a0a156b.
		// Build QR code.
		__obf_d827976a41e20008()
	__obf_703c48bee433c373 := // Minimum pixels (both width and height) required.
		__obf_a486326c9a0a156b.

			// Variable size support.
			__obf_87d3969cdacfe5ff.__obf_5077244fac9670a3

	if __obf_5077244fac9670a3 < 0 {
		__obf_5077244fac9670a3 = __obf_5077244fac9670a3 * -1 * __obf_703c48bee433c373
	}

	// Actual pixels available to draw the symbol. Automatically increase the
	// image size if it's not large enough.
	if __obf_5077244fac9670a3 < __obf_703c48bee433c373 {
		__obf_5077244fac9670a3 = __obf_703c48bee433c373
	}
	__obf_58372b1da9bc6a22 := // Output image.
		image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{__obf_5077244fac9670a3,

			// Saves a few bytes to have them in this order
			__obf_5077244fac9670a3}}
	__obf_4b93c521d976e862 := color.Palette([]color.Color{__obf_a486326c9a0a156b.BackgroundColor, __obf_a486326c9a0a156b.ForegroundColor})
	__obf_b93a248f1e5212c7 := image.NewPaletted(__obf_58372b1da9bc6a22, __obf_4b93c521d976e862)
	__obf_1f861d8ac351fe95 := uint8(__obf_b93a248f1e5212c7.Palette.Index(__obf_a486326c9a0a156b.ForegroundColor))
	__obf_0d0a40eb3c5b3d00 := // QR code bitmap.
		__obf_a486326c9a0a156b.

			// Map each image pixel to the nearest QR code module.
			__obf_87d3969cdacfe5ff.__obf_0d0a40eb3c5b3d00()
	__obf_80c312e581650d7e := float64(__obf_703c48bee433c373) / float64(__obf_5077244fac9670a3)
	for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < __obf_5077244fac9670a3; __obf_8eaa1145d0e6ec9a++ {
		__obf_597c19901f833f4c := int(float64(__obf_8eaa1145d0e6ec9a) * __obf_80c312e581650d7e)
		for __obf_e3f2eace849c7041 := 0; __obf_e3f2eace849c7041 < __obf_5077244fac9670a3; __obf_e3f2eace849c7041++ {
			__obf_8007a28d1a40d6ee := int(float64(__obf_e3f2eace849c7041) * __obf_80c312e581650d7e)
			__obf_3591b555f7f18238 := __obf_0d0a40eb3c5b3d00[__obf_597c19901f833f4c][__obf_8007a28d1a40d6ee]

			if __obf_3591b555f7f18238 {
				__obf_cc1fa433d6dd306d := __obf_b93a248f1e5212c7.PixOffset(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a)
				__obf_b93a248f1e5212c7.
					Pix[__obf_cc1fa433d6dd306d] = __obf_1f861d8ac351fe95
			}
		}
	}

	return __obf_b93a248f1e5212c7
}

// PNG returns the QR Code as a PNG image.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
func (__obf_a486326c9a0a156b *QRCode) PNG(__obf_5077244fac9670a3 int) ([]byte, error) {
	__obf_b93a248f1e5212c7 := __obf_a486326c9a0a156b.Image(__obf_5077244fac9670a3)
	__obf_c275e7e500011398 := png.Encoder{CompressionLevel: png.BestCompression}

	var __obf_b58cd35067d92ce7 bytes.Buffer
	__obf_643ee744984d6f9b := __obf_c275e7e500011398.Encode(&__obf_b58cd35067d92ce7, __obf_b93a248f1e5212c7)

	if __obf_643ee744984d6f9b != nil {
		return nil, __obf_643ee744984d6f9b
	}

	return __obf_b58cd35067d92ce7.Bytes(), nil
}

// Write writes the QR Code as a PNG image to io.Writer.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_a486326c9a0a156b *QRCode) Write(__obf_5077244fac9670a3 int, __obf_7d6e3ab1e312bea2 io.Writer) error {
	var png []byte

	png, __obf_643ee744984d6f9b := __obf_a486326c9a0a156b.PNG(__obf_5077244fac9670a3)

	if __obf_643ee744984d6f9b != nil {
		return __obf_643ee744984d6f9b
	}
	_, __obf_643ee744984d6f9b = __obf_7d6e3ab1e312bea2.Write(png)
	return __obf_643ee744984d6f9b
}

// WriteFile writes the QR Code as a PNG image to the specified file.
//
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a
// variable sized image to be written: See the documentation for Image().
func (__obf_a486326c9a0a156b *QRCode) WriteFile(__obf_5077244fac9670a3 int, __obf_db953d3270cbb769 string) error {
	var png []byte

	png, __obf_643ee744984d6f9b := __obf_a486326c9a0a156b.PNG(__obf_5077244fac9670a3)

	if __obf_643ee744984d6f9b != nil {
		return __obf_643ee744984d6f9b
	}

	return os.WriteFile(__obf_db953d3270cbb769, png, os.FileMode(0644))
}

// encode completes the steps required to encode the QR Code. These include
// adding the terminator bits and padding, splitting the data into blocks and
// applying the error correction, and selecting the best data mask.
func (__obf_a486326c9a0a156b *QRCode) __obf_d827976a41e20008() {
	__obf_52c948bfc84a0bb4 := __obf_a486326c9a0a156b.__obf_264d1de845406472.__obf_7f895ab7638986af(__obf_a486326c9a0a156b.__obf_75065e802bd2136f.Len())
	__obf_a486326c9a0a156b.__obf_a135402cf59533c0(__obf_52c948bfc84a0bb4)
	__obf_a486326c9a0a156b.__obf_844c090ef4f6ba49()
	__obf_fe987f8387318f16 := __obf_a486326c9a0a156b.__obf_db8a699ad81c1d77()

	const __obf_4ba3d84a74703f7c int = 8
	__obf_d40e98e6d42d4c72 := 0

	for __obf_8833ff9f1beb17df := 0; __obf_8833ff9f1beb17df < __obf_4ba3d84a74703f7c; __obf_8833ff9f1beb17df++ {
		var __obf_0f0d15721c6081af *__obf_87d3969cdacfe5ff
		var __obf_643ee744984d6f9b error
		__obf_0f0d15721c6081af, __obf_643ee744984d6f9b = __obf_2e60da472094edd0(__obf_a486326c9a0a156b.__obf_264d1de845406472, __obf_8833ff9f1beb17df, __obf_fe987f8387318f16, !__obf_a486326c9a0a156b.Borderless)

		if __obf_643ee744984d6f9b != nil {
			log.Panic(__obf_643ee744984d6f9b.Error())
		}
		__obf_20c8c0869c8b5b37 := __obf_0f0d15721c6081af.__obf_20c8c0869c8b5b37()
		if __obf_20c8c0869c8b5b37 != 0 {
			log.Panicf("bug: numEmptyModules is %d (expected 0) (version=%d)", __obf_20c8c0869c8b5b37, __obf_a486326c9a0a156b.
				VersionNumber)
		}
		__obf_4b93c521d976e862 := __obf_0f0d15721c6081af.

			//log.Printf("mask=%d p=%3d p1=%3d p2=%3d p3=%3d p4=%d\n", mask, p, s.penalty1(), s.penalty2(), s.penalty3(), s.penalty4())
			__obf_a62e72882ae00260()

		if __obf_a486326c9a0a156b.__obf_87d3969cdacfe5ff == nil || __obf_4b93c521d976e862 < __obf_d40e98e6d42d4c72 {
			__obf_a486326c9a0a156b.__obf_87d3969cdacfe5ff = __obf_0f0d15721c6081af
			__obf_a486326c9a0a156b.__obf_8833ff9f1beb17df = __obf_8833ff9f1beb17df
			__obf_d40e98e6d42d4c72 = __obf_4b93c521d976e862
		}
	}
}

// addTerminatorBits adds final terminator bits to the encoded data.
//
// The number of terminator bits required is determined when the QR Code version
// is chosen (which itself depends on the length of the data encoded). The
// terminator bits are thus added after the QR Code version
// is chosen, rather than at the data encoding stage.
func (__obf_a486326c9a0a156b *QRCode) __obf_a135402cf59533c0(__obf_52c948bfc84a0bb4 int) {
	__obf_a486326c9a0a156b.__obf_75065e802bd2136f.
		AppendNumBools(__obf_52c948bfc84a0bb4, false)
}

// encodeBlocks takes the completed (terminated & padded) encoded data, splits
// the data into blocks (as specified by the QR Code version), applies error
// correction to each block, then interleaves the blocks together.
//
// The QR Code's final data sequence is returned.
func (__obf_a486326c9a0a156b *QRCode) __obf_db8a699ad81c1d77() *bitset.Bitset {
	// Split into blocks.
	type __obf_0241e6f86e52a031 struct {
		__obf_75065e802bd2136f *bitset.Bitset
		__obf_f5daadaf42ec702f int
	}
	__obf_d754684ce53f7200 := make([]__obf_0241e6f86e52a031, __obf_a486326c9a0a156b.__obf_264d1de845406472.__obf_3d59d9d22bc9936c())
	__obf_0fb8459f0ac45df2 := 0
	__obf_5fb9e8aaba600698 := 0
	__obf_ecf58eac7ed5a8af := 0

	for _, __obf_b58cd35067d92ce7 := range __obf_a486326c9a0a156b.__obf_264d1de845406472.__obf_d754684ce53f7200 {
		for __obf_e8fc5f84cbc468cf := 0; __obf_e8fc5f84cbc468cf < __obf_b58cd35067d92ce7.__obf_3d59d9d22bc9936c; __obf_e8fc5f84cbc468cf++ {
			__obf_0fb8459f0ac45df2 = __obf_5fb9e8aaba600698
			__obf_5fb9e8aaba600698 = __obf_0fb8459f0ac45df2 + __obf_b58cd35067d92ce7.__obf_896247f00fa0ee79*8
			__obf_228f69ef6b536527 := // Apply error correction to each block.
				__obf_b58cd35067d92ce7.__obf_9be363f84174ff9f - __obf_b58cd35067d92ce7.__obf_896247f00fa0ee79
			__obf_d754684ce53f7200[__obf_ecf58eac7ed5a8af].__obf_75065e802bd2136f = reedsolomon.Encode(__obf_a486326c9a0a156b.__obf_75065e802bd2136f.Substr(__obf_0fb8459f0ac45df2, __obf_5fb9e8aaba600698), __obf_228f69ef6b536527)
			__obf_d754684ce53f7200[__obf_ecf58eac7ed5a8af].__obf_f5daadaf42ec702f = __obf_5fb9e8aaba600698 - __obf_0fb8459f0ac45df2
			__obf_ecf58eac7ed5a8af++
		}
	}
	__obf_0f51954847d61555 := // Interleave the blocks.

		bitset.New()
	__obf_a77245ab1b77f6ce := // Combine data blocks.
		true
	for __obf_f1022e2071655fce := 0; __obf_a77245ab1b77f6ce; __obf_f1022e2071655fce += 8 {
		__obf_a77245ab1b77f6ce = false

		for __obf_e8fc5f84cbc468cf, __obf_b58cd35067d92ce7 := range __obf_d754684ce53f7200 {
			if __obf_f1022e2071655fce >= __obf_d754684ce53f7200[__obf_e8fc5f84cbc468cf].__obf_f5daadaf42ec702f {
				continue
			}
			__obf_0f51954847d61555.
				Append(__obf_b58cd35067d92ce7.__obf_75065e802bd2136f.Substr(__obf_f1022e2071655fce, __obf_f1022e2071655fce+8))
			__obf_a77245ab1b77f6ce = true
		}
	}
	__obf_a77245ab1b77f6ce = // Combine error correction blocks.
		true
	for __obf_f1022e2071655fce := 0; __obf_a77245ab1b77f6ce; __obf_f1022e2071655fce += 8 {
		__obf_a77245ab1b77f6ce = false

		for __obf_e8fc5f84cbc468cf, __obf_b58cd35067d92ce7 := range __obf_d754684ce53f7200 {
			__obf_9dd48857b229d5dd := __obf_f1022e2071655fce + __obf_d754684ce53f7200[__obf_e8fc5f84cbc468cf].__obf_f5daadaf42ec702f
			if __obf_9dd48857b229d5dd >= __obf_d754684ce53f7200[__obf_e8fc5f84cbc468cf].__obf_75065e802bd2136f.Len() {
				continue
			}
			__obf_0f51954847d61555.
				Append(__obf_b58cd35067d92ce7.__obf_75065e802bd2136f.Substr(__obf_9dd48857b229d5dd, __obf_9dd48857b229d5dd+8))
			__obf_a77245ab1b77f6ce = true
		}
	}
	__obf_0f51954847d61555.

		// Append remainder bits.
		AppendNumBools(__obf_a486326c9a0a156b.__obf_264d1de845406472.__obf_059895faaf46eda6, false)

	return __obf_0f51954847d61555
}

// addPadding pads the encoded data upto the full length required.
func (__obf_a486326c9a0a156b *QRCode) __obf_844c090ef4f6ba49() {
	__obf_35f302a87ce4a7f1 := __obf_a486326c9a0a156b.__obf_264d1de845406472.__obf_35f302a87ce4a7f1()

	if __obf_a486326c9a0a156b.__obf_75065e802bd2136f.Len() == __obf_35f302a87ce4a7f1 {
		return
	}
	__obf_a486326c9a0a156b.

		// Pad to the nearest codeword boundary.
		__obf_75065e802bd2136f.
		AppendNumBools(__obf_a486326c9a0a156b.__obf_264d1de845406472.__obf_d5e0eb0d36587ec2(__obf_a486326c9a0a156b.__obf_75065e802bd2136f.Len()), false)
	__obf_7d09c59f1bfa356f := // Pad codewords 0b11101100 and 0b00010001.
		[2]*bitset.Bitset{
			bitset.New(true, true, true, false, true, true, false, false),
			bitset.New(false, false, false, true, false, false, false, true),
		}
	__obf_f1022e2071655fce := // Insert pad codewords alternately.
		0
	for __obf_35f302a87ce4a7f1-__obf_a486326c9a0a156b.__obf_75065e802bd2136f.Len() >= 8 {
		__obf_a486326c9a0a156b.__obf_75065e802bd2136f.
			Append(__obf_7d09c59f1bfa356f[__obf_f1022e2071655fce])
		__obf_f1022e2071655fce = 1 - __obf_f1022e2071655fce // Alternate between 0 and 1.
	}

	if __obf_a486326c9a0a156b.__obf_75065e802bd2136f.Len() != __obf_35f302a87ce4a7f1 {
		log.Panicf("BUG: got len %d, expected %d", __obf_a486326c9a0a156b.__obf_75065e802bd2136f.Len(), __obf_35f302a87ce4a7f1)
	}
}

// ToString produces a multi-line string that forms a QR-code image.
func (__obf_a486326c9a0a156b *QRCode) ToString(__obf_69114b2883e65cdd bool) string {
	__obf_eab1d90e2b2cc6ca := __obf_a486326c9a0a156b.Bitmap()
	var __obf_6be124f4f98276cf bytes.Buffer
	for __obf_8eaa1145d0e6ec9a := range __obf_eab1d90e2b2cc6ca {
		for __obf_e3f2eace849c7041 := range __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a] {
			if __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a][__obf_e3f2eace849c7041] != __obf_69114b2883e65cdd {
				__obf_6be124f4f98276cf.
					WriteString("  ")
			} else {
				__obf_6be124f4f98276cf.
					WriteString("██")
			}
		}
		__obf_6be124f4f98276cf.
			WriteString("\n")
	}
	return __obf_6be124f4f98276cf.String()
}

// ToSmallString produces a multi-line string that forms a QR-code image, a
// factor two smaller in x and y then ToString.
func (__obf_a486326c9a0a156b *QRCode) ToSmallString(__obf_69114b2883e65cdd bool) string {
	__obf_eab1d90e2b2cc6ca := __obf_a486326c9a0a156b.Bitmap()
	var __obf_6be124f4f98276cf bytes.Buffer
	// if there is an odd number of rows, the last one needs special treatment
	for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < len(__obf_eab1d90e2b2cc6ca)-1; __obf_8eaa1145d0e6ec9a += 2 {
		for __obf_e3f2eace849c7041 := range __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a] {
			if __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a][__obf_e3f2eace849c7041] == __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a+1][__obf_e3f2eace849c7041] {
				if __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a][__obf_e3f2eace849c7041] != __obf_69114b2883e65cdd {
					__obf_6be124f4f98276cf.
						WriteString(" ")
				} else {
					__obf_6be124f4f98276cf.
						WriteString("█")
				}
			} else {
				if __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a][__obf_e3f2eace849c7041] != __obf_69114b2883e65cdd {
					__obf_6be124f4f98276cf.
						WriteString("▄")
				} else {
					__obf_6be124f4f98276cf.
						WriteString("▀")
				}
			}
		}
		__obf_6be124f4f98276cf.
			WriteString("\n")
	}
	// special treatment for the last row if odd
	if len(__obf_eab1d90e2b2cc6ca)%2 == 1 {
		__obf_8eaa1145d0e6ec9a := len(__obf_eab1d90e2b2cc6ca) - 1
		for __obf_e3f2eace849c7041 := range __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a] {
			if __obf_eab1d90e2b2cc6ca[__obf_8eaa1145d0e6ec9a][__obf_e3f2eace849c7041] != __obf_69114b2883e65cdd {
				__obf_6be124f4f98276cf.
					WriteString(" ")
			} else {
				__obf_6be124f4f98276cf.
					WriteString("▀")
			}
		}
		__obf_6be124f4f98276cf.
			WriteString("\n")
	}
	return __obf_6be124f4f98276cf.String()
}
