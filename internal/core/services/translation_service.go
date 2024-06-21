package services

import (
	"context"
	"leonardodelira/go-clean-template/internal/core/domain"
	"leonardodelira/go-clean-template/internal/core/ports"
)

type service struct {
	translationRepo   ports.TranslationRepository
	translatorGateway ports.TranslatorGateway
}

func NewTranslationService(translationRepo ports.TranslationRepository, translatorGateway ports.TranslatorGateway) ports.TranslationService {
	return &service{
		translationRepo:   translationRepo,
		translatorGateway: translatorGateway,
	}
}

func (s *service) DoTranslation(ctx context.Context, input domain.TranslationInput) (*domain.Translation, error) {
	result, err := s.translatorGateway.Translate(ctx, input.Text, input.LanguageDestination)
	if err != nil {
		return nil, err
	}

	id, err := s.translationRepo.SaveTranslation(ctx, result)
	if err != nil {
		return nil, err
	}

	result.ID = &id

	return result, nil
}

func (s *service) GetTranslation(ctx context.Context) ([]domain.Translation, error) {
	translations, err := s.translationRepo.GetTranslations(ctx)
	if err != nil {
		return nil, err
	}

	return translations, nil
}
