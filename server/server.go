package server

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"fmt"

	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app"
	"github.com/suusan2go/familog/app/domain"
	"google.golang.org/grpc/metadata"
)

type FamilogServer struct {
	familog.FamilogServer
	Registry *app.Registry
}

// return domain.Auth from gRPC metadata x-endpoint-api-userinfo
func (s *FamilogServer) currentAuth(ctx context.Context) (*domain.Auth, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to load metadata")
	}
	info := md["x-endpoint-api-userinfo"][0]
	b, err := base64.StdEncoding.DecodeString(info)
	if err != nil {
		return nil, fmt.Errorf("failed to decode auth jwt: %v", err)
	}
	var auth *domain.Auth
	if err = json.Unmarshal(b, auth); err != nil {
		return nil, fmt.Errorf("failed to unmarshal authjson: %v", err)
	}
	return auth, nil
}
