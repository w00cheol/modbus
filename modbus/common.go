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
	COIL_DEFAULT  ReturnTypeCoil = 0
	COIL_DISCRETE ReturnTypeCoil = 1
	COIL_STRING   ReturnTypeCoil = 2

	REGISTER_DEFAULT  ReturnTypeRegister = 0
	REGISTER_DISCRETE ReturnTypeRegister = 1
)

type ONOFF uint16

const (
	ON  ONOFF = 0xff00
	OFF ONOFF = 0x0000
)
