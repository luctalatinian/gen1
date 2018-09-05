package input

// Input represents a bitfield of button states.
type Input uint8

// Enumeration of Input
const (
	InputStart  Input = 1
	InputSelect       = 2
	InputUp           = 4
	InputDown         = 8
	InputLeft         = 16
	InputRight        = 32
	InputA            = 64
	InputB            = 128
)

// Scanner defines the most basic interface for reading inputs.
type Scanner interface {
	Scan() Input
}
