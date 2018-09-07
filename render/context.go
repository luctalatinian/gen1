package render

import "luctalatinian/gen1/render/img"

// Context defines the most basic interface for rendering.
type Context interface {
	Frame() *img.Image
	Draw()
}
