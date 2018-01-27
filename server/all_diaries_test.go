package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestAllDiaries(t *testing.T) {
	ctx := context.Background()
	req := &familog.AllDiariesRequest{}

	res, err := cli.AllDiaries(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
