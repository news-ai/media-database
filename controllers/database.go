package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/news-ai/pitch/models"
	"github.com/news-ai/pitch/sync"

	"github.com/news-ai/web/middleware"

	pitchModels "github.com/news-ai/pitch/models"
	tabulaeModels "github.com/news-ai/tabulae/models"

	"github.com/news-ai/api/search"
)

/*
* Private methods
 */

type createMediaDatabaseContact struct {
	Email string `json:"email"`

	ContactInfo        pitchModels.ContactInfo        `json:"contactInfo"`
	Demographics       pitchModels.Demographic        `json:"demographics"`
	Photos             []pitchModels.Photo            `json:"photos"`
	SocialProfiles     []pitchModels.SocialProfile    `json:"socialProfiles"`
	DigitalFootprint   pitchModels.DigitalFootprint   `json:"digitalFootprint"`
	Organizations      []pitchModels.Organization     `json:"organizations"`
	WritingInformation pitchModels.WritingInformation `json:"writingInformation"`
}

/*
* Get methods
 */

func getMediaDatabaseProfile(c context.Context, r *http.Request, email string) (models.MediaDatabaseProfile, error) {
	contactProfile, err := search.SearchContactInMediaDatabase(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, err
	}

	if contactProfile.Data.Status != 200 {
		return models.MediaDatabaseProfile{}, errors.New("Contact does not exist in Enhance")
	}

	return contactProfile, err
}

func getMediaDatabaseContactSocialNetworkUsername(c context.Context, r *http.Request, contact models.MediaDatabaseProfile, socialNetwork string) string {
	username := ""
	for i := 0; i < len(contact.Data.SocialProfiles); i++ {
		if contact.Data.SocialProfiles[i].TypeID == socialNetwork {
			username = contact.Data.SocialProfiles[i].Username
		}
	}

	return username
}

/*
* Public methods
 */

/*
* Get methods
 */

func GetSchemaForContacts(c context.Context, r *http.Request) (interface{}, interface{}, error) {
	mapping, err := search.GetMediaDatabaseContactsSchema(c)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, err
	}

	return mapping, nil, nil
}

func GetMediaDatabaseProfiles(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	contacts, hits, total, err := search.SearchESMediaDatabase(c, r)
	if err != nil {
		log.Errorf(c, "%v", err)
		return contacts, nil, 0, 0, err
	}

	return contacts, nil, hits, total, nil
}

func GetMediaDatabaseProfile(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	contactProfile, err := getMediaDatabaseProfile(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}

	return contactProfile.Data, nil, nil
}

/*
* Search methods
 */

func SearchContactsInMediaDatabase(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	buf, _ := ioutil.ReadAll(r.Body)
	decoder := ffjson.NewDecoder()
	var searchQuery search.SearchMediaDatabaseQuery
	err := decoder.Decode(buf, &searchQuery)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, 0, 0, err
	}

	contacts, hits, total, err := search.SearchContactsInESMediaDatabase(c, r, searchQuery)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, 0, 0, err
	}

	return contacts, nil, hits, total, nil
}

/*
* RSS methods
 */

func GetHeadlinesForContact(c context.Context, r *http.Request, email string) ([]search.Headline, interface{}, int, int, error) {
	contact, err := getMediaDatabaseProfile(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, 0, 0, err
	}

	if len(contact.Data.WritingInformation.RSS) == 0 {
		return nil, nil, 0, 0, errors.New("This contact has no RSS feeds")
	}

	headlines, total, err := search.SearchHeadlinesByResourceId(c, r, []tabulaeModels.Feed{}, contact.Data.WritingInformation.RSS)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, 0, 0, err
	}

	return headlines, nil, len(headlines), total, nil
}

/*
* Twitter methods
 */

func GetTweetsForContact(c context.Context, r *http.Request, email string) (interface{}, interface{}, int, int, error) {
	contact, err := getMediaDatabaseProfile(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, 0, 0, err
	}

	twitterUsername := getMediaDatabaseContactSocialNetworkUsername(c, r, contact, "twitter")
	tweets, total, err := search.SearchTweetsByUsername(c, r, twitterUsername)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, 0, 0, err
	}

	return tweets, nil, len(tweets), total, nil
}

func GetTwitterProfileForContact(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	contact, err := getMediaDatabaseProfile(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, err
	}

	twitterUsername := getMediaDatabaseContactSocialNetworkUsername(c, r, contact, "twitter")
	twitterProfile, err := search.SearchProfileByUsername(c, r, twitterUsername)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, err
	}

	return twitterProfile, nil, nil
}

