package __obf_b0f3614e51931450

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
func NewPNGImage(__obf_71f0387c46e538dc image.Image) *PNGImage {
	return &PNGImage{
		__obf_71f0387c46e538dc,
	}
}

// Get retrieves the original image
func (__obf_587b9fceec62883e *PNGImage) Get() image.Image {
	return __obf_587b9fceec62883e.Image
}

// SaveToFile is to save PNG as a file
func (__obf_587b9fceec62883e *PNGImage) SaveToFile(__obf_18f0ff27836b66c6 string) error {
	if __obf_587b9fceec62883e.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_06449da585078edf *os.File
		__obf_369d77920507ae91 error
	)

	__obf_369d77920507ae91 = os.MkdirAll(path.Dir(__obf_18f0ff27836b66c6), os.ModePerm)
	if __obf_369d77920507ae91 != nil {
		return __obf_369d77920507ae91
	}

	if _, __obf_369d77920507ae91 = os.Stat(__obf_18f0ff27836b66c6); os.IsNotExist(__obf_369d77920507ae91) {
		__obf_06449da585078edf, __obf_369d77920507ae91 = os.Create(__obf_18f0ff27836b66c6)
	} else {
		__obf_06449da585078edf, __obf_369d77920507ae91 = os.OpenFile(__obf_18f0ff27836b66c6, os.O_RDWR, 0666)
	}
	if __obf_369d77920507ae91 != nil {
		return __obf_369d77920507ae91
	}
	defer __obf_06449da585078edf.Close()

	return png.Encode(__obf_06449da585078edf, __obf_587b9fceec62883e.Image)
}

// ToBytes converts the PNG image to a byte array
func (__obf_587b9fceec62883e *PNGImage) ToBytes() ([]byte, error) {
	if __obf_587b9fceec62883e.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_587b9fceec62883e.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_587b9fceec62883e *PNGImage) ToBase64() (string, error) {
	if __obf_587b9fceec62883e.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_587b9fceec62883e.Image)
}
