package controllers

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/log"

	"github.com/news-ai/pitch/models"

	"github.com/news-ai/api/search"
)

/*
* Public methods
 */

/*
* Get methods
 */

func GetMediaDatabaseProfiles(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	return nil, nil, 0, 0, nil
}

func GetMediaDatabaseProfile(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	contactProfile, err := search.SearchContactInMediaDatabase(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}

	return contactProfile, nil, err
}

/*
* Create methods
 */

func CreateContactInMediaDatabase(c context.Context, r *http.Request) (models.MediaDatabaseProfile, interface{}, int, int, error) {
	// Get contact from Enhance Full Contact
	contactProfile, err := search.SearchContactDatabase(c, r, "")
	if err != nil {
		return models.MediaDatabaseProfile{}, nil, 0, 0, err
	}

	if contactProfile.Data.Status != 200 {
		return models.MediaDatabaseProfile{}, nil, 0, 0, errors.New("Could not retrieve contact data from Enhance")
	}
}

/*
* Update methods
 */

func UpdateContactInMediaDatabase(c context.Context, r *http.Request) {

}

/*
* Delete methods
 */

func DeleteContactFromMediaDatabase(c context.Context, r *http.Request) {

}
