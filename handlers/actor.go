package handlers

import (
	"musical_wiki/models"
	"musical_wiki/request"
	"musical_wiki/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
