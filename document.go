package main

type Document struct {
	fileName string
	content  string
}

func NewDocument(fileName string, content string) *Document {
	return &Document{
		fileName,
		content,
	}
}
