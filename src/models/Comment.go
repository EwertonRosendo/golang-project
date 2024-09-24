package models

type Comment struct {
	ID        uint64 `json:"id,omitempty"`
	Comment   string `json:"comment,omitempty"`
	User      User   `json:"user,omitempty"`
	Review    Review `json:"review,omitempty"`
	CreatedAt string `json:"CreatedAt,omitempty"`
}
