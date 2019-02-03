package structures

type Game struct {
	ID string `json:"ID,omitempty"`
	GameTable [3][3]int `json:"GameTable"` // the game storage will consist of a 3X3 matrix containing 0 or 1
	LastMove string `json:"LastMove"`
	TotalMovesMade int `json:"TotalMovesMade"`
	Status string `json:"Status"`
	Winner string `json:"Winner"`
}

type Message struct {
	Msg string `json:"Msg,omitempty"`
}

type JsonModel struct {
	ID string `json:"ID"`
	M00 string `json:"M00"`
	M01 string `json:"M01"`
	M02 string `json:"M02"`
	M10 string `json:"M10"`
	M11 string `json:"M11"`
	M12 string `json:"M12"`
	M20 string `json:"M20"`
	M21 string `json:"M21"`
	M22 string `json:"M22"`
	LastMove string `json:"LastMove"`
	Status string `json:"Status"`
	Winner string `json:"Winner"`
}

type Payload struct {
	Row string
	Column string
}