package logic

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 5; i++ {
		round := PlayRound(i)
		fmt.Println(i)
		switch i {
		case 0:
			fmt.Println("PIEDRA")

		case 1:
			fmt.Println("PAPEL")

		case 2:
			fmt.Println("TIJERA")

		case 3:
			fmt.Println("LAGARTO")

		case 4:
			fmt.Println("SPOCK")

		}
		fmt.Printf("Message: %s \n", round.Message)
		fmt.Printf("ComputerChoice: %s \n", round.ComputerChoice)
		fmt.Printf("RoundResult: %s \n", round.RoundResult)
		fmt.Printf("ComputerChoiceInt: %d \n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s \n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s \n", round.PlayerScore)

		fmt.Println("----------------------------")
	}

}
