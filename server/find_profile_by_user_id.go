package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) FindProfileByUserId(ctx context.Context, r *familog.FindProfileByUserIdRequest) (*familog.FindProfileByUserIdResponse, error) {
	return nil, errors.New("not yet implemented")
}
