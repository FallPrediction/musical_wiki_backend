package handlers

import (
	"errors"
	"math"
	"musical_wiki/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func handleErrorAndReturn(err error, c *gin.Context, onSuccess func()) {
	if err != nil {
		handleError(err, c)
		return
	}
	onSuccess()
}

func handleError(err error, c *gin.Context) {
	switch {
	case errors.As(err, &(validator.ValidationErrors{})):
		sendResponse(c, http.StatusUnprocessableEntity, err.(validator.ValidationErrors).Translate(global.Translator), nil)
	case errors.Is(err, gorm.ErrRecordNotFound):
		sendResponse(c, http.StatusNotFound, "伺服器找不到請求的資源", nil)
	default:
		global.Logger.Error("db error", err)
		sendResponse(c, http.StatusUnprocessableEntity, "系統錯誤", nil)
	}
}

func sendResponse(c *gin.Context, statusCode int, message interface{}, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func sendResponseWithPagination(c *gin.Context, statusCode int, message interface{}, data interface{}, currentPage int, perPage int, total int) {
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
		"pagination": gin.H{
			"currentPage": currentPage,
			"lastPage":    lastPage,
			"total":       total,
			"perPage":     perPage,
		},
	})
}
