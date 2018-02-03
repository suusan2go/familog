package server

import (
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app/usecase"
	context "golang.org/x/net/context"
)

func (s FamilogServer) RegisterDevice(ctx context.Context, r *familog.RegisterDeviceRequest) (*familog.RegisterDeviceResponse, error) {
	u := usecase.NewRegisterDeviceUsecase(
		s.Registry.DeviceRepository(),
		s.Registry.UserRepository(),
	)
	if _, err := u.Register(r.GetDeviceId()); err != nil {
		return nil, err
	}

	res := &familog.RegisterDeviceResponse{}
	return res, nil
}
