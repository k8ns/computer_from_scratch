package component

import (
	"github.com/k8ns/computer_from_scratch/gate"
	"github.com/k8ns/computer_from_scratch/wire"
)

type Memory struct {
	I  *wire.Pin
	S  *wire.Pin
	O  *wire.Pin
}

func NewMemory() *Memory {
	m := &Memory{
		I:  &wire.Pin{},
		S:  &wire.Pin{},
		O:  &wire.Pin{},
	}

	g1 := gate.NewNANDGate()
	g2 := gate.NewNANDGate()
	g3 := gate.NewNANDGate()
	g4 := gate.NewNANDGate()

	wire.Connect(g1.O, g2.A, g3.A)
	wire.Connect(g2.O, g4.B)
	wire.Connect(g4.O, g3.B)
	wire.Connect(g3.O, g4.A)

	wire.Connect(m.I, g1.A)
	wire.Connect(m.S, g1.B, g2.B)
	wire.Connect(g3.O, m.O)
	return m
}
