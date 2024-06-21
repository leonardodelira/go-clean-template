package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"leonardodelira/go-clean-template/internal/core/domain"
	"leonardodelira/go-clean-template/internal/core/ports"
	"net/http"
	"os"
	"time"
)

type deeplTranslatorRepo struct {
	basePath string
	apiKey   string
}

// todo: colocar essas structs em outro arquivo
type deeplEntity struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

type outputDeepl struct {
	Translations []deeplEntity
}

type inputDeepl struct {
	Text        []string `json:"text"`
	Target_lang string   `json:"target_lang"`
}

func NewTranslatorGatewayDeepl() ports.TranslatorGateway {
	apiKeyDeepl := os.Getenv("DEEPL_KEY")

	return &deeplTranslatorRepo{
		basePath: "https://api-free.deepl.com",
		apiKey:   apiKeyDeepl,
	}
}

func (d *deeplTranslatorRepo) Translate(ctx context.Context, text string, targetLang string) (*domain.Translation, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	url := d.basePath + "/v2/translate"

	input := &inputDeepl{
		Text:        []string{text},
		Target_lang: targetLang,
	}

	inputMarshalled, _ := json.Marshal(input)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(inputMarshalled))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", d.apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("TranslatorExternalDeepl - Translate - GET: %w", err)
	}

	if resp.StatusCode != 200 {
		//todo: check others status code
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("TranslatorExternalDeepl - Translate - ReadAll: %w", err)
	}
	defer resp.Body.Close()

	var deeplResult outputDeepl
	err = json.Unmarshal(body, &deeplResult)
	if err != nil {
		return nil, fmt.Errorf("TranslatorExternalDeepl - Translate - Unmarshal: %w", err)
	}

	result := externalToDomain(deeplResult, text)
	result.LanguageDestination = targetLang

	return result, nil
}

func externalToDomain(deeplResult outputDeepl, text string) *domain.Translation {
	return &domain.Translation{
		Text:                   text,
		Translation:            deeplResult.Translations[0].Text,
		LanguageOriginDetected: deeplResult.Translations[0].DetectedSourceLanguage,
	}
}
