package __obf_244ef50d49151021

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
func NewJPEGImage(__obf_4b8ec8035512b29f image.Image) *JPEGImage {
	return &JPEGImage{__obf_4b8ec8035512b29f}
}

// Get retrieves the original image
func (__obf_dbbef660ec878016 *JPEGImage) Get() image.Image {
	return __obf_dbbef660ec878016.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_dbbef660ec878016 *JPEGImage) SaveToFile(__obf_6f182e746b9b2f31 string, __obf_a4ac91c30367b179 int) error {
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
	return jpeg.Encode(__obf_1ad0cb034eca86b6, __obf_dbbef660ec878016.Image, &jpeg.Options{Quality: __obf_a4ac91c30367b179})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_dbbef660ec878016 *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_dbbef660ec878016.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_dbbef660ec878016.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_dbbef660ec878016 *JPEGImage) ToBytesWithQuality(__obf_d60be5ff4c7c9dd0 int) ([]byte, error) {
	if __obf_dbbef660ec878016.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_d60be5ff4c7c9dd0 <= QualityNone && __obf_d60be5ff4c7c9dd0 >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_dbbef660ec878016.Image, __obf_d60be5ff4c7c9dd0)
	}
	return util.EncodeJPEGToByte(__obf_dbbef660ec878016.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_dbbef660ec878016 *JPEGImage) ToBase64() (string, error) {
	if __obf_dbbef660ec878016.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_dbbef660ec878016.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_dbbef660ec878016 *JPEGImage) ToBase64WithQuality(__obf_d60be5ff4c7c9dd0 int) (string, error) {
	if __obf_dbbef660ec878016.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_d60be5ff4c7c9dd0 <= QualityNone && __obf_d60be5ff4c7c9dd0 >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_dbbef660ec878016.Image, __obf_d60be5ff4c7c9dd0)
	}
	return util.EncodeJPEGToBase64(__obf_dbbef660ec878016.Image, QualityNone)
}
