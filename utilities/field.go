package utilities

type Field struct {
	Width            uint8 `json:"width"`
	Height           uint8 `json:"height"`
	Position         Position `json:"position"`
	OpponentPosition Position `json:"opponentPosition"`
	Barriers         ObstacleArray `json:"barriers"`
}

//func NewField(width, height uint8, position, opponentPosition Position, barriers [][4][2]uint8) *Field {
//	return &Field{
//		Width:            width,
//		Height:           height,
//		Position:         position,
//		OpponentPosition: opponentPosition,
//		Barriers:         barriers,
//	}
//}
