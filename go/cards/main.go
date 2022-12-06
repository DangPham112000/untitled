package main

func main() {
	cards := newDeck()

	hand, remainingCards := cards.deal(7)

	hand.print()
	remainingCards.print()
	cards.print()
}
