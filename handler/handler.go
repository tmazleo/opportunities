// as funções das rotas estão separadas em seus arq respectivos
package handler

import (
	"github.com/tmazleo/opportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}