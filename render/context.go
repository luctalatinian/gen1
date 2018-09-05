package render

import "luctalatinian/gen1/render/image"

// Context defines the most basic interface for rendering.
type Context interface {
	Frame() *image.Image
	Draw()
}
