package generated

// This code generated by go generate.
// DO NOT EDIT BY HAND!
import (
	"html/template"
	"strings"
)

type TranslationEn struct {
	TemplateFuncMap template.FuncMap
}
func (l *TranslationEn) Welcome(name string) string {
	t := template.New("")
	if l.TemplateFuncMap != nil {
		t = t.Funcs(l.TemplateFuncMap)
	}
	t, err := t.Parse("Welcome {{.name | ToUpper}}!")
	if err != nil {
		return "welcome"
	}
	var b strings.Builder
	err = t.Execute(&b, map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return "welcome"
	}
	return b.String()
}