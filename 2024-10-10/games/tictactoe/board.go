package ticTacToe

import (
	"errors"
	"strconv"
)

type Board struct {
	symbols     [9]Symbol
	currentMove int
}

func NewBoard() *Board {
	var tempBoard [9]Symbol
	for i := 0; i < 9; i++ {
		tempBoard[i] = NewEmptySymbol()
	}
	return &Board{
		symbols:     tempBoard,
		currentMove: 0,
	}
}

func (board *Board) PlayMove(cellNumber int) error {
	err := validateCellNumber(cellNumber)
	if err != nil {
		return err
	}
	err1 := board.validateIfOccupied(cellNumber)
	if err1 != nil {
		return err1
	}
	if board.currentMove%2 == 0 {
		board.symbols[cellNumber] = NewXSymbol()
	} else {
		board.symbols[cellNumber] = NewOSymbol()
	}
	board.currentMove++
	return nil

}

func (board Board) getCurrentBoard() string {
	tempBoardObject := ""
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			tempBoardObject += "\n"
		}
		if board.symbols[i].getValue() == EMPTY_SYMBOL {
			position := strconv.Itoa(i)
			tempBoardObject += " " + position + " "
		} else {
			tempBoardObject += " " + board.symbols[i].getValue() + " "

		}

	}
	return tempBoardObject
}

func (board *Board) CheckForWinner() (bool, int) {
	lastUsedSymbol := board.getLastUsedSymbol().getValue()
	for i := 0; i < 3; i++ {
		if board.symbols[i*3+0].getValue() == lastUsedSymbol &&
			board.symbols[i*3+1].getValue() == lastUsedSymbol &&
			board.symbols[i*3+2].getValue() == lastUsedSymbol {
			return true, board.getWinner()
		}
		if board.symbols[i+0*3].getValue() == lastUsedSymbol &&
			board.symbols[i+1*3].getValue() == lastUsedSymbol &&
			board.symbols[i+2*3].getValue() == lastUsedSymbol {
			return true, board.getWinner()
		}

	}
	if board.symbols[0].getValue() == lastUsedSymbol &&
		board.symbols[1*3+1].getValue() == lastUsedSymbol &&
		board.symbols[2*3+2].getValue() == lastUsedSymbol {
		return true, board.getWinner()
	}

	if board.symbols[0+2].getValue() == lastUsedSymbol &&
		board.symbols[1*3+1].getValue() == lastUsedSymbol &&
		board.symbols[2*3].getValue() == lastUsedSymbol {
		return true, board.getWinner()
	}

	return false, -1
}

func (board *Board) IsBoardFull() bool {
	return board.currentMove == 9
}

func (board *Board) getLastUsedSymbol() Symbol {
	if board.currentMove%2 == 0 {
		return NewOSymbol()
	}
	return NewXSymbol()
}

func (board *Board) getWinner() int {
	return (board.currentMove - 1) % 2
}

func (board *Board) validateIfOccupied(cellNumber int) error {

	if board.symbols[cellNumber].getValue() != EMPTY_SYMBOL {
		return errors.New("This Cell is Already Occupied")
	}
	return nil

}

func validateCellNumber(cellNumber int) error {
	if cellNumber >= 9 {
		return errors.New("Cell Number Cannot be more than 8")
	}
	if cellNumber < 0 {
		return errors.New("Cell Number Cannot be negative")
	}
	return nil
}

type Symbol struct {
	value string
}

func (s Symbol) getValue() string {
	return s.value
}

func NewXSymbol() Symbol {
	return Symbol{
		value: "X",
	}
}

func NewOSymbol() Symbol {
	return Symbol{
		value: "O",
	}
}

var EMPTY_SYMBOL = ""

func NewEmptySymbol() Symbol {
	return Symbol{
		value: EMPTY_SYMBOL,
	}
}
