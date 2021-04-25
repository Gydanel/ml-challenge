package service

import (
	"errors"
	"ml-challenge/domain/models"
)

type SatelliteService struct {
	db    [3]models.Satellite
	count int
}

func NewSatelliteService() *SatelliteService {
	s := &SatelliteService{}
	s.count = 3
	s.db = [3]models.Satellite{}
	s.db[0] = models.Satellite{
		Name: "kenobi",
		PosX: -500,
		PosY: -200,
	}
	s.db[1] = models.Satellite{
		Name: "skywalker",
		PosX: 100,
		PosY: -100,
	}
	s.db[2] = models.Satellite{
		Name: "sato",
		PosX: 500,
		PosY: 100,
	}
	return s
}

func (s *SatelliteService) ByName(name string) (*models.Satellite, error) {
	for _, i := range s.db {
		if i.Name == name {
			return &i, nil
		}
	}
	return nil, errors.New("not found")
}
