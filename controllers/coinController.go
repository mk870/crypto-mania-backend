package controllers

import (
	"cryptoApi/middleware"
	"cryptoApi/services"

	"github.com/gin-gonic/gin"
)

func CreateCoin(router *gin.Engine) {
	router.POST("/coin", middleware.AuthValidator, services.CreateCoinService)
}

func GetCoins(router *gin.Engine) {
	router.GET("/coins", middleware.AuthValidator, services.GetCoinsService)
}

func DeleteCoin(router *gin.Engine) {
	router.DELETE("/coin/:id", middleware.AuthValidator, services.DeleteCoinService)
}
