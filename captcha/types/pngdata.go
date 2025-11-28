package __obf_154ce31cd9492d61

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
func NewPNGImage(__obf_7e8dc6da215ff393 image.Image) *PNGImage {
	return &PNGImage{
		__obf_7e8dc6da215ff393,
	}
}

// Get retrieves the original image
func (__obf_ad3a8dc916d9ec04 *PNGImage) Get() image.Image {
	return __obf_ad3a8dc916d9ec04.Image
}

// SaveToFile is to save PNG as a file
func (__obf_ad3a8dc916d9ec04 *PNGImage) SaveToFile(__obf_2287c49fb91d4310 string) error {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_78396bf8ad30e393 *os.File
		__obf_ffc118ace13b322b error
	)

	__obf_ffc118ace13b322b = os.MkdirAll(path.Dir(__obf_2287c49fb91d4310), os.ModePerm)
	if __obf_ffc118ace13b322b != nil {
		return __obf_ffc118ace13b322b
	}

	if _, __obf_ffc118ace13b322b = os.Stat(__obf_2287c49fb91d4310); os.IsNotExist(__obf_ffc118ace13b322b) {
		__obf_78396bf8ad30e393, __obf_ffc118ace13b322b = os.Create(__obf_2287c49fb91d4310)
	} else {
		__obf_78396bf8ad30e393, __obf_ffc118ace13b322b = os.OpenFile(__obf_2287c49fb91d4310, os.O_RDWR, 0666)
	}
	if __obf_ffc118ace13b322b != nil {
		return __obf_ffc118ace13b322b
	}
	defer __obf_78396bf8ad30e393.Close()

	return png.Encode(__obf_78396bf8ad30e393, __obf_ad3a8dc916d9ec04.Image)
}

// ToBytes converts the PNG image to a byte array
func (__obf_ad3a8dc916d9ec04 *PNGImage) ToBytes() ([]byte, error) {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_ad3a8dc916d9ec04.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_ad3a8dc916d9ec04 *PNGImage) ToBase64() (string, error) {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_ad3a8dc916d9ec04.Image)
}
