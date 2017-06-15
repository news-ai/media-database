package controllers

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/log"

	"github.com/news-ai/pitch/models"

	"github.com/news-ai/api/search"
)

func GetMediaDatabaseProfile(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	contactProfile, err := search.SearchContactDatabaseForMediaDatabase(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}

	return contactProfile.Data, nil, err
}
