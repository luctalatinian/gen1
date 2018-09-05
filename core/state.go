package core

import (
	"luctalatinian/gen1/input"
	"luctalatinian/gen1/render"
)

// State represents a game state. In general, Update() and Render() should be called once per frame.
type State interface {
	Update(input.Input, interface{}) State
	Render(render.Context, interface{})
}
