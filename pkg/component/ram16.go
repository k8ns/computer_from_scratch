package component

// Random Access Memory 16 bytes
type RAM struct {
	SA *Pin
	S  *Pin
	E  *Pin
}

func NewRam16(bus *Bus) *RAM {
	ram := &RAM{
		SA: &Pin{},
		S:  &Pin{},
		E:  &Pin{},
	}

	addressRegister := NewRegister()
	bus.ConnectRegisterInputs(addressRegister)
	Connect(ram.SA, addressRegister.S)
	addressRegister.E.Voltage(true)

	d1 := NewDecoder2x4()
	d2 := NewDecoder2x4()

	Connect(addressRegister.O1, d1.I1)
	Connect(addressRegister.O2, d1.I2)
	Connect(addressRegister.O3, d2.I1)
	Connect(addressRegister.O4, d2.I2)

	for _, dOut1 := range d1.Outputs() {
		for _, dOut2 := range d2.Outputs() {
			addressGate := NewANDGate()
			Connect(dOut1, addressGate.A)
			Connect(dOut2, addressGate.B)

			enableGate := NewANDGate()
			setGate := NewANDGate()
			Connect(addressGate.O, enableGate.A, setGate.A)
			Connect(ram.E, enableGate.B)
			Connect(ram.S, setGate.B)

			register := NewRegister()
			Connect(enableGate.O, register.E)
			Connect(setGate.O, register.S)

			bus.ConnectRegister(register)
		}
	}

	return ram
}
