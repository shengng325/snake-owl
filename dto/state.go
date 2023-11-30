package dto

type State struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  Fruit  `json:"fruit"`
	Snake  Snake  `json:"snake"`
}

type Fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

type Tick struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

type SnakePos struct {
	X int
	Y int
}
