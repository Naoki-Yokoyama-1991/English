package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/naoki/english/app/models"
)

type Phrase interface {
	Create(phrase *models.Phrase)
}

type Repository struct {
	Phrase
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Phrase: NewPharse(db),
	}
}
