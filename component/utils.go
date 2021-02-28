package component

import "strings"

func ByteToString(bits []bool) string {
	bitsStr := make([]string, 0)
	for _, bit := range bits {
		bStr := "0"
		if bit {
			bStr = "1"
		}
		bitsStr = append(bitsStr, bStr)
	}

	return strings.Join(bitsStr, "")
}

func StringToByte(bitsStr string) []bool {
	bits := make([]bool, 0)
	for i, _ := range bitsStr {
		bit := true
		if bitsStr[i] == 48 { // 48 is ASCII code for 0
			bit = false
		}
		bits = append(bits, bit)
	}
	return bits
}
