package controller

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"

	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type TravelController struct {
	Interactor usecase.TravelInteractor
}

func NewTravelController(
	sqlHandler database.SqlHandler,
	tokenHandler token.TokenHandler,
) *TravelController {
	return &TravelController{
		Interactor: usecase.TravelInteractor{
			TravelRepository: &database.TravelRepository{
				SqlHandler: sqlHandler,
			},
			Tokenizer: &token.TokenizerImpl{
				TokenHandler: tokenHandler,
			},
		},
	}
}

//ルーティングのハンドラ
func (controller *TravelController) Add(c *gin.Context) {
	// Get token from request header
	var header domain.HeaderWithToken
	err := c.BindHeader(&header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}

	// Verify token
	tokenString := domain.Token(header.Authorization)
	id, err := controller.Interactor.Authenticate(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}
	fmt.Printf("JWTから解析したID : %v", id)

	var travelJson domain.Travel

	if err := c.ShouldBindJSON(&travelJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONのデータ構造が間違ってる"})
		return
	}

	travel := domain.Travel{}
	travel.UserId = id
	travel.MaxPerson = travelJson.MaxPerson
	travel.MaxDay = travelJson.MaxDay
	travel.DepartureId = travelJson.DepartureId
	travel.DestinationId = travelJson.DestinationId

	travelId, err := controller.Interactor.Add(travel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "データが登録されていません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"travel_id": travelId})
}

func (controller *TravelController) Update(c *gin.Context) {
	// Get token from request header
	var header domain.HeaderWithToken
	err := c.BindHeader(&header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}

	// Verify token
	tokenString := domain.Token(header.Authorization)
	id, err := controller.Interactor.Authenticate(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}

	var travelJson domain.Travel

	if err := c.ShouldBindJSON(&travelJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONのデータ構造が間違ってる"})
		return
	}

	travelJson.UserId = id

	result := controller.Interactor.Update(travelJson)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "データの更新に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "データ更新成功"})
}
