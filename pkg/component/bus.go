package component

type BusInput struct {
	I1 *Pin
	I2 *Pin
	I3 *Pin
	I4 *Pin
	I5 *Pin
	I6 *Pin
	I7 *Pin
	I8 *Pin
}

func NewBusInput() *BusInput {
	return &BusInput{
		I1: &Pin{},
		I2: &Pin{},
		I3: &Pin{},
		I4: &Pin{},
		I5: &Pin{},
		I6: &Pin{},
		I7: &Pin{},
		I8: &Pin{},
	}
}

func (i *BusInput) Inputs() []*Pin {
	return []*Pin{i.I1, i.I2, i.I3, i.I4, i.I5, i.I6, i.I7, i.I8}
}

func (i *BusInput) Set(bits string) {
	for n, bit := range StringToByte(bits) {
		i.Inputs()[n].Voltage(bit)
	}
}

func (i *BusInput) Reset() {
	for _, input := range i.Inputs() {
		input.Voltage(false)
	}
}

type Bus struct {
	W1 *Wire
	W2 *Wire
	W3 *Wire
	W4 *Wire
	W5 *Wire
	W6 *Wire
	W7 *Wire
	W8 *Wire
}

func NewBus() *Bus {
	b := &Bus{
		W1: &Wire{},
		W2: &Wire{},
		W3: &Wire{},
		W4: &Wire{},
		W5: &Wire{},
		W6: &Wire{},
		W7: &Wire{},
		W8: &Wire{},
	}
	return b
}

func (b *Bus) Wires() []*Wire {
	return []*Wire{b.W1, b.W2, b.W3, b.W4, b.W5, b.W6, b.W7, b.W8}
}

func (b *Bus) Connect(pins ...*Pin) {
	for i, p := range pins {
		b.Wires()[i].Add(p)
	}
}

func (b *Bus) ConnectRegister(registers ...*Register) {
	for _, r := range registers {
		b.Connect(r.Inputs()...)
		b.Connect(r.Outputs()...)
	}
}

func (b *Bus) ConnectRegisterInputs(registers ...*Register) {
	for _, r := range registers {
		b.Connect(r.Inputs()...)
	}
}

func (b *Bus) Get() []bool {
	bits := make([]bool, 0)
	for _, w := range b.Wires() {
		bits = append(bits, w.Get())
	}
	return bits
}
