package __obf_0e73be9e4c3ea3fd

import (
	"image"
)

// jpegImageDta struct for JPEG image data
type JPEGImage struct {
	image.Image
}

// NewJPEGImage creates a new JPEG image data instance
func NewJPEGImage(__obf_d1b36837c286905e image.Image) *JPEGImage {
	return &JPEGImage{
		__obf_d1b36837c286905e,
	}
}

// Get retrieves the original image
func (__obf_a5de954533461107 *JPEGImage) Get() image.Image {
	return __obf_a5de954533461107.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_a5de954533461107 *JPEGImage) SaveToFile(__obf_674636158f7d2792 string, __obf_01f2bb93c2c9908f int) error {
	if __obf_a5de954533461107.Image == nil {
		return ImageMissingDataErr
	}

	return __obf_fe2051ae37fdf99e(__obf_a5de954533461107.Image, __obf_674636158f7d2792, false, __obf_01f2bb93c2c9908f)
}

// ToBytes converts the JPEG image to a byte array
func (__obf_a5de954533461107 *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_a5de954533461107.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return EncodeJPEGToByte(__obf_a5de954533461107.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_a5de954533461107 *JPEGImage) ToBytesWithQuality(__obf_d77346d9383af876 int) ([]byte, error) {
	if __obf_a5de954533461107.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_d77346d9383af876 <= QualityNone && __obf_d77346d9383af876 >= QualityLevel5 {
		return EncodeJPEGToByte(__obf_a5de954533461107.Image, __obf_d77346d9383af876)
	}
	return EncodeJPEGToByte(__obf_a5de954533461107.Image, QualityNone)
}

// ToBase64Data converts the JPEG image to Base64 data (without prefix)
func (__obf_a5de954533461107 *JPEGImage) ToBase64Data() (string, error) {
	if __obf_a5de954533461107.Image == nil {
		return "", ImageEmptyErr
	}

	return EncodeJPEGToBase64Data(__obf_a5de954533461107.Image, QualityNone)
}

// ToBase64DataWithQuality converts the JPEG image to Base64 data with specified quality (without prefix)
func (__obf_a5de954533461107 *JPEGImage) ToBase64DataWithQuality(__obf_d77346d9383af876 int) (string, error) {
	if __obf_a5de954533461107.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_d77346d9383af876 <= QualityNone && __obf_d77346d9383af876 >= QualityLevel5 {
		return EncodeJPEGToBase64Data(__obf_a5de954533461107.Image, __obf_d77346d9383af876)
	}
	return EncodeJPEGToBase64Data(__obf_a5de954533461107.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_a5de954533461107 *JPEGImage) ToBase64() (string, error) {
	if __obf_a5de954533461107.Image == nil {
		return "", ImageEmptyErr
	}

	return EncodeJPEGToBase64(__obf_a5de954533461107.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_a5de954533461107 *JPEGImage) ToBase64WithQuality(__obf_d77346d9383af876 int) (string, error) {
	if __obf_a5de954533461107.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_d77346d9383af876 <= QualityNone && __obf_d77346d9383af876 >= QualityLevel5 {
		return EncodeJPEGToBase64(__obf_a5de954533461107.Image, __obf_d77346d9383af876)
	}
	return EncodeJPEGToBase64(__obf_a5de954533461107.Image, QualityNone)
}
