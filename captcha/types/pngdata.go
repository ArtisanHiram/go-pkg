package __obf_738b46210fdb4199

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
func NewPNGImage(__obf_0912be1955a21fd1 image.Image) *PNGImage {
	return &PNGImage{__obf_0912be1955a21fd1}
}

// Get retrieves the original image
func (__obf_cf060e7b58cbd8c5 *PNGImage) Get() image.Image {
	return __obf_cf060e7b58cbd8c5.Image
}

// SaveToFile is to save PNG as a file
func (__obf_cf060e7b58cbd8c5 *PNGImage) SaveToFile(__obf_c2783f7dad1650bd string) error {
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

	return png.Encode(__obf_70803e315fbd275e,

		// ToBytes converts the PNG image to a byte array
		__obf_cf060e7b58cbd8c5.Image)
}

func (__obf_cf060e7b58cbd8c5 *PNGImage) ToBytes() ([]byte, error) {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return []byte{}, ImageEmptyErr
	}
	return util.EncodePNGToByte(__obf_cf060e7b58cbd8c5.Image)
}

// ToBase64 converts the PNG image to a Base64 string
func (__obf_cf060e7b58cbd8c5 *PNGImage) ToBase64() (string, error) {
	if __obf_cf060e7b58cbd8c5.Image == nil {
		return "", ImageEmptyErr
	}
	return util.EncodePNGToBase64(__obf_cf060e7b58cbd8c5.Image)
}
