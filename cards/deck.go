package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() (cards deck) {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func saveToFile(d deck) {
	if err := ioutil.WriteFile("/Users/ricardo.souza/cards.txt", []byte(d.toString()), 0664); err != nil {
		msg := "Erro ao gravar arquivo: " + err.Error()
		println(msg)
		panic(msg)
	}
}

func newDeckFromFile() deck {
	if bytes, err := ioutil.ReadFile("/Users/ricardo.souza/cards.txt"); err != nil {
		msg := "Erro ao ler arquivo: " + err.Error()
		println(msg)
		panic(msg)
	} else {
		arch := string(bytes)
		return deck(strings.Split(arch, "\n"))
	}
}
