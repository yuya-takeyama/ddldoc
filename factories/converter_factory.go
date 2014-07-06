package factories

import (
	"github.com/codegangsta/cli"

	"github.com/yuya-takeyama/ddldoc/converters"
)

type ConverterFactory struct {
	cliContext *cli.Context
}

func NewConverterFactory(cliContext *cli.Context) *ConverterFactory {
	return &ConverterFactory{
		cliContext,
	}
}

func (self *ConverterFactory) Create() converters.Converter {
	return converters.NewSQLConverter()
}
