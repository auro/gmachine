// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// OpNOOP No Operation
const OpNOOP = 0

// OpHALT Stops the Machine
const OpHALT = 1

// MachineStruct is a struct type defined wich represents Program Counter (P) and Memory allocation slots
type MachineStruct struct {
	P      uint64
	Memory []uint64
}

// initialize m with MachineStruct type ad default values
var m = MachineStruct{}

// New initializes m struct with predefined values
func New() *MachineStruct {
	memSlice := make([]uint64, 1024)
	for i := range memSlice {
		memSlice[i] = uint64(0)
	}
	m.P = 0
	m.Memory = memSlice
	return &m
}

// Run must loop and return next memory slot pointer
func Run() *MachineStruct {
	fetchNext := m.P
	m.P = m.Memory[fetchNext]
	switch fetchNext {
	case OpHALT:

	case OpNOOP:

	default:
	}
	return &m
}
