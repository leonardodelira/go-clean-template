package translationhdl

import (
	"fmt"
	"leonardodelira/go-clean-template/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Text                string `json:"text" required:"true"`
	LanguageDestination string `json:"language_destination" required:"true"`
}

func ParseRequest(c *gin.Context) (domain.TranslationInput, error) {
	body := requestBody{}
	if err := c.Bind(&body); err != nil {
		return domain.TranslationInput{}, fmt.Errorf("error on bind body")
	}

	return domain.TranslationInput{
		Text:                body.Text,
		LanguageDestination: body.LanguageDestination,
	}, nil
}
