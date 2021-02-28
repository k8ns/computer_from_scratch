package gate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNandGate_InitState_OutIsOn(t *testing.T) {
	gate := NewNANDGate()
	assert.True(t, gate.O.State)
	assert.True(t, gate.O.VoltageSupply)
}

func TestNandGate_With_A_On_B_On_OutIsOff(t *testing.T) {
	gate := NewNANDGate()
	gate.A.Voltage(true)
	gate.B.Voltage(true)
	assert.False(t, gate.O.State)
	assert.False(t, gate.O.VoltageSupply)
}

func TestNandGate_A_Off_B_Off_OutIsOn(t *testing.T) {
	gate := NewNANDGate()
	gate.A.Voltage(false)
	gate.B.Voltage(false)
	assert.True(t, gate.O.State)
	assert.True(t, gate.O.VoltageSupply)
}

func TestNandGate_With_A_On_B_Off_OutIsOn(t *testing.T) {
	gate := NewNANDGate()
	gate.A.Voltage(true)
	gate.B.Voltage(false)
	assert.True(t, gate.O.State)
	assert.True(t, gate.O.VoltageSupply)
}

func TestNandGate_With_A_Off_B_On_OutIsOn(t *testing.T) {
	gate := NewNANDGate()
	gate.A.Voltage(false)
	gate.B.Voltage(true)
	assert.True(t, gate.O.State)
	assert.True(t, gate.O.VoltageSupply)
}
