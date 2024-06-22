package ports

import (
	"context"
	"leonardodelira/go-clean-template/internal/core/domain"
)

type TranslationService interface {
	GetTranslations(ctx context.Context) ([]domain.Translation, error)
	DoTranslation(ctx context.Context, input domain.TranslationInput) (*domain.Translation, error)
}
