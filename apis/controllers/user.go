package controllers

import (
	"log"
	"rices/apis/entities"
	"rices/apis/resources"
	"rices/core/services"

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
		log.Println("error")
		u.reso.Error(ctx, err)
		return
	}

	u.reso.CreatedSuccess(ctx)
}
