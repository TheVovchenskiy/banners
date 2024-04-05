package model

type Feature struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt" format:"date-time"`
	UpdatedAt   string `json:"updatedAt" format:"date-time"`
}
