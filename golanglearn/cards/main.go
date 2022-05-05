package main

func main() {
	// cards := newDeck()
	// cards.print()
	// c := color("Red")
	// fmt.Println(c.describe("is an awesome color"))

	// hand, remaingCard := deal(cards, 2)

	// hand.print()
	// fmt.Println("Baru")
	// remaingCard.print()
	// fmt.Println(stringVal)
	// stringVal := cards.toString()

	// cards.saveToFile("testsimpan")
	cards := newDeckFromFile("testsimpan")
	// cards := newDeckFromFile("testsasasimpan")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "hello guys"
}
