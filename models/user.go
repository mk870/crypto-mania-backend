package models

import (
	"time"

	"gorm.io/gorm"
)

type MyModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	MyModel
	Id                int               `json:"id" gorm:"primary_key"`
	FirstName         string            `json:"firstName" validate:"required,min=2,max=50"`
	LastName          string            `json:"lastName" validate:"required,min=2,max=50"`
	Email             string            `json:"email" gorm:"unique" validate:"email,required"`
	Password          string            `json:"password" validate:"required,min=2,max=50"`
	RefreshToken      string            `json:"refreshToken"`
	IsActive          bool              `json:"isActive"`
	Coins             []Coin            `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RegistrationToken VerificationToken `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
