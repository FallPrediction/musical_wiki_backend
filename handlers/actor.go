package handlers

import (
	"musical_wiki/models"
	"musical_wiki/request"
	"musical_wiki/resource"
	"musical_wiki/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActorHandler struct {
	baseHandler Handler
	service     service.ActorService
}

func (handler *ActorHandler) Index(c *gin.Context) {
	var request request.IndexActor
	err := c.ShouldBind(&request)
	if err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	if request.CurrentPage == nil {
		page := 1
		request.CurrentPage = &page
	}
	if request.PerPage == nil {
		perPage := 12
		request.PerPage = &perPage
	}

	actors, count, err := handler.service.Index(*request.CurrentPage, *request.PerPage)

	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		resource := resource.ActorSliceResource{ModelSlice: actors}
		handler.baseHandler.sendResponseWithPagination(c, http.StatusOK, "成功", resource.ToSliceResource(), *request.CurrentPage, *request.PerPage, int(count))
	})
}

func (handler *ActorHandler) Show(c *gin.Context) {
	id := c.Param("id")
	actor, err := handler.service.Show(id)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		resource := resource.ActorResource{Model: actor}
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", resource.ToMap())
	})
}

func (handler *ActorHandler) Store(c *gin.Context) {
	var request request.Actor
	err := c.ShouldBind(&request)
	if err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	actor, err := handler.service.Store(&request)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "成功", map[string]models.Actor{"actor": actor})
	})
}

func (handler *ActorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var request request.Actor
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	err := handler.service.Update(id, &request)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", nil)
	})
}

func (handler *ActorHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.Destroy(id)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", nil)
	})
}

func NewActorHandler(baseHandler Handler, service service.ActorService) ActorHandler {
	return ActorHandler{baseHandler: baseHandler, service: service}
}
