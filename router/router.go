package router

import (
	"musical_wiki/handlers"
	"musical_wiki/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(actorHandler handlers.ActorHandler, creditHandler handlers.CreditHandler, imageHandler handlers.ImageHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	apis := router.Group("api")
	{
		actors := apis.Group("actor")
		{
			actors.GET("", actorHandler.Index)
			actors.GET("/:id", actorHandler.Show)
			actors.POST("", actorHandler.Store)
			actors.PUT("/:id", actorHandler.Update)
			actors.DELETE("/:id", actorHandler.Destroy)
		}
		credits := apis.Group("credit")
		{
			credits.POST("", creditHandler.Store)
			credits.PUT("/:id", creditHandler.Update)
			credits.DELETE("/:id", creditHandler.Destroy)
		}
		images := apis.Group("image")
		{
			images.POST("/avatar", imageHandler.StoreAvatar)
			images.POST("/gallery", imageHandler.StoreGallery)
			images.DELETE("/:id", imageHandler.Destroy)
		}
	}
	router.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	return router
}
