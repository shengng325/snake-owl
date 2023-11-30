package validator

import (
	"testing"

	"snake/dto"
)

type testArgs struct {
	state       dto.ValidationDto
	expectError bool
}

func TestValidateOutOfBounds(t *testing.T) {
	testCases := []testArgs{}
	// case 1
	state1 := dto.ValidationDto{}
	state1.Height = 10
	state1.Width = 10
	state1.SnakePosTrace = []dto.SnakePos{
		{X: 8, Y: 7},
		{X: 8, Y: 8},
		{X: 8, Y: 9},
		{X: 8, Y: 10},
	}
	testCases = append(testCases, testArgs{state1, true})
	// case 2
	state2 := dto.ValidationDto{}
	state2.Height = 10
	state2.Width = 10
	state2.SnakePosTrace = []dto.SnakePos{
		{X: 8, Y: 7},
		{X: 8, Y: 8},
		{X: 8, Y: 9},
		{X: 9, Y: 9},
	}
	testCases = append(testCases, testArgs{state2, false})
	for _, tc := range testCases {
		err, _ := validateOutOfBounds(tc.state)
		if tc.expectError != (err != nil) {
			t.Errorf("Expect error to be %v, got error = %v", tc.expectError, err)
		}
	}
}

func TestValidateFruitIsFound(t *testing.T) {
	testCases := []testArgs{}
	// case 1
	state1 := dto.ValidationDto{}
	state1.Fruit.X = 8
	state1.Fruit.Y = 10
	state1.SnakePosTrace = []dto.SnakePos{
		{X: 8, Y: 8},
		{X: 8, Y: 9},
		{X: 9, Y: 9},
	}
	testCases = append(testCases, testArgs{state1, true})
	// case 2
	state2 := dto.ValidationDto{}
	state2.Fruit.X = 8
	state2.Fruit.Y = 10
	state2.SnakePosTrace = []dto.SnakePos{
		{X: 8, Y: 8},
		{X: 8, Y: 9},
		{X: 8, Y: 10},
	}
	testCases = append(testCases, testArgs{state2, false})
	for _, tc := range testCases {
		err, _ := validateFruitIsFound(tc.state)
		if tc.expectError != (err != nil) {
			t.Errorf("Expect error to be %v, got error = %v", tc.expectError, err)
		}
	}
}

func TestValidateMovements(t *testing.T) {
	testCases := []testArgs{}
	// case 1
	state1 := dto.ValidationDto{}
	state1.SnakePosTrace = []dto.SnakePos{
		{X: 5, Y: 5},
		{X: 5, Y: 6},
		{X: 5, Y: 7},
		{X: 5, Y: 6},
	}
	testCases = append(testCases, testArgs{state1, true})
	// case 2
	state2 := dto.ValidationDto{}
	state2.SnakePosTrace = []dto.SnakePos{
		{X: 5, Y: 5},
		{X: 4, Y: 5},
		{X: 3, Y: 5},
		{X: 4, Y: 5},
	}
	testCases = append(testCases, testArgs{state2, true})
	// case 3
	state3 := dto.ValidationDto{}
	state3.SnakePosTrace = []dto.SnakePos{
		{X: 5, Y: 5},
		{X: 4, Y: 5},
		{X: 3, Y: 5},
		{X: 3, Y: 4},
	}
	testCases = append(testCases, testArgs{state3, false})
	for _, tc := range testCases {
		err, _ := validateMovements(tc.state)
		if tc.expectError != (err != nil) {
			t.Errorf("Expect error to be %v, got error = %v", tc.expectError, err)
		}
	}
}
