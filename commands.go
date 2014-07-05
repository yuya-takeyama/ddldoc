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
	},
}

func doGenerate(c *cli.Context) {
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
		var ddl string
		sql := fmt.Sprintf("SHOW CREATE TABLE `%s`", name)
		db.QueryRow(sql).Scan(&table, &ddl)

		file, err := os.OpenFile(fmt.Sprintf("%s.sql", name), os.O_CREATE | os.O_WRONLY, 0644)
		DieIfError(err, "Failed to open file")

		defer file.Close()

		_, err = file.WriteString(ddl)
		DieIfError(err, "Failed to write on file")
	}

	fmt.Println("Generated successfully")
}

func DieIfError(err error, m string) {
	if err != nil {
		fmt.Println(m)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
