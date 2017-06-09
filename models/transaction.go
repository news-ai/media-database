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

type Transaction struct {
	apiModels.Base

	Wallet int64 `json:"wallet"`

	Value          int64 `json:"cost"`
	RunningBalance int64 `json:"runningbalance"` // Balance of the wallet after
}

/*
* Public methods
 */

func (t *Transaction) Key(c context.Context) *datastore.Key {
	return t.BaseKey(c, "Transaction")
}

/*
* Create methods
 */

func (t *Transaction) Create(c context.Context, r *http.Request, currentUser apiModels.User) (*Transaction, error) {
	t.CreatedBy = currentUser.Id
	t.Created = time.Now()

	_, err := t.Save(c, r)
	return t, err
}

/*
* Update methods
 */

// Function to save a new wallet into App Engine
func (t *Transaction) Save(c context.Context, r *http.Request) (*Transaction, error) {
	// Update the Updated time
	t.Updated = time.Now()

	k, err := nds.Put(c, t.BaseKey(c, "Transaction"), t)
	if err != nil {
		log.Errorf(c, "%v", err)
		return nil, err
	}
	t.Id = k.IntID()
	return t, nil
}
