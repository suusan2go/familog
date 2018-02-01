package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/lileio/lile"
	"github.com/lileio/lile/fromenv"
	"github.com/lileio/pubsub"
	"github.com/sirupsen/logrus"
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app"
	"github.com/suusan2go/familog/familog/cmd"
	"github.com/suusan2go/familog/server"
	"google.golang.org/grpc"
)

func main() {
	s := &server.FamilogServer{}
	dsc, err := configureDatastoreDB("dummy")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	s.Registry = app.NewRegistry(dsc)

	lile.Name("familog")
	lile.Server(func(g *grpc.Server) {
		familog.RegisterFamilogServer(g, s)
		wrappedServer := grpcweb.WrapServer(g)
		handler := func(resp http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/" {
				wrappedServer.ServeHttp(resp, req)
				logrus.Infof("%s", req.URL.Path)
				return
			}
			resp.WriteHeader(http.StatusOK)
			fmt.Fprintf(resp, "ok")
		}
		http.HandleFunc("/", handler)
		logrus.Infof("grpcwebproxy runs at :9000")
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
	})

	cmd.Execute()
}

func configureDatastoreDB(projectID string) (*datastore.Client, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}
