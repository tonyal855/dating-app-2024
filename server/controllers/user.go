package controllers

import (
	"dating-app-dealls/server/helper"
	"dating-app-dealls/server/models"
	"dating-app-dealls/server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlUserList struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	IsPremium bool   `json:"is_premium"`
}

func WriteJsonResponse(ctx *gin.Context, payload *helper.Response) {
	ctx.JSON(payload.Status, payload)
}

type UserController struct {
	userRepo repositories.UserRepo
}

func NewUserController(userRepo repositories.UserRepo) *UserController {
	return &UserController{userRepo: userRepo}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req models.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Check email and username
	_, errU := u.userRepo.GetUserByUser(req.Username)
	if errU == nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "username is exist",
		})
		return
	}

	_, errE := u.userRepo.GetUserByEmail(req.Email)
	if errE == nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "email is exist",
		})
		return
	}

	req.Password = helper.GeneratePasswordBrypt(req.Password)
	errs := u.userRepo.CreateUser(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_USER_FAIL",
			Error:   errs.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusCreated,
		Message: "CREATE_USER_SUCCESS",
		Payload: "",
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var req models.ReqLogin
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, errU := u.userRepo.GetUserByEmail(req.Email)
	if errU != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong Email",
		})
		return
	}

	comparePass := helper.ComparePwd([]byte(user.Password), []byte(req.Password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong Password",
		})
		return
	}
	token, errT := helper.GenerateToken(user.ID, user.Email)
	if errT != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errT.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status: http.StatusOK,
		Token:  token,
	})
}

func (u *UserController) GetListUsers(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	getUsers, err := u.userRepo.GetListUsers(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	println(len(*getUsers))
	if len(*getUsers) <= 0 {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusOK,
			Payload: make([]string, 0),
		})
		return
	}
	var data []handlUserList
	for _, user := range *getUsers {
		data = append(data, handlUserList{Id: user.Id, Username: user.Username, IsPremium: user.IsPremium})
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Payload: data,
	})
}

func (u *UserController) UpgradePremium(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)
	err := u.userRepo.UpgradePremium(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PREMIUM_SUCCESS",
		Payload: "",
	})
}
