package main

const spanish = "Hola"
const englishHelloPrifix = "Hello"

func greating(language string) string {
	switch language {
	case "spanish":
		return spanish
	case "English":
		return englishHelloPrifix
	default:
		return englishHelloPrifix
	}
}
func Hello(name string, language string) string {
	if name == " " {
		name = "World"
	}
	return greating(language) + ", " + name

}

func main() {

}
