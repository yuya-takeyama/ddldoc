package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/yuya-takeyama/ddldoc/components"
	"github.com/yuya-takeyama/ddldoc/factories"
)

var Generate = cli.Command{
	Name:  "generate",
	Usage: "Generates documentation files from DDL",
	Description: `
	Generates documentation files from DDL
`,
	Action: doGenerate,
	Flags: []cli.Flag{
		cli.StringFlag{"dsn", "", "Data source name"},
		cli.StringFlag{"dir", "", "Target directory where the document is generated into"},
		cli.BoolFlag{"with-auto-increment", "Whether DDL contains AUTO_INCREMENT count (omitted by default)"},
	},
}

func doGenerate(c *cli.Context) {
	ddlFactory := factories.NewDDLFactory(factories.NewDDLOptionFactory(c))
	ddlFetcher := components.NewDDLFetcher(c.String("dsn"))
	converter := factories.NewConverterFactory(c).Create()
	documentFileGenerator := factories.NewDocumentFileGeneratorFactory(c).Create()

	err := ddlFetcher.Fetch(func(tableName string, ddlContent string) {
		ddl := ddlFactory.Create(tableName, ddlContent)
		document := converter.Convert(ddl)
		err := documentFileGenerator.Generate(document)

		dieIfError(err, "Failed to generate document file")

		fmt.Printf("Generated %s from %s\n", document.GetFileName(), ddl.GetTableName())
	})

	dieIfError(err, "Failed to fetch DDL")

	fmt.Println("Finished successfully")
}

func dieIfError(err error, m string) {
	if err != nil {
		fmt.Println(m)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
