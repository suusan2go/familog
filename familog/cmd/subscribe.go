package cmd

import (
	"github.com/lileio/pubsub"
	"github.com/spf13/cobra"
	"github.com/suusan2go/familog/subscribers"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to and process queue messages",
	Run: func(cmd *cobra.Command, args []string) {
		pubsub.Subscribe(&subscribers.FamilogServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
}
