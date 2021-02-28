package component

import (
	"github.com/k8ns/computer_from_scratch/gate"
	"github.com/k8ns/computer_from_scratch/wire"
)

type Memory struct {
	I  *wire.Pin
	S  *wire.Pin
	O  *wire.Pin
	g1 *gate.NANDGate
	g2 *gate.NANDGate
	g3 *gate.NANDGate
	g4 *gate.NANDGate
}

func NewMemory() *Memory {
	m := &Memory{
		I:  &wire.Pin{},
		S:  &wire.Pin{},
		O:  &wire.Pin{},
		g1: gate.NewNANDGate(),
		g2: gate.NewNANDGate(),
		g3: gate.NewNANDGate(),
		g4: gate.NewNANDGate(),
	}

	wire.Connect(m.g1.O, m.g2.A, m.g3.A)
	wire.Connect(m.g2.O, m.g4.B)
	wire.Connect(m.g4.O, m.g3.B)
	wire.Connect(m.g3.O, m.g4.A)

	wire.Connect(m.I, m.g1.A)
	wire.Connect(m.S, m.g1.B, m.g2.B)
	wire.Connect(m.g3.O, m.O)
	return m
}
