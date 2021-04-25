package api

import "ml-challenge/domain/models"

type MessageApi struct {
	Receiver SatelliteApi
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type DecodedMessageApi struct {
	Pos     PositionApi `json:"position" binding:"required"`
	Message string      `json:"message" binding:"required"`
}

func (m *MessageApi) Init(x models.Message) {
	m.Message = x.Message
	m.Receiver = SatelliteApi{
		Name: x.Receiver.Name,
		PosX: x.Receiver.PosX,
		PosY: x.Receiver.PosY,
	}
	m.Distance = x.Distance
}
