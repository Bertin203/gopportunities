package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default() // inicializa o Router utilizando as configurações Default do gin

  	router.GET("/ping", func(c *gin.Context) { // define uma rota
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
