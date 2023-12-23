package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"musical_wiki/global"
	"net/http"
	"time"
)

type Actor struct {
	Id             uint32         `json:"id"`
	Name           string         `json:"name" validate:"required"`
	TranslatedName string         `json:"translated_name"`
	NickName       string         `json:"nick_name"`
	Nationality    string         `json:"nationality"`
	Born           string         `json:"born"`
	ImageId        *uint32        `json:"image_id"`
	Content        string         `json:"content"`
	Socials        datatypes.JSON `json:"socials" validate:"json,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type ActorHandler struct{}

func (handler *ActorHandler) Index(c *gin.Context) {
	var actors []Actor
	err := global.Db.Find(&actors).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
	}
	c.JSON(http.StatusOK, gin.H{"actors": actors})
}

func (handler *ActorHandler) Show(c *gin.Context) {
	id := c.Param("id")
	var actor Actor
	err := global.Db.Where("id = ?", id).First(&actor).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}

func (handler *ActorHandler) Store(c *gin.Context) {
	var actor Actor
	err := c.ShouldBind(&actor)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = global.Db.Create(&actor).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}

func (handler *ActorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var actor Actor
	err := global.Db.Where("id = ?", id).First(&actor).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
		return
	}
	err = c.ShouldBind(&actor)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = global.Db.Save(&actor).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}

func (hanler ActorHandler) Destroy(c *gin.Context) {
	id := c.Param("id")
	var actor Actor
	err := global.Db.Where("id = ?", id).First(&actor).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "actor not found"})
		return
	}
	err = global.Db.Delete(&actor).Error
	if err != nil {
		global.Logger.Error("db error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actor": actor})
}
