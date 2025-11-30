package __obf_738b46210fdb4199

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
func NewJPEGImage(__obf_0912be1955a21fd1 image.Image) *JPEGImage {
	return &JPEGImage{__obf_0912be1955a21fd1}
}

// Get retrieves the original image
func (__obf_cf060e7b58cbd8c5 *JPEGImage) Get() image.Image {
	return __obf_cf060e7b58cbd8c5.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_cf060e7b58cbd8c5 *JPEGImage) SaveToFile(__obf_c2783f7dad1650bd string, __obf_36d89c55cc4e8880 int) error {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_70803e315fbd275e *os.File
		__obf_95d84a130cd584e5 error
	)
	__obf_95d84a130cd584e5 = os.MkdirAll(path.Dir(__obf_c2783f7dad1650bd), os.ModePerm)
	if __obf_95d84a130cd584e5 != nil {
		return __obf_95d84a130cd584e5
	}

	if _, __obf_95d84a130cd584e5 = os.Stat(__obf_c2783f7dad1650bd); os.IsNotExist(__obf_95d84a130cd584e5) {
		__obf_70803e315fbd275e, __obf_95d84a130cd584e5 = os.Create(__obf_c2783f7dad1650bd)
	} else {
		__obf_70803e315fbd275e, __obf_95d84a130cd584e5 = os.OpenFile(__obf_c2783f7dad1650bd, os.O_RDWR, 0666)
	}
	if __obf_95d84a130cd584e5 != nil {
		return __obf_95d84a130cd584e5
	}
	defer __obf_70803e315fbd275e.Close()
	return jpeg.Encode(__obf_70803e315fbd275e, __obf_cf060e7b58cbd8c5.Image, &jpeg.Options{Quality: __obf_36d89c55cc4e8880})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_cf060e7b58cbd8c5 *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_cf060e7b58cbd8c5.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_cf060e7b58cbd8c5 *JPEGImage) ToBytesWithQuality(__obf_aaad69feaf1cb657 int) ([]byte, error) {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_aaad69feaf1cb657 <= QualityNone && __obf_aaad69feaf1cb657 >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_cf060e7b58cbd8c5.Image, __obf_aaad69feaf1cb657)
	}
	return util.EncodeJPEGToByte(__obf_cf060e7b58cbd8c5.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_cf060e7b58cbd8c5 *JPEGImage) ToBase64() (string, error) {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_cf060e7b58cbd8c5.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_cf060e7b58cbd8c5 *JPEGImage) ToBase64WithQuality(__obf_aaad69feaf1cb657 int) (string, error) {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_aaad69feaf1cb657 <= QualityNone && __obf_aaad69feaf1cb657 >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_cf060e7b58cbd8c5.Image, __obf_aaad69feaf1cb657)
	}
	return util.EncodeJPEGToBase64(__obf_cf060e7b58cbd8c5.Image, QualityNone)
}
