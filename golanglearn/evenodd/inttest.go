package main

import "fmt"

type inttest []int

func newInt() inttest {
	cards := inttest{}
	cardSuits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, suit := range cardSuits {
		cards = append(cards, suit)
	}
	return cards
}

func (d inttest) print() {
	for i, card := range d {
		if card%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
