package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseController struct {
	*gin.Context
}

func NewBaseController() *baseController {
	return &baseController{}
}
func (u *baseController) GetParamTypeNumber() (int64, bool) {
	return 0, false
}

func (u *baseController) Bind(req interface{}) bool {
	if err := u.ShouldBindBodyWithJSON(req); err != nil {
		u.JSON(http.StatusBadRequest, err)
		return false
	}

	return true
}
