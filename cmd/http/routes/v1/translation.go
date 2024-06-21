package v1

import (
	"leonardodelira/go-clean-template/cmd/http/handlers/translationhdl"
	"leonardodelira/go-clean-template/dependencies"

	"github.com/gin-gonic/gin"
)

func NewRouterTranslations(handler *gin.RouterGroup) {
	hdl := translationhdl.NewTranslationHandler(dependencies.TranslationService)

	h := handler.Group("/translation")
	{
		h.GET("/", hdl.GetTranslations)
		h.POST("/", hdl.DoTranslation)
	}
}
