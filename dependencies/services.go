package dependencies

import "leonardodelira/go-clean-template/internal/core/services"

func initServices() {
	TranslationService = services.NewTranslationService(translationRepository, translatorGateway)
}
