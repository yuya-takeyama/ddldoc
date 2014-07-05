package main

type SQLConverter struct {
}

func (self *SQLConverter) Convert(ddl *DDL) *Document {
	fileName := ddl.name + ".sql"

	return NewDocument(fileName, ddl.GetContent())
}
