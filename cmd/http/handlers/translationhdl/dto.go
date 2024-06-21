package translationhdl

import (
	"fmt"
	"leonardodelira/go-clean-template/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Text                string `json:"text" binding:"required"`
	LanguageDestination string `json:"language_destination" binding:"required"`
}

func ParseRequest(c *gin.Context) (*domain.TranslationInput, error) {
	var body requestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		return nil, fmt.Errorf("error on bind body")
	}

	return &domain.TranslationInput{
		Text:                body.Text,
		LanguageDestination: body.LanguageDestination,
	}, nil
}
