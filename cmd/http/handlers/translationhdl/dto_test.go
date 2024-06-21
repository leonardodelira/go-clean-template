package translationhdl

import (
	"bytes"
	"encoding/json"
	"leonardodelira/go-clean-template/internal/core/domain"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParseRequest(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	requestBody := requestBody{
		Text:                "Hello, World!",
		LanguageDestination: "pt",
	}

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(requestBody)

	c.Request = httptest.NewRequest("POST", "/", &buf)
	c.Request.Header.Set("Content-Type", "application/json")

	translationInput, err := ParseRequest(c)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedTranslationInput := domain.TranslationInput{
		Text:                requestBody.Text,
		LanguageDestination: requestBody.LanguageDestination,
	}

	if expectedTranslationInput != *translationInput {
		t.Errorf("expected translation input %v, got %v", expectedTranslationInput, translationInput)
	}
}
