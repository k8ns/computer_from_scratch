package gate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrGate_InitState_OutIsOff(t *testing.T) {
	gate := NewORGate()
	assert.False(t, gate.O.State)
}

func TestOrGate_With_A_Off_B_Off_OutIsOff(t *testing.T) {
	gate := NewORGate()
	gate.A.Voltage(false)
	gate.B.Voltage(false)
	assert.False(t, gate.O.State)
}

func TestOrGate_With_A_On_B_On_OutIsOn(t *testing.T) {
	gate := NewORGate()
	gate.A.Voltage(true)
	gate.B.Voltage(true)
	assert.True(t, gate.O.State)
}

func TestOrGate_With_A_On_B_Off_OutIsOn(t *testing.T) {
	gate := NewORGate()
	gate.A.Voltage(true)
	gate.B.Voltage(false)
	assert.True(t, gate.O.State)
}

func TestOrGate_With_A_Off_B_On_OutIsOn(t *testing.T) {
	gate := NewORGate()
	gate.A.Voltage(false)
	gate.B.Voltage(true)
	assert.True(t, gate.O.State)
}
