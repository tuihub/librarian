package page

import "gorm.io/gorm"

type Builder struct {
	db *gorm.DB
}

func NewBuilder(db *gorm.DB) *Builder {
	return &Builder{db: db}
}
