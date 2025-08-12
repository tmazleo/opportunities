//utilizando o gorm para criar tabelas no banco.
//gorm.Model necessario para definir os campos
package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
