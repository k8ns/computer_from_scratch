package component

type NANDGate struct {
	A *Pin
	B *Pin
	O *Pin
}

func NewNANDGate() *NANDGate {
	gate := &NANDGate{
		A: &Pin{},
		B: &Pin{},
		O: &Pin{},
	}

	gate.A.Component = gate
	gate.B.Component = gate

	gate.Update()
	return gate
}

func (g *NANDGate) Update() {
	if !g.A.State || !g.B.State {
		g.O.Voltage(true)
	} else {
		g.O.Voltage(false)
	}
}
