package modbus

import (
	"errors"
	"time"

	"github.com/goburrow/modbus"
)

type TransportType uint

const (
	TCP   TransportType = 0
	RTU   TransportType = 1
	ASCII TransportType = 2
)

type Client struct {
	modbus.Client
}

func NewTCPClientHandler(address string, to time.Duration, slaveId byte) *modbus.TCPClientHandler {
	tcpHandler := modbus.NewTCPClientHandler(address)
	tcpHandler.Timeout = to
	tcpHandler.SlaveId = slaveId
	NewLogger().Debug("Init TCP Handler")

	return tcpHandler
}

func NewRTUClientHandler(address string, to time.Duration, slaveId byte) *modbus.RTUClientHandler {
	rtuHandler := modbus.NewRTUClientHandler(address)
	rtuHandler.Timeout = to
	rtuHandler.SlaveId = slaveId
	NewLogger().Debug("Init RTU Handler")

	return rtuHandler
}

func NewASCIIClientHandler(address string, to time.Duration, slaveId byte) *modbus.ASCIIClientHandler {
	asciiHandler := modbus.NewASCIIClientHandler(address)
	asciiHandler.Timeout = to
	asciiHandler.SlaveId = slaveId
	NewLogger().Debug("Init ASCII Handler")

	return asciiHandler
}

func NewClient(tt TransportType, address string, to time.Duration, slaveId byte) (*Client, error) {
	var handler modbus.ClientHandler
	// define TrasportType
	if tt == TCP {
		tcpHandler := NewTCPClientHandler(address, to, slaveId)
		if err := tcpHandler.Connect(); err != nil {
			return nil, err
		}
		handler = tcpHandler
		NewLogger().Debug("Handler TCP connect")

	} else if tt == RTU {
		rtuHandler := NewRTUClientHandler(address, to, slaveId)
		if err := rtuHandler.Connect(); err != nil {
			return nil, err
		}
		handler = rtuHandler
		NewLogger().Debug("Handler RTU connect")

	} else if tt == ASCII {
		asciiHandler := NewASCIIClientHandler(address, to, slaveId)
		if err := asciiHandler.Connect(); err != nil {
			return nil, err
		}
		handler = asciiHandler
		NewLogger().Debug("Handler ASCII connect")

	} else {
		NewLogger().Error(errors.New("Invalid TransportType"))
		return nil, errors.New("Invalid TransportType")
	}

	return &Client{
		Client: modbus.NewClient(handler),
	}, nil
}

func ParseToUint16Array(bytes []byte) ([]uint16, error) {
	data := make([]uint16, 0)

	for i := range bytes {
		if i%2 == 1 {
			firstInt := uint16(bytes[i-1])
			secondInt := uint16(bytes[i])

			data = append(data, firstInt*256+secondInt)
		}
	}

	return data, nil
}

func ParseToBitArray(bytes []byte) ([]byte, error) {
	data := make([]byte, 0)

	for _, b := range bytes {
		data = append(data, reverseByte(b)...)
	}

	return data, nil
}

func reverseByte(b byte) []byte {
	bits := make([]uint8, 8)

	for i := range bits {
		bits[i] = b & 00000001
		b >>= 1
	}

	return bits
}
