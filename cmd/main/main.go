package main

import (
	"luctalatinian/gen1/core"
	"luctalatinian/gen1/core/debug"
	"luctalatinian/gen1/input"
	"luctalatinian/gen1/render/opengl"
)

func main() {
	rctx, err := opengl.NewRenderer()
	if err != nil {
		panic(err)
	}

	scanner := input.NewGLFWScanner(rctx.Window())

	var state core.State = debug.Instance
	for {
		input := scanner.Scan()
		next := state.Update(input, nil)
		state.Render(rctx, nil)

		if next != nil {
			state = next
		}
	}
}
