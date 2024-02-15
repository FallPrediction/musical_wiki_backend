package router

import (
	"musical_wiki/handlers"
	"musical_wiki/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	apis := router.Group("api")
	{
		actors := apis.Group("actor")
		actorHandler := handlers.ActorHandler{}
		{
			actors.GET("", actorHandler.Index)
			actors.GET("/:id", actorHandler.Show)
			actors.POST("", actorHandler.Store)
			actors.PUT("/:id", actorHandler.Update)
			actors.DELETE("/:id", actorHandler.Destroy)
		}
		credits := apis.Group("credit")
		creditHandler := handlers.CreditHandler{}
		{
			credits.POST("", creditHandler.Store)
			credits.PUT("/:id", creditHandler.Update)
			credits.DELETE("/:id", creditHandler.Destroy)
		}
	}
	router.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	return router
}
