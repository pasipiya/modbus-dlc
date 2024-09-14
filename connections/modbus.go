package modbusconnection

import (
	"time"
	"github.com/grid-x/modbus"
	"fmt"
    "sync"
    "log"
    "os"
)

var (
    clientInstance modbus.Client
    handlerInstance *modbus.RTUClientHandler
    once           sync.Once
)

type ModbusConfig struct {
    Port       string
    BaudRate   int
    DataBits   int
    StopBits   int
    Parity     string
    Timeout    time.Duration
    SlaveID    byte
}

func isDeviceAvailable(devicePath string) bool {
	_, err := os.Stat(devicePath)
	return !os.IsNotExist(err)
}

// Helper function to connect or reconnect to the Modbus client
func connectClient(config ModbusConfig) error {

    for !isDeviceAvailable(config.Port) {
		log.Printf("Device %s not available. Retrying...", config.Port)
		time.Sleep(3 * time.Second) // Wait and retry
	}

    handler := modbus.NewRTUClientHandler(config.Port)
    handler.BaudRate = config.BaudRate
    handler.DataBits = config.DataBits
    handler.StopBits = config.StopBits
    handler.Parity = config.Parity
    handler.Timeout = config.Timeout
    handler.SlaveID = config.SlaveID

    err := handler.Connect()
    if err != nil {
        return fmt.Errorf("failed to connect to Modbus device: %w", err)
    }

    // Assign handler instance for reconnection
    handlerInstance = handler
    clientInstance = modbus.NewClient(handler)

    return nil
}

func getClient(config ModbusConfig) (modbus.Client, error) {
    var err error
    once.Do(func() {
        err = connectClient(config)
        if err != nil {
            return
        }
    })

    // Check if the device was disconnected, try reconnecting if needed
    if clientInstance == nil || handlerInstance == nil {
        err = connectClient(config)
        if err != nil {
            return nil, err
        }
    }

    return clientInstance, err
}

// Automatically reconnect if write or read operation fails
func reconnectAndRetry(config ModbusConfig) error {
    err := connectClient(config)
    if err != nil {
        return fmt.Errorf("reconnection failed: %w", err)
    }
    return nil
}

func GetModbusClient(SlaveID byte) (modbus.Client, error) {
    config := ModbusConfig{
        Port:       "/dev/ttyusb0",
        BaudRate:   19200,
        DataBits:   8,
        StopBits:   1,
        Parity:     "N",
        Timeout:    10 * time.Second,
        SlaveID:    SlaveID,
    }
    return getClient(config)
}

func WriteRegister(slaveID byte, registerAddress uint16, value uint16) error {
    client, err := GetModbusClient(slaveID)
    if err != nil {
        return err
    }

    _, err = client.WriteSingleRegister(registerAddress, value)
    if err != nil {
        log.Printf("Write failed: %v, attempting reconnection", err)
        
        config := ModbusConfig{
            Port:       "/dev/ttyusb0",
            BaudRate:   19200,
            DataBits:   8,
            StopBits:   1,
            Parity:     "N",
            Timeout:    10 * time.Second,
            SlaveID:    slaveID,
        }
        
        reconnectErr := reconnectAndRetry(config)
        if reconnectErr != nil {
            return reconnectErr
        }
        
        // Retry the operation after reconnecting
        _, err = client.WriteSingleRegister(registerAddress, value)
        if err != nil {
            return fmt.Errorf("failed to write to register after reconnection: %w", err)
        }
    }

    return nil
}

func ReadRegister(slaveID byte, registerAddress uint16, registerLength uint16) ([]byte, error) {
    client, err := GetModbusClient(slaveID)
    if err != nil {
        return nil, err
    }

    results, err := client.ReadHoldingRegisters(registerAddress, registerLength)
    if err != nil {
        log.Printf("Read failed: %v, attempting reconnection", err)

        config := ModbusConfig{
            Port:       "/dev/ttyusb0",
            BaudRate:   19200,
            DataBits:   8,
            StopBits:   1,
            Parity:     "N",
            Timeout:    10 * time.Second,
            SlaveID:    slaveID,
        }

        reconnectErr := reconnectAndRetry(config)
        if reconnectErr != nil {
            return nil, reconnectErr
        }

        // Retry the operation after reconnecting
        results, err = client.ReadHoldingRegisters(registerAddress, registerLength)
        if err != nil {
            return nil, fmt.Errorf("failed to read registers after reconnection: %w", err)
        }
    }

    return results, nil
}