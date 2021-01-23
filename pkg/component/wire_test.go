package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanWirePointToPoint(t *testing.T) {
	in := &Pin{}
	out := &Pin{}
	Connect(in, out)

	in.Voltage(true)
	assert.True(t, out.State)

	in.Voltage(false)
	assert.False(t, out.State)
}

func TestCanWireVFormWithTwoInsAndOneOut(t *testing.T) {
	in1 := &Pin{}
	in2 := &Pin{}
	out := &Pin{}

	Connect(in1, out)
	Connect(in2, out)

	in1.Voltage(true)
	assert.True(t, out.State)

	in2.Voltage(true)
	assert.True(t, out.State)

	in1.Voltage(false)
	assert.True(t, out.State)

	in2.Voltage(false)
	assert.False(t, out.State)
}

func TestCanWireMultipleInsAndOuts(t *testing.T) {
	in1 := &Pin{}
	in2 := &Pin{}
	in3 := &Pin{}
	out1 := &Pin{}
	out2 := &Pin{}
	out3 := &Pin{}
	Connect(in1, out1)
	Connect(in2, out2)
	Connect(in3, out3)
	Connect(in1, out2)
	Connect(in2, out3)

	outs := []*Pin{out1, out2, out3}

	in1.Voltage(true)
	for _, o := range outs {
		assert.True(t, o.State)
	}

	in1.Voltage(false)
	for _, o := range outs {
		assert.False(t, o.State)
	}

	in3.Voltage(true)
	for _, o := range outs {
		assert.True(t, o.State)
	}

	in2.Voltage(true)
	for _, o := range outs {
		assert.True(t, o.State)
	}

	in3.Voltage(false)
	for _, o := range outs {
		assert.True(t, o.State)
	}
}

func TestTwoComponentsCanBeWiredUp(t *testing.T) {

	gate1 := NewNANDGate()
	gate2 := NewNANDGate()

	Connect(gate1.O, gate2.A, gate2.B)

	gate1.A.Voltage(true)
	gate1.B.Voltage(true)

	assert.True(t, gate2.O.State)
}

type JoinComponent struct {
	I *Pin
	O *Pin
}

func NewJoinComponent() *JoinComponent {
	c := &JoinComponent{I: &Pin{}, O: &Pin{}}
	c.I.Component = c
	return c
}

func (g *JoinComponent) Update() {
	g.O.Voltage(g.I.State)
}

func TestMultipleComponentsCanBeWiredUp(t *testing.T) {
	c1 := NewJoinComponent()
	c1.I.Voltage(true)
	c2 := NewJoinComponent()
	c3 := NewJoinComponent()
	c4 := NewJoinComponent()
	c5 := NewJoinComponent()

	Connect(c1.O, c2.I)
	Connect(c2.O, c3.I)
	Connect(c2.O, c4.I)
	Connect(c4.O, c5.I)
	assert.True(t, c5.O.State)
	assert.True(t, c4.O.State)

}
