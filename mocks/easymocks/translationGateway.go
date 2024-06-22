package easymocks

import "leonardodelira/go-clean-template/internal/core/domain"

func TranslationGatewayMock() *domain.Translation {
	return &domain.Translation{
		ID:                     1,
		Text:                   "Hello World",
		Translation:            "Olá Mundo",
		LanguageDestination:    "PT",
		LanguageOriginDetected: "EN",
	}
}
