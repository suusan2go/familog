package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) SubscribeDiary(ctx context.Context, r *familog.SubscribeDiaryRequest) (*familog.SubscribeDiaryResponse, error) {
	return nil, errors.New("not yet implemented")
}
