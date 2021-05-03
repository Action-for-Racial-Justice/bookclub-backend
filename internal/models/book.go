package models

type BookRequest struct {
}

// Book is a book data entry
// swagger:model Book
//Book ...
type Book struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	IsActive bool   `json:"is_active" db:"isActive"`
}

//BookDataRequest ...
type BookDataRequest struct {
	EntryID string `json:"entry_id"`
}
