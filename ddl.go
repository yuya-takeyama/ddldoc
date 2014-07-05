package main

type DDL struct {
	content string
	context *DDLOption
}

type DDLOption struct {
	withAutoIncrement bool
}

func NewDDL(content string, ddlOption *DDLOption) *DDL {
	return &DDL{
		content,
		ddlOption,
	}
}
