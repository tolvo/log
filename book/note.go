package book

import (
	"time"
)

type Note struct {
	Number   	int 
	Title    	string
	Content  	string
	Timestamp 	time.Time
	Archived	bool
	Book 		*Book
}

func NewNote(title, content string, book *Book) *Note {
	
	if title == "" {
		panic("Note title cannot be empty")
	}
	
	if content == "" {
		panic("Note content cannot be empty")
	}
	
	return &Note{
		Title:     title,
		Content:   content,
		Timestamp: time.Now(),
		Archived:  false,
		Book:      &Book{
			Name:       book.Name,
			Notes:      book.Notes,
			LastNoteNumber: book.LastNoteNumber,
			CreatedAt:  book.CreatedAt,
			UpdatedAt:  book.UpdatedAt,
			Shelter:    book.Shelter,
		},
	}
}