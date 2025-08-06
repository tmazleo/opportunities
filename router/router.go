package router

import "github.com/gin-gonic/gin"


func Initialize() {
	//inicializar rota com configs default do gin
	router := gin.Default()
	//define uma rota
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message" : "pong",
		})
	})
	//rodar api
	router.Run(":8080")
}