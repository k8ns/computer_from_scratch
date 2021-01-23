package component

type NOTGate struct {
	I        *Pin
	O        *Pin
	nandGate *NANDGate
}

func NewNOTGate() *NOTGate {
	gate := &NOTGate{
		I:        &Pin{},
		O:        &Pin{},
		nandGate: NewNANDGate(),
	}

	//gate.I.Component = gate

	Connect(gate.I, gate.nandGate.A, gate.nandGate.B)
	Connect(gate.nandGate.O, gate.O)

	//	gate.Update()
	return gate
}

func (g *NOTGate) Update() {
	//g.nandGate.Update()
}
