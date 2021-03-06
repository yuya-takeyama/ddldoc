package entities

import (
	"regexp"
)

type DDL struct {
	name    string
	content string
	option  *DDLOption
}

func NewDDL(name string, content string, ddlOption *DDLOption) *DDL {
	return &DDL{
		name,
		content,
		ddlOption,
	}
}

func (self *DDL) GetTableName() string {
	return self.name
}

func (self *DDL) GetContent() string {
	var result string

	if self.option.IsWithAutoIncrement() == false {
		re := regexp.MustCompile("AUTO_INCREMENT=\\d+ ")
		result = re.ReplaceAllLiteralString(self.content, "")
	} else {
		result = self.content
	}

	result += "\n"

	return result
}
