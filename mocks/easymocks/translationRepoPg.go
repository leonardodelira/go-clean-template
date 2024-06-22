package easymocks

import "leonardodelira/go-clean-template/internal/core/domain"

func TranslationRepoPGMock() []domain.Translation {
	return []domain.Translation{
		{
			ID:                     1,
			Text:                   "Hello",
			Translation:            "Olá",
			LanguageDestination:    "PT",
			LanguageOriginDetected: "EN",
		},
	}
}
