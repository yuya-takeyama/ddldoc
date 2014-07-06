package converters

import (
	"github.com/yuya-takeyama/ddldoc/entity"
)

type Converter interface {
	Convert(*entity.DDL) *entity.Document
}
