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

type Wallet struct {
	apiModels.Base

	CurrentBalance int64 `json:"currentbalance"`
}

/*
* Public methods
 */

func (w *Wallet) Key(c context.Context) *datastore.Key {
	return w.BaseKey(c, "Wallet")
}

/*
* Create methods
 */

func (w *Wallet) Create(c context.Context, r *http.Request, currentUser apiModels.User) (*Wallet, error) {
	w.CreatedBy = currentUser.Id
	w.Created = time.Now()

	_, err := w.Save(c, r)
	return w, err
}

/*
* Update methods
 */

// Function to save a new wallet into App Engine
func (w *Wallet) Save(c context.Context, r *http.Request) (*Wallet, error) {
	// Update the Updated time
	w.Updated = time.Now()

	k, err := nds.Put(c, w.BaseKey(c, "Wallet"), w)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, err
	}
	w.Id = k.IntID()
	return w, nil
}
