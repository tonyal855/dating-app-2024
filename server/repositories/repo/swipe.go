package repo

import (
	"dating-app-dealls/server/models"
	"dating-app-dealls/server/repositories"
	"time"

	"gorm.io/gorm"
)

type swipeRepo struct {
	db *gorm.DB
}

func NewSwipeRepo(db *gorm.DB) repositories.SwipeRepo {
	return &swipeRepo{db: db}
}

func (s *swipeRepo) CreateSwipe(swipe *models.Swipe) error {
	return s.db.Create(swipe).Error
}

func (s *swipeRepo) CountLimit(id int) int {
	now := time.Now()
	formattedDate := now.Format("2006-01-02")
	startDate := formattedDate + " 00:00:00"
	endDate := formattedDate + " 23:59:59"
	var count int64

	s.db.Model(models.Swipe{}).Where("swiper_id = ? AND created_at >= ? AND created_at <= ?", id, startDate, endDate).Count(&count)
	return int(count)
}
