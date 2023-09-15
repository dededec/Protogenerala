package main

import "strconv"

type Card struct {
	value int
	suit  string
}

func (c Card) String() string {
	return strconv.Itoa(c.value) + "-" + c.suit
}
