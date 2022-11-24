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

type BudgetController struct {
	Interactor usecase.BudgetInteractor
	Toekn      usecase.TokenInteractor
}

func NewBudgetController(
	sqlHandler database.SqlHandler,
	tokenHandler token.TokenHandler,
) *BudgetController {
	return &BudgetController{
		Interactor: usecase.BudgetInteractor{
			BudgetRepository: &database.BudgetRepository{
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
func (controller *BudgetController) Add(c *gin.Context) {
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

	var budgetJson domain.Budget

	if err := c.ShouldBindJSON(&budgetJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONのデータ構造が間違ってる"})
		return
	}

	budget := domain.Budget{}
	budget.TravelId = budgetJson.TravelId
	budget.Sum = budgetJson.Sum
	budget.Transportations = budgetJson.Transportations
	budget.Accommodation = budgetJson.Accommodation
	budget.Meal = budgetJson.Meal
	budget.Sightseeing = budgetJson.Sightseeing

	budgetId, err := controller.Interactor.Add(budget)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "データが登録されていません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budget_id": budgetId})
}
