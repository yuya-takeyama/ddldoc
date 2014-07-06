package domain

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

func (self *Document) GetFileName() string {
	return self.fileName
}

func (self *Document) GetContent() string {
	return self.content
}
