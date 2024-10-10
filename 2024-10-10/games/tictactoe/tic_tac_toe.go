package ticTacToe

import (
	"errors"
)

var gameId = 0

type TicTacToeGame struct {
	gameId      int
	players     [2]*Player
	board       *Board
	isComplete  bool
	winnerIndex int
}

func NewTicTacToeGame(player1Name, player2Name string) (*TicTacToeGame, error) {
	if player1Name == player2Name {

	}
	var players [2]*Player
	tempPlayer1, err := NewPlayer(player1Name)
	if err != nil {
		return nil, err
	}
	tempPlayer2, err1 := NewPlayer(player2Name)
	if err1 != nil {
		return nil, err1
	}
	players[0] = tempPlayer1
	players[1] = tempPlayer2
	board := NewBoard()

	tempTicTacToeGameObject := &TicTacToeGame{
		gameId:      gameId,
		players:     players,
		board:       board,
		isComplete:  false,
		winnerIndex: -1,
	}
	gameId++
	return tempTicTacToeGameObject, nil
}

func (tic *TicTacToeGame) Play(parameter ...interface{}) (bool, error) {
	cellNumber, errx := getCellNumberFromParameter(parameter)
	if errx != nil {
		return false, errx
	}

	err := tic.validateIfOngoing()
	if err != nil {
		return false, err
	}
	err1 := tic.board.PlayMove(cellNumber)
	if err1 != nil {
		return false, err1
	}
	isWinner, winnerNumber := tic.board.CheckForWinner()
	if isWinner {
		tic.winnerIndex = winnerNumber
		tic.setGameAsComplete()
		return true, nil
	}
	if tic.board.IsBoardFull() {
		tic.setGameAsComplete()
		return true, nil
	}

	return false, nil
}

func (tic *TicTacToeGame) GetCurrentBoard() string {
	return tic.board.getCurrentBoard()

}

func (tic *TicTacToeGame) GetResult() (string, error) {
	err := tic.validateIfCompleted()
	if err != nil {
		return "", err
	}
	finalVerdict := "The Game is A Draw"
	if tic.winnerIndex != -1 {
		finalVerdict = tic.players[tic.winnerIndex].getName() + " won the Game"
	}
	return finalVerdict, nil

}

func (tic *TicTacToeGame) setGameAsComplete() {
	tic.isComplete = true
}

func getCellNumberFromParameter(value []interface{}) (int, error) {
	if len(value) != 1 {
		return -1, errors.New("Please Pass only Cell Number")
	}
	cellNumber, intValidation := value[0].(int)
	if !intValidation {
		return -1, errors.New("Please pass a Valid Cell Number")
	}
	return cellNumber, nil
}

func (tic *TicTacToeGame) validateIfOngoing() error {
	if tic.isComplete {
		return errors.New("The Game Is Already Completed")
	}
	return nil
}

func (tic *TicTacToeGame) validateIfCompleted() error {
	if !tic.isComplete {
		return errors.New("The Game Is Still Going On")
	}
	return nil
}
