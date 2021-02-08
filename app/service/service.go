package service

import (
	"github.com/naoki/english/app/models"
	"github.com/naoki/english/app/repository"
)

type Phrase interface {
	Create(phrase *models.Phrase) (*models.Phrase, error)
}

type Service struct {
	Phrase
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Phrase: NewPhraseService(repos.Phrase),
	}
}
