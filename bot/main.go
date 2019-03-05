package main

import "fmt"

type (
	englishBot struct{}
	spanishBot struct{}

	bot interface {
		getGreeting() string
	}
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
