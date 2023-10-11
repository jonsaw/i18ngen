package generated

// This code generated by go generate.
// DO NOT EDIT BY HAND!

type Base interface {
	// Confirmation sent to message
	//
	// In en, this message translates to: Confirmation sent to {{.email}}
	ConfirmationSentToEmail(email string) string
	// No definition provided for @home
	//
	// In en, this message translates to: Home
	Home() string
}

// Load loads the translation for the given language.
// Defaults to base language if not found.
func Load(lang string) Base {
	switch lang {
	case "en":
		return &TranslationEn{}
	case "ms":
		return &TranslationMs{}
	}
	return &TranslationEn{}
}
