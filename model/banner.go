package model

import "encoding/json"

type Banners struct {
	Items []Banner `json:"items"`
}

type Banner struct {
	Id        uint            `json:"id"`
	FeatureId uint            `json:"featureId"`
	Tags      []Tag           `json:"tags"`
	Content   json.RawMessage `json:"content" swaggertype:"string" example:"{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}"`
	IsActive  bool            `json:"isActive"`
	CreatedAt string          `json:"createdAt,omitempty" format:"date-time"`
	UpdatedAt string          `json:"updatedAt,omitempty" format:"date-time"`
}
