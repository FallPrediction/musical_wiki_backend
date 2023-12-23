package handlers

import (
	"github.com/gin-gonic/gin"
	"musical_wiki/global"
	"musical_wiki/request"
	"musical_wiki/service"
	"net/http"
)

type ActorHandler struct {
	service service.ActorService
}

func (handler *ActorHandler) Index(c *gin.Context) {
	actors, err := handler.service.Index()
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
	}
	c.JSON(http.StatusOK, gin.H{"actors": actors})
}

func (handler *ActorHandler) Show(c *gin.Context) {
	id := c.Param("id")
	actor, err := handler.service.Show(id)
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}

func (handler *ActorHandler) Store(c *gin.Context) {
	var request request.Actor
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	actor, err := handler.service.Store(&request)
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}

func (handler *ActorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var request request.Actor
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	actor, err := handler.service.Update(id, &request)
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}

func (handler *ActorHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.Destroy(id)
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sccuess"})
}
