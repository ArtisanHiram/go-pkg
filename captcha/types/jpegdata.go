package __obf_154ce31cd9492d61

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
func NewJPEGImage(__obf_7e8dc6da215ff393 image.Image) *JPEGImage {
	return &JPEGImage{
		__obf_7e8dc6da215ff393,
	}
}

// Get retrieves the original image
func (__obf_ad3a8dc916d9ec04 *JPEGImage) Get() image.Image {
	return __obf_ad3a8dc916d9ec04.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_ad3a8dc916d9ec04 *JPEGImage) SaveToFile(__obf_2287c49fb91d4310 string, __obf_6782c488bc09bbb0 int) error {
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
	return jpeg.Encode(__obf_78396bf8ad30e393, __obf_ad3a8dc916d9ec04.Image, &jpeg.Options{Quality: __obf_6782c488bc09bbb0})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_ad3a8dc916d9ec04 *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_ad3a8dc916d9ec04.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_ad3a8dc916d9ec04 *JPEGImage) ToBytesWithQuality(__obf_585ca1d090652320 int) ([]byte, error) {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_585ca1d090652320 <= QualityNone && __obf_585ca1d090652320 >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_ad3a8dc916d9ec04.Image, __obf_585ca1d090652320)
	}
	return util.EncodeJPEGToByte(__obf_ad3a8dc916d9ec04.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_ad3a8dc916d9ec04 *JPEGImage) ToBase64() (string, error) {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_ad3a8dc916d9ec04.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_ad3a8dc916d9ec04 *JPEGImage) ToBase64WithQuality(__obf_585ca1d090652320 int) (string, error) {
	if __obf_ad3a8dc916d9ec04.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_585ca1d090652320 <= QualityNone && __obf_585ca1d090652320 >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_ad3a8dc916d9ec04.Image, __obf_585ca1d090652320)
	}
	return util.EncodeJPEGToBase64(__obf_ad3a8dc916d9ec04.Image, QualityNone)
}
