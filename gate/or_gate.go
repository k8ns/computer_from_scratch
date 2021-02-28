package gate

import "github.com/k8ns/computer_from_scratch/wire"

type ORGate struct {
	A        *wire.Pin
	B        *wire.Pin
	O        *wire.Pin
}

func NewORGate() *ORGate {
	g := &ORGate{
		A:        &wire.Pin{},
		B:        &wire.Pin{},
		O:        &wire.Pin{},
	}

    notGateA := NewNOTGate()
	notGateB := NewNOTGate()
	nandGate := NewNANDGate()

	wire.Connect(notGateA.O, nandGate.A)
	wire.Connect(notGateB.O, nandGate.B)

	wire.Connect(g.A, notGateA.I)
	wire.Connect(g.B, notGateB.I)
	wire.Connect(nandGate.O, g.O)

	return g
}
