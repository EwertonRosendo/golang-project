package models

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Title         string   `json:"title"`
			Subtitle      string   `json:"subtitle"`
			Description   string   `json:"description"`
			Authors       []string `json:"authors"`
			PublishedDate string   `json:"publishedDate"`
			Publisher     string   `json:"publisher"`
			ImageLinks    struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	} `json:"items"`
}
