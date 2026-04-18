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
	if language == "" {
		prefix = enginlishHelloPrefix
	} else if language == spanish {
		prefix = spanishHelloPrefix
	} else if language == french {
		prefix = frenchHelloPrefix
	} else {
		prefix = enginlishHelloPrefix
	}

	return prefix + name
}
