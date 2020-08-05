package main

import (
	"fmt"

	"github.com/fomiller/Go-Coverletter-CLI/docs"

	"github.com/fomiller/Go-Coverletter-CLI/drive"
)

func main() {
	templateCopyId := drive.CreateTemplateCopy()
	fmt.Println("This is the template Copy id: ", templateCopyId)
	templateId := docs.UpdateTemplateFile(templateCopyId)

	fmt.Println("Successfully updated document: ", templateId)
}
