package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/naoki/english/app/models"
)

type PhrasePostgres struct {
	db *gorm.DB
}

func NewPharse(db *gorm.DB) *PhrasePostgres {
	return &PhrasePostgres{db: db}
}

func (r *PhrasePostgres) Create(pharse *models.Phrase) (*models.Phrase, error) {
	tx := r.db.Begin()

	if !tx.NewRecord(pharse) {
		tx.Rollback()
		return nil, errors.New("errors record")
	}

	if err := tx.Debug().Create(pharse).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return pharse, tx.Commit().Error
}
