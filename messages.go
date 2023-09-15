package main

import (
	"fmt"
)

const FullDelimiterLine = "\n---------------------------------------------\n"
const StartTitleDelimiter = "--------------- "
const EndTitleDelimiter = " ---------------"

// Menu messages
const GameTitle = FullDelimiterLine + StartTitleDelimiter + "PROTOGENERALA" + EndTitleDelimiter + FullDelimiterLine
const WelcomeMsg = "\nWelcome to my little game written in Go!\n"
const MenuTitle = StartTitleDelimiter + "MENU" + EndTitleDelimiter
const MenuMsg = "1 - Play\n" +
	"2 - Instructions\n" +
	"3 - Exit"
const InstructionsMsg = "You get a hand of 5 cards.\n" +
	"The objective is to achieve the highest sum of card values.\n" +
	"To do so, during 3 rounds, you can change whichever cards you want for another random one from the deck."
const ExitingMsg = "Exiting..."

// Game Messages
const ShuffleMsg = "Shuffling Deck..."
const TurnMsg = StartTitleDelimiter + " Turn "
const HandMsg = "Here's your current hand. Please choose which cards you wish to change ('X' to finish): "
const ChooseCardMsg = "Card chosen for change"
const ScoreMsg = "Current Score: "
const EndGameMsg = StartTitleDelimiter + " FINAL RESULTS " + EndTitleDelimiter

// Errors
const WrongInputMsg = "Wrong input, please try again!"

func println(s string) {
	fmt.Println(s)
}

func print(s string) {
	fmt.Print(s)
}

func printMenu() {
	println(MenuTitle)
	println(MenuMsg)
}
