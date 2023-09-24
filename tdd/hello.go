package main

import "fmt"

type LangConstants struct {
	prefix,	suffix string
}

func main() {
	fmt.Println(Hello("world", "en"))
}

func Hello(name string, lang string) string {

	langMap := map[string]LangConstants{
		"en": {"Hello, ", "World"},
		"es": {"Hola, ", "Mundo"},
		"fr": {"Bonjour, ", "Monde"},
	}
	
	c, ok := langMap[lang]

	if !ok {
		 c = langMap["en"]
	}

	if name == "" {
		name = c.suffix
	}

	return c.prefix + name
}