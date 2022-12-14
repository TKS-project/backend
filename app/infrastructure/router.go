package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	controllers "github.com/Kantaro0829/clean-architecture-in-go/interfaces/api"
	"github.com/gin-gonic/gin"
)

func Init() {

	router := gin.Default()
	userController := controllers.NewUserController(
		NewSqlHandler(),
		NewTokenHandler(), //jwt用のDI
	)

	scrapeController := controllers.NewScrapeController(
		NewGoQueryHandler(),
	)

	hotelApiController := controllers.NewHotelApiRequestController(
		NewApiRequestHandler(),
	)

	prefectureAndCityController := controllers.NewPrefectureAndCityController(
		NewSqlHandler(),
	)
	travelController := controllers.NewTravelController(
		NewSqlHandler(),
		NewTokenHandler(),
	)
	budgetController := controllers.NewBudgetController(
		NewSqlHandler(),
		NewTokenHandler(),
	)
	visitController := controllers.NewVisitController(
		NewSqlHandler(),
		NewTokenHandler(),
	)
	transportsController := controllers.NewTransportsController(
		NewSqlHandler(),
		NewTokenHandler(),
	)
	sightseeingController := controllers.NewSightseeingController(
		NewSqlHandler(),
	)

	router.GET("/users", func(c *gin.Context) {
		userController.GetUser(c)
		return
	})

	router.POST("/users", func(c *gin.Context) {
		userController.Create(c)
		return
	})

	router.POST("/users/update", func(c *gin.Context) {
		var userJson domain.User
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(userJson)
		message, err, isValidated := userController.UpdateByMail(userJson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
			return
		}
		if isValidated != true {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": message})
			return
		}
		//userController.Update(c)
		c.JSON(http.StatusOK, gin.H{"message": message})
		return
	})

	router.POST("users/login", func(c *gin.Context) {
		var userJson domain.User
		//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mail, password := userJson.Mail, userJson.Password
		//以下Tokenを追加で受け取る
		token, result, err := userController.Login(mail, password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "dbサーバーのエラー"})
			return
		}

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "JWTtokenの発行失敗"})
			return
		}

		if result {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "ログイン完了", "token": token})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "パスワードかメールアドレスが違います"})
		return
	})

	router.GET("users/authenticate", func(c *gin.Context) {
		result := userController.Authenticate(c)
		if result != nil {
			fmt.Printf(":エラー内容：%v", result)
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "JWT認証失敗",
				})
			return
		}
		c.JSON(
			http.StatusAccepted,
			gin.H{
				"statsu":  http.StatusAccepted,
				"message": "JWT認証成功",
			})
		return
	})

	router.GET("/word/:word", func(c *gin.Context) {
		scrapeController.GetMeaning(c)
		return
	})

	router.POST("/scrape/transports", func(c *gin.Context) {
		scrapeController.GetOneTransports(c)
		return
	})

	router.GET("/scrape", func(c *gin.Context) {
		scrapeController.GetTransports(c)
		return
	})

	router.GET("/hotel/:prefecture/:city/:checkin/:checkout/:adultNum/:maxcharge", func(c *gin.Context) {
		hotelApiController.GetVacantHotelAndRoom(c)
		return
	})

	router.GET("/prefectures", func(c *gin.Context) {
		prefectureAndCityController.GetAllPrefectures(c)
		return
	})

	router.GET("/cities", func(c *gin.Context) {
		prefectureAndCityController.GetALlCities(c)
		return
	})

	router.GET("/cities/:prefectureId", func(c *gin.Context) {
		prefectureAndCityController.GetCitiesByPrefectureId(c)
		return
	})

	router.GET("/detailed", func(c *gin.Context) {
		prefectureAndCityController.GetAllDetailedCities(c)
		return
	})

	router.GET("/detailed/:cityId", func(c *gin.Context) {
		prefectureAndCityController.GetDetailedCitiesByCityId(c)
		return
	})

	router.GET("/detailed/exist/:cityId", func(c *gin.Context) {
		prefectureAndCityController.CheckIsDcityExisting(c)
		return
	})

	router.GET("/user/:mail", func(c *gin.Context) {
		userController.GetNamePass(c)
		return
	})

	router.POST("/travel", func(c *gin.Context) {
		travelController.Add(c)
		return
	})

	router.PUT("/travel", func(c *gin.Context) {
		travelController.Update(c)
		return
	})

	router.POST("/budget", func(c *gin.Context) {
		budgetController.Add(c)
		return
	})
	router.PUT("/budget", func(c *gin.Context) {
		budgetController.Update(c)
		return
	})
	router.POST("/visit", func(c *gin.Context) {
		visitController.Add(c)
		return
	})
	router.POST("/transports", func(c *gin.Context) {
		transportsController.Add(c)
		return
	})
	router.GET("/sightseeing/:id", func(c *gin.Context) {
		sightseeingController.GetInfo(c)
		return
	})

	router.Run(":3000")
}
