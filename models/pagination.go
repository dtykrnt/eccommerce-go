package models

type Paginate struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
