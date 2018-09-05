package input

import "github.com/go-gl/glfw/v3.2/glfw"

// GLFWScanner implements the Scanner interface using a GLFW window.
type GLFWScanner struct {
	window *glfw.Window
}

// NewGLFWScanner creates a scanner from a glfw window.
func NewGLFWScanner(w *glfw.Window) Scanner {
	return &GLFWScanner{
		window: w,
	}
}

// Scan implements Scanner::Scan() for GLFWScanner.
func (g *GLFWScanner) Scan() Input {
	var Inputs Input
	if g.window.GetKey(glfw.KeySpace) == glfw.Press {
		Inputs |= InputStart
	}
	if g.window.GetKey(glfw.KeyEnter) == glfw.Press {
		Inputs |= InputSelect
	}
	if g.window.GetKey(glfw.KeyW) == glfw.Press {
		Inputs |= InputUp
	}
	if g.window.GetKey(glfw.KeyS) == glfw.Press {
		Inputs |= InputDown
	}
	if g.window.GetKey(glfw.KeyA) == glfw.Press {
		Inputs |= InputLeft
	}
	if g.window.GetKey(glfw.KeyD) == glfw.Press {
		Inputs |= InputRight
	}
	if g.window.GetKey(glfw.KeyZ) == glfw.Press {
		Inputs |= InputA
	}
	if g.window.GetKey(glfw.KeyX) == glfw.Press {
		Inputs |= InputB
	}
	return Inputs
}
