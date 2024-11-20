package usecase

import (
	"github.com/SangBejoo/parking-space-monitor/entity"
	"github.com/SangBejoo/parking-space-monitor/repository"
)

type MapPlaceUsecase interface {
	Create(mapPlace entity.MapPlace) error
	GetAll() ([]entity.MapPlace, error)
}

type mapPlaceUsecase struct {
	mapPlaceRepo repository.MapPlaceRepository
}

func NewMapPlaceUsecase(repo repository.MapPlaceRepository) MapPlaceUsecase {
	return &mapPlaceUsecase{
		mapPlaceRepo: repo,
	}
}

func (u *mapPlaceUsecase) Create(mapPlace entity.MapPlace) error {
	return u.mapPlaceRepo.Create(mapPlace)
}

func (u *mapPlaceUsecase) GetAll() ([]entity.MapPlace, error) {
	return u.mapPlaceRepo.GetAll()
}