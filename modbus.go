package main

import (
	"fmt"
	"log"
	"time"

	"main/modbus"
)

func main() {
	// Modbus RTU
	var address string = "x.x.x.x:502"
	var timeout time.Duration = 3 * time.Second
	var slaveId byte = 1

	client, err := modbus.NewClient(modbus.TCP, address, timeout, slaveId, modbus.COIL_DISCRETE, modbus.REGISTER_DISCRETE)
	if err != nil {
		log.Fatal(err)
	}

	results, err := client.WriteSingleCoil(0, uint16(modbus.ON))
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("write: %+v\n", results)

	results, err = client.ReadCoils(0, 10)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("read: %+v\n", results)
}
