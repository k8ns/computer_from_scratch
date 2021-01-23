package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnabler_InitState(t *testing.T) {
	e := NewEnabler()

	for _, output := range e.Outputs() {
		assert.False(t, output.State)
	}

	e.E.Voltage(true)
	for _, output := range e.Outputs() {
		assert.False(t, output.State)
	}
}

func TestEnabler_CanOnAndOff(t *testing.T) {
	e := NewEnabler()

	for _, input := range e.Inputs() {
		input.Voltage(true)
	}

	e.E.Voltage(true)
	for _, output := range e.Outputs() {
		assert.True(t, output.State)
	}

	e.E.Voltage(false)
	for _, output := range e.Outputs() {
		assert.False(t, output.State)
	}
}
