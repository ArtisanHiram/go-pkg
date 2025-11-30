package __obf_54406b1a1de84196

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
func NewJPEGImage(__obf_10d21d6285312279 image.Image) *JPEGImage {
	return &JPEGImage{__obf_10d21d6285312279}
}

// Get retrieves the original image
func (__obf_a26ee860b3cee35b *JPEGImage) Get() image.Image {
	return __obf_a26ee860b3cee35b.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_a26ee860b3cee35b *JPEGImage) SaveToFile(__obf_0a97a8569823a471 string, __obf_d1e65134900a349e int) error {
	if __obf_a26ee860b3cee35b.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_4279250641db8149 *os.File
		__obf_d99742a1139768da error
	)
	__obf_d99742a1139768da = os.MkdirAll(path.Dir(__obf_0a97a8569823a471), os.ModePerm)
	if __obf_d99742a1139768da != nil {
		return __obf_d99742a1139768da
	}

	if _, __obf_d99742a1139768da = os.Stat(__obf_0a97a8569823a471); os.IsNotExist(__obf_d99742a1139768da) {
		__obf_4279250641db8149, __obf_d99742a1139768da = os.Create(__obf_0a97a8569823a471)
	} else {
		__obf_4279250641db8149, __obf_d99742a1139768da = os.OpenFile(__obf_0a97a8569823a471, os.O_RDWR, 0666)
	}
	if __obf_d99742a1139768da != nil {
		return __obf_d99742a1139768da
	}
	defer __obf_4279250641db8149.Close()
	return jpeg.Encode(__obf_4279250641db8149, __obf_a26ee860b3cee35b.Image, &jpeg.Options{Quality: __obf_d1e65134900a349e})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_a26ee860b3cee35b *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_a26ee860b3cee35b.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_a26ee860b3cee35b.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_a26ee860b3cee35b *JPEGImage) ToBytesWithQuality(__obf_4e494bdce6819a7e int) ([]byte, error) {
	if __obf_a26ee860b3cee35b.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_4e494bdce6819a7e <= QualityNone && __obf_4e494bdce6819a7e >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_a26ee860b3cee35b.Image, __obf_4e494bdce6819a7e)
	}
	return util.EncodeJPEGToByte(__obf_a26ee860b3cee35b.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_a26ee860b3cee35b *JPEGImage) ToBase64() (string, error) {
	if __obf_a26ee860b3cee35b.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_a26ee860b3cee35b.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_a26ee860b3cee35b *JPEGImage) ToBase64WithQuality(__obf_4e494bdce6819a7e int) (string, error) {
	if __obf_a26ee860b3cee35b.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_4e494bdce6819a7e <= QualityNone && __obf_4e494bdce6819a7e >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_a26ee860b3cee35b.Image, __obf_4e494bdce6819a7e)
	}
	return util.EncodeJPEGToBase64(__obf_a26ee860b3cee35b.Image, QualityNone)
}
