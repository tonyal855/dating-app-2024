package repo

import (
	"dating-app-dealls/server/models"
	"dating-app-dealls/server/repositories"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(orders *models.User) error {
	return u.db.Create(orders).Error
}

func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := u.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetUserByUser(username string) (*models.User, error) {
	user := models.User{}
	err := u.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetUserById(id int) (*models.User, error) {
	user := models.User{}
	err := u.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetListUsers(id int) (*[]models.UserList, error) {
	var usersData []models.UserList
	err := u.db.Raw(`SELECT users.id, users.username, users.is_premium FROM users 
    LEFT JOIN swipes ON users.id = swipes.swipee_id AND swipes.swiper_id = ?
    WHERE swipes.swipee_id IS NULL AND NOT users.id = ?
    LIMIT 10`, id, id).Scan(&usersData).Error

	if err != nil {
		return nil, err
	}

	return &usersData, nil
}

func (u *userRepo) UpgradePremium(id int) error {
	return u.db.Model(&models.User{}).Where("id = ?", id).Update("is_premium", true).Error
}
