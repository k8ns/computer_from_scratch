package component

import (
	"github.com/k8ns/computer_from_scratch/gate"
	"github.com/k8ns/computer_from_scratch/wire"
)

type Decoder2x4 struct {
	I1 *wire.Pin
	I2 *wire.Pin
	O1 *wire.Pin
	O2 *wire.Pin
	O3 *wire.Pin
	O4 *wire.Pin
}

func NewDecoder2x4() *Decoder2x4 {
	d := &Decoder2x4{
		I1: &wire.Pin{},
		I2: &wire.Pin{},
		O1: &wire.Pin{},
		O2: &wire.Pin{},
		O3: &wire.Pin{},
		O4: &wire.Pin{},
	}

	not1 := gate.NewNOTGate()
	not2 := gate.NewNOTGate()
	and1 := gate.NewANDGate()
	and2 := gate.NewANDGate()
	and3 := gate.NewANDGate()
	and4 := gate.NewANDGate()

	wire.Connect(d.I1, not1.I, and3.A, and4.A)
	wire.Connect(d.I2, not2.I, and2.B, and4.B)
	wire.Connect(not1.O, and1.A, and2.A)
	wire.Connect(not2.O, and1.B, and3.B)
	wire.Connect(and1.O, d.O1)
	wire.Connect(and2.O, d.O2)
	wire.Connect(and3.O, d.O3)
	wire.Connect(and4.O, d.O4)

	return d
}

func (d *Decoder2x4) Outputs() []*wire.Pin {
	return []*wire.Pin{d.O1, d.O2, d.O3, d.O4}
}
