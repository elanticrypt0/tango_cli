package filemaker

import "strings"

type FileMaker struct {
	Filename         string   `json:"filename"`
	Filetypes        []string `json:"filetypes"`
	Filenames        []string `json:"filenames"`
	RootPath         string   `json:"root_path"`
	ForcedMode       bool     `json:"forced_mode"`
	TemplateSelected string   `json:"template_selected"`
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
