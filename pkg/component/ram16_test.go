package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRAM16_InitState(t *testing.T) {
	bus := NewBus()
	ram16 := NewRam16(bus)

	ram16.E.Voltage(true)
	assert.Equal(t, "00000000", ByteToString(bus.Get()))
}

func TestRAM16_CanStoreAtAddress(t *testing.T) {

	bus := NewBus()
	ram := NewRam16(bus)
	busInput := NewBusInput()
	bus.Connect(busInput.Inputs()...)

	var SetAddressRam = func(addr string) {
		busInput.Set(addr)
		ram.SA.Voltage(true)
		ram.SA.Voltage(false)
		busInput.Reset()
	}

	var StoreToRAM = func(addr, val string) {
		SetAddressRam(addr)

		busInput.Set(val)
		ram.S.Voltage(true)
		ram.S.Voltage(false)
		busInput.Reset()
	}

	var GetFromRAM = func(addr string) string {
		SetAddressRam(addr)
		ram.E.Voltage(true)

		defer ram.E.Voltage(false)
		return ByteToString(bus.Get())
	}

	addr1 := "0101____"
	addr2 := "0010____"
	addr3 := "1110____"

	val1 := "10101011"
	val2 := "10111101"
	val3 := "00100100"

	StoreToRAM(addr1, val1)
	StoreToRAM(addr2, val2)
	StoreToRAM(addr3, val3)

	for i := 0; i < 2; i++ {
		assert.Equal(t, val1, GetFromRAM(addr1))
		assert.Equal(t, val2, GetFromRAM(addr2))
		assert.Equal(t, val3, GetFromRAM(addr3))
		assert.Equal(t, "00000000", GetFromRAM("1111____"))
	}
}
