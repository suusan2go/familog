package domain

// DeviceRepository device persistence and find
type DeviceRepository interface {
	Save(*Device) error
	FindByDeviceID(deviceID string) (*Device, error)
}
