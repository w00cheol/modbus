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
	case COIL_BIT:
		return &Datum{Data: ToBitArray(data)}
	case COIL_STRING:
		return &Datum{Data: string(ToBitArray(data))}
	default:
		return &Datum{Data: data}
	}
}

func (mb *Client) returnRegisterByType(data []byte) (results *Datum) {
	switch mb.ReturnTypeRegister {
	case REGISTER_INT16:
		return &Datum{Data: ToInt16Array(data)}
	case REGISTER_UINT16:
		return &Datum{Data: ToUint16Array(data)}
	case REGISTER_INT32:
		return &Datum{Data: ToInt32Array(data)}
	case REGISTER_UINT32:
		return &Datum{Data: ToUint32Array(data)}
	case REGISTER_INT64:
		return &Datum{Data: ToInt64Array(data)}
	case REGISTER_UINT64:
		return &Datum{Data: ToUint64Array(data)}
	case REGISTER_FLOAT32:
		return &Datum{Data: ToFloat32Array(data)}
	case REGISTER_FLOAT64:
		return &Datum{Data: ToFloat64Array(data)}
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

	return &Datum{Data: data}, nil
}

func (mb *Client) WriteMultipleCoils(address uint16, quantity uint16, value []byte) (results *Datum, err error) {
	data, err := mb.client.WriteMultipleCoils(address, quantity, value)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return &Datum{Data: data}, nil
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

	return &Datum{Data: data}, nil
}

func (mb *Client) WriteMultipleRegisters(address uint16, quantity uint16, value []byte) (results *Datum, err error) {
	data, err := mb.client.WriteMultipleRegisters(address, quantity, value)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return &Datum{Data: data}, nil
}

func (mb *Client) ReadInputRegisters(address uint16, quantity uint16) (results *Datum, err error) {
	data, err := mb.client.ReadInputRegisters(address, quantity)
	if err != nil {
		NewLogger().Error(err)
		return nil, err
	}

	return mb.returnRegisterByType(data), nil
}
