package modbus

type Datum struct {
	Data any
}

type TransportType uint

const (
	TCP   TransportType = 0
	RTU   TransportType = 1
	ASCII TransportType = 2
)

type ReturnTypeCoil uint
type ReturnTypeRegister uint

const (
	COIL_DEFAULT ReturnTypeCoil = 0
	COIL_BIT     ReturnTypeCoil = 1
	COIL_STRING  ReturnTypeCoil = 2

	REGISTER_DEFAULT ReturnTypeRegister = 108
	REGISTER_INT16   ReturnTypeRegister = 16
	REGISTER_UINT16  ReturnTypeRegister = 116
	REGISTER_INT32   ReturnTypeRegister = 32
	REGISTER_UINT32  ReturnTypeRegister = 132
	REGISTER_INT64   ReturnTypeRegister = 64
	REGISTER_UINT64  ReturnTypeRegister = 164
	REGISTER_FLOAT32 ReturnTypeRegister = 232
	REGISTER_FLOAT64 ReturnTypeRegister = 264
)

type ONOFF uint16

const (
	ON  ONOFF = 0xff00
	OFF ONOFF = 0x0000
)
