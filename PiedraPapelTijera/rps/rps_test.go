package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("Player chose PIEDRA")
		case 1:
			fmt.Println("Player chose PAPEL")
		case 2:
			fmt.Println("Player chose TIJERAS")
		}

		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("Computer chose: %s\n", round.ComputerChoice)
		fmt.Printf("Round result: %s\n", round.RoundResult)
		fmt.Printf("Computer choice int: %d\n", round.ComputerChoiceInt)
		fmt.Printf("Scores - Player: %s, Computer: %s\n", round.PlayerScore, round.ComputerScore)
		fmt.Println("-----")

	}
}
