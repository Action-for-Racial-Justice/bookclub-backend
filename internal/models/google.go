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
		Authors     []string          `json:"authors"`
		Description string            `json:"description"`
		ImageLinks  map[string]string `json:"imageLinks"`
	} `json:"volumeInfo"`
}
