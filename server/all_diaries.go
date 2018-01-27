package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) AllDiaries(ctx context.Context, r *familog.AllDiariesRequest) (*familog.AllDiariesResponse, error) {
	return nil, errors.New("not yet implemented")
}
