package models

type Review struct {
	ID     uint64 `json:"id,omitempty"`
	User   User   `json:"user,omitempty"`
	Book   Book   `json:"book,omitempty"`
	Status string `json:"status,omitempty"`
	Rating string `json:"rating,omitempty"`
	Review string `json:"review,omitempty"`
}
