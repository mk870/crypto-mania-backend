package repositories

import (
	"errors"

	"cryptoApi/db"
	"cryptoApi/models"

	"gorm.io/gorm"
)

func CreateCoinRepository(user *models.User, coin *models.Coin) bool {
	err := db.DB.Model(user).Association("Coins").Append(coin)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetCoinsRepository(id int) []models.Coin {
	var user = models.User{}
	err := db.DB.Preload("Coins").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Coins
}

func GetCoinRepository(userId int, coinId string) models.Coin {
	var user = models.User{}
	err := db.DB.Preload("Coins", "id=?", coinId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Coins[0]
}

func GetUserWithCoinsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Coins").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateCoinRepository(user *models.User, updateList []models.Coin) bool {
	user.Coins = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteCoinByIdRepository(user *models.User, coin models.Coin) bool {
	db.DB.Model(&user).Unscoped().Association("Coins").Delete(coin)
	return true
}
