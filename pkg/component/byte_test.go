package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByte_InitState_AllBitsOff(t *testing.T) {
	b := NewByte()

	assert.False(t, b.O1.State)
	assert.False(t, b.O2.State)
	assert.False(t, b.O3.State)
	assert.False(t, b.O4.State)
	assert.False(t, b.O5.State)
	assert.False(t, b.O6.State)
	assert.False(t, b.O7.State)
	assert.False(t, b.O8.State)
}

func TestByte_CanSetAllBitsOn(t *testing.T) {
	b := NewByte()

	b.S.Voltage(true)
	b.I1.Voltage(true)
	b.I2.Voltage(true)
	b.I3.Voltage(true)
	b.I4.Voltage(true)
	b.I5.Voltage(true)
	b.I6.Voltage(true)
	b.I7.Voltage(true)
	b.I8.Voltage(true)
	b.S.Voltage(false)

	assert.True(t, b.O1.State)
	assert.True(t, b.O2.State)
	assert.True(t, b.O3.State)
	assert.True(t, b.O4.State)
	assert.True(t, b.O5.State)
	assert.True(t, b.O6.State)
	assert.True(t, b.O7.State)
	assert.True(t, b.O8.State)
}

func TestByte_CanSetRandomBits(t *testing.T) {
	b := NewByte()

	b.S.Voltage(true)

	b.I1.Voltage(true)
	b.I2.Voltage(false)
	b.I3.Voltage(false)
	b.I4.Voltage(true)
	b.I5.Voltage(true)
	b.I6.Voltage(false)
	b.I7.Voltage(true)
	b.I8.Voltage(true)

	b.S.Voltage(false)

	assert.True(t, b.O1.State)
	assert.False(t, b.O2.State)
	assert.False(t, b.O3.State)
	assert.True(t, b.O4.State)
	assert.True(t, b.O5.State)
	assert.False(t, b.O6.State)
	assert.True(t, b.O7.State)
	assert.True(t, b.O8.State)
}
