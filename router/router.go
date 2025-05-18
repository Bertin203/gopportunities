package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default() // initialize Router

	initializeRoutes(router) // inicialize Routess

	router.Run(":8080") // run the server
}
	