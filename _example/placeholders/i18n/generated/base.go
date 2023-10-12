package generated

// This code generated by go generate.
// DO NOT EDIT BY HAND!
import "html/template"

type Base interface {
	// Confirmation sent to message
	//
	// In en, this message translates to: Confirmation sent to {{.email}}
	ConfirmationSentToEmail(email string) string
}

// Load loads the translation for the given language.
// Defaults to base language if not found.
func Load(lang string, funcMap template.FuncMap) Base {
	switch lang {
	case "en":
		return &TranslationEn{TemplateFuncMap: funcMap}
	case "ms":
		return &TranslationMs{TemplateFuncMap: funcMap}
	}
	return &TranslationEn{TemplateFuncMap: funcMap}
}