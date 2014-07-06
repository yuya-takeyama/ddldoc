package converters

import (
	"github.com/yuya-takeyama/ddldoc/domain"
)

type SQLConverter struct {
}

func (c *SQLConverter) Convert(ddl *domain.DDL) *domain.Document {
	fileName := ddl.GetTableName() + ".sql"

	return domain.NewDocument(fileName, ddl.GetContent())
}
