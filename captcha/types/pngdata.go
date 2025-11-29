package __obf_244ef50d49151021

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
func NewPNGImage(__obf_4b8ec8035512b29f image.Image) *PNGImage {
	return &PNGImage{__obf_4b8ec8035512b29f}
}

// Get retrieves the original image
func (__obf_dbbef660ec878016 *PNGImage) Get() image.Image {
	return __obf_dbbef660ec878016.Image
}

// SaveToFile is to save PNG as a file
func (__obf_dbbef660ec878016 *PNGImage) SaveToFile(__obf_6f182e746b9b2f31 string) error {
	if __obf_dbbef660ec878016.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_1ad0cb034eca86b6 *os.File
		__obf_45f5b0f1f05cc095 error
	)
	__obf_45f5b0f1f05cc095 = os.MkdirAll(path.Dir(__obf_6f182e746b9b2f31), os.ModePerm)
	if __obf_45f5b0f1f05cc095 != nil {
		return __obf_45f5b0f1f05cc095
	}

	if _, __obf_45f5b0f1f05cc095 = os.Stat(__obf_6f182e746b9b2f31); os.IsNotExist(__obf_45f5b0f1f05cc095) {
		__obf_1ad0cb034eca86b6, __obf_45f5b0f1f05cc095 = os.Create(__obf_6f182e746b9b2f31)
	} else {
		__obf_1ad0cb034eca86b6, __obf_45f5b0f1f05cc095 = os.OpenFile(__obf_6f182e746b9b2f31, os.O_RDWR, 0666)
	}
	if __obf_45f5b0f1f05cc095 != nil {
		return __obf_45f5b0f1f05cc095
	}
	defer __obf_1ad0cb034eca86b6.Close()

	return png.Encode(__obf_1ad0cb034eca86b6,

		// ToBytes converts the PNG image to a byte array
		__obf_dbbef660ec878016.Image)
}

func (__obf_dbbef660ec878016 *PNGImage) ToBytes() ([]byte, error) {
	if __obf_dbbef660ec878016.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_dbbef660ec878016.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_dbbef660ec878016 *PNGImage) ToBase64() (string, error) {
	if __obf_dbbef660ec878016.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_dbbef660ec878016.Image)
}
