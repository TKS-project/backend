package controller

import (
	"net/http"

	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"

	api "github.com/Kantaro0829/clean-architecture-in-go/interfaces/rakutenApi"
	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type HotelApiController struct {
	Interactor usecase.HotelApiRequestInteractor
}

func NewHotelApiRequestController(
	rakutenApiHandler api.ApiRequestHandler,
) *HotelApiController {
	return &HotelApiController{
		Interactor: usecase.HotelApiRequestInteractor{
			ApiRequestRepository: &api.ApiRequestRepository{
				ApiRequestHandler: rakutenApiHandler,
			},
		},
	}
}

func (controller *HotelApiController) GetVacantHotelAndRoom(c *gin.Context) {
	//englishword := c.Param("word")
	hotelsAndRooms := controller.Interactor.VacantHotel()

	c.JSON(http.StatusOK, gin.H{"message": hotelsAndRooms})
}
