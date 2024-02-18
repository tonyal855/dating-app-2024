package server

import (
	"dating-app-dealls/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userRouter  *controllers.UserController
	swipeRouter *controllers.SwipeController
}

func NewRouter(useRouter *controllers.UserController, swipeRouter *controllers.SwipeController) *Router {
	return &Router{
		userRouter:  useRouter,
		swipeRouter: swipeRouter,
	}
}

func (r *Router) Start(port string) {
	router := gin.Default()
	router.POST("/users/register", r.userRouter.CreateUser)
	router.POST("/login", r.userRouter.Login)
	router.GET("/users/get", CheckAuth, r.userRouter.GetListUsers)
	router.POST("/users/upgrade-premium", CheckAuth, r.userRouter.UpgradePremium)

	router.POST("/swipe", CheckAuth, r.swipeRouter.CreateSwipe)

	router.Run(port)
}
