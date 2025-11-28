package __obf_deb7e65d40f46bd0

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/v2/util"
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
func NewPNGImage(__obf_7e8297a5326972ed image.Image) *PNGImage {
	return &PNGImage{
		__obf_7e8297a5326972ed,
	}
}

// Get retrieves the original image
func (__obf_f012bb226e8e4a2d *PNGImage) Get() image.Image {
	return __obf_f012bb226e8e4a2d.Image
}

// SaveToFile is to save PNG as a file
func (__obf_f012bb226e8e4a2d *PNGImage) SaveToFile(__obf_68a9f2fdfea43cce string) error {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_a148c9a511cdb474 *os.File
		__obf_cdc2a8664de7ba0a error
	)

	__obf_cdc2a8664de7ba0a = os.MkdirAll(path.Dir(__obf_68a9f2fdfea43cce), os.ModePerm)
	if __obf_cdc2a8664de7ba0a != nil {
		return __obf_cdc2a8664de7ba0a
	}

	if _, __obf_cdc2a8664de7ba0a = os.Stat(__obf_68a9f2fdfea43cce); os.IsNotExist(__obf_cdc2a8664de7ba0a) {
		__obf_a148c9a511cdb474, __obf_cdc2a8664de7ba0a = os.Create(__obf_68a9f2fdfea43cce)
	} else {
		__obf_a148c9a511cdb474, __obf_cdc2a8664de7ba0a = os.OpenFile(__obf_68a9f2fdfea43cce, os.O_RDWR, 0666)
	}
	if __obf_cdc2a8664de7ba0a != nil {
		return __obf_cdc2a8664de7ba0a
	}
	defer __obf_a148c9a511cdb474.Close()

	return png.Encode(__obf_a148c9a511cdb474, __obf_f012bb226e8e4a2d.Image)
}

// ToBytes converts the PNG image to a byte array
func (__obf_f012bb226e8e4a2d *PNGImage) ToBytes() ([]byte, error) {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_f012bb226e8e4a2d.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_f012bb226e8e4a2d *PNGImage) ToBase64() (string, error) {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_f012bb226e8e4a2d.Image)
}
