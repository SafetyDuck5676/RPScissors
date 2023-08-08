package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	playerChoice := getPlayerChoice()
	computerChoice := getComputerChoice()

	fmt.Println("Player chose:", playerChoice)
	fmt.Println("Computer chose:", computerChoice)

	winner := getWinner(playerChoice, computerChoice)
	fmt.Println("Winner:", winner)
}

func getPlayerChoice() string {
	var choice string
	fmt.Print("Enter your choice (rock, paper, or scissors): ")
	fmt.Scanln(&choice)
	return choice
}

func getComputerChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(choices))
	return choices[randomIndex]
}

func getWinner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "Draw!"
	}

	switch playerChoice {
	case "rock":
		if computerChoice == "scissors" {
			return "Player"
		} else {
			return "Computer"
		}
	case "paper":
		if computerChoice == "rock" {
			return "Player"
		} else {
			return "Computer"
		}
	case "scissors":
		if computerChoice == "paper" {
			return "Player"
		} else {
			return "Computer"
		}
	default:
		return "Invalid choice!"
	}
}
