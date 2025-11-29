package __obf_bda21a78cb74016a

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
func NewPNGImage(__obf_65fb97e6a605e786 image.Image) *PNGImage {
	return &PNGImage{__obf_65fb97e6a605e786}
}

// Get retrieves the original image
func (__obf_6ac27ae0c7f24e19 *PNGImage) Get() image.Image {
	return __obf_6ac27ae0c7f24e19.Image
}

// SaveToFile is to save PNG as a file
func (__obf_6ac27ae0c7f24e19 *PNGImage) SaveToFile(__obf_0494682892f936e3 string) error {
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

	return png.Encode(__obf_a20d63a4458733fc,

		// ToBytes converts the PNG image to a byte array
		__obf_6ac27ae0c7f24e19.Image)
}

func (__obf_6ac27ae0c7f24e19 *PNGImage) ToBytes() ([]byte, error) {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_6ac27ae0c7f24e19.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_6ac27ae0c7f24e19 *PNGImage) ToBase64() (string, error) {
	if __obf_6ac27ae0c7f24e19.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_6ac27ae0c7f24e19.Image)
}
