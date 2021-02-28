package component

import (
	"github.com/k8ns/computer_from_scratch/gate"
	"github.com/k8ns/computer_from_scratch/wire"
)

type Enabler struct {
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
	E  *wire.Pin
}

func NewEnabler() *Enabler {
	e := &Enabler{
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
		E:  &wire.Pin{},
	}

	inputs := e.Inputs()
	outputs := e.Outputs()

	for i := 0; i < 8; i++ {
		andGate := gate.NewANDGate()
		wire.Connect(e.E, andGate.B)
		wire.Connect(inputs[i], andGate.A)
		wire.Connect(andGate.O, outputs[i])
	}

	return e
}

func (e *Enabler) Inputs() []*wire.Pin {
	return []*wire.Pin{e.I1, e.I2, e.I3, e.I4, e.I5, e.I6, e.I7, e.I8}
}

func (e *Enabler) Outputs() []*wire.Pin {
	return []*wire.Pin{e.O1, e.O2, e.O3, e.O4, e.O5, e.O6, e.O7, e.O8}
}
