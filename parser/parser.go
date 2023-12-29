package parser

type Parser struct {
	nameInput    string
	NamePlural   string
	NameSingular string
	FirstChar    string
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Read(name string) {

	p.nameInput = name

	// Convert the input
	p.ConvertToPlural()
	p.ConvertToSnakeCase()
	p.GetFirstChar()

}

func (p *Parser) ConvertToPlural() {

	// todo
	// si termina en vocal agrega S
	// si termina en Y agrega "ies"

}

func (p *Parser) ConvertToSnakeCase() {

	// Todo

}

func (p *Parser) GetFirstChar() {
	p.FirstChar = p.nameInput[0:1]
}
