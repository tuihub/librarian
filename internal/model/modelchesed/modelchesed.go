package modelchesed

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Image struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"column:user_image;index"`
	FileID      model.InternalID `gorm:"column:file_image;index"`
	Name        string
	Description string
	Status      ImageStatus
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *model.User `gorm:"foreignKey:OwnerID"`
	File        *File       `gorm:"foreignKey:FileID"`
}

func (Image) TableName() string {
	return "images"
}

type File struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID   model.InternalID `gorm:"column:user_file;index"`
	Name      string
	Size      int64
	Type      string
	Sha256    []byte
	UpdatedAt time.Time
	CreatedAt time.Time
	Owner     *model.User `gorm:"foreignKey:OwnerID"`
}

func (File) TableName() string {
	return "files"
}

type ImageStatus int

const (
	ImageStatusUnspecified ImageStatus = iota
	ImageStatusUploaded
	ImageStatusScanned
)

func (s ImageStatus) Value() (driver.Value, error) {
	switch s {
	case ImageStatusUploaded:
		return "uploaded", nil
	case ImageStatusScanned:
		return "scanned", nil
	default:
		return "", nil
	}
}

func (s *ImageStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for ImageStatus")
	}
	switch v {
	case "uploaded":
		*s = ImageStatusUploaded
	case "scanned":
		*s = ImageStatusScanned
	default:
		*s = ImageStatusUnspecified
	}
	return nil
}

type OCRResponse struct {
	Msg     string         `json:"msg"`
	Results [][]OCRResults `json:"results"`
	Status  string         `json:"status"`
}
type OCRResults struct {
	Confidence float64 `json:"confidence"`
	Text       string  `json:"text"`
	TextRegion [][]int `json:"text_region"`
}
