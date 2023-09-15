package main

import (
	"fmt"
)

func main() {

	print(GameTitle)
	println(WelcomeMsg)

mainLoop:
	for {
		printMenu()
		var input string
		fmt.Scan(&input)
		switch input {
		case "1":
			gameLoop()
		case "2":
			println(InstructionsMsg)
		case "3":
			println(ExitingMsg)
			break mainLoop
		default:
			println(WrongInputMsg)
		}
	}
}
