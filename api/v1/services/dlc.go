package services

import (
	// "context"
	// "database/sql"
	"time"

	// M "github.com/atharvbhadange/go-api-template/models"
	T "github.com/atharvbhadange/go-api-template/types"
	// "github.com/gofiber/fiber/v2"
	// "github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/grid-x/modbus"
	"connections/modbusservice"
	"fmt"
	"log"
)

// type ProductBody struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Price       int    `json:"price"`
// }

func GetDlc() (string, *T.ServiceError) {
	// products, err := M.Products().All(ctx, dbTrx)

	// if err != nil {
	// 	return nil, &T.ServiceError{
	// 		Message: "Unable to get products",
	// 		Error:   err,
	// 		Code:    fiber.StatusInternalServerError,
	// 	}
	// }

	// return &T.ServiceError{
	// 	Message: "Unable to get product",
	// 	Error:   err,
	// 	Code:    fiber.StatusInternalServerError,
	// }
		// Connect to Modbus RTU device
		// handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
		// handler.BaudRate = 19200
		// handler.DataBits = 8
		// handler.StopBits = 1
		// handler.Parity = "N"
		// // handler.SlaveID = 1
		// handler.Timeout = 10 * time.Second

		slaveIDs := []byte{1, 2} 
	
		// // Connect to the Modbus RTU device
		// err := handler.Connect()
		// if err != nil {
		// 	log.Fatalf("Failed to connect to Modbus device: %v", err)
		// }
		// defer handler.Close()
	
		// // Create a Modbus client
		// client := modbus.NewClient(handler)
	
		// Example: Read Holding Registers (starting at address 0, read 10 registers)
		for _, slaveID := range slaveIDs {

			results, err := modbusservice.ReadRegister(slaveID, 131, 1)
			if err != nil {
				log.Printf("Error reading register: %v", err)
			} else {
				fmt.Printf("Register value: %v\n", results)
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