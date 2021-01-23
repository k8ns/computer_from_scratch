package component

type Register struct {
	I1 *Pin
	I2 *Pin
	I3 *Pin
	I4 *Pin
	I5 *Pin
	I6 *Pin
	I7 *Pin
	I8 *Pin
	O1 *Pin
	O2 *Pin
	O3 *Pin
	O4 *Pin
	O5 *Pin
	O6 *Pin
	O7 *Pin
	O8 *Pin
	S  *Pin
	E  *Pin
}

func NewRegister() *Register {
	r := &Register{
		I1: &Pin{},
		I2: &Pin{},
		I3: &Pin{},
		I4: &Pin{},
		I5: &Pin{},
		I6: &Pin{},
		I7: &Pin{},
		I8: &Pin{},
		O1: &Pin{},
		O2: &Pin{},
		O3: &Pin{},
		O4: &Pin{},
		O5: &Pin{},
		O6: &Pin{},
		O7: &Pin{},
		O8: &Pin{},
		S:  &Pin{},
		E:  &Pin{},
	}

	b := NewByte()
	e := NewEnabler()

	Connect(r.S, b.S)
	Connect(r.E, e.E)
	for i := 0; i < 8; i++ {
		Connect(b.Outputs()[i], e.Inputs()[i])
		Connect(r.Inputs()[i], b.Inputs()[i])
		Connect(e.Outputs()[i], r.Outputs()[i])
	}
	return r
}

func (r *Register) Inputs() []*Pin {
	return []*Pin{r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8}
}

func (r *Register) Outputs() []*Pin {
	return []*Pin{r.O1, r.O2, r.O3, r.O4, r.O5, r.O6, r.O7, r.O8}
}
