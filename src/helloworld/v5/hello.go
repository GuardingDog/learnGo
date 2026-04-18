package main

const enginlishHelloPrefix = "Hello, World! "

func Hello(name string) string {
	if name == "" {
		name = "Bob"
	}
	return enginlishHelloPrefix + name
}
