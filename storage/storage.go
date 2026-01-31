package storage

import (
	"os"
	"path/filepath"
)

type Shelter struct {
	StoragePath string
}

func NewShelter(storagePath string) *Shelter {
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	fullPath := filepath.Join(dir, storagePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.MkdirAll(fullPath, os.ModePerm)
	}

	return &Shelter{
		StoragePath: fullPath,
	}
}
