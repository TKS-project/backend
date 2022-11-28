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

type TransportsController struct {
	Interactor usecase.TransportsInteractor
	Token      usecase.TokenInteractor
}

func NewTransportsController(
	sqlHandler database.SqlHandler,
	tokenHandler token.TokenHandler,
) *TransportsController {
	return &TransportsController{
		Interactor: usecase.TransportsInteractor{
			TransportsRepository: &database.TransportsRepository{
				SqlHandler: sqlHandler,
			},
		},
		Token: usecase.TokenInteractor{
			Tokenizer: &token.TokenizerImpl{
				TokenHandler: tokenHandler,
			},
		},
	}
}

//ルーティングのハンドラ
func (controller *TransportsController) Add(c *gin.Context) {
	//Get token from request header
	var header domain.HeaderWithToken
	err := c.BindHeader(&header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}

	// Verify token
	tokenString := domain.Token(header.Authorization)
	id, err := controller.Token.GetId(tokenString)
	//id, err := controller.Interactor.Authenticate(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証できませんでした"})
		return
	}
	fmt.Printf("JWTから解析したID : %v", id)

	var transportsJson domain.Transportation

	if err := c.ShouldBindJSON(&transportsJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONのデータ構造が間違ってる"})
		return
	}

	transports := domain.Transportation{}
	transports.BudgetId = transportsJson.BudgetId
	transports.Destination = transportsJson.Destination
	transports.Departure = transportsJson.Departure
	transports.Type = transportsJson.Type
	transports.Charge = transportsJson.Charge
	transports.OrderNo = transportsJson.OrderNo
	transports.WayThere = transportsJson.WayThere

	err = controller.Interactor.Add(transports)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "データが登録されていません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登録完了"})
}
