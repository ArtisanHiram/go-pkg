package __obf_0e73be9e4c3ea3fd

import (
	"image"
)

// PNGImage struct for PNG image data
type PNGImage struct {
	image.Image
}

// NewPNGImage creates a new PNG image data instance
func NewPNGImage(__obf_d1b36837c286905e image.Image) *PNGImage {
	return &PNGImage{
		__obf_d1b36837c286905e,
	}
}

// Get retrieves the original image
func (__obf_a5de954533461107 *PNGImage) Get() image.Image {
	return __obf_a5de954533461107.Image
}

// SaveToFile is to save PNG as a file
func (__obf_a5de954533461107 *PNGImage) SaveToFile(__obf_674636158f7d2792 string) error {
	if __obf_a5de954533461107.Image == nil {
		return ImageMissingDataErr
	}

	return __obf_fe2051ae37fdf99e(__obf_a5de954533461107.Image, __obf_674636158f7d2792, true, QualityNone)
}

// ToBytes converts the PNG image to a byte array
func (__obf_a5de954533461107 *PNGImage) ToBytes() ([]byte, error) {
	if __obf_a5de954533461107.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return EncodePNGToByte(__obf_a5de954533461107.Image)
}

// ToBase64Data converts the PNG image to Base64 data (without prefix)
func (__obf_a5de954533461107 *PNGImage) ToBase64Data() (string, error) {
	if __obf_a5de954533461107.Image == nil {
		return "", ImageEmptyErr
	}
	return EncodePNGToBase64Data(__obf_a5de954533461107.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_a5de954533461107 *PNGImage) ToBase64() (string, error) {
	if __obf_a5de954533461107.Image == nil {
		return "", ImageEmptyErr
	}
	return EncodePNGToBase64(__obf_a5de954533461107.Image)
}
