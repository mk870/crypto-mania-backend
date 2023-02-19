package services

import (
	"net/http"

	"cryptoApi/models"
	"cryptoApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateCoinService(c *gin.Context) {
	var coin models.Coin
	err := c.ShouldBindJSON(&coin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not bind the request body",
		})
		return
	}
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find this user",
		})
		return
	}

	isCoinCreated := repositories.CreateCoinRepository(user, &coin)
	if !isCoinCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the coin",
		})
		return
	}
	c.String(http.StatusOK, "coin successfully added to your watchlist")
}

func GetCoinsService(c *gin.Context) {
	loggedStudent := c.MustGet("user").(*models.User)
	email := loggedStudent.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	coinList := repositories.GetCoinsRepository(user.Id)
	c.JSON(http.StatusOK, coinList)
}

func GetCoinService(c *gin.Context) {
	coin_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	coin := repositories.GetCoinRepository(user.Id, coin_id)
	c.JSON(http.StatusOK, coin)
}

func DeleteCoinService(c *gin.Context) {
	coin_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	coin := repositories.GetCoinRepository(user.Id, coin_id)
	isDeleted := repositories.DeleteCoinByIdRepository(user, coin)
	if isDeleted {
		c.JSON(http.StatusOK, "coin successfully deleted")
		return
	}
}
