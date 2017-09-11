package sync

import (
	"encoding/json"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/pubsub"
)

func sync(r *http.Request, data map[string]string, topicName string) error {
	c := appengine.NewContext(r)
	PubsubClient, err := configurePubsub(r)
	if err != nil {
		log.Errorf(c, "%v", err)
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Errorf(c, "%v", err)
		return err
	}

	topic := PubsubClient.Topic(topicName)
	defer topic.Stop()

	var results []*pubsub.PublishResult
	res := topic.Publish(c, &pubsub.Message{Data: jsonData})
	results = append(results, res)
	for _, result := range results {
		id, err := result.Get(c)
		if err != nil {
			log.Infof(c, "%v", err)
			continue
		}
		log.Infof(c, "Published a message with a message ID: %s\n", id)
	}

	return nil
}
func TwitterSync(r *http.Request, twitterUser string) error {
	// Create an map with twitter username
	twitterUser = strings.ToLower(twitterUser)
	data := map[string]string{
		"username": twitterUser,
	}

	return sync(r, data, TwitterTopicID)
}
