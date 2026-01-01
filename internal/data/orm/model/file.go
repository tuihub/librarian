package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type File struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID   model.InternalID `gorm:"column:user_file"` // Inferred
	Name      string
	Size      int64
	Type      string
	Sha256    []byte
	UpdatedAt time.Time
	CreatedAt time.Time
	Owner     *User  `gorm:"foreignKey:OwnerID"`
	Image     *Image `gorm:"foreignKey:FileID"` // Check Image schema
}

func (File) TableName() string {
	return "files"
}
