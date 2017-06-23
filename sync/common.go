package sync

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/pubsub"
)

var (
	PubsubClient   *pubsub.Client
	TwitterTopicID = "process-twitter-md-feed"
	projectID      = "newsai-1166"
)

func configurePubsub(r *http.Request) (*pubsub.Client, error) {
	if PubsubClient != nil {
		return PubsubClient, nil
	}
	c := appengine.NewContext(r)
	PubsubClient, err := pubsub.NewClient(c, projectID)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, err
	}

	// Create the topic for twitter if it doesn't exist.
	if exists, err := PubsubClient.Topic(TwitterTopicID).Exists(c); err != nil {
		log.Errorf(c, "%v", err)
		return nil, err
	} else if !exists {
		if _, err := PubsubClient.CreateTopic(c, TwitterTopicID); err != nil {
			log.Errorf(c, "%v", err)
			return nil, err
		}
	}

	return PubsubClient, nil
}
