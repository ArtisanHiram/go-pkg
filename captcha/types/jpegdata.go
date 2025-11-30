package __obf_6dcb1d06bd949756

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
func NewJPEGImage(__obf_587078e38cdb4111 image.Image) *JPEGImage {
	return &JPEGImage{__obf_587078e38cdb4111}
}

// Get retrieves the original image
func (__obf_6b14c339de9bea34 *JPEGImage) Get() image.Image {
	return __obf_6b14c339de9bea34.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_6b14c339de9bea34 *JPEGImage) SaveToFile(__obf_b62b8c1a0e618495 string, __obf_0baf5d2b0b6ed705 int) error {
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
	return jpeg.Encode(__obf_6e11aba9569a8248, __obf_6b14c339de9bea34.Image, &jpeg.Options{Quality: __obf_0baf5d2b0b6ed705})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_6b14c339de9bea34 *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_6b14c339de9bea34.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_6b14c339de9bea34.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_6b14c339de9bea34 *JPEGImage) ToBytesWithQuality(__obf_b51052a6e71a1227 int) ([]byte, error) {
	if __obf_6b14c339de9bea34.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_b51052a6e71a1227 <= QualityNone && __obf_b51052a6e71a1227 >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_6b14c339de9bea34.Image, __obf_b51052a6e71a1227)
	}
	return util.EncodeJPEGToByte(__obf_6b14c339de9bea34.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_6b14c339de9bea34 *JPEGImage) ToBase64() (string, error) {
	if __obf_6b14c339de9bea34.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_6b14c339de9bea34.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_6b14c339de9bea34 *JPEGImage) ToBase64WithQuality(__obf_b51052a6e71a1227 int) (string, error) {
	if __obf_6b14c339de9bea34.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_b51052a6e71a1227 <= QualityNone && __obf_b51052a6e71a1227 >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_6b14c339de9bea34.Image, __obf_b51052a6e71a1227)
	}
	return util.EncodeJPEGToBase64(__obf_6b14c339de9bea34.Image, QualityNone)
}
