package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestReadDiaryEntry(t *testing.T) {
	ctx := context.Background()
	req := &familog.ReadDiaryEntryRequest{}

	res, err := cli.ReadDiaryEntry(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
