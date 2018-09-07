package image

// All images are represented as RGBA.
const imgDepth = 4

// Image represents a 2D image.
type Image struct {
	Pixels []byte

	Stride int
	Width  int
	Height int
}

// NewImage constructs a blank image with width w and height h.
func NewImage(w, h int) *Image {
	return &Image{
		Pixels: make([]byte, imgDepth*w*h),

		Stride: imgDepth * w,
		Width:  w,
		Height: h,
	}
}

// Draw maps image src onto tgt at the position (x,y).
func Draw(src, tgt *Image, x, y int) {
	// don't bother to draw if src is out of bounds for tgt
	// -y > src.Height is technically redundant as that case is caught when calculating the column offset
	// (the copy loop won't execute at all) but we impose it to avoid extra multiplications
	if x > tgt.Width || y > tgt.Height || -x > src.Width || -y > src.Height {
		return
	}

	rows := min(src.Height, tgt.Height-y)
	cols := min(src.Stride, tgt.Stride-imgDepth*x)

	var colOffset, rowOffset int
	if x < 0 {
		colOffset = imgDepth * -x
		cols -= colOffset
	}
	if y < 0 {
		rowOffset = -y
	}

	spos := colOffset
	tpos := tgt.Stride*(y+rowOffset) + imgDepth*x + colOffset
	for i := 0; i < rows-rowOffset; i++ {
		copy(
			tgt.Pixels[tpos:tpos+cols],
			src.Pixels[spos:spos+cols],
		)
		tpos += tgt.Stride
		spos += src.Stride
	}
}

// DrawSubregion maps the subregion of image src starting at (x0, y0), with width w and height h,
// onto tgt at the position (x,y).
func DrawSubregion(src, tgt *Image, x0, y0, w, h, uint, x, y int) {
	if x > tgt.Width || y > tgt.Height || -x > w || -y > h {
		return
	}

	subStride := imgDepth * w
	rows := min(h, tgt.Height-y)
	cols := min(subStride, tgt.Stride-imgDepth*x)

	var colOffset, rowOffset int
	if x < 0 {
		colOffset = imgDepth * -x
		cols -= colOffset
	}
	if y < 0 {
		rowOffset = -y
	}

	spos := subStride*(y0+rowOffset) + imgDepth*x0 + colOffset
	tpos := tgt.Stride*(y+rowOffset) + imgDepth*x + colOffset
	for i := 0; i < rows-rowOffset; i++ {
		copy(
			tgt.Pixels[tpos:tpos+cols],
			src.Pixels[spos:spos+cols],
		)
		tpos += tgt.Stride
		spos += src.Stride
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
