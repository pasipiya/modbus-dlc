package services

import (
	"time"
	T "modbus-dlc/types"
	// "github.com/gofiber/fiber/v2"
	// "github.com/volatiletech/sqlboiler/v4/boil"
	//"github.com/grid-x/modbus"
	"modbus-dlc/connections"
	"fmt"
	"log"
)

func GetDlc() (string, *T.ServiceError) {

		slaveIDs := []byte{1, 2} 
	
		for _, slaveID := range slaveIDs {

			log.Printf("Device ID: %v", slaveID)

			registerAddress := uint16(131)
			registerLength := uint16(1)
			valueToWrite := uint16(620)

			err := modbusconnection.WriteRegister(slaveID, registerAddress, valueToWrite)
			if err != nil {
				log.Printf("Error writing register: %v", err)
			}

			time.Sleep(300 * time.Millisecond)

			results, err := modbusconnection.ReadRegister(slaveID, registerAddress, registerLength)
			if err != nil {
				log.Printf("Error reading register: %v", err)
			} else {
				fmt.Printf("Register value: %v\n", results)
				value := uint16(results[0])<<8 | uint16(results[1])  // Combine high and low bytes
				fmt.Printf("Register Value at address %d: %d\n", registerAddress, value)
			}

			// err := modbusservice.WriteRegister(1, 131, 6500)
			// if err != nil {
			// 	log.Printf("Error writing register: %v", err)
			// }

			// handler.SetSlave(slaveID)

			// registerAddress := uint16(131)
			// registerLength := uint16(1)
			// valueToWrite := uint16(640)

			// fmt.Printf("Device ID : %v\n",slaveID)
	
			// _, err = client.WriteSingleRegister(registerAddress, valueToWrite)
			// if err != nil {
			// 	fmt.Printf("Failed to write to holding register: %v\n", err)
			// }
			// time.Sleep(300 * time.Millisecond)
	
			// results, err := client.ReadHoldingRegisters(registerAddress, registerLength)
	
			// // results, err := client.ReadHoldingRegisters(0, 10)
			// if err != nil {
			// 	log.Fatalf("Failed to read registers: %v", err)
			// }

			// time.Sleep(100 * time.Millisecond)
		
			// value := uint16(results[0])<<8 | uint16(results[1])  // Combine high and low bytes
	
			// fmt.Printf("Register Value at address %d: %d\n", registerAddress, value)
		}

	return "dlc starting", nil
}