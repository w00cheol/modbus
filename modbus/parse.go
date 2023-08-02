package modbus

import (
	"encoding/binary"
	"fmt"
	"math"
)

func ToInt16Array(bytes []byte) []int16 {
	if len(bytes)%2 != 0 {
		ERROR.Println("cannot be typed (want: 2*N byte)")
		return nil
	}

	data := make([]int16, len(bytes)/2)

	for i := range data {
		data[i] = int16(binary.BigEndian.Uint16(bytes[(2 * i) : 2*(i+1)]))
	}

	return data
}

func ToUint16Array(bytes []byte) []uint16 {
	if len(bytes)%2 != 0 {
		ERROR.Println("cannot be typed (want: 2*N byte)")
		return nil
	}

	data := make([]uint16, len(bytes)/2)

	for i := range data {
		data[i] = binary.BigEndian.Uint16(bytes[(2 * i) : 2*(i+1)])
	}

	return data
}

func ToInt32Array(bytes []byte) []int32 {
	if len(bytes)%4 != 0 {
		ERROR.Println("cannot be typed (want: 4*N byte)")
		return nil
	}

	data := make([]int32, len(bytes)/4)

	for i := range data {
		data[i] = int32(binary.BigEndian.Uint32(bytes[(4 * i) : 4*(i+1)]))
	}

	return data
}

func ToUint32Array(bytes []byte) []uint32 {
	if len(bytes)%4 != 0 {
		ERROR.Println("cannot be typed (want: 4*N byte)")
		return nil
	}

	data := make([]uint32, len(bytes)/4)

	for i := range data {
		data[i] = binary.BigEndian.Uint32(bytes[(4 * i) : 4*(i+1)])
	}

	return data
}

func ToInt64Array(bytes []byte) []int64 {
	if len(bytes)%8 != 0 {
		ERROR.Println("cannot be typed (want: 8*N byte)")
		return nil
	}

	data := make([]int64, len(bytes)/8)

	for i := range data {
		data[i] = int64(binary.BigEndian.Uint64(bytes[(8 * i) : 8*(i+1)]))
	}

	return data
}

func ToUint64Array(bytes []byte) []uint64 {
	if len(bytes)%8 != 0 {
		ERROR.Println("cannot be typed (want: 8*N byte)")
		return nil
	}

	data := make([]uint64, len(bytes)/8)

	for i := range data {
		data[i] = binary.BigEndian.Uint64(bytes[(8 * i) : 8*(i+1)])
	}

	return data
}

func ToFloat32Array(bytes []byte) []float32 {
	if len(bytes)%4 != 0 {
		ERROR.Println("cannot be typed (want: 4*N byte)")
		return nil
	}

	data := make([]float32, len(bytes)/4)

	for i := range data {
		data[i] = math.Float32frombits(binary.BigEndian.Uint32(bytes[(4 * i) : 4*(i+1)]))
	}

	return data
}

func ToFloat64Array(bytes []byte) []float64 {
	if len(bytes)%8 != 0 {
		ERROR.Println("cannot be typed (want: 8*N byte)")
		return nil
	}

	data := make([]float64, len(bytes)/8)

	for i := range data {
		data[i] = math.Float64frombits(binary.BigEndian.Uint64(bytes[(8 * i) : 8*(i+1)]))
		fmt.Printf("data: %v\n", data)
	}

	return data
}

func ToBitArray(bytes []byte) []byte {
	data := make([]byte, 8*len(bytes))

	for _, b := range data {
		data = append(data, reverseByte(b)...)
	}

	return data
}

func ToString(bytes []byte) string {
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
