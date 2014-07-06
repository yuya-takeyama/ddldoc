package commands

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/codegangsta/cli"

	"github.com/yuya-takeyama/ddldoc/converters"
	"github.com/yuya-takeyama/ddldoc/entities"
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
	converter := getConverter(c)

	generateDocumentFiles(c, converter, func(converter converters.Converter, ddl *entities.DDL) {
		document := converter.Convert(ddl)

		file, err := os.OpenFile(getFilePath(c, document.GetFileName()), os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
		dieIfError(err, "Failed to open file")

		defer file.Close()

		_, err = file.WriteString(document.GetContent())
		dieIfError(err, "Failed to write on file")

		fmt.Printf("Generated %s from %s\n", document.GetFileName(), ddl.GetTableName())
	})

	fmt.Println("Finished successfully")
}

func generateDocumentFiles(c *cli.Context, converter converters.Converter, f func(converters.Converter, *entities.DDL)) {
	dsn := c.String("dsn")
	db, err := sql.Open("mysql", dsn)
	dieIfError(err, "Failed to connect to database")

	rows, err := db.Query("SHOW TABLES")
	dieIfError(err, "Failed to fetch table list")

	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		dieIfError(err, "Failed to fetch table name")

		var table string
		var ddlString string
		sql := fmt.Sprintf("SHOW CREATE TABLE `%s`", name)
		db.QueryRow(sql).Scan(&table, &ddlString)

		ddl := entities.NewDDL(table, ddlString, entities.NewDDLOption(c.Bool("with-auto-increment")))

		go f(converter, ddl)
	}
}

func dieIfError(err error, m string) {
	if err != nil {
		fmt.Println(m)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getFilePath(c *cli.Context, fileName string) string {
	var dir string
	if len(c.String("dir")) > 0 {
		dir = c.String("dir")
	} else {
		dir = "."
	}

	return fmt.Sprintf("%s/%s", dir, fileName)
}

func getConverter(c *cli.Context) converters.Converter {
	var converter converters.Converter = &converters.SQLConverter{}

	return converter
}
