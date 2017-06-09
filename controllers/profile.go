package controllers

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/qedus/nds"

	"github.com/news-ai/pitch/models"

	apiControllers "github.com/news-ai/api/controllers"
	apiModels "github.com/news-ai/api/models"

	"github.com/news-ai/web/utilities"
)

func GetUserProfile(c context.Context, r *http.Request, id string) (interface{}, interface{}, error) {
	user := apiModels.User{}
	err := errors.New("")

	switch id {
	case "me":
		user, err = apiControllers.GetCurrentUser(c, r)
		if err != nil {
			log.Errorf(c, "%v", err)
			return nil, nil, err
		}
	default:
		userId, err := utilities.StringIdToInt(id)
		if err != nil {
			log.Errorf(c, "%v", err)
			return nil, nil, err
		}
		user, _, err = apiControllers.GetUserById(c, r, userId)
		if err != nil {
			log.Errorf(c, "%v", err)
			return nil, nil, err
		}
	}

	if user.Profile == 0 {
		if user.Type == "journalist" {
			return models.JournalistProfile{}, nil, nil
		}
		return models.PRProfile{}, nil, nil
	}

	if user.Type == "journalist" {
		var userProfile models.JournalistProfile
		userProfileId := datastore.NewKey(c, "JournalistProfile", "", user.Profile, nil)

		err := nds.Get(c, userProfileId, &userProfile)
		if err != nil {
			log.Errorf(c, "%v", err)
			return models.JournalistProfile{}, nil, err
		}

		if !userProfile.Created.IsZero() {
			userProfile.Format(userProfileId, "profiles")
			return userProfile, nil, nil
		}
	} else {
		var userProfile models.PRProfile
		userProfileId := datastore.NewKey(c, "PRProfile", "", user.Profile, nil)

		err := nds.Get(c, userProfileId, &userProfile)
		if err != nil {
			log.Errorf(c, "%v", err)
			return models.JournalistProfile{}, nil, err
		}

		if !userProfile.Created.IsZero() {
			userProfile.Format(userProfileId, "profiles")
			return userProfile, nil, nil
		}
	}

	return models.JournalistProfile{}, nil, errors.New("No profile by this id")
}

func CreateUserProfile(c context.Context, r *http.Request, id string) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func UpdateUserProfile(c context.Context, r *http.Request, id string) (interface{}, interface{}, error) {
	return nil, nil, nil
}
