package main

import (
	"fmt"
	"log"
	"time"

	"main/modbus"
)

func main() {
	// Modbus RTU
	var address string = "xxx.xxx.xxx.xxx:502"
	var timeout time.Duration = 10 * time.Second
	var slaveId byte = 1

	client, err := modbus.NewClient(modbus.TCP, address, timeout, slaveId)
	if err != nil {
		log.Fatal(err)
	}

	results, err := client.WriteSingleRegister(3, 511)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("results: %+v\n", results)

	results, err = client.ReadHoldingRegisters(0, 10)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("results: %+v\n", results)
}
