package router

import (
	"github.com/gin-gonic/gin"
	"musical_wiki/handlers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	apis := router.Group("api")
	{
		actors := apis.Group("actor")
		actorHandler := handlers.ActorHandler{}
		{
			actors.GET("/", actorHandler.Index)
			actors.GET("/:id", actorHandler.Show)
			actors.POST("/", actorHandler.Store)
			actors.PUT("/:id", actorHandler.Update)
			actors.DELETE("/:id", actorHandler.Destroy)
		}
	}
	return router
}
