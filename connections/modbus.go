package connections

import (
	"time"
	"github.com/grid-x/modbus"
	"fmt"
	"log"
)

var (
    clientInstance *modbus.Client
    once           sync.Once
)

type ModbusConfig struct {
    Port       string
    BaudRate   int
    DataBits   int
    StopBits   int
    Parity     string
    Timeout    time.Duration
}

func getClient(config ModbusConfig) (*modbus.Client, error) {
    var err error
    once.Do(func() {
        handler := modbus.NewRTUClientHandler(config.Port)
        handler.BaudRate = config.BaudRate
        handler.DataBits = config.DataBits
        handler.StopBits = config.StopBits
        handler.Parity = config.Parity
        handler.Timeout = config.Timeout

        err = handler.Connect()
        if err != nil {
            return
        }

        clientInstance = modbus.NewClient(handler)
    })

    return clientInstance, err
}

func GetModbusClient() (*modbus.Client, error) {
    config := ModbusConfig{
        Port:       "/dev/ttyUSB0",
        BaudRate:   19200,
        DataBits:   8,
        StopBits:   1,
        Parity:     "N",
        Timeout:    10 * time.Second,
    }
    return getClient(config)
}

func WriteRegister(slaveID byte, registerAddress uint16, value uint16) error {
    client, err := GetModbusClient()
    if err != nil {
        return err
    }

    handler := client.GetHandler()
    handler.SetSlave(slaveID)

    _, err = client.WriteSingleRegister(registerAddress, value)
    if err != nil {
        return fmt.Errorf("failed to write to register: %w", err)
    }

    return nil
}

func ReadRegister(slaveID byte, registerAddress uint16, registerLength uint16) ([]byte, error) {
    client, err := GetModbusClient()
    if err != nil {
        return nil, err
    }

    handler := client.GetHandler()
    handler.SetSlave(slaveID)

    results, err := client.ReadHoldingRegisters(registerAddress, registerLength)
    if err != nil {
        return nil, fmt.Errorf("failed to read registers: %w", err)
    }

    return results, nil
}