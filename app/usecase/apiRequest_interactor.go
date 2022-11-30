package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type HotelApiRequestInteractor struct {
	ApiRequestRepository ApiRequestRepository
}

func (interactor *HotelApiRequestInteractor) VacantHotel(
	prefecture string,
	city string,
	checkinDate string,
	checkoutDate string,
	adultNum string,
	maxCharge string,
) domain.VacantHotels {

	hotelsAndRooms := interactor.ApiRequestRepository.FindRoom(prefecture, city, checkinDate, checkoutDate, adultNum, maxCharge)
	return hotelsAndRooms
}
