package main

const enginlishHelloPrefix = "Hello, World! "
const spanishHelloPrefix = "Hola, Mundo! "
const frenchHelloPrefix = "Bonjour, Monde! "

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "Bob"
	}
	prefix := ""
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = enginlishHelloPrefix
	}
	return prefix + name
}
