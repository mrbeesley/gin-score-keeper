package controllers

import (
	"fmt"
	"strconv"
	"strings"

	models "github.com/mrbeesley/gin_score_keeper/models"
)

type GameManager struct {
	CurrentGame  *models.Game
	CurrentStep  int
	MoveStep     bool
	StepCounter  int
	NextPlayerID int
	MoveBack     bool
}

func NewGameManager() *GameManager {
	return &GameManager{
		CurrentStep:  0,
		StepCounter:  0,
		MoveStep:     false,
		NextPlayerID: 1,
		CurrentGame:  models.NewGame(),
		MoveBack:     false,
	}
}

func ManageGame(gm *GameManager, msg string) (bool, string) {
	retMSG := ""
	readOnly := false

	if strings.ToLower(msg) == "exit" {
		return true, "break"
	}

	switch gm.CurrentStep {
	case 0:
		readOnly, retMSG = gm.StartGame(msg)
	case 1:
		readOnly, retMSG = gm.GetPlayers(msg)
	case 2:
		readOnly, retMSG = gm.KeepScore(msg)
	case 3:
		readOnly, retMSG = gm.GetScores(msg)
	case 4:
		readOnly, retMSG = gm.EndGame(msg)
	default:
		retMSG = "break"
	}

	if gm.MoveStep {
		gm.CurrentStep++
		gm.MoveStep = false
	} else if gm.MoveBack {
		gm.CurrentStep--
		gm.MoveBack = false
	}
	return readOnly, retMSG

}

func (gm *GameManager) StartGame(msg string) (bool, string) {
	msg1 := "Lets get a Gin 13 game started!"
	msg2 := "Now we need to know who is playing."
	gm.StepCounter = 0
	gm.MoveStep = true
	return true, fmt.Sprint(msg1, "\n", msg2)
}

func (gm *GameManager) GetPlayers(msg string) (bool, string) {
	retMSG := ""
	altMSG := ""
	readOnly := false

	if gm.StepCounter != 0 && (msg == "done" || msg == "") {
		if len(gm.CurrentGame.Players) > 1 {
			gm.StepCounter = 0
			gm.MoveStep = true
			retMSG = "All players added, Lets go!"
			readOnly = true
		} else {
			altMSG = "You must enter at least 2 players \n"
		}
	} else if gm.StepCounter == 0 {
		gm.StepCounter++
		retMSG = fmt.Sprint("Enter the name of Player ", gm.NextPlayerID, " or type done to finish adding players. (then press enter)")
	} else {
		plr := models.NewPlayer(gm.NextPlayerID, msg)
		gm.CurrentGame.Players = append(gm.CurrentGame.Players, plr)
		gm.NextPlayerID++
		retMSG = fmt.Sprint("Enter the name of Player ", gm.NextPlayerID, " or type done to finish adding players. (then press enter)")
	}
	return readOnly, fmt.Sprint(altMSG, retMSG)
}

func (gm *GameManager) KeepScore(msg string) (bool, string) {
	round := gm.CurrentGame.Round
	wildCard := gm.CurrentGame.Wilds[round]

	roundInfo := fmt.Sprint("Round: ", round, " \n", "Wild Card: ", wildCard)
	scoreInfo := gm.CurrentGame.BuildScoreCard()
	gm.MoveStep = true

	if gm.CurrentGame.GameOver {
		return true, scoreInfo
	}

	return true, fmt.Sprint(roundInfo, "\n", scoreInfo)
}

func (gm *GameManager) GetScores(msg string) (bool, string) {
	retMSG := ""
	readOnly := false
	hasValidScore := false
	playersCount := len(gm.CurrentGame.Players)

	if gm.CurrentGame.GameOver {
		gm.StepCounter = 0
		gm.MoveStep = true
		return true, ""
	}

	if gm.StepCounter > 0 && gm.StepCounter <= playersCount {
		score, err := strconv.Atoi(msg)
		if err == nil {
			gm.CurrentGame.Players[gm.StepCounter-1].AddScores(score, gm.CurrentGame.Round)
			hasValidScore = true
		}
	}

	if gm.StepCounter == 0 {
		retMSG = fmt.Sprint("Enter the score for ", gm.CurrentGame.Players[gm.StepCounter].Name, ": \n")
		gm.StepCounter++
	} else if gm.StepCounter >= playersCount {
		gm.MoveBack = true
		gm.StepCounter = 0
		gm.CurrentGame.CloseRound()
		readOnly = true
		retMSG = "All scores Entered, Starting next round. Good luck!"
	} else if hasValidScore {
		retMSG = fmt.Sprint("Enter the score for ", gm.CurrentGame.Players[gm.StepCounter].Name, ": \n")
		gm.StepCounter++
	} else {
		retMSG = "Please enter a valid score \n"
	}

	return readOnly, retMSG
}

func (gm *GameManager) EndGame(msg string) (bool, string) {
	winner := gm.CurrentGame.GetWinner()
	gm.MoveStep = true
	return false, fmt.Sprint("Game Over! The winner was: ", winner)
}
