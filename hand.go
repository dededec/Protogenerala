package main

import (
	"strconv"
	"strings"
)

type Hand []Card

func (h Hand) String() string {
	result := ""
	for i, card := range h {
		result += strconv.Itoa(i) + ": " + card.String() + "\n"
	}
	return result
}

func (h Hand) Serialize() []byte {
	result := ""
	for _, card := range h {
		result += card.String() + ","
	}
	result = strings.TrimSuffix(result, ",")
	return []byte(result)
}

func (h Hand) Deserialize(b []byte) {
	var cards []string = strings.Split(string(b), ",")

	for _, card := range cards {
		args := strings.Split(card, "-")
		value, _ := strconv.Atoi(args[0])
		h = append(h, Card{value: value, suit: args[1]})
	}
}
