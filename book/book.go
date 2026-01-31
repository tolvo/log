package book

import (
	"logme/storage"
	"time"
)

type Book struct {
	Name 			string
	Notes			[]Note
	LastNoteNumber	int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	Shelter 		*storage.Shelter
}

func NewBook(name string) *Book {
	return &Book{
		Name:      			name,
		Notes:     			[]Note{},
		LastNoteNumber: 	0,
		CreatedAt: 			time.Now(),
		UpdatedAt: 			time.Now(),
		Shelter:       		storage.NewShelter(".log_shelter/"),
	}
}

func (b *Book) GetNotes() []Note {
	return b.Notes
}

func (b *Book) AddNote(title, content string) {
	note := NewNote(title, content, *book)
}
