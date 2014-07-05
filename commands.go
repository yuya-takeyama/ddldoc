package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGenerate,
}

var commandGenerate = cli.Command{
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
	converter := GetConverter(c)

	GenerateDocumentFiles(c, converter, func(converter *SQLConverter, ddl *DDL) {
		document := converter.Convert(ddl)

		file, err := os.OpenFile(FilePath(c, document.fileName), os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
		DieIfError(err, "Failed to open file")

		defer file.Close()

		_, err = file.WriteString(document.content)
		DieIfError(err, "Failed to write on file")

		fmt.Printf("Generated %s from %s\n", document.fileName, ddl.name)
	})

	fmt.Println("Generated successfully")
}

func GenerateDocumentFiles(c *cli.Context, converter *SQLConverter, f func(*SQLConverter, *DDL)) {
	dsn := c.String("dsn")
	db, err := sql.Open("mysql", dsn)
	DieIfError(err, "Failed to connect to database")

	rows, err := db.Query("SHOW TABLES")
	DieIfError(err, "Failed to fetch table list")

	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		DieIfError(err, "Failed to fetch table name")

		var table string
		var ddlString string
		sql := fmt.Sprintf("SHOW CREATE TABLE `%s`", name)
		db.QueryRow(sql).Scan(&table, &ddlString)

		ddl := NewDDL(table, ddlString, NewDDLOption(c))

		go f(converter, ddl)
	}
}

func DieIfError(err error, m string) {
	if err != nil {
		fmt.Println(m)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func FilePath(c *cli.Context, fileName string) string {
	var dir string
	if len(c.String("dir")) > 0 {
		dir = c.String("dir")
	} else {
		dir = "."
	}

	return fmt.Sprintf("%s/%s", dir, fileName)
}

func NewDDLOption(c *cli.Context) *DDLOption {
	return &DDLOption{
		c.Bool("with-auto-increment"),
	}
}

func GetConverter(c *cli.Context) *SQLConverter {
	return &SQLConverter{}
}
