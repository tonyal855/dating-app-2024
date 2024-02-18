package repositories

import "dating-app-dealls/server/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUser(username string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetListUsers(id int) (*[]models.UserList, error)
	UpgradePremium(id int) error
}

type SwipeRepo interface {
	CreateSwipe(swipe *models.Swipe) error
	CountLimit(id int) int
}
