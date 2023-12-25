package handlers

import (
	"errors"
	"musical_wiki/global"
	"musical_wiki/models"
	"musical_wiki/request"
	"musical_wiki/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ActorHandler struct {
	service service.ActorService
}

func (handler *ActorHandler) Index(c *gin.Context) {
	actors, err := handler.service.Index()
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusOK, "成功", map[string][]models.Actor{"actors": actors})
	})
}

func (handler *ActorHandler) Show(c *gin.Context) {
	id := c.Param("id")
	actor, err := handler.service.Show(id)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusOK, "成功", map[string]models.Actor{"actor": actor})
	})
}

func (handler *ActorHandler) Store(c *gin.Context) {
	var request request.Actor
	err := c.ShouldBind(&request)
	if err != nil {
		handleError(err, c)
		return
	}
	actor, err := handler.service.Store(&request)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusCreated, "成功", map[string]models.Actor{"actor": actor})
	})
}

func (handler *ActorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var request request.Actor
	if err := c.ShouldBind(&request); err != nil {
		handleError(err, c)
		return
	}
	err := handler.service.Update(id, &request)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusOK, "成功", nil)
	})
}

func (handler *ActorHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.Destroy(id)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusOK, "成功", nil)
	})
}

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
