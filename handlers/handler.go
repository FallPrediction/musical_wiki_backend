package handlers

import (
	"errors"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	logger     *zap.SugaredLogger
	translator ut.Translator
}

func (handler *Handler) handleErrorAndReturn(err error, c *gin.Context, onSuccess func()) {
	if err != nil {
		handler.handleError(err, c)
		return
	}
	onSuccess()
}

func (handler *Handler) handleError(err error, c *gin.Context) {
	switch {
	case errors.As(err, &(validator.ValidationErrors{})):
		handler.sendResponse(c, http.StatusUnprocessableEntity, err.(validator.ValidationErrors).Translate(handler.translator), nil)
	case errors.Is(err, gorm.ErrRecordNotFound):
		handler.sendResponse(c, http.StatusNotFound, "伺服器找不到請求的資源", nil)
	default:
		handler.logger.Error("db error", err)
		handler.sendResponse(c, http.StatusUnprocessableEntity, "系統錯誤", nil)
	}
}

func (handler *Handler) sendResponse(c *gin.Context, statusCode int, message interface{}, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func (handler *Handler) sendResponseWithPagination(c *gin.Context, statusCode int, message interface{}, data interface{}, currentPage int, perPage int, total int) {
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

func NewBaseHandler(logger *zap.SugaredLogger, translator ut.Translator) Handler {
	return Handler{logger: logger, translator: translator}
}
