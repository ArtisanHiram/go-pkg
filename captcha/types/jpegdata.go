package __obf_bda21a78cb74016a

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
func NewJPEGImage(__obf_65fb97e6a605e786 image.Image) *JPEGImage {
	return &JPEGImage{__obf_65fb97e6a605e786}
}

// Get retrieves the original image
func (__obf_6ac27ae0c7f24e19 *JPEGImage) Get() image.Image {
	return __obf_6ac27ae0c7f24e19.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_6ac27ae0c7f24e19 *JPEGImage) SaveToFile(__obf_0494682892f936e3 string, __obf_485194ed9df1c4bb int) error {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_a20d63a4458733fc *os.File
		__obf_eb4c912f039240f3 error
	)
	__obf_eb4c912f039240f3 = os.MkdirAll(path.Dir(__obf_0494682892f936e3), os.ModePerm)
	if __obf_eb4c912f039240f3 != nil {
		return __obf_eb4c912f039240f3
	}

	if _, __obf_eb4c912f039240f3 = os.Stat(__obf_0494682892f936e3); os.IsNotExist(__obf_eb4c912f039240f3) {
		__obf_a20d63a4458733fc, __obf_eb4c912f039240f3 = os.Create(__obf_0494682892f936e3)
	} else {
		__obf_a20d63a4458733fc, __obf_eb4c912f039240f3 = os.OpenFile(__obf_0494682892f936e3, os.O_RDWR, 0666)
	}
	if __obf_eb4c912f039240f3 != nil {
		return __obf_eb4c912f039240f3
	}
	defer __obf_a20d63a4458733fc.Close()
	return jpeg.Encode(__obf_a20d63a4458733fc, __obf_6ac27ae0c7f24e19.Image, &jpeg.Options{Quality: __obf_485194ed9df1c4bb})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_6ac27ae0c7f24e19 *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_6ac27ae0c7f24e19.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_6ac27ae0c7f24e19 *JPEGImage) ToBytesWithQuality(__obf_8969d16176aa07af int) ([]byte, error) {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_8969d16176aa07af <= QualityNone && __obf_8969d16176aa07af >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_6ac27ae0c7f24e19.Image, __obf_8969d16176aa07af)
	}
	return util.EncodeJPEGToByte(__obf_6ac27ae0c7f24e19.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_6ac27ae0c7f24e19 *JPEGImage) ToBase64() (string, error) {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_6ac27ae0c7f24e19.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_6ac27ae0c7f24e19 *JPEGImage) ToBase64WithQuality(__obf_8969d16176aa07af int) (string, error) {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_8969d16176aa07af <= QualityNone && __obf_8969d16176aa07af >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_6ac27ae0c7f24e19.Image, __obf_8969d16176aa07af)
	}
	return util.EncodeJPEGToBase64(__obf_6ac27ae0c7f24e19.Image, QualityNone)
}
