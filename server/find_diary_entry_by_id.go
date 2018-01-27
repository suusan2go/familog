package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) FindDiaryEntryById(ctx context.Context, r *familog.FindDiaryEntryByIdRequest) (*familog.FindDiaryEntryByIdResponse, error) {
	return nil, errors.New("not yet implemented")
}
