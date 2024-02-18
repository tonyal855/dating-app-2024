package db

import (
	"dating-app-dealls/server/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")
	fmt.Println(host, port, user, pass, dbname)

	confPg := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, errDb := gorm.Open(postgres.Open(confPg))
	if errDb != nil {
		panic(errDb)
	}
	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.Swipe{})

	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		// Insert seed data if no users exist
		if err := insertSeedData(db); err != nil {
			panic(err)
		}
	}
	return db
}

func insertSeedData(db *gorm.DB) error {
	seedUsers := []models.User{
		{Username: "user1", Email: "user1@example.com", Password: "password1"},
		{Username: "user2", Email: "user2@example.com", Password: "password2"},
		{Username: "user3", Email: "user3@example.com", Password: "password2"},
		{Username: "user4", Email: "user4@example.com", Password: "password2"},
		{Username: "user5", Email: "user5@example.com", Password: "password2"},
		{Username: "user6", Email: "user6@example.com", Password: "password2"},
		{Username: "user7", Email: "user7@example.com", Password: "password2"},
		{Username: "user8", Email: "user8@example.com", Password: "password2"},
		{Username: "user9", Email: "user9@example.com", Password: "password2"},
		{Username: "user10", Email: "user10@example.com", Password: "password2"},
		{Username: "user11", Email: "user11@example.com", Password: "password2"},
		{Username: "user12", Email: "user12@example.com", Password: "password2"},
		{Username: "user13", Email: "user13@example.com", Password: "password2"},
		{Username: "user14", Email: "user14@example.com", Password: "password2"},
		{Username: "user15", Email: "user15@example.com", Password: "password2"},
	}

	if err := db.Create(&seedUsers).Error; err != nil {
		return err
	}

	return nil
}
