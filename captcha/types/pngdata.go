package __obf_6dcb1d06bd949756

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
func NewPNGImage(__obf_587078e38cdb4111 image.Image) *PNGImage {
	return &PNGImage{__obf_587078e38cdb4111}
}

// Get retrieves the original image
func (__obf_6b14c339de9bea34 *PNGImage) Get() image.Image {
	return __obf_6b14c339de9bea34.Image
}

// SaveToFile is to save PNG as a file
func (__obf_6b14c339de9bea34 *PNGImage) SaveToFile(__obf_b62b8c1a0e618495 string) error {
	if __obf_6b14c339de9bea34.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_6e11aba9569a8248 *os.File
		__obf_d349b3521516d4ce error
	)
	__obf_d349b3521516d4ce = os.MkdirAll(path.Dir(__obf_b62b8c1a0e618495), os.ModePerm)
	if __obf_d349b3521516d4ce != nil {
		return __obf_d349b3521516d4ce
	}

	if _, __obf_d349b3521516d4ce = os.Stat(__obf_b62b8c1a0e618495); os.IsNotExist(__obf_d349b3521516d4ce) {
		__obf_6e11aba9569a8248, __obf_d349b3521516d4ce = os.Create(__obf_b62b8c1a0e618495)
	} else {
		__obf_6e11aba9569a8248, __obf_d349b3521516d4ce = os.OpenFile(__obf_b62b8c1a0e618495, os.O_RDWR, 0666)
	}
	if __obf_d349b3521516d4ce != nil {
		return __obf_d349b3521516d4ce
	}
	defer __obf_6e11aba9569a8248.Close()

	return png.Encode(__obf_6e11aba9569a8248,

		// ToBytes converts the PNG image to a byte array
		__obf_6b14c339de9bea34.Image)
}

func (__obf_6b14c339de9bea34 *PNGImage) ToBytes() ([]byte, error) {
	if __obf_6b14c339de9bea34.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_6b14c339de9bea34.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_6b14c339de9bea34 *PNGImage) ToBase64() (string, error) {
	if __obf_6b14c339de9bea34.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_6b14c339de9bea34.Image)
}
