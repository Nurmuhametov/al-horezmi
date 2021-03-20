package utilities

type Position [2]uint8
type ObstacleArray [][4][2]uint8

type Obstacle struct {
	From1 Position
	To1   Position
	From2 Position
	To2   Position
}

func (o Obstacle) isStepOver(from, to Position) bool {
	if (from [0] ==  o.From1[0] && from [1] ==  o.From1[1] && to [0] ==  o.To1[0] && to [1] ==  o.To1[1]) ||
		(from [0] ==  o.From2[0] && from [1] ==  o.From2[1] && to [0] ==  o.To2[0] && to [1] ==  o.To2[1]) ||
		(from [0] ==  o.To1[0] && from [1] ==  o.To1[1] && to [0] ==  o.From1[0] && to [1] ==  o.From1[1]) ||
		(from [0] ==  o.To2[0] && from [1] ==  o.To2[1] && to [0] ==  o.From2[0] && to [1] ==  o.From2[1]) { return true }
	return false
}

func (o Obstacle) isValidObstacle(width, height uint8) bool {
	if o.From1[0] < 0 || o.From1[0] >= height || o.From1[1] < 0 || o.From1[1] >= width ||
		o.From2[0] < 0 || o.From2[0] >= height || o.From2[1] < 0 || o.From2[1] >= width ||
		o.To1[0] < 0 || o.To1[0] >= height || o.To1[1] < 0 || o.To1[1] >= width ||
		o.To2[0] < 0 || o.To2[0] >= height || o.To2[1] < 0 || o.To2[1] >= width {
		return false
	}
	return true
}

func GetObstacles(array ObstacleArray) []Obstacle {
	var res []Obstacle
	for _, value := range array{
		res = append(res, Obstacle{
			From1: Position{value[0][0], value[0][1]},
			To1:   Position{value[1][0], value[1][1]},
			From2: Position{value[2][0], value[2][1]},
			To2:   Position{value[3][0], value[3][1]},
		})
	}
	return res
}

type Message struct {
	Msg string `json:"MESSAGE"`
}

type LobbyInfo struct {
	ID string `json:"_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	GameBarrierCount int `json:"gameBarrierCount"`
	PlayerBarrierCount int `json:"playerBarrierCount"`
	Name string `json:"name"`
	PlayersCount int `json:"players_count"`
}

type JoinResponse struct {
	Data LobbyInfo `json:"DATA"`
	Success bool `json:"SUCCESS"`
}

type StartResponse struct {
	Move bool `json:"move"`
	Width uint8 `json:"width"`
	Height uint8 `json:"height"`
	Position Position `json:"position"`
	OpponentPosition Position `json:"opponentPosition"`
	Barriers ObstacleArray `json:"barriers"`
}