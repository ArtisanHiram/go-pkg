package __obf_154ce31cd9492d61

import (
	"errors"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
	MoveCaptchat
	DragCaptchat
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

type Captcha[T any] interface {
	GetPrimary() CaptchaImage
	GetSecondary() CaptchaImage
	GetData() T
	Type() CaptchaType
	Verify(__obf_3a3460dc570a7769 T, __obf_fb55e2058b66a9a8 int) bool
}

// RangeVal .
type Range struct {
	Min, Max int
}

// Size .
type Size struct {
	Width, Height int
}

// Point .
type Point struct {
	X, Y int
}

// AreaRect struct for defining a rectangular area
type AreaRect struct {
	MinX, MaxX, MinY, MaxY int
}

// MakeAreaRect creates an area rectangle
func MakeAreaRect(__obf_ebfbc480f77d6b4e, __obf_f8ada7bda495de21, __obf_0fbf2ddf76c8b701, __obf_2a9ea8d456f07c20 int) *AreaRect {
	return &AreaRect{
		MinX: __obf_ebfbc480f77d6b4e,
		MaxX: __obf_0fbf2ddf76c8b701,
		MinY: __obf_f8ada7bda495de21,
		MaxY: __obf_2a9ea8d456f07c20,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_f03bf0f36c5833f9, __obf_5c507cf47bf5267e int) *PositionRect {
	return &PositionRect{
		X:      __obf_c87d6c7ba4a30d82,
		Y:      __obf_f8fead3d682d659f,
		Height: __obf_f03bf0f36c5833f9,
		Width:  __obf_5c507cf47bf5267e,
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
	X int `json:"x"`
	Y int `json:"y"`
	// Display x,y
	DX     int         `json:"dx"`
	DY     int         `json:"dy"`
	Width  int         `json:"width" msgpack:"-"`
	Height int         `json:"height" msgpack:"-"`
	Angle  int         `json:"angle" msgpack:"-"`
	Shape  image.Image `json:"shape" msgpack:"-"`
}
