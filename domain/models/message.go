package models

type Message struct {
	Receiver Satellite
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}
