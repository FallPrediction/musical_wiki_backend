package main

import (
	"musical_wiki/initialize"
	"musical_wiki/router"
)

func main() {
	initialize.InitLogger()
	initialize.InitDatabase()
	initialize.InitTranslator()
	initialize.InitRedis()

	r := router.InitRouter()

	r.Run()
}
