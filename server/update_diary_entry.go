package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) UpdateDiaryEntry(ctx context.Context, r *familog.UpdateDiaryEntryRequest) (*familog.UpdateDiaryEntryResponse, error) {
	return nil, errors.New("not yet implemented")
}