func GetTwitterTimeseriesForContact(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	// Get the details of the current user
	contact, err := getMediaDatabaseProfile(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, err
	}

	twitterUsername := getMediaDatabaseContactSocialNetworkUsername(c, r, contact, "twitter")
	twitterTimeseries, _, err := search.SearchTwitterTimeseriesByUsername(c, r, twitterUsername)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, nil, err
	}

	return twitterTimeseries, nil, nil
}

/*
* Location methods
 */

func GetLocationsForContacts(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	// variables to define country, state, and city
	country := middleware.GetParambyId(r, "country")
	state := middleware.GetParambyId(r, "state")
	city := middleware.GetParambyId(r, "city")

	if country != "" && state == "" && city == "" {
		countryLocations, hits, total, err := search.ESCountryLocation(c, r, country)
		if err != nil {
			log.Errorf(c, "%v", err)
			return nil, nil, 0, 0, err
		}

		return countryLocations, nil, hits, total, nil
	} else if country != "" && state != "" && city == "" {
		stateLocations, hits, total, err := search.ESStateLocation(c, r, state, country)
		if err != nil {
			log.Errorf(c, "%v", err)
			return nil, nil, 0, 0, err
		}

		return stateLocations, nil, hits, total, nil
	} else if country != "" && state != "" && city != "" {
		cityLocations, hits, total, err := search.ESCityLocation(c, r, city, state, country)
		if err != nil {
			log.Errorf(c, "%v", err)
			return nil, nil, 0, 0, err
		}

		return cityLocations, nil, hits, total, nil
	}

	return nil, nil, 0, 0, nil
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

	createContact.Email = strings.ToLower(createContact.Email)

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
	contactProfile.Data.Created = time.Now()
	contactProfile.Data.Updated = time.Now()
	contactProfile.Data.ToUpdate = false
	contactProfile.Data.Email = createContact.Email
	contactProfile.Data.WritingInformation = createContact.WritingInformation

	// Add contact to Media Database with approved flag off
	contextWithTimeout, _ := context.WithTimeout(c, time.Second*15)
	client := urlfetch.Client(contextWithTimeout)

	ContactProfile, err := json.Marshal(contactProfile)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, 0, 0, err
	}
	contactProfileJson := bytes.NewReader(ContactProfile)
	req, _ := http.NewRequest("POST", "https://enhance.newsai.org/md", contactProfileJson)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.MediaDatabaseProfile{}, nil, 0, 0, errors.New("Fail to POST data to Enhance")
	}

	twitterUsername := getMediaDatabaseContactSocialNetworkUsername(c, r, contactProfile, "twitter")
	err = sync.TwitterSync(r, twitterUsername)
	if err != nil {
		log.Errorf(c, "%v", err)
	}

	return contactProfile.Data, nil, 1, 0, nil
}

/*
* Update methods
 */

func UpdateContactInMediaDatabase(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	buf, _ := ioutil.ReadAll(r.Body)
	decoder := ffjson.NewDecoder()
	var createContact createMediaDatabaseContact
	err := decoder.Decode(buf, &createContact)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}

	contactProfile, err := search.SearchContactInMediaDatabase(c, r, email)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}

	// Alter contact details before writing it to Media Database
	contactProfile.Data.Updated = time.Now()
	contactProfile.Data.ContactInfo = createContact.ContactInfo
	contactProfile.Data.Demographics = createContact.Demographics
	contactProfile.Data.SocialProfiles = createContact.SocialProfiles
	contactProfile.Data.DigitalFootprint = createContact.DigitalFootprint
	contactProfile.Data.Organizations = createContact.Organizations
	contactProfile.Data.Photos = createContact.Photos
	contactProfile.Data.WritingInformation = createContact.WritingInformation

	// Add contact to Media Database with approved flag off
	contextWithTimeout, _ := context.WithTimeout(c, time.Second*15)
	client := urlfetch.Client(contextWithTimeout)

	ContactProfile, err := json.Marshal(contactProfile)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}
	contactProfileJson := bytes.NewReader(ContactProfile)
	req, _ := http.NewRequest("POST", "https://enhance.newsai.org/md", contactProfileJson)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf(c, "%v", err)
		return models.MediaDatabaseProfile{}, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.MediaDatabaseProfile{}, nil, errors.New("Fail to POST data to Enhance")
	}

	twitterUsername := getMediaDatabaseContactSocialNetworkUsername(c, r, contactProfile, "twitter")
	err = sync.TwitterSync(r, twitterUsername)
	if err != nil {
		log.Errorf(c, "%v", err)
	}

	return contactProfile.Data, nil, nil
}

/*
* Delete methods
 */

func DeleteContactFromMediaDatabase(c context.Context, r *http.Request, email string) (interface{}, interface{}, error) {
	return nil, nil, nil
}
