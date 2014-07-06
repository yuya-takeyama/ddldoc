package converters

import (
	"github.com/yuya-takeyama/ddldoc/domain"
)

type Converter interface {
	Convert(*domain.DDL) *domain.Document
}
