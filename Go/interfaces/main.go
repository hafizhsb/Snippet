package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spainBot struct{}

func main() {
	eb := englishBot{}
	sb := spainBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spainBot) getGreeting() string {
	return "Hola!"
}
