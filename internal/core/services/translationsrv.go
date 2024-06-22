package services

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("error on translate in the gateway: %v", err.Error())
	}

	id, err := s.translationRepo.SaveTranslation(ctx, result)
	if err != nil {
		return nil, fmt.Errorf("error on save translation on database: %v", err.Error())
	}

	result.ID = id

	return result, nil
}

func (s *service) GetTranslations(ctx context.Context) ([]domain.Translation, error) {
	translations, err := s.translationRepo.GetTranslations(ctx)
	if err != nil {
		return nil, fmt.Errorf("error on get all translations: %v", err.Error())
	}

	return translations, nil
}
