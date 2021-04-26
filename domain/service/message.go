package service

import (
	"errors"
	"ml-challenge/domain/models"
)

type MessageService struct {
	/*
		En este caso el siguiente paso era implementar una db a nivel app para guardar los models pero me encontre
		con problemas al configurar el package gorm y sqlite y si seguia no llegaba con el deadline que habia prometido ademas
		supuse que una db para mantener 6 registros (3 messages y 3 satelites) era un poco de overkill
	*/
	db []*models.Message
}

func NewMessageService() *MessageService {
	return &MessageService{db: make([]*models.Message, 0, 3)}
}

func (s *MessageService) BySatelliteName(name string) (*models.Message, error) {
	for _, m := range s.db {
		if m != nil && m.Receiver.Name == name {
			return m, nil
		}
	}
	return nil, errors.New("not found")
}

func (s *MessageService) Add(d float32, m []string, sa models.Satellite) {
	s.db = append(s.db, &models.Message{
		Receiver: sa,
		Distance: d,
		Message:  m,
	})
}

func (s *MessageService) UpdateBySatelliteName(distance float32, message []string, name string) *models.Message {
	r, _ := s.BySatelliteName(name)
	r.Distance = distance
	r.Message = message
	return r
}

func (s *MessageService) All() []*models.Message {
	return s.db
}
