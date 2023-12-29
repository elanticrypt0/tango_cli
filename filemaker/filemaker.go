package filemaker

import "strings"

type FileMaker struct {
	Filename         string
	Filetypes        []string
	Filenames        []string
	RootPath         string
	ForcedMode       bool
	TemplateSelected string
}

func New() *FileMaker {
	return &FileMaker{
		ForcedMode: false,
	}
}

func (fm *FileMaker) setFileNames() {

	ftypes := []string{
		"templ",
		"go",
	}

	fm.Filetypes = ftypes

}

func (fm *FileMaker) setFileTypes() {

	ftypes := []string{
		"templ",
		"go",
	}

	fm.Filetypes = ftypes

}
func (fm *FileMaker) setFilename(filename string) {
	fm.Filename = strings.ToLower(filename)
}

func (fm *FileMaker) SetRootPath(path string) {
	// todo
}

func (fm *FileMaker) SetForcedMode(mode bool) {
	fm.ForcedMode = mode
}

func (fm *FileMaker) CheckIfFileExists(filepath string) {}

func (fm *FileMaker) Maker(filepath, template string) {}

func (fm *FileMaker) SelectTemplate(template string) {}
