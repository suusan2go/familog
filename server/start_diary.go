package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) StartDiary(ctx context.Context, r *familog.StartDiaryRequest) (*familog.StartDiaryResponse, error) {
	return nil, errors.New("not yet implemented")
}
