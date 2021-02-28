package gate

import "github.com/k8ns/computer_from_scratch/wire"

type NANDGate struct {
	A *wire.Pin
	B *wire.Pin
	O *wire.Pin
}

func NewNANDGate() *NANDGate {
	g := &NANDGate{
		A: &wire.Pin{},
		B: &wire.Pin{},
		O: &wire.Pin{},
	}

	g.A.Component = g
	g.B.Component = g

	g.Update()
	return g
}

func (g *NANDGate) Update() {
	if !g.A.State || !g.B.State {
		g.O.Voltage(true)
	} else {
		g.O.Voltage(false)
	}
}
