package controller

import (

	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"

	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type PrefectureAndCityController struct {
	Interactor usecase.PrefectureAndCityInteractor
}

func NewPrefectureAndCityController(
	sqlHandler database.SqlHandler,
) *PrefectureAndCityController {
	return &PrefectureAndCityController{
		Interactor: usecase.PrefectureAndCityInteractor{
			PrefectureRepository: &database.PrefectureRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PrefectureAndCityController) GetAllPrefectures(c *gin.Context) {
	res, err := controller.Interactor.FindAllPrefecture()
	if err != nil {
		//エラーハンドリング
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *PrefectureAndCityController) GetALlCities(c *gin.Context) {
	res, err := controller.Interactor.FindAllCities()
	if err != nil {
		//エラーハンドリング
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
