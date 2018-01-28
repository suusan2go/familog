package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestUpdateDiaryEntry(t *testing.T) {
	ctx := context.Background()
	req := &familog.UpdateDiaryEntryRequest{}

	res, err := cli.UpdateDiaryEntry(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
