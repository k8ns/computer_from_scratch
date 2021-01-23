package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAndGate_InitState_OutIsOff(t *testing.T) {
	gate := NewANDGate()
	assert.False(t, gate.O.State)
}

func TestAndGate_With_A_On_B_On_OutIsOn(t *testing.T) {
	gate := NewANDGate()
	gate.A.Voltage(true)
	gate.B.Voltage(true)
	assert.True(t, gate.O.State)
}

func TestAndGate_A_Off_B_Off_OutIsOff(t *testing.T) {
	gate := NewANDGate()
	gate.A.Voltage(false)
	gate.B.Voltage(false)
	assert.False(t, gate.O.State)
}

func TestAndGate_With_A_On_B_Off_OutIsOff(t *testing.T) {
	gate := NewANDGate()
	gate.A.Voltage(true)
	gate.B.Voltage(false)
	assert.False(t, gate.O.State)
}

func TestAndGate_With_A_Off_B_On_OutIsOff(t *testing.T) {
	gate := NewANDGate()
	gate.A.Voltage(false)
	gate.B.Voltage(true)
	assert.False(t, gate.O.State)
}
