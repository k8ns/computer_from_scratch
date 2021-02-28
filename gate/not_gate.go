package gate

import "github.com/k8ns/computer_from_scratch/wire"

type NOTGate struct {
	I        *wire.Pin
	O        *wire.Pin
	nandGate *NANDGate
}

func NewNOTGate() *NOTGate {
	gate := &NOTGate{
		I:        &wire.Pin{},
		O:        &wire.Pin{},
		nandGate: NewNANDGate(),
	}

	wire.Connect(gate.I, gate.nandGate.A, gate.nandGate.B)
	wire.Connect(gate.nandGate.O, gate.O)
	return gate
}
