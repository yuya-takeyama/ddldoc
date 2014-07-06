package components

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DDLFetcher struct {
	dsn string
}

func NewDDLFetcher(dsn string) *DDLFetcher {
	return &DDLFetcher{
		dsn,
	}
}

func (self *DDLFetcher) Fetch(handler func(tableName string, ddlContent string)) error {
	db, err := sql.Open("mysql", self.dsn)

	if err != nil {
		return err
	}

	rows, err := db.Query("SHOW TABLES")

	if err != nil {
		return err
	}

	for rows.Next() {
		var name string
		err := rows.Scan(&name)

		if err != nil {
			return err
		}

		var tableName string
		var ddlContent string
		sql := fmt.Sprintf("SHOW CREATE TABLE `%s`", name)
		db.QueryRow(sql).Scan(&tableName, &ddlContent)

		go handler(tableName, ddlContent)
	}

	return nil
}
