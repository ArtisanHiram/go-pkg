package __obf_40cc6e54bc08fb7a

import (
	core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"
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

// Block defines the block data for the slide CAPTCHA
type Block struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Angle  int `json:"angle"`
	// Display x,y
	DX    int         `json:"dx"`
	DY    int         `json:"dy"`
	Shape image.Image `json:"-" msgpack:"-"`
}

// DrawImageParams defines the parameters for drawing the main image
type DrawImageParam struct {
	Width      int
	Height     int
	Background image.Image
	Alpha      float32
	DrawBlocks []*Block
}

// DrawTplImageParams defines the parameters for drawing the template image (tile)
type DrawTplImageParam struct {
	Width      int
	Height     int
	Background image.Image
	Mask       image.Image
	Alpha      float32
	Block      *Block
}

// Primary defines the main image configuration
type Primary struct {
	Size  core.Size // Image size
	Alpha float32   // Image alpha
}

// GraphConfig defines the graph configuration
type Secondary struct {
	CountRange           core.Range              // Number of graphs to generate
	SizeRange            core.Range              // Graph size range
	AnglePosRange        []core.Range            // Angle position range
	DeadZoneDirections   []DeadZoneDirectionType // Dead zone directions
	EnableVerticalRandom bool                    // Enable vertical random positioning
}

// Options defines the configuration options for the captcha
type Options struct {
	Primary
	Secondary
}

type Option func(*Options)
