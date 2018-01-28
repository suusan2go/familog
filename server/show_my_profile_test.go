package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestShowMyProfile(t *testing.T) {
	ctx := context.Background()
	req := &familog.ShowMyProfileRequest{}

	res, err := cli.ShowMyProfile(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
