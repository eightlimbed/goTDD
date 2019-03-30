package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const swahiliPrefix = "Hujambo, "
const spanish = "Spanish"
const french = "French"
const swahili = "Swahili"

// Hello returns a "Hello, world" string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

// greetingPrefix takes a language and returns it's prefix for Hello()
// This function has a named return value ("prefix")
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	case swahili:
		prefix = swahiliPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
