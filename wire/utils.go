package wire

func Connect(pins ...*Pin) {
	w := &Wire{}
	for _, p := range pins {
		w.Add(p)
	}
	w.Update()
}
