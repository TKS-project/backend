package controller

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/scrape"

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
