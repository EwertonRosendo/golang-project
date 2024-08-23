package models


type Book struct {
	Title         string   `json:"title,omitempty"`
	Subtitle      string   `json:"subtitle,omitempty"`
	Description   string   `json:"description,omitempty"`
	PublishedDate string   `json:"publishedDate,omitempty"`
	Publisher     string   `json:"publisher,omitempty"`
	Thumbnail     string   `json:"thumbnail,omitempty"`
	Authors       []string `json:"authors,omitempty"`
}
