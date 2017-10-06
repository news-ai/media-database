package models

import (
	// "net/http"
	// "time"

	apiModels "github.com/news-ai/api-v1/models"
)

type Pitch struct {
	apiModels.Base
}

/*
* Public methods
 */

// func (p *Pitch) Key(c context.Context) *datastore.Key {
// 	return p.BaseKey(c, "Pitch")
// }

/*
* Create methods
 */

// func (p *Pitch) Create(c context.Context, r *http.Request, currentUser apiModels.User) (*Pitch, error) {
// 	p.CreatedBy = currentUser.Id
// 	p.Created = time.Now()

// 	_, err := p.Save(c, r)
// 	return p, err
// }

/*
* Update methods
 */

// // Function to save a new wallet into App Engine
// func (p *Pitch) Save(c context.Context, r *http.Request) (*Pitch, error) {
// 	// Update the Updated time
// 	p.Updated = time.Now()

// 	k, err := nds.Put(c, p.BaseKey(c, "Pitch"), p)
// 	if err != nil {
// 		log.Errorf(c, "%v", err)
// 		return nil, err
// 	}
// 	p.Id = k.IntID()
// 	return p, nil
// }
