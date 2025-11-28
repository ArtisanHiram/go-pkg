package __obf_b99521ec2c18257a

import core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"

type Captcha interface {
	GetPrimary() *core.JPEGImage
	GetSecondary() *core.PNGImage
	GetData() any
	Verify(__obf_1d1a492bc2eec0f9 any, __obf_aa2c635eda2639bb int) bool
}
