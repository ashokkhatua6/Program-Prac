package main

func main() {
	//v)ar card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card := "Five of Diamonds"
	//card = "Ashok"
	//card := newCard()

	//fmt.Println(card)

	//slice
	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	cards := newDeck()

	//for i, card := range cards {
	//fmt.Println(i, card)
	// }
	//fmt.Println(cards)
	//cards.print()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

//func newCard() string {
//return "Five of Diamonds"
//}
