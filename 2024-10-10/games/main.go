package main

import (
	"fmt"
	"games/game"
	ticTacToe "games/tictactoe"
)

func main() {
	Player1 := "Tanmai"
	Player2 := "Tanmay"
	var ticTacGame game.Game
	ticTacGame, err := ticTacToe.NewTicTacToeGame(Player1, Player2)
	if err != nil {
		panic(err)
	}
	gameComplete := false
	var inp int

	for !gameComplete {
		fmt.Println("Enter Your Choice: ")
		fmt.Scanln(&inp)
		tempGameCompleteValidation, err1 := ticTacGame.Play(inp)
		gameComplete = tempGameCompleteValidation
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println("Current Board:")
		fmt.Println(ticTacGame.GetCurrentBoard())
	}
	result, err2 := ticTacGame.GetResult()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result)

}
