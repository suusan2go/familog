package domain

import (
	"time"
)

// Device device user uses
type Device struct {
	ID           uint64
	OwnerUserID  UserID
	DeviceID     string
	RegisteredAt time.Time
}

// NewDevice create new device Struct
func NewDevice(deviceID string) Device {
	return Device{DeviceID: deviceID, RegisteredAt: time.Now()}
}
