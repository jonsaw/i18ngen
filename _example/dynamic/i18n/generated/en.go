package generated

// This code generated by go generate.
// DO NOT EDIT BY HAND!

import "html/template"

type TranslationEn struct {
	TemplateFuncMap template.FuncMap
}
func (l *TranslationEn) Email() string {
	return TranslationEnMap.MustGet("email")
}
func (l *TranslationEn) InvalidAlreadyInUse(label string) string {
	return TranslationEnMap.MustGetTemplated(
		"invalidAlreadyInUse",
		map[string]interface{}{
			"label": label,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) InvalidAlreadyRegistered(label string) string {
	return TranslationEnMap.MustGetTemplated(
		"invalidAlreadyRegistered",
		map[string]interface{}{
			"label": label,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) InvalidGreaterThan(label string, settings []string) string {
	return TranslationEnMap.MustGetTemplated(
		"invalidGreaterThan",
		map[string]interface{}{
			"label": label,
			"settings": settings,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) InvalidMinLength(label string, settings []string) string {
	return TranslationEnMap.MustGetTemplated(
		"invalidMinLength",
		map[string]interface{}{
			"label": label,
			"settings": settings,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) InvalidRequired(label string) string {
	return TranslationEnMap.MustGetTemplated(
		"invalidRequired",
		map[string]interface{}{
			"label": label,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) Password() string {
	return TranslationEnMap.MustGet("password")
}

var TranslationEnMap = LangMap{
	"email": "Email",
	"invalidAlreadyInUse": "{{.label}} already in use",
	"invalidAlreadyRegistered": "{{.label}} already registered",
	"invalidGreaterThan": "{{.label}} must be greater than {{index .settings 0}}",
	"invalidMinLength": "{{.label}} must have min length {{index .settings 0}}",
	"invalidRequired": "{{.label}} is required",
	"password": "Password",
}