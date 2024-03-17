package handlers

import (
	"musical_wiki/request"
	"musical_wiki/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	service service.ImageService
}

func (handler *ImageHandler) StoreAvatar(c *gin.Context) {
	var request request.Image
	err := c.ShouldBind(&request)
	if err != nil {
		handleError(err, c)
		return
	}
	_, err = handler.service.UpdateAvatar(&request)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusCreated, "成功", nil)
	})
}

func (handler *ImageHandler) StoreGallery(c *gin.Context) {
	var request request.Image
	err := c.ShouldBind(&request)
	if err != nil {
		handleError(err, c)
		return
	}
	_, err = handler.service.StoreGallery(&request)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusCreated, "成功", nil)
	})
}

func (handler *ImageHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.Destroy(id)
	handleErrorAndReturn(err, c, func() {
		sendResponse(c, http.StatusOK, "成功", nil)
	})
}
