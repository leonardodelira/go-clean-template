package ports

import (
	"context"
	"leonardodelira/go-clean-template/internal/core/domain"
)

type TranslationRepository interface {
	GetTranslations(ctx context.Context) ([]domain.Translation, error)
	SaveTranslation(ctx context.Context, translation *domain.Translation) (int32, error)
}
