package lib

func ByteIsDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func IntFromByte(b byte) int {
	return int(b) - 48
}

func ByteIsPeriod(b byte) bool {
	return b == 46
}

func ByteIsGear(b byte) bool {
	return b == 42
}

func CharToByte(char string) byte {
	return []byte(char)[0]
}

// Probably don't need this
func ByteToChar(b byte) string {
	return string(b)
}
