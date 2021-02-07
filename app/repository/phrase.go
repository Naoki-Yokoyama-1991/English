package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/naoki/english/app/models"
)

type PhrasePostgres struct {
	db *gorm.DB
}

func NewPharse(db *gorm.DB) *PhrasePostgres {
	return &PhrasePostgres{db: db}
}

func (r *PhrasePostgres) Create(pharse *models.Phrase) {

}
