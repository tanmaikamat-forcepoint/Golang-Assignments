package game

type Game interface {
	Play(...interface{}) (bool, error)
	GetResult() (string, error)
	GetCurrentBoard() string
}
