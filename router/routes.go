//pacote com grupo de rotas acessados com o .Group() que define as paths de versao
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazleo/opportunities/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	//agrupation routes with version
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}