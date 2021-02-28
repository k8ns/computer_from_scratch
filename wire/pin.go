package wire

type Pin struct {
	VoltageSupply bool
	State         bool
	Wire          *Wire
	Component     Component
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
		if x, ok := i.Component.(interface{ Update()}); ok {
			x.Update()
		}
	}
}

