package rakutenapi

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository struct {
	ApiRequestHandler
}

func (request *ApiRequestRepository) FindRoom() domain.VacantHotels {

	hotelAndRooms := request.FindAllVacantHotel()
	return hotelAndRooms
}
