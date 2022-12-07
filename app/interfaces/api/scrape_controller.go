package controller

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/scrape"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type ScrapeController struct {
	Interactor usecase.GetMeaningInteractor
}

func NewScrapeController(
	goQueryHandler scrape.ScrapeHandler,
	//sqlHandler database.SqlHandler,
	//以下追加
	//tokenHandler token.TokenHandler,
) *ScrapeController {
	return &ScrapeController{
		Interactor: usecase.GetMeaningInteractor{
			ScrapeRepository: &scrape.ScrapeRepository{
				ScrapeHandler: goQueryHandler,
			},
		},
	}
}

func (controller *ScrapeController) GetMeaning(c *gin.Context) {
	englishword := c.Param("word")
	res := controller.Interactor.GetTranslatedMeaning(englishword)
	fmt.Println("取得した単語の意味:%v", res)

	c.JSON(http.StatusOK, gin.H{"message": res})
}

func (controller *ScrapeController) GetTransports(c *gin.Context) {
	res := controller.Interactor.GetTransports()
	c.JSON(http.StatusOK, gin.H{"message": res})
}

func (controller *ScrapeController) GetOneTransports(c *gin.Context) {

	var transportRequest domain.TransportsRequest

	if err := c.ShouldBindJSON(&transportRequest); err != nil {
		fmt.Println(transportRequest)
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONのデータ構造が間違ってる"})
		return
	}

	res := controller.Interactor.GetOneTransports(transportRequest)
	c.JSON(http.StatusOK, res)
}
