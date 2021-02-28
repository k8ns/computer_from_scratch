package gate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotGate_InitState_OutIsOn(t *testing.T) {
	gate := NewNOTGate()
	assert.True(t, gate.O.State)
}

func TestNotGate_With_In_Off_OutIsOn(t *testing.T) {
	gate := NewNOTGate()
	gate.I.Voltage(false)
	assert.True(t, gate.O.State)
}

func TestNotGate_With_In_On_OutIsOff(t *testing.T) {
	gate := NewNOTGate()
	gate.I.Voltage(true)
	assert.False(t, gate.O.State)
}
