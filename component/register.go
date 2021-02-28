package component

import "github.com/k8ns/computer_from_scratch/wire"

type Register struct {
	I1 *wire.Pin
	I2 *wire.Pin
	I3 *wire.Pin
	I4 *wire.Pin
	I5 *wire.Pin
	I6 *wire.Pin
	I7 *wire.Pin
	I8 *wire.Pin
	O1 *wire.Pin
	O2 *wire.Pin
	O3 *wire.Pin
	O4 *wire.Pin
	O5 *wire.Pin
	O6 *wire.Pin
	O7 *wire.Pin
	O8 *wire.Pin
	S  *wire.Pin
	E  *wire.Pin
}

func NewRegister() *Register {
	r := &Register{
		I1: &wire.Pin{},
		I2: &wire.Pin{},
		I3: &wire.Pin{},
		I4: &wire.Pin{},
		I5: &wire.Pin{},
		I6: &wire.Pin{},
		I7: &wire.Pin{},
		I8: &wire.Pin{},
		O1: &wire.Pin{},
		O2: &wire.Pin{},
		O3: &wire.Pin{},
		O4: &wire.Pin{},
		O5: &wire.Pin{},
		O6: &wire.Pin{},
		O7: &wire.Pin{},
		O8: &wire.Pin{},
		S:  &wire.Pin{},
		E:  &wire.Pin{},
	}

	b := NewByte()
	e := NewEnabler()

	wire.Connect(r.S, b.S)
	wire.Connect(r.E, e.E)
	for i := 0; i < 8; i++ {
		wire.Connect(b.Outputs()[i], e.Inputs()[i])
		wire.Connect(r.Inputs()[i], b.Inputs()[i])
		wire.Connect(e.Outputs()[i], r.Outputs()[i])
	}
	return r
}

func (r *Register) Inputs() []*wire.Pin {
	return []*wire.Pin{r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8}
}

func (r *Register) Outputs() []*wire.Pin {
	return []*wire.Pin{r.O1, r.O2, r.O3, r.O4, r.O5, r.O6, r.O7, r.O8}
}
