package gate

import "github.com/k8ns/computer_from_scratch/wire"

type XORGate struct {
	A        *wire.Pin
	B        *wire.Pin
	O        *wire.Pin
}

func NewXORGate() *ORGate {
	g := &ORGate{
		A:        &wire.Pin{},
		B:        &wire.Pin{},
		O:        &wire.Pin{},
	}

    gateA := NewNOTGate()
	gateB := NewNOTGate()
	gateC := NewNANDGate()
	gateD := NewNANDGate()
	gateE := NewNANDGate()

	wire.Connect(g.A, gateA.I, gateD.A)
	wire.Connect(g.B, gateB.I, gateC.B)

	wire.Connect(gateA.O, gateC.A)
	wire.Connect(gateB.O, gateD.B)

	wire.Connect(gateC.O, gateE.A)
	wire.Connect(gateD.O, gateE.B)

	wire.Connect(gateE.O, g.O)

	return g
}
