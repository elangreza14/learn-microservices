package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of 'deck'
// which is slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamond", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, "-", card)
	}
}

// fmt.Println(cards[0])   // dibaca dari index 0
// fmt.Println(cards[0:3])  // dibaca dari index 0,1,2
// fmt.Println(cards[:3])  // dibaca dari index 0,1,2
// fmt.Println(cards[4:7]) // dibaca dari index 4,5,6
// fmt.Println(cards[12:]) // dibaca dari index 12,13,14,15

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// opt 1 log err and create new deck
		// opt 2 log err and quit program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	start := time.Now()                        // generate new time setup
	source := rand.NewSource(start.UnixNano()) // generate int 64 and put in source
	r := rand.New(source)                      // generate new random source
	for i := range d {
		newPosition := r.Intn(len(d) - 1) // generate new random
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// type color string

// func (c color) describe(description string) string {
// 	return string(c) + " " + description
// }
