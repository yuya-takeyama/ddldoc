package main

import (
	"regexp"
)

type DDL struct {
	content string
	option  *DDLOption
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

func (self *DDL) GetContent() string {
	var result string

	if self.option.withAutoIncrement == false {
		re := regexp.MustCompile("AUTO_INCREMENT=\\d+ ")
		result = re.ReplaceAllLiteralString(self.content, "")
	} else {
		result = self.content
	}

	result += "\n"

	return result
}
