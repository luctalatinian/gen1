package debug

import (
	"luctalatinian/gen1/core"
	"luctalatinian/gen1/input"
	"luctalatinian/gen1/render"
	"luctalatinian/gen1/render/image"
)

// Instance is the singleton instance of Debugger.
var Instance *Debugger

func init() {
	Instance = &Debugger{
		sprite: image.NewImage(64, 64),
	}
	fill(Instance.sprite, 255)
}

// Debugger implements a state for testing image rendering.
type Debugger struct {
	sprite *image.Image
	x, y   int
}

// Update implements core.State::Update() for Debugger.
func (d *Debugger) Update(in input.Input, i interface{}) core.State {
	if in&input.InputUp != 0 {
		d.y--
	}
	if in&input.InputDown != 0 {
		d.y++
	}
	if in&input.InputLeft != 0 {
		d.x--
	}
	if in&input.InputRight != 0 {
		d.x++
	}

	return nil
}

// Render implements core.State::Render() for Debugger.
func (d *Debugger) Render(rctx render.Context, i interface{}) {
	frame := rctx.Frame()
	fill(frame, 127)

	image.Draw(d.sprite, frame, d.x, d.y)

	rctx.Draw()
}

func fill(im *image.Image, p byte) {
	for i := 0; i < len(im.Pixels); i++ {
		im.Pixels[i] = p
	}
}
