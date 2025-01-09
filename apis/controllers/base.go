package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type baseController struct {
	*gin.Context
}

func NewBaseController() *baseController {
	return &baseController{}
}

func (bc *baseController) GetParamTypeNumber(param string) (int64, bool) {
	paramValue := bc.Param(param)
	if paramValue == "" {
		return 0, false
	}

	num, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		return 0, false
	}

	return num, true
}

func (u *baseController) Bind(req interface{}) bool {
	if err := u.ShouldBindBodyWithJSON(req); err != nil {
		u.JSON(http.StatusBadRequest, err)
		return false
	}

	return true
}
