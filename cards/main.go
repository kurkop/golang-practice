package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// remainingCards.print()
}
