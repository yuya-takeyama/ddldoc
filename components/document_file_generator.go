package components

import (
	"fmt"
	"os"

	"github.com/yuya-takeyama/ddldoc/entities"
)

type DocumentFileGenerator struct {
	targetDirectory string
}

func NewDocumentFileGenerator(targetDirectory string) *DocumentFileGenerator {
	return &DocumentFileGenerator{
		targetDirectory,
	}
}

func (self *DocumentFileGenerator) Generate(document *entities.Document) error {
		file, err := os.OpenFile(self.getFilePath(document), os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)

		if err != nil {
			return err
		}

		defer file.Close()

		_, err = file.WriteString(document.GetContent())

		return err

}

func (self *DocumentFileGenerator) getFilePath(document *entities.Document) string {
	var dir string
	if len(self.targetDirectory) > 0 {
		dir = self.targetDirectory
	} else {
		dir = "."
	}

	return fmt.Sprintf("%s/%s", dir, document.GetFileName())
}
