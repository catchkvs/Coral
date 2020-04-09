package listener

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/catchkvs/Coral/pkg/config"
	"github.com/catchkvs/Coral/pkg/metrics"
	"github.com/catchkvs/Coral/pkg/model"
	"github.com/catchkvs/Coral/pkg/server"
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

var client *pubsub.Client
var ctx = context.Background()


func init() {
	projectID := config.GetProperty("coral.googlecloud.projectid")
	// Get a Firestore client.
	client, _ = pubsub.NewClient(ctx, projectID)

}

// Listen fact update messages from google pubsub.
func ListenFacts() {
	subID := config.GetProperty("coral.fact.update.subscription")
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		var fact model.FactEntity;
		json.Unmarshal(msg.Data, &fact)
		store := server.GetSessionStore();
		if store.IsFactChannelPresent(fact.DimensionId) {
			channel := store.GetFactChannel(fact.DimensionId)
			channel <- &fact
		} else {
			log.Println("No channel available for given dimension")
			metrics.MissingChannelCounter.With(prometheus.Labels{"missing_channel":fact.DimensionId}).Inc()
		}
		msg.Ack()
	})

	if err != nil {
		fmt.Errorf("Receive: %v", err)
	}
}