package controllers

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/news-ai/pitch/models"

	apiModels "github.com/news-ai/api/models"
)

func GetUserProfile(c context.Context, r *http.Request, user apiModels.User) (interface{}, error) {
	if user.Type == "journalist" {
		var userProfile models.JournalistProfile
		userProfileId := datastore.NewKey(c, "JournalistProfile", "", user.Profile, nil)
	} else {
		var userProfile models.PRProfile
		userProfileId := datastore.NewKey(c, "PRProfile", "", user.Profile, nil)
	}
}

func UpdateUserProfile(c context.Context, r *http.Request, user apiModels.User) (interface{}, error) {

}
