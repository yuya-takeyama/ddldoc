package factories

import (
	"github.com/yuya-takeyama/ddldoc/entities"
)

type DDLFactory struct {
	ddlOptionFactory *DDLOptionFactory
}

func NewDDLFactory(ddlOptionFactory *DDLOptionFactory) *DDLFactory {
	return &DDLFactory{
		ddlOptionFactory,
	}
}

func (self *DDLFactory) Create(tableName string, content string) *entities.DDL {
	return entities.NewDDL(tableName, content, self.ddlOptionFactory.Create())
}
