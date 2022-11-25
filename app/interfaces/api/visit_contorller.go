package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"

	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type VisitController struct {
	Interactor usecase.VisitInteractor
	Toekn      usecase.TokenInteractor
}

func NewVisitController(
	sqlHandler database.SqlHandler,
	tokenHandler token.TokenHandler,
) *VisitController {
	return &VisitController{
		Interactor: usecase.VisitInteractor{
			VisitRepository: &database.VisitRepository{
				SqlHandler: sqlHandler,
			},
		},
		Toekn: usecase.TokenInteractor{
			Tokenizer: &token.TokenizerImpl{
				TokenHandler: tokenHandler,
			},
		},
	}
}

//ルーティングのハンドラ
func (controller *VisitController) Add(c *gin.Context) {
	//Get token from request header
	var header domain.HeaderWithToken
	err := c.BindHeader(&header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}

	// Verify token
	tokenString := domain.Token(header.Authorization)
	id, err := controller.Toekn.GetId(tokenString)
	//id, err := controller.Interactor.Authenticate(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}
	fmt.Printf("JWTから解析したID : %v", id)

	var visitJson domain.VisitJson

	if err := c.ShouldBindJSON(&visitJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONのデータ構造が間違ってる"})
		return
	}
	layout := "2006/01/02"
	dateString := visitJson.VisitDay
	visit := domain.Visit{}

	visit.VisitDay, _ = time.Parse(layout, dateString)
	visit.BudgetId = visitJson.BudgetId
	visit.Address = visitJson.Address
	visit.Charge = visitJson.Charge
	visit.Info = visitJson.Info
	visit.Name = visitJson.Name
	visit.VisitType = visitJson.VisitType
	//visitday

	visitId, err := controller.Interactor.Add(visit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "データが登録されていません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budget_id": visitId})
}
