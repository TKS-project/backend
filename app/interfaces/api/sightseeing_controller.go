package controller

import (
	"net/http"
	"strconv"

	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"

	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type SightseeingController struct {
	Interactor usecase.SightseeingInteractor
}

func NewSightseeingController(
	sqlHandler database.SqlHandler,
) *SightseeingController {
	return &SightseeingController{
		Interactor: usecase.SightseeingInteractor{
			SightseeingRepository: &database.SightseeingRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SightseeingController) GetInfo(c *gin.Context) {
	id := c.Param("id")
	preId, err := strconv.Atoi(id)
	if err != nil {

	}

	results, err := controller.Interactor.GetSightseeingInfo(preId)
	c.JSON(http.StatusOK, gin.H{"sightdata": results})
}
