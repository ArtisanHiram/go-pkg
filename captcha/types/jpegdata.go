package __obf_89363901d8df63bc

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
func NewJPEGImage(__obf_7b74ec68f3d93f19 image.Image) *JPEGImage {
	return &JPEGImage{__obf_7b74ec68f3d93f19}
}

// Get retrieves the original image
func (__obf_8f1cd35f635f319d *JPEGImage) Get() image.Image {
	return __obf_8f1cd35f635f319d.Image
}

// SaveToFile saves the JPEG image to a file
func (__obf_8f1cd35f635f319d *JPEGImage) SaveToFile(__obf_b3e288f298122efa string, __obf_8e42bab5bbdfd280 int) error {
	if __obf_8f1cd35f635f319d.Image == nil {
		return ImageMissingDataErr
	}

	var (
		__obf_f49c961621401e38 *os.File
		__obf_7b78cb8578d9f1c2 error
	)
	__obf_7b78cb8578d9f1c2 = os.MkdirAll(path.Dir(__obf_b3e288f298122efa), os.ModePerm)
	if __obf_7b78cb8578d9f1c2 != nil {
		return __obf_7b78cb8578d9f1c2
	}

	if _, __obf_7b78cb8578d9f1c2 = os.Stat(__obf_b3e288f298122efa); os.IsNotExist(__obf_7b78cb8578d9f1c2) {
		__obf_f49c961621401e38, __obf_7b78cb8578d9f1c2 = os.Create(__obf_b3e288f298122efa)
	} else {
		__obf_f49c961621401e38, __obf_7b78cb8578d9f1c2 = os.OpenFile(__obf_b3e288f298122efa, os.O_RDWR, 0666)
	}
	if __obf_7b78cb8578d9f1c2 != nil {
		return __obf_7b78cb8578d9f1c2
	}
	defer __obf_f49c961621401e38.Close()
	return jpeg.Encode(__obf_f49c961621401e38, __obf_8f1cd35f635f319d.Image, &jpeg.Options{Quality: __obf_8e42bab5bbdfd280})

}

// ToBytes converts the JPEG image to a byte array
func (__obf_8f1cd35f635f319d *JPEGImage) ToBytes() ([]byte, error) {
	if __obf_8f1cd35f635f319d.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	return util.EncodeJPEGToByte(__obf_8f1cd35f635f319d.Image, QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (__obf_8f1cd35f635f319d *JPEGImage) ToBytesWithQuality(__obf_eff1b0c0f7f5c0bf int) ([]byte, error) {
	if __obf_8f1cd35f635f319d.Image == nil {
		return []byte{}, ImageEmptyErr
	}

	if __obf_eff1b0c0f7f5c0bf <= QualityNone && __obf_eff1b0c0f7f5c0bf >= QualityLevel5 {
		return util.EncodeJPEGToByte(__obf_8f1cd35f635f319d.Image, __obf_eff1b0c0f7f5c0bf)
	}
	return util.EncodeJPEGToByte(__obf_8f1cd35f635f319d.Image, QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (__obf_8f1cd35f635f319d *JPEGImage) ToBase64() (string, error) {
	if __obf_8f1cd35f635f319d.Image == nil {
		return "", ImageEmptyErr
	}

	return util.EncodeJPEGToBase64(__obf_8f1cd35f635f319d.Image, QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (__obf_8f1cd35f635f319d *JPEGImage) ToBase64WithQuality(__obf_eff1b0c0f7f5c0bf int) (string, error) {
	if __obf_8f1cd35f635f319d.Image == nil {
		return "", ImageEmptyErr
	}

	if __obf_eff1b0c0f7f5c0bf <= QualityNone && __obf_eff1b0c0f7f5c0bf >= QualityLevel5 {
		return util.EncodeJPEGToBase64(__obf_8f1cd35f635f319d.Image, __obf_eff1b0c0f7f5c0bf)
	}
	return util.EncodeJPEGToBase64(__obf_8f1cd35f635f319d.Image, QualityNone)
}
