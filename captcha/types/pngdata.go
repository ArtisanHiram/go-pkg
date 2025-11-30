package __obf_54406b1a1de84196

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"image"
	"image/png"
	"os"
	"path"
)

// PNGImage struct for PNG image data
type PNGImage struct {
	image.Image
}

// NewPNGImage creates a new PNG image data instance
func NewPNGImage(__obf_10d21d6285312279 image.Image) *PNGImage {
	return &PNGImage{__obf_10d21d6285312279}
}

// Get retrieves the original image
func (__obf_a26ee860b3cee35b *PNGImage) Get() image.Image {
	return __obf_a26ee860b3cee35b.Image
}

// SaveToFile is to save PNG as a file
func (__obf_a26ee860b3cee35b *PNGImage) SaveToFile(__obf_0a97a8569823a471 string) error {
	if __obf_a26ee860b3cee35b.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_4279250641db8149 *os.File
		__obf_d99742a1139768da error
	)
	__obf_d99742a1139768da = os.MkdirAll(path.Dir(__obf_0a97a8569823a471), os.ModePerm)
	if __obf_d99742a1139768da != nil {
		return __obf_d99742a1139768da
	}

	if _, __obf_d99742a1139768da = os.Stat(__obf_0a97a8569823a471); os.IsNotExist(__obf_d99742a1139768da) {
		__obf_4279250641db8149, __obf_d99742a1139768da = os.Create(__obf_0a97a8569823a471)
	} else {
		__obf_4279250641db8149, __obf_d99742a1139768da = os.OpenFile(__obf_0a97a8569823a471, os.O_RDWR, 0666)
	}
	if __obf_d99742a1139768da != nil {
		return __obf_d99742a1139768da
	}
	defer __obf_4279250641db8149.Close()

	return png.Encode(__obf_4279250641db8149,

		// ToBytes converts the PNG image to a byte array
		__obf_a26ee860b3cee35b.Image)
}

func (__obf_a26ee860b3cee35b *PNGImage) ToBytes() ([]byte, error) {
	if __obf_a26ee860b3cee35b.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_a26ee860b3cee35b.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_a26ee860b3cee35b *PNGImage) ToBase64() (string, error) {
	if __obf_a26ee860b3cee35b.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_a26ee860b3cee35b.Image)
}
