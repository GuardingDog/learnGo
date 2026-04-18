package main

import "fmt"

func Hello(name string) string{
	return fmt.Sprintf("Hello, World! %s", name)
}

func main() {
	fmt.Println(Hello("Qingzhi"))
}