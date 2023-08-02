package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"main/modbus"
)

func main() {
	// Set loggers
	modbus.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	modbus.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	// Modbus RTU
	var address string = "x.x.x.x:502"
	var timeout time.Duration = 3 * time.Second
	var slaveId byte = 1
	// Get client
	client, err := modbus.NewClient(modbus.TCP, address, timeout, slaveId, modbus.COIL_BIT, modbus.REGISTER_UINT16)
	if err != nil {
		log.Fatal(err)
	}

	// Call funtions
	results, err := client.WriteSingleRegister(5, 1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("write: %+v\n", results)

	results, err = client.ReadHoldingRegisters(0, 32)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("read: %+v\n", results)
}
