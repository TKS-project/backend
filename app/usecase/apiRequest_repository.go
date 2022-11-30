package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository interface {
	FindRoom(
		prefecture string,
		city string,
		checkinDate string,
		checkoutDate string,
		adultNum string,
		maxCharge string,
	) domain.VacantHotels
}
