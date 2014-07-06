package factories

import (
	"github.com/codegangsta/cli"

	"github.com/yuya-takeyama/ddldoc/entities"
)

type DDLOptionFactory struct {
	cliContext *cli.Context
}

func NewDDLOptionFactory(cliContext *cli.Context) *DDLOptionFactory {
	return &DDLOptionFactory{
		cliContext,
	}
}

func (factory *DDLOptionFactory) Create() *entities.DDLOption {
	return entities.NewDDLOption(factory.cliContext.Bool("with-auto-increment"))
}
