package main

import (
	"fmt"

	i18ngen "github.com/jonsaw/i18ngen/_example/placeholders/i18n/generated"
)

func main() {
	fmt.Println(i18ngen.Load("en", nil).ConfirmationSentToEmail("email@email.com"))
	fmt.Println(i18ngen.Load("ms", nil).ConfirmationSentToEmail("email@email.com"))
	fmt.Println(i18ngen.Load("en", nil).WelcomeMessage("John Doe", "My App"))
	fmt.Println(i18ngen.Load("ms", nil).WelcomeMessage("John Doe", "My App"))
}
