package dependencies

import (
	"leonardodelira/go-clean-template/internal/core/ports"
)

var (
	TranslationService ports.TranslationService
)

var (
	translationRepository ports.TranslationRepository
	translatorGateway     ports.TranslatorGateway
)

func Init() {
	initGateways()
	initRepositories()
	initServices()
}
