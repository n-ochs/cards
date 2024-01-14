package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // append returns a new slice. doesn't modify the original
		}
	}

	return cards
}

// Receiver function. Any variable that is type 'deck' has access to this method
// Usually call the variable a one or two letter variable
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// returns the portion of the slice
	// syntax [startingFrom : upToNotIncluding]. shorthand below
	// if you omit the startingFrom, it starts at 0
	// if you omit the upToNotIncluding, it goes to the end
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("*** [ERROR] ***", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // generates new int64 number every time program starts
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// swaps positions of cards
		// just reassigning the variables
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
