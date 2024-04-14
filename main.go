package main

import (
	"musical_wiki/config"
	"musical_wiki/handlers"
	"musical_wiki/initialize"
	"musical_wiki/repository"
	"musical_wiki/router"
	"musical_wiki/service"
	"musical_wiki/utils"
)

func main() {
	logger := initialize.NewLogger()
	redis := initialize.NewRedis(logger)
	db := initialize.NewDB(config.NewPg(), logger)
	s3 := initialize.NewS3(logger)
	translator := initialize.NewTranslator(logger)
	uploader := utils.NewUploader(s3, logger)

	baseHandler := handlers.NewBaseHandler(logger, translator)

	creditRepository := repository.NewCreditRepository(db)
	creditService := service.NewCreditService(creditRepository, logger, redis)
	creditHandler := handlers.NewCreditHandler(baseHandler, creditService)

	imageRepository := repository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepository, logger, redis, uploader)
	imageHandler := handlers.NewImageHandler(baseHandler, imageService)

	actorRepository := repository.NewActorRepository(db)
	actorService := service.NewActorService(actorRepository, logger, redis, creditService, imageService)
	actorHandler := handlers.NewActorHandler(baseHandler, actorService)

	r := router.NewRouter(actorHandler, creditHandler, imageHandler)

	r.Run()
}
