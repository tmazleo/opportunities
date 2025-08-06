package router

import "github.com/gin-gonic/gin"


func Initialize() {

	//initialize router with configs default gin
	router := gin.Default()
	
	//initialize routes
	initializeRoutes(router)

	//run api server
	router.Run(":8080")
}