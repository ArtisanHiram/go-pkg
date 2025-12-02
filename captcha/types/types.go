package __obf_89363901d8df63bc

import (
	"errors"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
)

var (
	ImageEmptyErr       = errors.New("image is empty")
	ImageMissingDataErr = errors.New("missing image data")
)

const (
	DistortNone = iota
	DistortLevel1
	DistortLevel2
	DistortLevel3
	DistortLevel4
	DistortLevel5
)

type CaptchaType int

const (
	ClickCaptchat CaptchaType = iota
	RotateCaptchat
	SlideCaptchat
)

const (
	QualityNone   = 100
	QualityLevel1 = 95
	QualityLevel2 = 85
	QualityLevel3 = 75
	QualityLevel4 = 65
	QualityLevel5 = 55
)

type CaptchaImage interface {
	Get() image.Image
	ToBytes() ([]byte, error)
	ToBase64() (string, error)
}

// Ensure JPEGImage and PNGImage implement CaptchaImage
var _ CaptchaImage = (*JPEGImage)(nil)
var _ CaptchaImage = (*PNGImage)(nil)

type CaptchaPayload struct {
	Type    CaptchaType
	Payload msgpack.RawMessage
}

func (__obf_8f1cd35f635f319d *CaptchaPayload) Verify(__obf_4a2d5e0db11864c5 any, __obf_35fc2be18bf12d23 int) bool {
	switch __obf_8f1cd35f635f319d.Type {
	case ClickCaptchat:
		// if len(clicks) != len(c.dots)*2 {
		// 	return false
		// }
		var __obf_2ab21b0572bef5b4 []*Dot
		__obf_2db4312e05b7f062 := // 配置 mapstructure
			&mapstructure.DecoderConfig{
				Result:           &__obf_2ab21b0572bef5b4,
				TagName:          "json",
				WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
			}
		__obf_28831695a0ffa90c, __obf_7b78cb8578d9f1c2 := mapstructure.NewDecoder(__obf_2db4312e05b7f062)
		if __obf_7b78cb8578d9f1c2 != nil {
			return false
		}

		if __obf_7b78cb8578d9f1c2 := __obf_28831695a0ffa90c.Decode(__obf_4a2d5e0db11864c5); __obf_7b78cb8578d9f1c2 != nil {
			return false
		}

		var __obf_ac1dfd7b84f6891c []*Dot
		if __obf_7b78cb8578d9f1c2 := msgpack.Unmarshal(__obf_8f1cd35f635f319d.Payload, &__obf_ac1dfd7b84f6891c); __obf_7b78cb8578d9f1c2 != nil {
			return false
		}

		if len(__obf_2ab21b0572bef5b4) != len(__obf_ac1dfd7b84f6891c) {
			return false
		}
		for __obf_1798a6750028be5f, __obf_b6881153086fd877 := range __obf_ac1dfd7b84f6891c {
			__obf_559cf55aa7e568ea := __obf_2ab21b0572bef5b4[__obf_1798a6750028be5f].X
			__obf_bd7620b16a5800e6 := __obf_2ab21b0572bef5b4[__obf_1798a6750028be5f].Y

			if __obf_559cf55aa7e568ea < __obf_b6881153086fd877.X-__obf_35fc2be18bf12d23 || __obf_559cf55aa7e568ea > __obf_b6881153086fd877.X+__obf_35fc2be18bf12d23 || __obf_bd7620b16a5800e6 < __obf_b6881153086fd877.Y-__obf_35fc2be18bf12d23 || __obf_bd7620b16a5800e6 > __obf_b6881153086fd877.Y+__obf_35fc2be18bf12d23 {
				return false
			}
		}
		return true
	case RotateCaptchat:
		// if angle, ok := input.(int); ok {
		// 	// Normalize angles to 0-360 range
		// 	totalAngle := (angle + c.angle) % 360
		// 	if totalAngle < 0 {
		// 		totalAngle += 360
		// 	}

		// 	// Check if total angle is close to 360 (or 0, which is equivalent)
		// 	minAngle := 360 - tolerance
		// 	maxAngle := tolerance

		// 	return totalAngle >= minAngle || totalAngle <= maxAngle
		// }
		// return false

		var __obf_1f7091190512fbb8 int
		// 使用类型选择来处理不同的输入类型
		switch __obf_9b2883182ebeb493 := __obf_4a2d5e0db11864c5.(type) {
		case int:
			__obf_1f7091190512fbb8 = __obf_9b2883182ebeb493
		case float64:
			__obf_1f7091190512fbb8 = // JSON 默认解析为 float64，这里将其转回 int
				int(__obf_9b2883182ebeb493)
		case int64:
			__obf_1f7091190512fbb8 = int(__obf_9b2883182ebeb493)
		case float32:
			__obf_1f7091190512fbb8 = int(__obf_9b2883182ebeb493)
		case string:
			// 如果前端传的是字符串 "90"，可能需要 strconv.Atoi 解析，视需求而定
			return false
		default:
			// 类型不匹配，直接验证失败
			return false
		}

		var __obf_5d98df60222101c3 int
		if __obf_7b78cb8578d9f1c2 := msgpack.Unmarshal(__obf_8f1cd35f635f319d.Payload, &__obf_5d98df60222101c3); __obf_7b78cb8578d9f1c2 != nil {
			return false
		}
		__obf_a349ab58f3b2fb7b := // Normalize angles to 0-360 range
			(__obf_1f7091190512fbb8 + __obf_5d98df60222101c3) % 360
		if __obf_a349ab58f3b2fb7b < 0 {
			__obf_a349ab58f3b2fb7b += 360
		}
		__obf_ac7981845901a6a8 := // Check if total angle is close to 360 (or 0, which is equivalent)
			360 - __obf_35fc2be18bf12d23
		__obf_4231cac591f2732d := __obf_35fc2be18bf12d23

		return __obf_a349ab58f3b2fb7b >= __obf_ac7981845901a6a8 || __obf_a349ab58f3b2fb7b <= __obf_4231cac591f2732d
	case SlideCaptchat:
		var __obf_df50034604593752 Point
		__obf_2db4312e05b7f062 := // 配置 mapstructure
			&mapstructure.DecoderConfig{
				Result:           &__obf_df50034604593752,
				TagName:          "json",
				WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
			}
		__obf_28831695a0ffa90c, __obf_7b78cb8578d9f1c2 := mapstructure.NewDecoder(__obf_2db4312e05b7f062)
		if __obf_7b78cb8578d9f1c2 != nil {
			return false
		}

		if __obf_7b78cb8578d9f1c2 := __obf_28831695a0ffa90c.Decode(__obf_4a2d5e0db11864c5); __obf_7b78cb8578d9f1c2 != nil {
			return false
		}

		var __obf_ed0636a5b9e4caee Point
		if __obf_7b78cb8578d9f1c2 := msgpack.Unmarshal(__obf_8f1cd35f635f319d.Payload, &__obf_ed0636a5b9e4caee); __obf_7b78cb8578d9f1c2 != nil {
			return false
		}
		__obf_ef02f5aef30dfa3c := __obf_ed0636a5b9e4caee.X - __obf_35fc2be18bf12d23
		__obf_014a5112e2c2a925 := __obf_ed0636a5b9e4caee.X + __obf_35fc2be18bf12d23
		__obf_081a97e9258e5ee5 := __obf_ed0636a5b9e4caee.Y - __obf_35fc2be18bf12d23
		__obf_f14c1ac944ab90a0 := __obf_ed0636a5b9e4caee.Y + __obf_35fc2be18bf12d23

		return __obf_df50034604593752.X >= __obf_ef02f5aef30dfa3c && __obf_df50034604593752.X <= __obf_014a5112e2c2a925 && __obf_df50034604593752.
			Y >= __obf_081a97e9258e5ee5 && __obf_df50034604593752.Y <= __obf_f14c1ac944ab90a0
	default:
		return false
	}
}

