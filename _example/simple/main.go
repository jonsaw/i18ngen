package main

import (
	"fmt"

	i18ngen "github.com/jonsaw/i18ngen/_example/simple/i18n/generated"
)

func main() {
	fmt.Println(i18ngen.Load("en", nil).OurStory())
	fmt.Println(i18ngen.Load("ms", nil).OurStory())
}
