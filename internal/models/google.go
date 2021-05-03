package models

type SearchQueryResponse struct {
	Kind  string        `json:"kind"`
	Count uint32        `json:"totalItems"`
	Items []*BookResult `json:"items"`
}

type BookResult struct {
	Id         string `json:"id"`
	VolumeInfo struct {
		Title       string            `json:"title"`
		Subtitle    string            `json:"subtitle"`
		Authors     []string          `json:"authors"`
		Publisher   string            `json:"publisher"`
		PublishDate string            `json:"publishedDate"`
		Description string            `json:"description"`
		PageCount   uint32            `json:"pageCount"`
		Categories  []string          `json:"categories"`
		ImageLinks  map[string]string `json:"imageLinks"`
	} `json:"volumeInfo"`
}
