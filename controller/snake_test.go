package controller

import (
	"reflect"
	"testing"

	"snake/dto"
)

func TestGenerateSnakePosTrace(t *testing.T) {
	right := dto.Tick{VelX: 1, VelY: 0}
	left := dto.Tick{VelX: -1, VelY: 0}
	up := dto.Tick{VelX: 0, VelY: 1}
	down := dto.Tick{VelX: 0, VelY: -1}
	initialPos := dto.SnakePos{X: 10, Y: 10}
	ticks := []dto.Tick{
		right, right, up, left, up, right, down, down, left,
	}
	expected := []dto.SnakePos{
		{X: 10, Y: 10},
		{X: 11, Y: 10},
		{X: 12, Y: 10},
		{X: 12, Y: 11},
		{X: 11, Y: 11},
		{X: 11, Y: 12},
		{X: 12, Y: 12},
		{X: 12, Y: 11},
		{X: 12, Y: 10},
		{X: 11, Y: 10},
	}

	outcome := generateSnakePosTrace(initialPos, ticks)
	if !reflect.DeepEqual(outcome, expected) {
		t.Errorf("Expected %v, got %v", expected, outcome)
	}
}
