package main

import (
	"fmt"
)

//Creat a new type of 'deck'
//which is a slice of strings

type deck []string //type deck ==== []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardvalues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardvalues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	//fmt.Println(d)
	return d[:handSize], d[handSize:]

}
