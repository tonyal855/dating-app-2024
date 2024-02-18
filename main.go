package main

import (
	"dating-app-dealls/db"
	"dating-app-dealls/server"
	"dating-app-dealls/server/controllers"
	"dating-app-dealls/server/repositories/repo"
	"os"
)

// import "dating-app-dealls/db"

func main() {
	db := db.ConnectDb()

	userRepo := repo.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo)

	swipeRepo := repo.NewSwipeRepo(db)
	swipeController := controllers.NewSwipeController(swipeRepo, userRepo)

	router := server.NewRouter(userController, swipeController)
	var PORT = os.Getenv("PORT")
	router.Start(":" + PORT)

}
