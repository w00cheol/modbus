package modbus

import (
	"errors"
	"time"

	"github.com/goburrow/modbus"
)

func NewTCPClientHandler(address string, to time.Duration, slaveId byte) *modbus.TCPClientHandler {
	tcpHandler := modbus.NewTCPClientHandler(address)
	tcpHandler.Timeout = to
	tcpHandler.SlaveId = slaveId
	DEBUG.Println("Init TCP Handler")

	return tcpHandler
}

func NewRTUClientHandler(address string, to time.Duration, slaveId byte) *modbus.RTUClientHandler {
	rtuHandler := modbus.NewRTUClientHandler(address)
	rtuHandler.Timeout = to
	rtuHandler.SlaveId = slaveId
	DEBUG.Println("Init RTU Handler")

	return rtuHandler
}

func NewASCIIClientHandler(address string, to time.Duration, slaveId byte) *modbus.ASCIIClientHandler {
	asciiHandler := modbus.NewASCIIClientHandler(address)
	asciiHandler.Timeout = to
	asciiHandler.SlaveId = slaveId
	DEBUG.Println("Init ASCII Handler")

	return asciiHandler
}

func NewClient(tt TransportType, address string, to time.Duration, slaveId byte, rtCoil ReturnTypeCoil, rtRegister ReturnTypeRegister) (*Client, error) {
	var handler modbus.ClientHandler
	// define TrasportType
	if tt == TCP {
		tcpHandler := NewTCPClientHandler(address, to, slaveId)
		if err := tcpHandler.Connect(); err != nil {
			return nil, err
		}
		handler = tcpHandler
		DEBUG.Println("Handler TCP connect")

	} else if tt == RTU {
		rtuHandler := NewRTUClientHandler(address, to, slaveId)
		if err := rtuHandler.Connect(); err != nil {
			return nil, err
		}
		handler = rtuHandler
		DEBUG.Println("Handler RTU connect")

	} else if tt == ASCII {
		asciiHandler := NewASCIIClientHandler(address, to, slaveId)
		if err := asciiHandler.Connect(); err != nil {
			return nil, err
		}
		handler = asciiHandler
		DEBUG.Println("Handler ASCII connect")

	} else {
		ERROR.Println("Invalid TransportType")
		return nil, errors.New("Invalid TransportType")
	}

	return &Client{
		client:             modbus.NewClient(handler),
		ReturnTypeCoil:     rtCoil,
		ReturnTypeRegister: rtRegister,
	}, nil
}
