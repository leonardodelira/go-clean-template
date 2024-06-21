package domain

type TranslationInput struct {
	Text                string `json:"text"`
	LanguageDestination string `json:"language_destination"`
}

type Translation struct {
	ID                     int32  `json:"id"`
	Text                   string `json:"text"`
	Translation            string `json:"translation"`
	LanguageDestination    string `json:"language_destination"`
	LanguageOriginDetected string `json:"language_origin_detected"`
}
