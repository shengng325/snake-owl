package dto

type ValidationDto struct {
	State
	Ticks         []Tick `json:"ticks"`
	SnakePosTrace []SnakePos
}
