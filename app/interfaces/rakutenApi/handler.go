package rakutenapi

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestHandler interface {
	//Find(word string) string
	FindAllVacantHotel() domain.VacantHotels
}
