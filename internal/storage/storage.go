package storage

import (
	"fmt"

	"github.com/b0gochort/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetByID(fileid uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileid]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileid)
	}
	return foundFile, nil
}
