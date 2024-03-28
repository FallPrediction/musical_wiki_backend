package handlers

import (
	"musical_wiki/models"
	"musical_wiki/request"
	"musical_wiki/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreditHandler struct {
	baseHandler Handler
	service     service.CreditService
}

func (handler *CreditHandler) Store(c *gin.Context) {
	var request request.Credit
	err := c.ShouldBind(&request)
	if err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	credit, err := handler.service.Store(&request)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "成功", map[string]models.Credit{"credit": credit})
	})
}

func (handler *CreditHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var request request.Credit
	err := c.ShouldBind(&request)
	if err != nil {
		handler.baseHandler.handleError(err, c)
		return
	}
	err = handler.service.Update(id, &request)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", nil)
	})
}

func (handler *CreditHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.Destroy(id)
	handler.baseHandler.handleErrorAndReturn(err, c, func() {
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", nil)
	})
}

func NewCreditHandler(baseHandler Handler, service service.CreditService) CreditHandler {
	return CreditHandler{baseHandler: baseHandler, service: service}
}
