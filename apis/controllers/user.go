package controllers

import (
	"demo_time_sheet_server/apis/entities"
	"demo_time_sheet_server/apis/resources"
	"demo_time_sheet_server/core/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user *services.UserService
	base *baseController
	reso *resources.Resource
}

func NewUserController(
	user *services.UserService,
	base *baseController,
	reso *resources.Resource,
) *UserController {
	return &UserController{
		user: user,
		base: base,
		reso: reso,
	}
}

func (u *UserController) Register(ctx *gin.Context) {
	var req entities.CreateUserRequest
	if !u.base.Bind(ctx, &req) {
		return
	}
	err := u.user.Register(ctx, &req)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.CreatedSuccess(ctx)
}

func (u *UserController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	token, err := u.user.Login(ctx, username, password)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.Response(ctx, token)
}

func (u *UserController) Profile(ctx *gin.Context) {

	userID, ok := u.base.GetUserID(ctx)
	if !ok {
		return
	}
	user, err := u.user.Profile(ctx, userID)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.Response(ctx, user)
}

func (u *UserController) LoginWithGG(ctx *gin.Context) {
	var req entities.LoginWithGGRequest
	if !u.base.Bind(ctx, &req) {
		return
	}
	token, err := u.user.LoginWithGG(ctx, req.Token)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.Response(ctx, token)
}
