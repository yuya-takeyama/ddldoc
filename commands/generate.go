package commands

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/codegangsta/cli"

	"github.com/yuya-takeyama/ddldoc/converters"
	"github.com/yuya-takeyama/ddldoc/factories"
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
	ddlFactory := factories.NewDDLFactory(factories.NewDDLOptionFactory(c))
	converter := getConverter(c)
	dsn := c.String("dsn")
	db, err := sql.Open("mysql", dsn)
	dieIfError(err, "Failed to connect to database")

	generateDocumentFiles(db, ddlFactory, func(ddl *entities.DDL) {
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

func generateDocumentFiles(db *sql.DB, ddlFactory *factories.DDLFactory, f func(*entities.DDL)) {
	rows, err := db.Query("SHOW TABLES")
	dieIfError(err, "Failed to fetch table list")

	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		dieIfError(err, "Failed to fetch table name")

		var tableName string
		var ddlContent string
		sql := fmt.Sprintf("SHOW CREATE TABLE `%s`", name)
		db.QueryRow(sql).Scan(&tableName, &ddlContent)

		ddl := ddlFactory.Create(tableName, ddlContent)

		go f(ddl)
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
