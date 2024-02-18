package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `json:"username" gorm:"unique;not null" binding:"required"`
	Email     string  `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password  string  `json:"password" gorm:"not null" binding:"required,min=6"`
	IsPremium bool    `json:"is_premium" gorm:"default:false"`
	Swipe     []Swipe `gorm:"foreignKey:SwiperID"`
}

type Swipe struct {
	gorm.Model
	SwiperID       int    `json:"swiper_id" gorm:"not null"`
	SwipeeID       int    `json:"swipee_id" gorm:"not null" binding:"required"`
	SwipeDirection string `json:"swipe_direction" gorm:"not null" binding:"required"`
}

type SwipeDirectionEnum string

const (
	Like    SwipeDirectionEnum = "like"
	Dislike SwipeDirectionEnum = "dislike"
)

type UserList struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	IsPremium bool   `json:"is_premium"`
}
