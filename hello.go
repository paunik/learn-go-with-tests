package main

import "fmt"

const german = "German"
const serbian = "Serbian"
const englishHelloPrefix = "Hello, "
const serbianHelloPrefix = "Pomaze Bog, "
const germanHelloPrefix = "Servus, "

// Hello returns hello world string
func Hello(name, language string) string {
	if "" == name {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case serbian:
		prefix = serbianHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("From Main", "English"))
}
