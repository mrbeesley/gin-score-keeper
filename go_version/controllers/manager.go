package controllers

import (
	"fmt"

	models "github.com/mrbeesley/gin_score_keeper/models"
)

type GameManager struct {
	CurrentGame  *models.Game
	CurrentStep  int
	MoveStep     bool
	StepCounter  int
	NextPlayerID int
}

func NewGameManager() *GameManager {
	return &GameManager{
		CurrentStep:  0,
		StepCounter:  0,
		MoveStep:     false,
		NextPlayerID: 1,
		CurrentGame:  models.NewGame(),
	}
}

func ManageGame(gm *GameManager, msg string) string {
	retMSG := ""
	switch gm.CurrentStep {
	case 0:
		retMSG = gm.StartGame(msg)
	case 1:
		retMSG = gm.GetPlayers(msg)
	case 2:
		retMSG = gm.KeepScore(msg)
	default:
		retMSG = "break"
	}

	if gm.MoveStep {
		gm.CurrentStep++
		gm.MoveStep = false
	}
	return retMSG

}

func (gm *GameManager) StartGame(msg string) string {
	msg1 := "Lets get a Gin 13 game started!"
	msg2 := "Now we need to know who is playing... (press enter)"
	gm.StepCounter = 0
	gm.MoveStep = true
	return fmt.Sprint(msg1, "\n", msg2)
}

func (gm *GameManager) GetPlayers(msg string) string {
	retMSG := fmt.Sprint("Enter the name of Player", gm.NextPlayerID, " or type done to finish adding players. (then press enter)")
	altMSG := ""
	if gm.StepCounter != 0 && (msg == "done" || msg == "") {
		if len(gm.CurrentGame.Players) > 1 {
			gm.StepCounter = 0
			gm.MoveStep = true
			retMSG = "All players added, Lets go! (press enter)"
		} else {
			altMSG = "You must enter at least 2 players \n"
		}
	} else if gm.StepCounter == 0 {
		gm.StepCounter++
	} else {
		plr := models.NewPlayer(gm.NextPlayerID, msg)
		gm.CurrentGame.Players = append(gm.CurrentGame.Players, plr)
		gm.NextPlayerID++
	}
	return fmt.Sprint(altMSG, retMSG)
}

func (gm *GameManager) KeepScore(msg string) string {
	round := gm.CurrentGame.Round
	wildCard := gm.CurrentGame.Wilds[round]

	roundInfo := fmt.Sprint("Round: ", round, " \n", "Wild Card: ", wildCard)
	scoreInfo := gm.CurrentGame.BuildScoreCard()

	return fmt.Sprint(roundInfo, "\n", scoreInfo)
}
