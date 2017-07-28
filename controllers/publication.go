package controllers

import (
	"net/http"

	"golang.org/x/net/context"

	gcontext "github.com/gorilla/context"

	pitchModels "github.com/news-ai/pitch/models"

	"google.golang.org/appengine/log"

	"github.com/news-ai/api/search"
)

/*
* Public methods
 */

/*
* Get methods
 */

func GetMediaDatabasePublications(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	queryField := gcontext.Get(r, "q").(string)
	if queryField != "" {
		publications, total, err := search.SearchPublicationInESMediaDatabase(c, r, queryField)
		if err != nil {
			return []pitchModels.Publication{}, nil, 0, 0, err
		}
		return publications, nil, len(publications), total, nil
	}

	contacts, hits, total, err := search.SearchESMediaDatabasePublications(c, r)
	if err != nil {
		log.Errorf(c, "%v", err)
		return contacts, nil, 0, 0, err
	}

	return contacts, nil, hits, total, nil
}
