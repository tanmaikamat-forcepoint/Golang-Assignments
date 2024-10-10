package ticTacToe

import "errors"

type Player struct {
	name string
}

func NewPlayer(name string) (*Player, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	return &Player{
		name: name,
	}, nil
}

func (player *Player) getName() string {
	return player.name
}
func validateName(name string) error {
	if len(name) < 3 {
		return errors.New("Invalid Name. Name should be of atleast length 3")
	}
	return nil
}