type CaptchaData interface {
	GetPrimary() CaptchaImage
	GetSecondary() CaptchaImage
	GetPayload() CaptchaPayload
}

type Captcha interface {
	Generate() (CaptchaData, error)
}

// RangeVal .
type Range struct {
	Min, Max int
}

// Point .
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// AreaRect struct for defining a rectangular area
type AreaRect struct {
	MinX, MaxX, MinY, MaxY int
}

// MakeAreaRect creates an area rectangle
func MakeAreaRect(__obf_081278319d4e8cbe, __obf_33020a883a9dc248, __obf_b3c292d03e4be02c, __obf_598f7b6bbc0cc3cd int) *AreaRect {
	return &AreaRect{
		MinX: __obf_081278319d4e8cbe,
		MaxX: __obf_b3c292d03e4be02c,
		MinY: __obf_33020a883a9dc248,
		MaxY: __obf_598f7b6bbc0cc3cd,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_64febf9a6792142f, __obf_672a1e7045b302e8 int) *PositionRect {
	return &PositionRect{
		X:      __obf_4d548ce157fe2d7b,
		Y:      __obf_e5fce044456d5b92,
		Height: __obf_64febf9a6792142f,
		Width:  __obf_672a1e7045b302e8,
	}
}

// DrawStringParams struct for string drawing parameters
type DrawStringParam struct {
	Color       color.Color
	FontSize    int
	Size        int
	FontDPI     int
	Text        string
	Font        *opentype.Font
	FontHinting font.Hinting
}

// Dot represents a single point (character or shape) in the captcha
type Dot struct {
	Index          int         `json:"index"`
	X              int         `json:"x"`
	Y              int         `json:"y"`
	Size           int         `json:"size"`
	Angle          int         `json:"angle" msgpack:"-"`
	Text           string      `json:"text" msgpack:"-"`
	Shape          image.Image `json:"shape" msgpack:"-"`
	FontSize       int         `json:"font-size" msgpack:"-"`
	PrimaryColor   string      `json:"primary-color" msgpack:"-"`
	SecondaryColor string      `json:"secondary-color" msgpack:"-"`
}

// Block defines the block data for the slide CAPTCHA
type Block struct {
	X      int         `json:"x"`
	Y      int         `json:"y"`
	Width  int         `json:"width" msgpack:"-"`
	Height int         `json:"height" msgpack:"-"`
	Angle  int         `json:"angle" msgpack:"-"`
	Shape  image.Image `json:"shape" msgpack:"-"`
}

// DeadZoneDirectionType defines the dead zone direction
type DeadZoneDirectionType int

const (
	DeadZoneDirectionTypeLeft DeadZoneDirectionType = iota
	DeadZoneDirectionTypeRight
	DeadZoneDirectionTypeTop
	DeadZoneDirectionTypeBottom
)

// Primary defines the main image configuration
type SlidePrimary struct {
	Width  int
	Height int
	Alpha  float32 // Image alpha
}

// GraphConfig defines the graph configuration
type SlideSecondary struct {
	CountRange           Range                   // Number of graphs to generate
	SizeRange            Range                   // Graph size range
	AnglePosRange        []Range                 // Angle position range
	DeadZoneDirections   []DeadZoneDirectionType // Dead zone directions
	EnableVerticalRandom bool                    // Enable vertical random positioning
}

// Options defines the configuration options for the captcha
type SlideOption struct {
	Primary   SlidePrimary
	Secondary SlideSecondary
}

// Primary defines the main image configuration
type RotatePrimary struct {
	Size          int     // Image size
	Alpha         float32 // Image alpha
	AnglePosRange []Range // Angle position range
}

// Secondary defines the thumbnail image configuration
type RotateSecondary struct {
	Alpha     float32 // Thumbnail alpha
	SizeRange []int   // Thumbnail size range
}

// Options defines the configuration options for the captcha
type RotateOption struct {
	Primary   RotatePrimary
	Secondary RotateSecondary
}

type ClickPrimary struct {
	Width         int
	Height        int
	LenRange      Range   // rangeLen
	AnglePosRange []Range // rangeAnglePos
	SizeRange     Range   // rangeSize
	Alpha         float32 // imageAlpha
	DotPadding    int
}

type ClickSecondary struct {
	Width          int
	Height         int
	VerifyLenRange Range // rangeVerifyLen
	SizeRange      Range // rangeThumbSize
	BgDistort      int   // thumbBgDistort
	BgCircles      int   // thumbBgCirclesNum
	BgSlimLines    int   // thumbBgSlimLineNum
	NonDeformAble  bool  // isThumbNonDeformAbility
	DisturbAlpha   float32
	DotPadding     int
}

type ClickOption struct {
	Primary     ClickPrimary
	Secondary   ClickSecondary
	FontDPI     int
	FontHinting font.Hinting
	ShowShadow  bool
	ShadowPoint Point
	UseRGBA     bool
}
