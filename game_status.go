package main

import (
	"strconv"
)

type GameStatus struct {
	score    int
	hand     Hand
	deck     Deck
	isChosen [HandSize]bool
	turn     int
}

func (gs *GameStatus) shuffleDeck() {
	println(ShuffleMsg)
	gs.deck.shuffle()
}

func (gs *GameStatus) dealHand() {
	gs.hand, gs.deck = gs.deck.deal(HandSize)
}

func (gs *GameStatus) calculateScore() {
	gs.score = 0
	for _, card := range gs.hand {
		gs.score += card.value
	}
}

func (gs *GameStatus) changeCards() {
	for i, card := range gs.hand {
		if gs.isChosen[i] {
			gs.hand[i] = gs.deck[0]
			gs.deck = append(gs.deck, card)
			gs.shuffleDeck()
		}
	}
}

func (gs *GameStatus) toggleCardChoice(index int) {
	if index >= 0 && index < HandSize {
		gs.isChosen[index] = !gs.isChosen[index]
	}
}

func (gs *GameStatus) startTurn() {
	gs.resetChosenCards()
	gs.turn++
	println(TurnMsg + strconv.Itoa(gs.turn) + "/" + strconv.Itoa(MaxTurns) + EndTitleDelimiter)
	println(HandMsg)
}

func (gs *GameStatus) endGame() {
	gs.resetChosenCards()
	gs.calculateScore()
	println(EndGameMsg)
	gs.Print()
}

func (gs *GameStatus) resetChosenCards() {
	for i := range gs.isChosen {
		gs.isChosen[i] = false
	}
}

func (gs GameStatus) Print() {
	for i, card := range gs.hand {
		if gs.isChosen[i] {
			print("X - ")
		}
		println(strconv.Itoa(i) + ": " + card.String())
	}
	println(ScoreMsg + strconv.Itoa(gs.score))
}
