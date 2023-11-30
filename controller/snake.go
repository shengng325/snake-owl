package controller

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"snake/dto"
	"snake/validator"

	"github.com/google/uuid"
)

type Snake struct {
	validator *validator.Validator
}

func NewSnakeController(v *validator.Validator) *Snake {
	return &Snake{
		validator: v,
	}
}

func (s *Snake) NewGame(w http.ResponseWriter, r *http.Request) {
	queryParamValues := r.URL.Query()
	width, err := getIntQueryParams(queryParamValues, "w")
	if err != nil {
		createErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	height, err := getIntQueryParams(queryParamValues, "h")
	if err != nil {
		createErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	newGameID := uuid.New().String()
	newFruitX := randomInt(0, width)
	newFruitY := randomInt(0, height)
	newGameState := &dto.State{
		GameID: newGameID,
		Width:  width,
		Height: height,
		Score:  0,
		Fruit: dto.Fruit{
			X: newFruitX,
			Y: newFruitY,
		},
		Snake: dto.Snake{
			X:    0,
			Y:    0,
			VelX: 1,
			VelY: 0,
		},
	}
	log.Printf("new game started with gameId: %v", newGameID)
	createHttpResponse(w, http.StatusOK, newGameState)
}

func (s *Snake) ValidateGame(w http.ResponseWriter, r *http.Request) {
	validationDto := dto.ValidationDto{}
	if err := json.NewDecoder(r.Body).Decode(&validationDto); err != nil {
		createErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validationDto.SnakePosTrace = generateSnakePosTrace(dto.SnakePos{X: validationDto.Snake.X, Y: validationDto.Snake.Y}, validationDto.Ticks)
	for _, validator := range s.validator.Validators {
		err, errCode := validator(validationDto)
		if err != nil {
			createErrorResponse(w, errCode, err)
			return
		}
	}

	log.Printf("valid score for gameId: %v, generating new fruit", validationDto.GameID)
	newScore := validationDto.Score + 1
	newFruitPosX := randomInt(0, validationDto.Width)
	newFruitPosY := randomInt(0, validationDto.Height)
	newGameState := &dto.State{
		GameID: validationDto.GameID,
		Width:  validationDto.Width,
		Height: validationDto.Height,
		Score:  newScore,
		Fruit: dto.Fruit{
			X: newFruitPosX,
			Y: newFruitPosY,
		},
		Snake: dto.Snake{
			X:    0,
			Y:    0,
			VelX: 1,
			VelY: 0,
		},
	}
	createHttpResponse(w, http.StatusOK, newGameState)
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// generateSnakePosTrace generates all the pos that the snake moved using ticks
func generateSnakePosTrace(initialSnakePos dto.SnakePos, ticks []dto.Tick) []dto.SnakePos {
	snakePosTrace := []dto.SnakePos{initialSnakePos}
	for _, tick := range ticks {
		nextX := snakePosTrace[len(snakePosTrace)-1].X + tick.VelX
		nextY := snakePosTrace[len(snakePosTrace)-1].Y + tick.VelY
		nextSnakePos := dto.SnakePos{X: nextX, Y: nextY}
		snakePosTrace = append(snakePosTrace, nextSnakePos)
	}
	log.Println(snakePosTrace)
	return snakePosTrace
}
