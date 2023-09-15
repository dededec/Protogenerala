package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const NumberRegexp = `\d`
const EndInput = "x"

func gameLoop() {
	var currentStatus GameStatus = GameStatus{score: 0, deck: newDeck(), turn: 0}

	currentStatus.shuffleDeck()
	currentStatus.dealHand()

	var input string
	for i := 0; i < MaxTurns; i++ {
		currentStatus.startTurn()
		currentStatus.calculateScore()
		currentStatus.Print()
		fmt.Scan(&input)

		for strings.ToLower(input) != EndInput {
			if regexp.MustCompile(NumberRegexp).MatchString(input) {
				index, _ := strconv.Atoi(input)
				currentStatus.toggleCardChoice(index)
			} else {
				println(WrongInputMsg)
			}

			currentStatus.Print()
			fmt.Scan(&input)
		}

		currentStatus.changeCards()
	}

	currentStatus.endGame()
}
