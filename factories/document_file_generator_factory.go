package factories

import (
	"github.com/codegangsta/cli"

	"github.com/yuya-takeyama/ddldoc/components"
)

type DocumentFileGeneratorFactory struct {
	cliContext *cli.Context
}

func NewDocumentFileGeneratorFactory(cliContext *cli.Context) *DocumentFileGeneratorFactory {
	return &DocumentFileGeneratorFactory{
		cliContext,
	}
}

func (self *DocumentFileGeneratorFactory) Create() *components.DocumentFileGenerator {
	return components.NewDocumentFileGenerator(self.cliContext.String("dir"))
}
