package controllers

/*
* Public methods
 */

/*
* Get methods
 */

func GetMediaDatabasePublications(c context.Context, r *http.Request) (interface{}, interface{}, int, int, error) {
	contacts, hits, total, err := search.SearchESMediaDatabasePublications(c, r)
	if err != nil {
		log.Errorf(c, "%v", err)
		return contacts, nil, 0, 0, err
	}

	return contacts, nil, hits, total, nil
}
