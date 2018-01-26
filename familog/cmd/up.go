package cmd

import (
	"github.com/lileio/lile"
	"github.com/lileio/pubsub"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/suusan2go/familog/subscribers"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC and pubub subscribers",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			logrus.Fatal(lile.Serve())

		}()

		pubsub.Subscribe(&subscribers.FamilogServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
