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
		"message": "Resource created successfully",
	})
}

func (u *Resource) DeletedSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Resource deleted successfully",
	})
}

func (u *Resource) UpdatedSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Resource updated successfully",
	})
}

func (u *Resource) ListAndCount(ctx *gin.Context, data interface{}, count int) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"count": count,
	})
}

func (u *Resource) Error(ctx *gin.Context, err *customerrors.CustomError) {
	ctx.JSON(err.Status, gin.H{
		"code":    err.Code,
		"message": err.Message,
	})
}
