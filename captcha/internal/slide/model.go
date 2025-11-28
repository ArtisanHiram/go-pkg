package __obf_7b43310cbd758abe

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	"image"
)

// Mode defines the slide CAPTCHA mode
type SlideType int

const (
	Move SlideType = iota // Move mode - slide to left
	Drag                  // Drag mode - drag to any direction
)

// DeadZoneDirectionType defines the dead zone direction
type DeadZoneDirectionType int

const (
	DeadZoneDirectionTypeLeft DeadZoneDirectionType = iota
	DeadZoneDirectionTypeRight
	DeadZoneDirectionTypeTop
	DeadZoneDirectionTypeBottom
)

// DrawImageParams defines the parameters for drawing the main image
type DrawImageParam struct {
	Width      int
	Height     int
	Background image.Image
	Alpha      float32
	DrawBlocks []*types.Block
}

// DrawTplImageParams defines the parameters for drawing the template image (tile)
type DrawTplImageParam struct {
	Width      int
	Height     int
	Background image.Image
	Mask       image.Image
	Alpha      float32
	Block      *types.Block
}

// Primary defines the main image configuration
type Primary struct {
	Size  types.Size // Image size
	Alpha float32    // Image alpha
}

// GraphConfig defines the graph configuration
type Secondary struct {
	CountRange           types.Range             // Number of graphs to generate
	SizeRange            types.Range             // Graph size range
	AnglePosRange        []types.Range           // Angle position range
	DeadZoneDirections   []DeadZoneDirectionType // Dead zone directions
	EnableVerticalRandom bool                    // Enable vertical random positioning
}

// Options defines the configuration options for the captcha
type Options struct {
	Primary
	Secondary
	Type SlideType
}

type Option func(*Options)
