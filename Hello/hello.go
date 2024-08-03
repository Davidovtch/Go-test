package hello

import "fmt"

const (
	englishP string = "Hello, "
	spanishP string = "Hola, "
	frenchP  string = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	// without switch
	// if lang == "ES" {
	// 	return spanishP + name
	// } else if lang == "FR" {
	// 	return frenchP + name
	// }
	// return englishP + name
	return prefix(lang) + name
}

func prefix(lang string) (pr string) {
	switch lang {
	case "ES":
		pr = spanishP
	case "FR":
		pr = frenchP
	default:
		pr = englishP
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
