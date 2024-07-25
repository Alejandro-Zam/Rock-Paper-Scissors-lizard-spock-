package logic

import "math/rand"

const (
	ROCK    = 0 // Piedra aplasta a lagarto y a tijera
	PAPER   = 1 // Papel envuelve piedra y desautoriza a spock
	TIJERA  = 2 // Tijera corta papel y decapita lagarto
	LAGARTO = 3 // Lagarto devora papel y envenena a Spock
	SPOCK   = 4 // Spock rompe tijera y vaporiza piedra
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bazzinga",
	"Our babies will be smart and beautiful",
	"It's a Saturn V rocket! I built it",
	"That's my spot",
	"I'm not crazy. My mother had me tested",
}

var loseMessages = []string{
	"Why do you hate me?",
	"I'm never going to catch a break, am I?",
	"Please, don't hurt my friend.",
	"I can't talk to women.",
}

var drawMessages = []string{
	"Mmm...",
	":c",
	"Here we go again",
}

var computerScore, playerScore int

func PlayRound(playerValue int) Round {

	computerValue := rand.Intn(5)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue{
		case ROCK:
			computerChoiceInt = ROCK
			computerChoice = "ChatGPT eligió PIEDRA"
		case PAPER:
			computerChoiceInt = PAPER
			computerChoice = "ChatGPT eligió PAPEL"
		case TIJERA:
			computerChoiceInt = TIJERA
			computerChoice = "ChatGPT eligió TIJERA"
		case LAGARTO:
			computerChoiceInt = LAGARTO
			computerChoice = "ChatGPT eligió LAGARTO"
		case SPOCK:
			computerChoiceInt = SPOCK
			computerChoice = "ChatGPT eligió SPOCK"
	}

	//Mensaje aleatorio
	messageInt := rand.Intn(5)

	var message string

	if playerValue == computerValue {
		roundResult = "Empate"

		message = drawMessages[messageInt]
	} else {
		message, roundResult = determinarGanador(computerValue, playerValue)
	}


}


func determinarGanador(computerValue, playerValue int)(string, string){

	var message, roundResult string
	//Mensaje aleatorio
	messageInt := rand.Intn(5)
	


	switch playerValue{
		case ROCK:
			if computerValue == LAGARTO || computerValue == TIJERA{
				playerScore++
				roundResult = "Jugador Gana"
				message = winMessages[messageInt]		
			}else{
				computerScore++
				roundResult = "ChatGPT Gana"
				message = loseMessages[messageInt]
			}
		case PAPER:
			if computerValue == ROCK || computerValue == SPOCK{
				playerScore++
				roundResult = "Jugador Gana"
				message = winMessages[messageInt]		
			}else{
				computerScore++
				roundResult = "ChatGPT Gana"
				message = loseMessages[messageInt]
			}
		case TIJERA:
			if computerValue == PAPER || computerValue == LAGARTO{
				playerScore++
				roundResult = "Jugador Gana"
				message = winMessages[messageInt]		
			}else{
				computerScore++
				roundResult = "ChatGPT Gana"
				message = loseMessages[messageInt]
			}
		case LAGARTO:
			if computerValue == PAPER || computerValue == SPOCK{
				playerScore++
				roundResult = "Jugador Gana"
				message = winMessages[messageInt]		
			}else{
				computerScore++
				roundResult = "ChatGPT Gana"
				message = loseMessages[messageInt]
			}
		case SPOCK:
			if computerValue == TIJERA || computerValue == ROCK{
				playerScore++
				roundResult = "Jugador Gana"
				message = winMessages[messageInt]		
			
			}else{
				computerScore++
				roundResult = "ChatGPT Gana"
				message = loseMessages[messageInt]
			}
	}

	return message, roundResult
}