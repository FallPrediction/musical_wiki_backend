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
	service service.ActorService
}

func (handler *ActorHandler) Index(c *gin.Context) {
	var request request.IndexActor
	err := c.ShouldBind(&request)
	if err != nil {
		handleError(err, c)
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

	handleErrorAndReturn(err, c, func() {
		resource := resource.ActorSliceResource{ModelSlice: actors}
		sendResponseWithPagination(c, http.StatusOK, "成功", resource.ToSliceResource(), *request.CurrentPage, *request.PerPage, int(count))
	})
}

func (handler *ActorHandler) Show(c *gin.Context) {
	id := c.Param("id")
	actor, err := handler.service.Show(id)
	handleErrorAndReturn(err, c, func() {
		resource := resource.ActorResource{Model: actor}
		sendResponse(c, http.StatusOK, "成功", resource.ToMap())
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
