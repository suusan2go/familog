package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) UpdateProfile(ctx context.Context, r *familog.UpdateProfileRequest) (*familog.UpdateProfileResponse, error) {
	return nil, errors.New("not yet implemented")
}
