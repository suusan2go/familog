package usecase

import (
	"github.com/suusan2go/familog/app/domain"
)

// RegisterDeviceUsecase register user's device
type RegisterDeviceUsecase struct {
	deviceRepository domain.DeviceRepository
	userRepository   domain.UserRepository
}

func NewRegisterDeviceUsecase(dr domain.DeviceRepository, ur domain.UserRepository) RegisterDeviceUsecase {
	return RegisterDeviceUsecase{
		deviceRepository: dr,
		userRepository:   ur,
	}
}

// Register find or register devie
func (usecase RegisterDeviceUsecase) Register(deviceID string) (*domain.Device, error) {
	device, err := usecase.deviceRepository.FindByDeviceID(deviceID)
	if err != nil {
		return nil, err
	}
	if device != nil {
		return device, nil
	}
	newUser := domain.NewUser()
	if err := usecase.userRepository.Save(&newUser); err != nil {
		return nil, err
	}
	newDevice := domain.NewDevice(deviceID)
	newDevice.OwnerUserID = newUser.ID
	if err = usecase.deviceRepository.Save(&newDevice); err != nil {
		return nil, err
	}
	return &newDevice, nil
}
