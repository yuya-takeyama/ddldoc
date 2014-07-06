package converters

import (
	"github.com/yuya-takeyama/ddldoc/entities"
)

type SQLConverter struct {
}

func (c *SQLConverter) Convert(ddl *entities.DDL) *entities.Document {
	fileName := ddl.GetTableName() + ".sql"

	return entities.NewDocument(fileName, ddl.GetContent())
}
