package main

import (
	"github.com/lileio/lile"
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/familog/cmd"
	"github.com/suusan2go/familog/server"
	"google.golang.org/grpc"
)

func main() {
	s := &server.FamilogServer{}

	lile.Name("familog")
	lile.Server(func(g *grpc.Server) {
		familog.RegisterFamilogServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
	})

	cmd.Execute()
}
