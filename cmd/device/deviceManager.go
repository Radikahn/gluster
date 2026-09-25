package device

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

// Device struct
// - Holds information for each registered device under the current Gluster instance
// - Not recommended to make your own instances of the `Device` struct, instead register them through `RegisterDevice`
type Device struct {
	id          string
	name        string
	os          string
	directiveId string
	totalLoad   int
	address     string
	active      bool
}

// Register a device to the Gluster instance.
// Automatically adds new device to `globalDevices` map
// Note: Registering a device will automcatically set it's `active` status to `false`
// Please ensure that you set the status to true if you intend to immediately use this cluster
//
// returns: newDevice *Device, deviceId string, error
func RegisterDevice(name string, os string, address string, globalDevices map[string]string) (*Device, string, error) {
	var id = uuid.NewString()

	//quick sanity check to prevent duplicate names for devices since we using a map
	if globalDevices[name] != "" {
		slog.Error("[DEVICE]: the name provided for this device already exists", "name", name)
		return nil, "", fmt.Errorf("Please provide a unique name, %s, is already taken", name)
	}

	var newDevice = &Device{
		id:        id,
		name:      name,
		os:        os,
		totalLoad: 0, //all machines should be initialized with zero load
		address:   address,
		active:    false,
	}

	globalDevices[name] = id

	slog.Info(
		fmt.Sprintf("[DEVICE]: device registered: %s with id: %s", name, id),
	)

	return newDevice, id, nil
}

// Assign a CommandDirective linked list ID to the given device
func AssignDirective(directiveId string, device *Device) {
	if !device.active {
		slog.Warn("[DEVICE]: device is NOT active, this task will NOT run", "name", device.name)
	}

	slog.Info("[DEVICE]: assigning command directive to device", "directiveId", directiveId, "name", device.name)
	device.directiveId = directiveId
}

// Set the provided device to active
func ActivateDevice(device *Device) {
	slog.Info("Activating device", "name", device.name)
	device.active = true
}

// Set the provided device to NOT active
func DeactivateDevice(device *Device) {
	slog.Info("Deactivating device", "name", device.name)
	device.active = false
}
