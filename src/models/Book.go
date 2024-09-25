package models

import (
	"strings"
)

type Book struct {
	ID           uint64 `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Subtitle     string `json:"subtitle,omitempty"`
	Description  string `json:"description,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Publisher    string `json:"publisher,omitempty"`
	Thumbnail    string `json:"thumbnail,omitempty"`
	Authors      string `json:"authors,omitempty"`
}

func (book *Book) FormatBook() {
	book.Title = strings.TrimSpace(book.Title)
	if len(book.Title) > 200 {
		book.Title = book.Title[0:199]
	}
	book.Subtitle = strings.TrimSpace(book.Subtitle)
	if len(book.Subtitle) > 200 {
		book.Subtitle = book.Subtitle[0:199]
	}
	book.Description = strings.TrimSpace(book.Description)
	if len(book.Description) > 500 {
		book.Description = book.Description[0:499]
	}
	book.Publisher = strings.TrimSpace(book.Publisher)
	if len(book.Publisher) > 99 {
		book.Publisher = book.Publisher[0:99]
	}
	book.Published_at = strings.TrimSpace(book.Published_at)
	book.Thumbnail = strings.TrimSpace(book.Thumbnail)
	//need to format the list of author

}
