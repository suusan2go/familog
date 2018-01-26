package subscribers

import "github.com/lileio/pubsub"

// See https://godoc.org/github.com/lileio/pubsub#Handler for an example
// of what an subscriber handler should look like

type FamilogServiceSubscriber struct{}

func (s *FamilogServiceSubscriber) Setup(c *pubsub.Client) {
	// // https://godoc.org/github.com/lileio/lile/pubsub#Client.On
	// c.On("service.event_name", "subscriber_name", s.CallbackMethod, 30*time.Second, true)
}
