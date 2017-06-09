package models

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	apiModels "github.com/news-ai/api/models"

	"github.com/qedus/nds"
)

type ProfileShare struct {
	// Bio & Title
	ProfileImage string `json:"profileimage"`
	Location     string `json:"location"`
	JobTitle     string `json:"jobtitle"`
	Bio          string `json:"bio"`

	Website  string `json:"website"`
	Blog     string `json:"blog"`
	Twitter  string `json:"twitter"`
	Linkedin string `json:"linkedin"`
}

type JournalistProfile struct {
	apiModels.Base
	ProfileShare

	// Topics Covered & Not Covered
	TopicsCovered    []string `json:"topicscovered"`
	TopicsNotCovered []string `json:"topicsnotcovered"`

	RSSFeeds []string `json:"rssfeeds"`

	NewsAIURL string `json:"newsaiurl"`

	Beats        []string `json:"beats"`
	Regions      []string `json:"regions"`
	MediaOutlets []int64  `json:"mediaoutlets"`

	Verified bool `json:"verified"`
}

type PRProfile struct {
	apiModels.Base
	ProfileShare
}

/*
* Public methods
 */

func (jp *JournalistProfile) Key(c context.Context) *datastore.Key {
	return jp.BaseKey(c, "JournalistProfile")
}

func (prp *PRProfile) Key(c context.Context) *datastore.Key {
	return prp.BaseKey(c, "PRProfile")
}

/*
* Create methods
 */

func (jp *JournalistProfile) Create(c context.Context, r *http.Request, currentUser apiModels.User) (*JournalistProfile, error) {
	jp.CreatedBy = currentUser.Id
	jp.Created = time.Now()

	_, err := jp.Save(c, r)
	return jp, err
}

func (prp *PRProfile) Create(c context.Context, r *http.Request, currentUser apiModels.User) (*PRProfile, error) {
	prp.CreatedBy = currentUser.Id
	prp.Created = time.Now()

	_, err := prp.Save(c, r)
	return prp, err
}

/*
* Update methods
 */

// Function to save a new journalist profile into App Engine
func (jp *JournalistProfile) Save(c context.Context, r *http.Request) (*JournalistProfile, error) {
	// Update the Updated time
	jp.Updated = time.Now()

	k, err := nds.Put(c, jp.BaseKey(c, "JournalistProfile"), jp)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, err
	}
	jp.Id = k.IntID()
	return jp, nil
}

// Function to save a new pr profile into App Engine
func (prp *PRProfile) Save(c context.Context, r *http.Request) (*PRProfile, error) {
	// Update the Updated time
	prp.Updated = time.Now()

	k, err := nds.Put(c, prp.BaseKey(c, "PRProfile"), prp)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, err
	}
	prp.Id = k.IntID()
	return prp, nil
}
