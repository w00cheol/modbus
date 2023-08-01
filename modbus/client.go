package modbus

import (
	"github.com/goburrow/modbus"
)

type Client struct {
	client             modbus.Client
	ReturnTypeCoil     ReturnTypeCoil
	ReturnTypeRegister ReturnTypeRegister
}

func (mb *Client) returnCoilByType(data []byte) (results *Datum) {
	switch mb.ReturnTypeCoil {
	case COIL_DISCRETE:
		return &Datum{Data: ParseToBitArray(data)}
	case COIL_STRING:
		return &Datum{Data: string(ParseToBitArray(data))}
	default:
		return &Datum{Data: data}
	}
}

func (mb *Client) returnRegisterByType(data []byte) (results *Datum) {
	switch mb.ReturnTypeRegister {
	case REGISTER_DISCRETE:
		return &Datum{Data: ParseToUint16Array(data)}
	default:
		return &Datum{Data: data}
	}
}

func (mb *Client) ReadCoils(address uint16, quantity uint16) (results *Datum, err error) {
	data, err := mb.client.ReadCoils(address, quantity)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnCoilByType(data), nil
}

func (mb *Client) WriteSingleCoil(address uint16, value uint16) (results *Datum, err error) {
	data, err := mb.client.WriteSingleCoil(address, value)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnCoilByType(data), nil
}

func (mb *Client) WriteMultipleCoils(address uint16, quantity uint16, value []byte) (results *Datum, err error) {
	data, err := mb.client.WriteMultipleCoils(address, quantity, value)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnCoilByType(data), nil
}

func (mb *Client) ReadDiscreteInputs(address uint16, quantity uint16) (results *Datum, err error) {
	data, err := mb.client.ReadDiscreteInputs(address, quantity)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnCoilByType(data), nil
}

func (mb *Client) ReadHoldingRegisters(address uint16, quantity uint16) (results *Datum, err error) {
	data, err := mb.client.ReadHoldingRegisters(address, quantity)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnRegisterByType(data), nil
}

func (mb *Client) WriteSingleRegister(address uint16, value uint16) (results *Datum, err error) {
	data, err := mb.client.WriteSingleRegister(address, value)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnRegisterByType(data), nil
}

func (mb *Client) WriteMultipleRegisters(address uint16, quantity uint16, value []byte) (results *Datum, err error) {
	data, err := mb.client.WriteMultipleRegisters(address, quantity, value)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnRegisterByType(data), nil
}

func (mb *Client) ReadInputRegisters(address uint16, quantity uint16) (results *Datum, err error) {
	data, err := mb.client.ReadInputRegisters(address, quantity)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnRegisterByType(data), nil
}
