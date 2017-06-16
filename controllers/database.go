package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/log"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/news-ai/pitch/models"

	"github.com/news-ai/api/search"
)

/*
* Private methods
 */

type createMediaDatabaseContact struct {
	Email string `json:"email"`

	WritingInformation struct {
		Beats           []string `json:"beats"`
		OccasionalBeats []string `json:"occasionalBeats"`
		IsFreelancer    bool     `json:"isFreelancer"`
	} `json:"writingInformation"`
}

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

func CreateContactInMediaDatabase(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	buf, _ := ioutil.ReadAll(r.Body)
	decoder := ffjson.NewDecoder()
	var createContact createMediaDatabaseContact
	err := decoder.Decode(buf, &createContact)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, 0, 0, err
	}

	// Check if contact already exists in Media Database
	_, err = search.SearchContactInMediaDatabase(c, r, createContact.Email)
	if err == nil {
		return models.MediaDatabaseProfile{}, nil, 0, 0, errors.New("Contact already exists in Media Database")
	}

	// Get contact from Enhance Full Contact
	contactProfile, err := search.SearchContactDatabaseForMediaDatbase(c, r, createContact.Email)
	if err != nil {
		return models.MediaDatabaseProfile{}, nil, 0, 0, err
	}

	if contactProfile.Data.Status != 200 {
		return models.MediaDatabaseProfile{}, nil, 0, 0, errors.New("Could not retrieve contact data from Enhance")
	}

	// Alter contact details before writing it to Media Database
	contactProfile.Data.WritingInformation = createContact.WritingInformation

	log.Infof(c, "%v", contactProfile)

	// Add contact to Media Database with approved flag off

	return contactProfile, nil, 1, 0, nil
}

/*
* Update methods
 */

func UpdateContactInMediaDatabase(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	return nil, nil, nil
}

/*
* Delete methods
 */

func DeleteContactFromMediaDatabase(c context.Context, r *http.Request) {

}
