package validator

import (
	"errors"
	"net/http"

	"snake/dto"
)

type ValidationFunc func(state dto.ValidationDto) (error, int)
type Validator struct {
	Validators []ValidationFunc
}

func NewValidator() *Validator {
	return &Validator{
		Validators: []ValidationFunc{
			validateOutOfBounds,
			validateMovements,
			validateFruitIsFound,
		},
	}
}

// validate all pos does not go out of bounds
func validateOutOfBounds(state dto.ValidationDto) (error, int) {
	maxWidth := state.Width
	maxHeight := state.Height
	for _, pos := range state.SnakePosTrace {
		if pos.X < 0 || pos.X >= maxWidth || pos.Y < 0 || pos.Y >= maxHeight {
			return errors.New("ouch! the snake hit the bounds!"), http.StatusTeapot
		}
	}
	return nil, 0
}

// validate that the snake reaches the fruit position
func validateFruitIsFound(state dto.ValidationDto) (error, int) {
	lastPos := state.SnakePosTrace[len(state.SnakePosTrace)-1]
	if lastPos.X != state.Fruit.X || lastPos.Y != state.Fruit.Y {
		return errors.New("no fruit for the snake :("), http.StatusNotFound
	}
	return nil, 0
}

// validate no invalid movements by checking 3 consecutive pos
// when prev-prev pos is equal to current pos, it means that it is an invalid 180 degree turn
func validateMovements(state dto.ValidationDto) (error, int) {
	if len(state.SnakePosTrace) < 3 {
		return nil, 0
	}
	prevPrevPos := state.SnakePosTrace[0]
	for i, currentPos := range state.SnakePosTrace {
		if i < 2 {
			continue
		}
		if currentPos == prevPrevPos {
			return errors.New("invalid movement"), http.StatusTeapot
		}
		prevPrevPos = state.SnakePosTrace[i-1]
	}
	return nil, 0
}
