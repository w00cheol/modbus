package modbus

func ParseToUint16Array(bytes []byte) []uint16 {
	data := make([]uint16, 0)

	for i := range bytes {
		if i%2 == 1 {
			firstInt := uint16(bytes[i-1])
			secondInt := uint16(bytes[i])

			data = append(data, firstInt*256+secondInt)
		}
	}

	return data
}

func ParseToBitArray(bytes []byte) []byte {
	data := make([]byte, 0)

	for _, b := range bytes {
		data = append(data, reverseByte(b)...)
	}

	return data
}

func ParseToString(bytes []byte) string {
	return string(bytes)
}

func reverseByte(b byte) []byte {
	bits := make([]uint8, 8)

	for i := range bits {
		bits[i] = b & 00000001
		b >>= 1
	}

	return bits
}
