package component

type ORGate struct {
	A        *Pin
	B        *Pin
	O        *Pin
	notGateA *NOTGate
	notGateB *NOTGate
	nandGate *NANDGate
}

func NewORGate() *ORGate {
	g := &ORGate{
		A:        &Pin{},
		B:        &Pin{},
		O:        &Pin{},
		notGateA: NewNOTGate(),
		notGateB: NewNOTGate(),
		nandGate: NewNANDGate(),
	}

	Connect(g.notGateA.O, g.nandGate.A)
	Connect(g.notGateB.O, g.nandGate.B)

	Connect(g.A, g.notGateA.I)
	Connect(g.B, g.notGateB.I)
	Connect(g.nandGate.O, g.O)

	return g
}

func (g *ORGate) Update() {
	//g.notGateA.Update()
	//g.notGateB.Update()
	//g.nandGate.Update()
}
