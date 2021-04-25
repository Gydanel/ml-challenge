package api

type TopSecretApi struct {
	Pos     PositionApi `json:"position"`
	Message string      `json:"message"`
}

type PositionApi struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
