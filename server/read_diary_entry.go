package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) ReadDiaryEntry(ctx context.Context, r *familog.ReadDiaryEntryRequest) (*familog.ReadDiaryEntryResponse, error) {
	return nil, errors.New("not yet implemented")
}
