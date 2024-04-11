package domain

type Tag struct {
	Id        uint   `json:"id"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty" format:"date-time"`
	UpdatedAt string `json:"updatedAt,omitempty" format:"date-time"`
}
