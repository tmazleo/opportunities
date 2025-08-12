//pacote princippal tendo a função de inicializar as demais
package main

import (

	"github.com/tmazleo/opportunities/config"
	"github.com/tmazleo/opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize config
	err := config.Init()
	if err != nil {
		logger.ErrorF("Config initialization error: %v",err)
		return
	}

	//initialize router
	router.Initialize()

}