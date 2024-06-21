package ports

import (
	"context"
	"leonardodelira/go-clean-template/internal/core/domain"
)

type TranslatorGateway interface {
	Translate(ctx context.Context, text string, target_lang string) (*domain.Translation, error)
}
