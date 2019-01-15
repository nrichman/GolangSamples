package main

func main() {
	cards := newDeck()
	cards.saveToFile("mycards")
	cards.shuffle()
	cards.print()
}
