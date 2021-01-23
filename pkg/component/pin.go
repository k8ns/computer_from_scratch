package component

import "fmt"

type Pin struct {
	VoltageSupply bool
	State         bool
	Wire          *Wire
	Component     Component
}

func (i *Pin) String() string {
	s := 0
	if i.State {
		s = 1
	}
	v := 0
	if i.VoltageSupply {
		v = 1
	}

	if i.VoltageSupply {
		return fmt.Sprintf("%d%d", s, v)
	}

	return fmt.Sprintf("%d", s)
}

func (i *Pin) Voltage(s bool) {
	i.VoltageSupply = s
	i.Update(s)
}

func (i *Pin) Update(s bool) {
	if i.State == s {
		return
	}

	i.State = s
	if i.Wire != nil {
		i.Wire.Update()
	}

	if i.Component != nil {
		i.Component.Update()
	}
}
