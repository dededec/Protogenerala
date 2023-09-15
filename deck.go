package main

import (
	"math/rand"
	"strconv"
	"strings"
)

type Deck []Card

var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}

func newDeck() Deck {
	cards := Deck{}

	for _, suit := range cardSuits {
		for i := 1; i <= 13; i++ {
			cards = append(cards, Card{value: i, suit: suit})
		}
	}

	return cards
}

func (d Deck) deal(num int) (Hand, Deck) {
	if num < 0 {
		num = 0
	}
	return Hand(d[:num]), d[num:]
}

func (d Deck) shuffle() { // Reminder: Slice is a ref. type, so no need for pointer receiver
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d Deck) String() string {
	result := ""
	for _, card := range d {
		result += card.String() + ","
	}
	result = strings.TrimSuffix(result, ",")
	return result
}

func (d Deck) Serialize() []byte {
	return []byte(d.String())
}

func (d Deck) Deserialize(b []byte) Serializer {
	var cards []string = strings.Split(string(b), ",")

	for _, card := range cards {
		args := strings.Split(card, "-")
		value, _ := strconv.Atoi(args[0])
		d = append(d, Card{value: value, suit: args[1]})
	}

	return d
}
