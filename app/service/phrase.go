package service

import (
	"github.com/naoki/english/app/models"
	"github.com/naoki/english/app/repository"
)

type PhraseService struct {
	repo repository.Phrase
}

func NewPhraseService(repo repository.Phrase) *PhraseService {
	return &PhraseService{repo: repo}
}

func (s *PhraseService) Create(pharse *models.Phrase) (*models.Phrase, error) {
	eg_phrase, err := s.repo.Create(pharse)
	if err != nil {
		return nil, err
	}

	return eg_phrase, nil
}
