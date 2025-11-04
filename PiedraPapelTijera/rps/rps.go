package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 //Piedra vence a tijera. (tijera + 1) % 3 = 0
	PAPER    = 1 //Papel vence a piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 //Tijera vence a papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessage = []string{
	"¡Bien echo!",
	"¡Buen trabajo!",
	"Que suertetienes.",
}

var loseMessage = []string{
	"¡Que lástima!",
	"¡Sigue intentando!",
	"Hoy no es tu día.",
}

var drawMessage = []string{
	"Los grandes piensan igual.",
	"Oh no. Sigue intentando.",
	"Nadie gana, suerte para la proxima.",
}

var ComputureScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var commputerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		commputerChoice = "La computadora eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		commputerChoice = "La computadora eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		commputerChoice = "La computadora eligió TIJERA"
	}

	messageInt := rand.Intn(3)
	var message string

	switch playerValue {
	case computerValue:
		roundResult = "Es un empate"

		message = drawMessage[messageInt]
	case (computerValue + 1) % 3:
		PlayerScore++
		roundResult = "¡El Juagador gana!"
		message = winMessage[messageInt]
	default:
		ComputureScore++
		roundResult = "¡La computadora gana!"
		message = loseMessage[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    commputerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputureScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
