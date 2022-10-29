package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type HotelApiRequestInteractor struct {
	ApiRequestRepository ApiRequestRepository
}

func (interactor *HotelApiRequestInteractor) VacantHotel() domain.VacantHotels {

	hotelsAndRooms := interactor.ApiRequestRepository.FindRoom()
	return hotelsAndRooms
}
