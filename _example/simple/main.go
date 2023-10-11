package main

import (
	"fmt"

	i18ngen "github.com/jonsaw/i18ngen/_example/simple/i18n/generated"
)

func main() {
	fmt.Println(i18ngen.Load("en").ConfirmationSentToEmail("email@email.com"))
	fmt.Println(i18ngen.Load("ms").ConfirmationSentToEmail("email@email.com"))
}
