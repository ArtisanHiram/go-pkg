package __obf_b0f3614e51931450

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"image"
	"image/jpeg"
	"os"
	"path"
)

// jpegImageDta struct for JPEG image data
type JPEGImage struct {
	image.Image
}

// NewJPEGImage creates a new JPEG image data instance
func NewJPEGImage(__obf_71f0387c46e538dc image.Image) *JPEGImage {
	return &JPEGImage{
		__obf_71f0387c46e538dc,
	}
}

// Get retrieves the original image
func (__obf_587b9fceec62883e *JPEGImage) Get() image.Image {
	return __obf_587b9fceec62883e.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_587b9fceec62883e *JPEGImage) SaveToFile(__obf_18f0ff27836b66c6 string, __obf_b5f41116292c66d6 int) error {
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
	return jpeg.Encode(__obf_06449da585078edf, __obf_587b9fceec62883e.Image, &jpeg.Options{Quality: __obf_b5f41116292c66d6})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_587b9fceec62883e *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_587b9fceec62883e.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_587b9fceec62883e.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_587b9fceec62883e *JPEGImage) ToBytesWithQuality(__obf_f276064e93bacc8b int) ([]byte, error) {
	if __obf_587b9fceec62883e.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_f276064e93bacc8b <= QualityNone && __obf_f276064e93bacc8b >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_587b9fceec62883e.Image, __obf_f276064e93bacc8b)
	}
	return util.EncodeJPEGToByte(__obf_587b9fceec62883e.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_587b9fceec62883e *JPEGImage) ToBase64() (string, error) {
	if __obf_587b9fceec62883e.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_587b9fceec62883e.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_587b9fceec62883e *JPEGImage) ToBase64WithQuality(__obf_f276064e93bacc8b int) (string, error) {
	if __obf_587b9fceec62883e.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_f276064e93bacc8b <= QualityNone && __obf_f276064e93bacc8b >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_587b9fceec62883e.Image, __obf_f276064e93bacc8b)
	}
	return util.EncodeJPEGToBase64(__obf_587b9fceec62883e.Image, QualityNone)
}
