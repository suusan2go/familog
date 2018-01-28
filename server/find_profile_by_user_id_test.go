package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestFindProfileByUserId(t *testing.T) {
	ctx := context.Background()
	req := &familog.FindProfileByUserIdRequest{}

	res, err := cli.FindProfileByUserId(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
