package generated

// This code generated by go generate.
// DO NOT EDIT BY HAND!

import "html/template"

type TranslationEn struct {
	TemplateFuncMap template.FuncMap
}
func (l *TranslationEn) ConfirmationSentToEmail(email string) string {
	return TranslationEnMap.MustGetTemplated(
		"confirmationSentToEmail",
		map[string]interface{}{
			"email": email,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) WelcomeMessage(name string, appName string) string {
	return TranslationEnMap.MustGetTemplated(
		"welcomeMessage",
		map[string]interface{}{
			"Name": name,
			"appName": appName,
		},
		l.TemplateFuncMap,
	)
}
func (l *TranslationEn) YourOTPIs(otp string) string {
	return TranslationEnMap.MustGetTemplated(
		"yourOTPIs",
		map[string]interface{}{
			"OTP": otp,
		},
		l.TemplateFuncMap,
	)
}

var TranslationEnMap = LangMap{
	"confirmationSentToEmail": "Confirmation sent to {{.email}}",
	"welcomeMessage": "Dear {{.Name}}, welcome to {{.appName}}",
	"yourOTPIs": "Your OTP is {{.OTP}}",
}
