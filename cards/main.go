package main

import "fmt"

func main() {
	cards := newDeck()

	hand, _ := deal(cards, 3)

	fmt.Println("Cards on the hand:")
	hand.print()
	//remeaningCards.print()
	saveToFile(hand)

	fmt.Println("New deck:")
	newDeckFromFile().print()
}
