package modelbinah

import (
	"errors"

	"github.com/tuihub/librarian/internal/model"
)

type FileMetadata struct {
	ID     model.InternalID `json:"id,string"`
	Name   string
	Size   int64
	Type   FileType
	Sha256 []byte
}

type FileType int

const (
	FileTypeUnspecified FileType = iota
	FileTypeGeburaSave
	FileTypeChesedImage
)

const MaxFileSize = 256 << 20

func (f FileMetadata) Check() error {
	if len(f.Name) == 0 {
		return errors.New("empty file name")
	}
	if f.Size <= 0 || f.Size >= MaxFileSize {
		return errors.New("invalid file size")
	}
	if f.Type == FileTypeUnspecified {
		return errors.New("invalid file type")
	}
	return nil
}
