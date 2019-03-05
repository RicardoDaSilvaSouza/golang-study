package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

func (d deck) saveToFile(file string) {
	if err := ioutil.WriteFile(file, []byte(d.toString()), 0664); err != nil {
		msg := "Erro ao gravar arquivo: " + err.Error()
		println(msg)
		panic(msg)
	}
}

func newDeckFromFile(file string) deck {
	if bytes, err := ioutil.ReadFile(file); err != nil {
		msg := "Erro ao ler arquivo: " + err.Error()
		println(msg)
		panic(msg)
	} else {
		arch := string(bytes)
		return deck(strings.Split(arch, "\n"))
	}
}

func (d deck) shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := range d {
		newInd := r.Intn(len(d) - 1)
		d[i], d[newInd] = d[newInd], d[i]
	}
}
