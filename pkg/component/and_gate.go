package component

type ANDGate struct {
	A        *Pin
	B        *Pin
	O        *Pin
	nandGate *NANDGate
	notGate  *NOTGate
}

func NewANDGate() *ANDGate {
	gate := &ANDGate{
		A:        &Pin{},
		B:        &Pin{},
		O:        &Pin{},
		nandGate: NewNANDGate(),
		notGate:  NewNOTGate(),
	}

	//gate.A.Component = gate
	//gate.B.Component = gate

	Connect(gate.nandGate.O, gate.notGate.I)

	Connect(gate.A, gate.nandGate.A)
	Connect(gate.B, gate.nandGate.B)
	Connect(gate.notGate.O, gate.O)

	//gate.Update()
	return gate
}

func (g *ANDGate) Update() {
	//.nandGate.Update()
	//g.notGate.Update()
}
