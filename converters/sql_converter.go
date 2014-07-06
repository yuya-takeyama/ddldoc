package converters

import (
	"github.com/yuya-takeyama/ddldoc/entity"
)

type SQLConverter struct {
}

func (c *SQLConverter) Convert(ddl *entity.DDL) *entity.Document {
	fileName := ddl.GetTableName() + ".sql"

	return entity.NewDocument(fileName, ddl.GetContent())
}
