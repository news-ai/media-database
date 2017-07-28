package models

import (
	apiModels "github.com/news-ai/api/models"
)

type Publication struct {
	Id   string `json:"id"`
	Type string `json:"type"`

	OrganizationName string `json:"name"`
	URL              string `json:"url"`
}

func (p *Publication) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := apiModels.SetField(p, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
