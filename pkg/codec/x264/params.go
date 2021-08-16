package x264

import (
	"github.com/trevor403/mediadevices/pkg/codec"
	"github.com/trevor403/mediadevices/pkg/io/video"
	"github.com/trevor403/mediadevices/pkg/prop"
)

// Params stores libx264 specific encoding parameters.
type Params struct {
	codec.BaseParams

	// Faster preset has lower CPU usage but lower quality
	Preset  Preset
	Profile Profile
}

// Preset represents a set of default configurations from libx264
type Preset int

// Profile represents a set of default configurations from libx264
type Profile int

const (
	PresetUltrafast Preset = iota
	PresetSuperfast
	PresetVeryfast
	PresetFaster
	PresetFast
	PresetMedium
	PresetSlow
	PresetSlower
	PresetVeryslow
	PresetPlacebo
)

const (
	ProfileBaseline Profile = iota
	ProfileMain
	ProfileHigh
	ProfileHigh10
	ProfileHigh422
	ProfileHigh444
)

// NewParams returns default x264 codec specific parameters.
func NewParams() (Params, error) {
	return Params{
		BaseParams: codec.BaseParams{
			KeyFrameInterval: 60,
		},
	}, nil
}

// RTPCodec represents the codec metadata
func (p *Params) RTPCodec() *codec.RTPCodec {
	return codec.NewRTPH264Codec(90000)
}

// BuildVideoEncoder builds x264 encoder with given params
func (p *Params) BuildVideoEncoder(r video.Reader, property prop.Media) (codec.ReadCloser, error) {
	return newEncoder(r, property, *p)
}
