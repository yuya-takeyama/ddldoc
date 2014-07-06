package converters

import (
	"github.com/yuya-takeyama/ddldoc/entities"
)

type Converter interface {
	Convert(*entities.DDL) *entities.Document
}
