package routes

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"

	"github.com/julienschmidt/httprouter"
	"github.com/pquerna/ffjson/ffjson"

	"github.com/news-ai/pitch/controllers"

	"github.com/news-ai/web/api"
	nError "github.com/news-ai/web/errors"
)

func handleDatabaseContactAction(c context.Context, r *http.Request, email string, action string) (interface{}, error) {
	switch r.Method {
	case "GET":
		switch action {
		case "tweets":
			val, included, count, total, err := controllers.GetTweetsForContact(c, r, email)
			return api.BaseResponseHandler(val, included, count, total, err, r)
		case "headlines":
			val, included, count, total, err := controllers.GetHeadlinesForContact(c, r, email)
			return api.BaseResponseHandler(val, included, count, total, err, r)
		case "twitterprofile":
			return api.BaseSingleResponseHandler(controllers.GetTwitterProfileForContact(c, r, email))
		case "twittertimeseries":
			return api.BaseSingleResponseHandler(controllers.GetTwitterTimeseriesForContact(c, r, email))
		}
	}
	return nil, errors.New("method not implemented")
}

func handleDatabaseContact(c context.Context, r *http.Request, id string) (interface{}, error) {
	switch r.Method {
	case "GET":
		if id == "locations" {
			val, included, count, total, err := controllers.GetLocationsForContacts(c, r)
			return api.BaseResponseHandler(val, included, count, total, err, r)
		} else if id == "_mapping" {
			return api.BaseSingleResponseHandler(controllers.GetSchemaForContacts(c, r))
		}
		return api.BaseSingleResponseHandler(controllers.GetMediaDatabaseProfile(c, r, id))
	case "PATCH":
		return api.BaseSingleResponseHandler(controllers.UpdateContactInMediaDatabase(c, r, id))
	case "DELETE":
		return api.BaseSingleResponseHandler(controllers.DeleteContactFromMediaDatabase(c, r, id))
	case "POST":
		switch id {
		case "search":
			val, included, count, total, err := controllers.SearchContactsInMediaDatabase(c, r)
			return api.BaseResponseHandler(val, included, count, total, err, r)
		}
	}
	return nil, errors.New("method not implemented")
}

func handleDatabaseContacts(c context.Context, r *http.Request) (interface{}, error) {
	switch r.Method {
	case "GET":
		val, included, count, total, err := controllers.GetMediaDatabaseProfiles(c, r)
		return api.BaseResponseHandler(val, included, count, total, err, r)
	case "POST":
		val, included, count, total, err := controllers.CreateContactInMediaDatabase(c, r)
		return api.BaseResponseHandler(val, included, count, total, err, r)
	}
	return nil, errors.New("method not implemented")
}

// Handler for when the user wants all the agencies.
func MediaDatabaseContactsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	c := appengine.NewContext(r)
	val, err := handleDatabaseContacts(c, r)

	if err == nil {
		err = ffjson.NewEncoder(w).Encode(val)
	}

	if err != nil {
		nError.ReturnError(w, http.StatusInternalServerError, "Media Database handling error", err.Error())
	}
	return
}

// Handler for when there is a key present after /users/<id> route.
func MediaDatabaseContactHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	c := appengine.NewContext(r)
	id := ps.ByName("id")
	val, err := handleDatabaseContact(c, r, id)

	if err == nil {
		err = ffjson.NewEncoder(w).Encode(val)
	}

	if err != nil {
		nError.ReturnError(w, http.StatusInternalServerError, "Media Database handling error", err.Error())
	}
	return
}

func MediaDatabaseContactActionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	c := appengine.NewContext(r)
	id := ps.ByName("id")
	action := ps.ByName("action")
	val, err := handleDatabaseContactAction(c, r, id, action)

	if err == nil {
		err = ffjson.NewEncoder(w).Encode(val)
	}

	if err != nil {
		nError.ReturnError(w, http.StatusInternalServerError, "Media Database handling error", err.Error())
	}
	return
}
