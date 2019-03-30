package main

import "fmt"

const englishPrefix = "Hello, "

// Hello returns a "Hello, world" string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
