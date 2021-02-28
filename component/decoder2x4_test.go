package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecoder2x4_InitState_Out1IsOn(t *testing.T) {
	d := NewDecoder2x4()
	assert.True(t, d.O1.State)
	assert.False(t, d.O2.State)
	assert.False(t, d.O3.State)
	assert.False(t, d.O4.State)
}

func TestDecoder2x4_Off_Off_Out1IsOn(t *testing.T) {
	d := NewDecoder2x4()
	d.I1.Voltage(false)
	d.I2.Voltage(false)

	assert.True(t, d.O1.State)
	assert.False(t, d.O2.State)
	assert.False(t, d.O3.State)
	assert.False(t, d.O4.State)
}

func TestDecoder2x4_Off_On_Out2IsOn(t *testing.T) {
	d := NewDecoder2x4()
	d.I1.Voltage(false)
	d.I2.Voltage(true)

	assert.False(t, d.O1.State)
	assert.True(t, d.O2.State)
	assert.False(t, d.O3.State)
	assert.False(t, d.O4.State)
}

func TestDecoder2x4_On_Off_Out3IsOn(t *testing.T) {
	d := NewDecoder2x4()
	d.I1.Voltage(true)
	d.I2.Voltage(false)

	assert.False(t, d.O1.State)
	assert.False(t, d.O2.State)
	assert.True(t, d.O3.State)
	assert.False(t, d.O4.State)
}

func TestDecoder2x4_On_On_Out4IsOn(t *testing.T) {
	d := NewDecoder2x4()
	d.I1.Voltage(true)
	d.I2.Voltage(true)

	assert.False(t, d.O1.State)
	assert.False(t, d.O2.State)
	assert.False(t, d.O3.State)
	assert.True(t, d.O4.State)
}
