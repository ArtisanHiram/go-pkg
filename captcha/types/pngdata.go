package __obf_89363901d8df63bc

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"image"
	"image/png"
	"os"
	"path"
)

// PNGImage struct for PNG image data
type PNGImage struct {
	image.Image
}

// NewPNGImage creates a new PNG image data instance
func NewPNGImage(__obf_7b74ec68f3d93f19 image.Image) *PNGImage {
	return &PNGImage{__obf_7b74ec68f3d93f19}
}

// Get retrieves the original image
func (__obf_8f1cd35f635f319d *PNGImage) Get() image.Image {
	return __obf_8f1cd35f635f319d.Image
}

// SaveToFile is to save PNG as a file
func (__obf_8f1cd35f635f319d *PNGImage) SaveToFile(__obf_b3e288f298122efa string) error {
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

	return png.Encode(__obf_f49c961621401e38,

		// ToBytes converts the PNG image to a byte array
		__obf_8f1cd35f635f319d.Image)
}

func (__obf_8f1cd35f635f319d *PNGImage) ToBytes() ([]byte, error) {
	if __obf_8f1cd35f635f319d.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_8f1cd35f635f319d.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_8f1cd35f635f319d *PNGImage) ToBase64() (string, error) {
	if __obf_8f1cd35f635f319d.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_8f1cd35f635f319d.Image)
}
