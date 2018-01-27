package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) ShowMyProfile(ctx context.Context, r *familog.ShowMyProfileRequest) (*familog.ShowMyProfileResponse, error) {
	return nil, errors.New("not yet implemented")
}
