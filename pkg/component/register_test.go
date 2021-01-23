package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister_InitState(t *testing.T) {
	r := NewRegister()

	for _, output := range r.Outputs() {
		assert.False(t, output.State)
	}

	r.E.Voltage(true)
	for _, output := range r.Outputs() {
		assert.False(t, output.State)
	}
}

func TestRegister_CanSetRememberEnableByteOfInfo(t *testing.T) {
	r := NewRegister()

	for _, input := range r.Inputs() {
		input.Voltage(true)
	}
	r.S.Voltage(true)
	r.S.Voltage(false)
	for _, input := range r.Inputs() {
		input.Voltage(false)
	}

	for _, output := range r.Outputs() {
		assert.False(t, output.State)
	}

	r.E.Voltage(true)
	for _, output := range r.Outputs() {
		assert.True(t, output.State)
	}
}
