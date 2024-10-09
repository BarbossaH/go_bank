package dto

type CustomerRes struct {
	Id       string `json:"customer_id,omitempty"`
	Name     string `json:"name,omitempty"`
	City     string `json:"city,omitempty"`
	Zipcode  string `json:"zipcode,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	Status   string `json:"status,omitempty"`
}