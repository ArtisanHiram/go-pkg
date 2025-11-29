package __obf_244ef50d49151021

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

type CaptchaData interface {
	GetPrimary() CaptchaImage
	GetSecondary() CaptchaImage
	GetData() any
	Type() CaptchaType
	Verify(__obf_6544c3ce66f054cc any, __obf_ca3c6f396acc09ca int) bool
}

type Captcha interface {
	Generate() (CaptchaData, error)
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
func MakeAreaRect(__obf_1e1cf7aec767392f, __obf_4135b4099e87ad95, __obf_953be905ac791795, __obf_8a6b79e05b444849 int) *AreaRect {
	return &AreaRect{
		MinX: __obf_1e1cf7aec767392f,
		MaxX: __obf_953be905ac791795,
		MinY: __obf_4135b4099e87ad95,
		MaxY: __obf_8a6b79e05b444849,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_e0d45be06f8e29d9, __obf_b0bdb05fb01e5fd8 int) *PositionRect {
	return &PositionRect{
		X:      __obf_e487b3618d36e102,
		Y:      __obf_8cbd87ef15403513,
		Height: __obf_e0d45be06f8e29d9,
		Width:  __obf_b0bdb05fb01e5fd8,
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
