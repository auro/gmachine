// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// MachineStruct is a struct type defined wich represents P register and memory allocation
type MachineStruct struct {
	P      uint64
	Memory []uint64
}

// New returns a struct with predefined values
func New() *MachineStruct {
	memSlice := make([]uint64, 1024)
	for i := range memSlice {
		memSlice[i] = uint64(0)
	}
	return &MachineStruct{0, memSlice}
}
