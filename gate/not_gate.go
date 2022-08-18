package gate

import "github.com/k8ns/computer_from_scratch/wire"

type NOTGate struct {
	I        *wire.Pin
	O        *wire.Pin
}

func NewNOTGate() *NOTGate {
	gate := &NOTGate{
		I:        &wire.Pin{},
		O:        &wire.Pin{},
	}

    nandGate := NewNANDGate()

	wire.Connect(gate.I, nandGate.A, nandGate.B)
	wire.Connect(nandGate.O, gate.O)
	return gate
}
