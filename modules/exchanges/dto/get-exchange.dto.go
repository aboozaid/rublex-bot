package dto

type Exchange struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Logo     *string `json:"logo,omitempty"`
	Website  *string `json:"website,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}
