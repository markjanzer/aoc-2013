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
