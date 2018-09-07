package opengl

import (
	"luctalatinian/gen1/render/img"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const screenWidth = 160
const screenHeight = 144

func init() {
	runtime.LockOSThread()
}

// Renderer implements a Context using GLFW and OpenGL.
type Renderer struct {
	window  *glfw.Window
	texture uint32

	screen *img.Image
}

// NewRenderer initializes the GLFW and OpenGL environments.
func NewRenderer() (*Renderer, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	window, err := glfw.CreateWindow(screenWidth, screenHeight, "Testing", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, err
	}

	r := &Renderer{
		window: window,

		screen: img.NewImage(screenWidth, screenHeight),
	}
	r.initGL()
	return r, nil
}

// Window returns the underlying glfw window.
func (r *Renderer) Window() *glfw.Window {
	return r.window
}

// Frame returns a pointer to the framebuffer.
func (r *Renderer) Frame() *img.Image {
	return r.screen
}

// Draw maps the frame into the context created via GLFW.
func (r *Renderer) Draw() {
	r.drawTexture()

	r.window.SwapBuffers()
	glfw.PollEvents()
}

func (r *Renderer) drawTexture() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.PixelStorei(gl.UNPACK_ROW_LENGTH, int32(r.screen.Width))
	gl.TexSubImage2D(
		gl.TEXTURE_2D,
		0,
		0,
		0,
		160,
		144,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		unsafe.Pointer(&r.screen.Pixels[0]),
	)
	gl.Begin(gl.TRIANGLE_STRIP)
	gl.TexCoord2f(0.0, 0.0)
	gl.Vertex3i(0, 0, 0)
	gl.TexCoord2f(160.0/256, 0.0) // TODO use r.screen.Width/Height
	gl.Vertex3i(160, 0, 0)
	gl.TexCoord2f(0.0, 144.0/256)
	gl.Vertex3i(0, 144, 0)
	gl.TexCoord2f(160.0/256, 144.0/256)
	gl.Vertex3i(160, 144, 0)
	gl.End()
}

// Destroy tears down the OpenGL context.
func (r *Renderer) Destroy() {
	glfw.Terminate()
}

func (r *Renderer) initGL() {
	gl.PushAttrib(gl.ENABLE_BIT)
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)
	gl.Enable(gl.TEXTURE_2D)

	gl.Viewport(0, 0, int32(r.screen.Width), int32(r.screen.Height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(
		1.0,
		float64(r.screen.Width-1),
		float64(r.screen.Height-1),
		1.0,
		0.0,
		1.0,
	)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.GenTextures(1, &r.texture)
	gl.BindTexture(gl.TEXTURE_2D, r.texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		256, // nearest power of 2 that fits our screen size
		256,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		nil,
	)
}
