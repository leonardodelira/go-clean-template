package dependencies

import (
	"leonardodelira/go-clean-template/internal/repositories"
	"leonardodelira/go-clean-template/pkg/database"
)

func initRepositories() {
	conn := database.NewPostgresConnection()

	translationRepository = repositories.NewTranslationPGRepository(conn)
}
