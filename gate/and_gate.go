package gate

import (
	"github.com/k8ns/computer_from_scratch/wire"
)

type ANDGate struct {
	A        *wire.Pin
	B        *wire.Pin
	O        *wire.Pin
}

func NewANDGate() *ANDGate {
	g := &ANDGate{
		A:        &wire.Pin{},
		B:        &wire.Pin{},
		O:        &wire.Pin{},
	}

	nandGate := NewNANDGate()
	notGate := NewNOTGate()

	wire.Connect(nandGate.O, notGate.I)

	wire.Connect(g.A, nandGate.A)
	wire.Connect(g.B, nandGate.B)
	wire.Connect(notGate.O, g.O)

	return g
}
