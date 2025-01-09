package resources

import (
	"net/http"
	customerrors "rices/core/custom_errors"

	"github.com/gin-gonic/gin"
)

type Resource struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewResource() *Resource {
	return &Resource{}
}

func (u *Resource) CreatedSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "created successfully",
	})
}

func (u *Resource) DeletedSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "deleted successfully",
	})
}

func (u *Resource) UpdatedSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "updated successfully",
	})
}

func (u *Resource) ListAndCount(ctx *gin.Context, data interface{}, count int) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"count": count,
		"data":  data,
	})
}

func (u *Resource) Error(ctx *gin.Context, err *customerrors.CustomError) {
	ctx.JSON(err.Status, gin.H{
		"code":    err.Code,
		"message": err.Message,
	})
}

func (u *Resource) Response(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
