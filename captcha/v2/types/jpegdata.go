package __obf_deb7e65d40f46bd0

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/v2/util"
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
func NewJPEGImage(__obf_7e8297a5326972ed image.Image) *JPEGImage {
	return &JPEGImage{
		__obf_7e8297a5326972ed,
	}
}

// Get retrieves the original image
func (__obf_f012bb226e8e4a2d *JPEGImage) Get() image.Image {
	return __obf_f012bb226e8e4a2d.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_f012bb226e8e4a2d *JPEGImage) SaveToFile(__obf_68a9f2fdfea43cce string, __obf_53063ca7a14c52fe int) error {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_a148c9a511cdb474 *os.File
		__obf_cdc2a8664de7ba0a error
	)

	__obf_cdc2a8664de7ba0a = os.MkdirAll(path.Dir(__obf_68a9f2fdfea43cce), os.ModePerm)
	if __obf_cdc2a8664de7ba0a != nil {
		return __obf_cdc2a8664de7ba0a
	}

	if _, __obf_cdc2a8664de7ba0a = os.Stat(__obf_68a9f2fdfea43cce); os.IsNotExist(__obf_cdc2a8664de7ba0a) {
		__obf_a148c9a511cdb474, __obf_cdc2a8664de7ba0a = os.Create(__obf_68a9f2fdfea43cce)
	} else {
		__obf_a148c9a511cdb474, __obf_cdc2a8664de7ba0a = os.OpenFile(__obf_68a9f2fdfea43cce, os.O_RDWR, 0666)
	}
	if __obf_cdc2a8664de7ba0a != nil {
		return __obf_cdc2a8664de7ba0a
	}
	defer __obf_a148c9a511cdb474.Close()
	return jpeg.Encode(__obf_a148c9a511cdb474, __obf_f012bb226e8e4a2d.Image, &jpeg.Options{Quality: __obf_53063ca7a14c52fe})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_f012bb226e8e4a2d *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_f012bb226e8e4a2d.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_f012bb226e8e4a2d *JPEGImage) ToBytesWithQuality(__obf_6891796f5f655569 int) ([]byte, error) {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_6891796f5f655569 <= QualityNone && __obf_6891796f5f655569 >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_f012bb226e8e4a2d.Image, __obf_6891796f5f655569)
	}
	return util.EncodeJPEGToByte(__obf_f012bb226e8e4a2d.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_f012bb226e8e4a2d *JPEGImage) ToBase64() (string, error) {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_f012bb226e8e4a2d.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_f012bb226e8e4a2d *JPEGImage) ToBase64WithQuality(__obf_6891796f5f655569 int) (string, error) {
	if __obf_f012bb226e8e4a2d.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_6891796f5f655569 <= QualityNone && __obf_6891796f5f655569 >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_f012bb226e8e4a2d.Image, __obf_6891796f5f655569)
	}
	return util.EncodeJPEGToBase64(__obf_f012bb226e8e4a2d.Image, QualityNone)
}
