package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Desk struct {
	DeskPath string
}

type Book struct {
	Name 			string
	Notes 			[]Note
	LastNotePage	int
	Desk 			*Desk
}

type Note struct {
	Number   	int
	Title    	string
	Content  	string
	CreatedAt  	time.Time
	Archived 	bool
	Book     	*Book
}

func NewDesk(deskPath string) *Desk {
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	fullPath := filepath.Join(dir, deskPath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.MkdirAll(fullPath, os.ModePerm)
	}

	return &Desk{
		DeskPath: fullPath,
	}
}

func (d *Desk) NewBook(name string) *Book {
	bookPath := filepath.Join(d.DeskPath, name)
	if _, err := os.Stat(bookPath); os.IsNotExist(err) {
		os.MkdirAll(bookPath, os.ModePerm)
	}

	return &Book{
		Name:         name,
		Notes:        []Note{},
		LastNotePage: 0,
		Desk:         d,
	}
}

func (d *Desk) GetBook(name string) (*Book, error) {
	bookPath := filepath.Join(d.DeskPath, name)
	if _, err := os.Stat(bookPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("book '%s' does not exist", name)
	}

	return &Book{
		Name:         name,
		Notes:        []Note{},
		LastNotePage: 0,
		Desk:         d,
	}, nil
}

func (b *Book) NewNote(title, content string) *Note {
	if title == "" {
		panic("Note title cannot be empty")
	}

	if content == "" {
		panic("Note content cannot be empty")
	}

	note := &Note{
		Number:    b.LastNotePage + 1,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		Archived:  false,
		Book:      b,
	}

	b.Notes = append(b.Notes, *note)
	b.LastNotePage++

	return note
}

func (n *Note) Archive() {
	n.Archived = true
}

func (n *Note) Unarchive() {
	n.Archived = false
}

func WriteNoteToFile(note *Note) error {
	// Implementation to write note to file goes here
	return nil
}

func ReadNoteFromFile(book *Book, noteNumber int) (*Note, error) {
	// Implementation to read note from file goes here
	return nil, nil
}
