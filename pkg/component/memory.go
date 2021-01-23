package component

type Memory struct {
	I  *Pin
	S  *Pin
	O  *Pin
	g1 *NANDGate
	g2 *NANDGate
	g3 *NANDGate
	g4 *NANDGate
}

func NewMemory() *Memory {
	m := &Memory{
		I:  &Pin{},
		S:  &Pin{},
		O:  &Pin{},
		g1: NewNANDGate(),
		g2: NewNANDGate(),
		g3: NewNANDGate(),
		g4: NewNANDGate(),
	}

	Connect(m.g1.O, m.g2.A, m.g3.A)
	Connect(m.g2.O, m.g4.B)
	Connect(m.g4.O, m.g3.B)
	Connect(m.g3.O, m.g4.A)

	Connect(m.I, m.g1.A)
	Connect(m.S, m.g1.B, m.g2.B)
	Connect(m.g3.O, m.O)
	return m
}
