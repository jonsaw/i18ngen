package generated

// This code generated by go generate.
// DO NOT EDIT BY HAND!
import (
	"html/template"
	"strings"
	"fmt"
)

type Base interface {
	// No definition provided for @email
	//
	// In en, this message translates to: Email
	Email() string
	// Confirmation sent to message
	//
	// In en, this message translates to: {{.label}} already in use
	InvalidAlreadyInUse(label string) string
	// Confirmation sent to message
	//
	// In en, this message translates to: {{.label}} already registered
	InvalidAlreadyRegistered(label string) string
	// Confirmation sent to message
	//
	// In en, this message translates to: {{.label}} must be greater than {{index .settings 0}}
	InvalidGreaterThan(label string, settings []string) string
	// Confirmation sent to message
	//
	// In en, this message translates to: {{.label}} must have min length {{index .settings 0}}
	InvalidMinLength(label string, settings []string) string
	// Confirmation sent to message
	//
	// In en, this message translates to: {{.label}} is required
	InvalidRequired(label string) string
	// No definition provided for @password
	//
	// In en, this message translates to: Password
	Password() string
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

// LoadLangMap loads the translation map for the given language.
// Typlically used to dynamically load translations by key.
// Defaults to base language if not found.
func LoadLangMap(lang string) LangMap {
	switch lang {
	case "en":
		return TranslationEnMap
	case "ms":
		return TranslationMsMap
	}
	return TranslationEnMap
}

type LangMap map[string]string

func (m LangMap) Get(key string) (string, error) {
	if val, ok := m[key]; ok {
		return val, nil
	}

	return "", fmt.Errorf("key %q not found", key)
}

func (m LangMap) MustGet(key string) string {
	val, err := m.Get(key)
	if err != nil {
		return key
	}

	return val
}

func (m LangMap) GetTemplated(key string, values interface{}, templateFunc template.FuncMap) (string, error) {
	val, err := m.Get(key)
	if err != nil {
		return "", err
	}

	t := template.New("")
	if templateFunc != nil {
		t = t.Funcs(templateFunc)
	}
	t, err = t.Parse(val)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	err = t.Execute(&b, values)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (m LangMap) MustGetTemplated(key string, values interface{}, templateFunc template.FuncMap) string {
	val, err := m.GetTemplated(key, values, templateFunc)
	if err != nil {
		return key
	}

	return val
}
