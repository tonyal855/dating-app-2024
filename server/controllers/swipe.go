package controllers

import (
	"dating-app-dealls/server/helper"
	"dating-app-dealls/server/models"
	"dating-app-dealls/server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SwipeController struct {
	swipeRepo repositories.SwipeRepo
	userRepo  repositories.UserRepo
}

func NewSwipeController(swipeRepo repositories.SwipeRepo, userRepo repositories.UserRepo) *SwipeController {
	return &SwipeController{
		swipeRepo: swipeRepo,
		userRepo:  userRepo,
	}
}

func (swp *SwipeController) CreateSwipe(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	var req models.Swipe
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Check premium
	getUser, errU := swp.userRepo.GetUserById(id)
	if errU != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_SWIPE_FAIL",
		})
		return
	}

	println(getUser.IsPremium)
	if !getUser.IsPremium {
		//Check limit per day
		limit := swp.swipeRepo.CountLimit(id)
		if limit >= 10 {
			WriteJsonResponse(ctx, &helper.Response{
				Status:  http.StatusInternalServerError,
				Message: "MAX_LIMIT_PER_DAY",
			})
			return
		}

	}

	req.SwiperID = id
	errs := swp.swipeRepo.CreateSwipe(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_SWIPE_FAIL",
			Error:   errs.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "CREATE_SWIPE_SUCCESS",
		Payload: req,
	})

}
