package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemory_InitState_OutIsOff(t *testing.T) {
	m := NewMemory()
	assert.False(t, m.O.State)
}

func TestMemory_CanSetBitOn(t *testing.T) {
	m := NewMemory()
	m.S.Voltage(true)
	m.I.Voltage(true)
	m.S.Voltage(false)

	assert.True(t, m.O.State)
}

func TestMemory_CanSetBitOff(t *testing.T) {
	m := NewMemory()
	m.S.Voltage(true)
	m.I.Voltage(false)
	m.S.Voltage(false)

	assert.False(t, m.O.State)
}

func TestMemory_WillNotSetBitWhileSetIsOff(t *testing.T) {
	m := NewMemory()
	m.S.Voltage(false)
	m.I.Voltage(true)
	assert.False(t, m.O.State)

	m.S.Voltage(true)
	m.I.Voltage(true)
	m.S.Voltage(false)

	m.S.Voltage(false)
	m.I.Voltage(false)
	assert.True(t, m.O.State)
}

func TestMemory_CanSetBitOnOffOn(t *testing.T) {
	m := NewMemory()

	m.S.Voltage(true)
	m.I.Voltage(true)
	m.S.Voltage(false)
	assert.True(t, m.O.State)

	m.S.Voltage(true)
	m.I.Voltage(false)
	m.S.Voltage(false)
	assert.False(t, m.O.State)

	m.S.Voltage(true)
	m.I.Voltage(true)
	m.S.Voltage(false)
	assert.True(t, m.O.State)
}
