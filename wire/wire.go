package wire

type Wire struct {
	Pins []*Pin
}

func (w *Wire) Add(p *Pin) {
	if p.Wire != nil {
		w.Connect(p.Wire)
		return
	}
	w.Pins = append(w.Pins, p)
	p.Wire = w
}

func (w *Wire) Connect(w2 *Wire) {
	for _, p := range w2.Pins {
		w.Pins = append(w.Pins, p)
		p.Wire = w
	}
}

func (w *Wire) Update() {
	s := false
	for _, p := range w.Pins {
		if p.VoltageSupply {
			s = p.VoltageSupply
			break
		}
	}
	for _, p := range w.Pins {
		p.Update(s)
	}
}

func (w *Wire) Get() bool {
	for _, p := range w.Pins {
		return p.State
	}
	return false
}
