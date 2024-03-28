package handlers

import (
	"musical_wiki/request"
	"musical_wiki/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	baseHandler Handler
	service     service.ImageService
}

func (handler *ImageHandler) StoreAvatar(c *gin.Context) {
	var request request.Image
	err := c.ShouldBind(&request)
	if err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	_, err = handler.service.UpdateAvatar(&request)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "成功", nil)
	})
}

func (handler *ImageHandler) StoreGallery(c *gin.Context) {
	var request request.Image
	err := c.ShouldBind(&request)
	if err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	_, err = handler.service.StoreGallery(&request)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "成功", nil)
	})
}

func (handler *ImageHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.Destroy(id)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", nil)
	})
}

func NewImageHandler(baseHandler Handler, service service.ImageService) ImageHandler {
	return ImageHandler{baseHandler: baseHandler, service: service}
}
