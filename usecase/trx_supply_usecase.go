package usecase

import (
    "github.com/SangBejoo/parking-space-monitor/repository"
    "github.com/SangBejoo/parking-space-monitor/entity"
)

type TrxSupplyUsecase interface {
    GetAll() ([]entity.TrxSupply, error)
    Create(supply entity.TrxSupply) error
}

type trxSupplyUsecase struct {
    repo      repository.TrxSupplyRepository
    placeRepo repository.MapPlaceRepository
}

func NewTrxSupplyUsecase(repo repository.TrxSupplyRepository, placeRepo repository.MapPlaceRepository) TrxSupplyUsecase {
    return &trxSupplyUsecase{
        repo:      repo,
        placeRepo: placeRepo,
    }
}

func (u *trxSupplyUsecase) GetAll() ([]entity.TrxSupply, error) {
    return u.repo.GetAll()
}

func (u *trxSupplyUsecase) Create(supply entity.TrxSupply) error {
    // Find matching place based on coordinates
    place, err := u.placeRepo.FindByPoint(supply.Latitude, supply.Longitude)
    if err != nil {
        return err
    }
    
    if place != nil {
        placeID := place.ID
        supply.PlaceID = &placeID
    } else {
        supply.PlaceID = nil
    }
    
    return u.repo.Create(supply)
}