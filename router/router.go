package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	//initialize router
	router := gin.Default()

	//initialize routes
	initializeRoutes(router)

	//run the server
	router.Run() // listen and serve on 0.0.0.0:8080
}
