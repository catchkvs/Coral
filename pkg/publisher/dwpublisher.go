package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"github.com/catchkvs/Coral/pkg/config"
	"log"
)

var client *pubsub.Client
var ctx = context.Background()


func init() {
	projectID := config.GetProjectId();
	client, _ = pubsub.NewClient(ctx, projectID)
}

// publish entities to topic to push it to Bigquery
func publishToDW(topicID string,  object interface{}, info DWDatasetInfo) (string,error) {
	log.Println("topic: ", topicID, " object: ", object, "datainfo: ", info)
	if config.PROFILE == "dev" {
		log.Println("Dev environment: logging message localy")
	} else {
		t := client.Topic(topicID)
		data, _ := json.Marshal(object)
		payload := b64.StdEncoding.EncodeToString(data)
		dwMessage := DWMessage{
			DatasetInfo: info,
			Payload:     payload,
		}
		dwMessageData, _ := json.Marshal(dwMessage)
		log.Println(dwMessageData)
		result := t.Publish(ctx, &pubsub.Message{
			Data: dwMessageData,
		})
		// Block until the result is returned and a server-generated
		// ID is returned for the published message.
		id, err := result.Get(ctx)
		log.Println("message published: ", id)
		if err != nil {
			return "", err
		}
		return id, nil
	}
	return "", nil

}

type DWMessage struct {
	DatasetInfo DWDatasetInfo
	Type string
	Payload string
}

type DWDatasetInfo struct {
	DatasetName string
	TableName string
}

